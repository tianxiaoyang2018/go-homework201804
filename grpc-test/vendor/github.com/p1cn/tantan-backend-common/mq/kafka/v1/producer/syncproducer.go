package producer

import (
	"sync"
	"sync/atomic"

	"github.com/Shopify/sarama"

	"errors"

	"github.com/p1cn/tantan-backend-common/config"
)

var (
	ErrProducerClosed = errors.New("producer is closed")
)

var defaultSyncProducer *KafkaSyncProducer

func DefaultSync() *KafkaSyncProducer {
	return defaultSyncProducer
}

func InitDefaultSyncProducer(conf *config.KafkaConfig) (err error) {
	defaultSyncProducer, err = NewKafkaSyncProducer(conf)
	return
}

type KafkaSyncProducer struct {
	BaseKafkaProducer
	producer       sarama.SyncProducer
	waitGroup      sync.WaitGroup
	producerStatus int32
}

// NewKafkaSyncProducer
// set configuration.Flush to a low value will speed up transmission
func NewKafkaSyncProducer(conf *config.KafkaConfig) (*KafkaSyncProducer, error) {
	async, err := sarama.NewSyncProducer(conf.Brokers, parseKafkaConfig(conf))
	if err != nil {
		return nil, err
	}
	pp := &KafkaSyncProducer{
		producer: async,
	}

	err = pp.Init(conf)
	if err != nil {
		async.Close()
		return nil, err
	}
	return pp, nil
}

// close producer until
func (k *KafkaSyncProducer) Close() error {
	k.waitGroup.Wait()

	if atomic.CompareAndSwapInt32(&k.producerStatus, 0, 1) {
		err := k.producer.Close()
		if err != nil {
			atomic.SwapInt32(&k.producerStatus, 0)
			return err
		}
	}

	return k.BaseKafkaProducer.Close()
}

func (k *KafkaSyncProducer) IsClosed() bool {
	return atomic.LoadInt32(&k.producerStatus) == 1
}

func (k *KafkaSyncProducer) SendMessage(msg *KafkaProducerMessage) (partition int32, offset int64, err error) {
	if atomic.LoadInt32(&k.producerStatus) == 1 {
		return -1, -1, ErrProducerClosed
	}

	k.waitGroup.Add(1)
	defer k.waitGroup.Done()

	p := k.TopicPartitions(msg.Topic)
	if len(p) == 0 {
		return -1, -1, ErrInvalidTopic
	}

	pp, offset, err := k.producer.SendMessage(toSaramaProducerMessage(msg, p))
	if err != nil {
		k.CounterInc(msg.Topic, COUNTER_FAILURE)
	} else {
		k.CounterInc(msg.Topic, COUNTER_SUCCESS)
	}
	return pp, offset, err
}

func (k *KafkaSyncProducer) SendMessages(msgs []*KafkaProducerMessage) error {
	if atomic.LoadInt32(&k.producerStatus) == 1 {
		return ErrProducerClosed
	}

	k.waitGroup.Add(1)
	defer k.waitGroup.Done()

	var ms []*sarama.ProducerMessage
	for _, m := range msgs {
		p := k.TopicPartitions(m.Topic)
		if len(p) == 0 {
			return ErrInvalidTopic
		}

		ms = append(ms, toSaramaProducerMessage(m, p))
	}
	err := k.producer.SendMessages(ms)
	if err != nil {
		k.CounterAdd(len(msgs), COUNTER_FAILURE)
	} else {
		k.CounterAdd(len(msgs), COUNTER_SUCCESS)
	}
	return err
}
