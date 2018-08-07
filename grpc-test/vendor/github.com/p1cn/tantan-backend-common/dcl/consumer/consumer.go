package consumer

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/golang/protobuf/proto"
	"github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/dcl/eventmeta"
	"github.com/p1cn/tantan-backend-common/dcl/processor"
	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/util/tracing"
	"github.com/p1cn/tantan-domain-schema/golang/event"
)

func printDebug(s string, args ...interface{}) {
	slog.Debug(s, args...)
}

// ConsumerConfig ...
type ConsumerConfig struct {
	ClientID    string // TODO:
	ServiceName string
	Topic       string
	Group       string
	KafkaConfig config.KafkaConfig
	Processor   processor.EventProcessor
	Errors      chan error
}

// NewConsumer ...
func NewConsumer(cfg ConsumerConfig) (*Consumer, error) {
	config := cluster.NewConfig()
	config.Group.Mode = cluster.ConsumerModePartitions
	config.Version = sarama.V0_10_0_0

	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	if cfg.KafkaConfig.Consumer != nil {
		if cfg.KafkaConfig.Consumer.Retry != nil {
			config.Consumer.Retry.Backoff = time.Duration(cfg.KafkaConfig.Consumer.Retry.Backoff)
		}
		if cfg.KafkaConfig.Consumer.Fetch != nil {
			config.Consumer.Fetch.Min = cfg.KafkaConfig.Consumer.Fetch.Min
			config.Consumer.Fetch.Default = cfg.KafkaConfig.Consumer.Fetch.Default
			config.Consumer.Fetch.Max = cfg.KafkaConfig.Consumer.Fetch.Max
		}
		if cfg.KafkaConfig.Consumer.MaxWaitTime != nil {
			config.Consumer.MaxWaitTime = time.Duration(*cfg.KafkaConfig.Consumer.MaxWaitTime)
		}
		if cfg.KafkaConfig.Consumer.MaxProcessingTime != nil {
			config.Consumer.MaxProcessingTime = time.Duration(*cfg.KafkaConfig.Consumer.MaxProcessingTime)
		}

		if cfg.KafkaConfig.Consumer.Offsets != nil {
			config.Consumer.Offsets.Initial = cfg.KafkaConfig.Consumer.Offsets.Initial
			config.Consumer.Offsets.CommitInterval = time.Duration(cfg.KafkaConfig.Consumer.Offsets.CommitInterval)
			config.Consumer.Offsets.Retention = time.Duration(cfg.KafkaConfig.Consumer.Offsets.Retention)
		}
	}
	if cfg.KafkaConfig.Group != nil {
		if cfg.KafkaConfig.Group.PartitionStrategy != nil {
			config.Group.PartitionStrategy = cluster.Strategy(*cfg.KafkaConfig.Group.PartitionStrategy)
		}
		if cfg.KafkaConfig.Group.Mode != nil {
			config.Group.Mode = cluster.ConsumerMode(*cfg.KafkaConfig.Group.Mode)
		}

		if cfg.KafkaConfig.Group.Offsets != nil {
			if cfg.KafkaConfig.Group.Offsets.Retry != nil {
				config.Group.Offsets.Retry.Max = cfg.KafkaConfig.Group.Offsets.Retry.Max
			}
			if cfg.KafkaConfig.Group.Offsets.Synchronization != nil {
				config.Group.Offsets.Synchronization.DwellTime = time.Duration(cfg.KafkaConfig.Group.Offsets.Synchronization.DwellTime)
			}
		}
		if cfg.KafkaConfig.Group.Session != nil {
			config.Group.Session.Timeout = time.Duration(cfg.KafkaConfig.Group.Session.Timeout)
		}
		if cfg.KafkaConfig.Group.Heartbeat != nil {
			config.Group.Heartbeat.Interval = time.Duration(cfg.KafkaConfig.Group.Heartbeat.Interval)
		}
		// if cfg.KafkaConfig.Group.Return != nil {
		// 	config.Group.Return.Notifications = cfg.KafkaConfig.Group.Return.Notifications
		// }
		if cfg.KafkaConfig.Group.Member != nil {
			config.Group.Member.UserData = cfg.KafkaConfig.Group.Member.UserData
		}
	}

	clusterConsumer, err := cluster.NewConsumer(cfg.KafkaConfig.Brokers, cfg.Group, []string{string(cfg.Topic)}, config)
	if err != nil {
		slog.Info("[DCL] [NewConsumer] Error: %+v\n", err)
		return nil, err
	}

	return &Consumer{
		c: newConsumer(cfg, clusterConsumer),
	}, nil
}

// Consumer ...
type Consumer struct {
	c *consumer
}

// Start ...
func (c *Consumer) Start() error {
	return c.c.consume(false)
}

// Close ...
func (c *Consumer) Close() {
	c.c.close()
}

func newConsumer(cfg ConsumerConfig, clusterConsumer *cluster.Consumer) *consumer {
	return &consumer{
		cfg:                 cfg,
		consumer:            clusterConsumer,
		errChan:             make(chan error),
		killChan:            make(chan struct{}, 100),
		partitionBufferSize: 100,
	}
}

type consumer struct {
	cfg                 ConsumerConfig
	consumer            *cluster.Consumer
	errChan             chan error
	endSessionWG        sync.WaitGroup
	killChan            chan struct{}
	closeFlag           int32
	partitionBufferSize int
}

type doneChanItem struct {
	done chan error
	msg  *sarama.ConsumerMessage
}

func (c *consumer) isClosed() bool {
	return atomic.LoadInt32(&c.closeFlag) == 1
}

func (c *consumer) close() {
	if c.isClosed() {
		slog.Info("IS CLOSED\n")
		return
	}
	atomic.CompareAndSwapInt32(&c.closeFlag, 0, 1)
	c.endSessionWG.Wait()
	c.killChan <- struct{}{}
}

func (c *consumer) consume(catchInterrupt bool) error {
	defer func() { slog.Info("DCL CLOSED\n") }()
	defer c.consumer.Close()
	defer func() { slog.Info("DCL CLOSING\n") }()

	go c.logNotifications()

	for {
		select {
		case partition, ok := <-c.consumer.Partitions():
			if !ok {
				return nil
			}
			go c.consumePartition(partition)
		case e := <-c.errChan:
			c.close()
			return e
		case <-c.killChan:
			return nil
		}
	}
}

func (c *consumer) consumePartition(pc cluster.PartitionConsumer) {
	slog.Info("[DCL] Start subscribe partition: %v topic: %v\n", pc.Partition(), pc.Topic())
	defer func() {
		slog.Info(fmt.Sprintf("[DCL] Close subscribe partition: %v, topic: %v\n", pc.Partition(), pc.Topic()))
	}()

	doneChans := make(chan doneChanItem, c.partitionBufferSize)
	defer close(doneChans)

	go func() {
		for done := range doneChans {
			err := <-done.done
			c.consumer.MarkOffset(done.msg, "")
			c.endSessionWG.Done()
			if err != nil {
				c.notifyError(err)
			}
		}
	}()

	for msg := range pc.Messages() {
		if c.isClosed() {
			slog.Info("consumer is closed")
			return
		}

		c.endSessionWG.Add(1)

		c.processMessage(msg, doneChans)
	}
}

func (c *consumer) processMessage(msg *sarama.ConsumerMessage, doneChans chan doneChanItem) {

	var event event.Event
	err := proto.Unmarshal(msg.Value, &event)
	if err != nil {
		err2 := errors.New(fmt.Sprintf("err : %v , topic : %s , msg : %s", err, msg.Topic, msg.Value))
		slog.Err("%+v", err2)
		c.notifyError(err2)
		return
	}

	event.Context = tracing.AppendServiceNameToServiceContext(event.Context, c.cfg.ServiceName)

	done := make(chan error)
	doneChans <- doneChanItem{
		done: done,
		msg:  msg,
	}

	ctx := tracing.SetServiceContext(context.Background(), event.Context)

	c.cfg.Processor.Process(ctx, &event, &eventmeta.EventMetaData{Timestamp: msg.Timestamp,
		Key:       msg.Key,
		Topic:     msg.Topic,
		Partition: msg.Partition,
		Offset:    msg.Offset,
	}, done)
}

func (c *consumer) logNotifications() {
	go func() {
		for err := range c.consumer.Errors() {
			c.notifyError(err)
		}
	}()

	for ntf := range c.consumer.Notifications() {
		slog.Info(fmt.Sprintf("[DCL] Rebalanced: %+v\n", ntf))
	}
}

func (c *consumer) notifyError(err error) {
	if err == nil {
		return
	}
	slog.Err("[DCL] Error: %s\n", err.Error())
	if c.cfg.Errors != nil {
		select {
		case c.cfg.Errors <- err:
		default:
		}
	}
}
