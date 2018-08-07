package producer

import (
	"sync"
	"time"

	"github.com/Shopify/sarama"

	"github.com/p1cn/tantan-backend-common/config"
	slog "github.com/p1cn/tantan-backend-common/log"
)

type KafkaClient struct {
	sync.RWMutex
	client          sarama.Client
	topicPartitions map[string][]int32
	quitChan        chan struct{}
}

func NewKafkaClient(conf *config.KafkaConfig) (*KafkaClient, error) {
	client, err := sarama.NewClient(conf.Brokers, parseKafkaConfig(conf))
	if err != nil {
		return nil, err
	}

	self := &KafkaClient{
		quitChan:        make(chan struct{}, 0),
		topicPartitions: make(map[string][]int32),
		client:          client,
	}
	self.updateTopics()

	go func() {
		for {
			select {
			case <-self.quitChan:
				return
			case <-time.After(5 * time.Minute):
				self.updateTopics()
			}
		}
	}()

	return self, nil
}

func (c *KafkaClient) Close() error {
	if c.client.Closed() {
		return nil
	}
	return c.client.Close()
}

func (c *KafkaClient) Closed() bool {
	return c.client.Closed()
}

func (c *KafkaClient) TopicPartitions(topic string) []int32 {
	c.RLock()
	defer c.RUnlock()
	return c.topicPartitions[topic]
}

func (c *KafkaClient) updateTopics() {
	topicPartition := make(map[string][]int32)
	topics, err := c.client.Topics()
	if err != nil {
		slog.Err("%v", err)
		return
	}
	for _, t := range topics {
		p, err := c.client.Partitions(t)
		if err != nil {
			slog.Err("%v", err)
			return
		}
		if len(p) == 0 {
			slog.Err("ERROR Topic(%v) has not partitions", t)
		}
		topicPartition[t] = p
	}
	if len(topicPartition) == 0 {
		slog.Err("ERROR empty topic and parition information")
		return
	}
	c.Lock()
	c.topicPartitions = topicPartition
	c.Unlock()
}
