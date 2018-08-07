package consumer

import (
	"github.com/p1cn/tantan-backend-common/config"

	"github.com/wvanbergen/kafka/consumergroup"
)

const (
	DEFAULT_BUFFER_SIZE = 10
)

// ConsumerGroupConfig
// ConsumerGroupConfig.Config.Offsets.Initial is initial offset(sarama.OffsetOldest (default) or sarama.OffsetNewest)
type KafkaConsumerGroupConfig struct {
	GroupName  string
	Topics     []string
	Zookeepers []string
	BufferSize int
	Config     *consumergroup.Config
}

//
func NewKafkaConsumerGroupConfig(groupName string, topics []string, kafkaConfig *config.KafkaConfig) *KafkaConsumerGroupConfig {
	cgConfig := consumergroup.NewConfig()
	return &KafkaConsumerGroupConfig{
		GroupName:  groupName,
		Topics:     topics,
		Zookeepers: kafkaConfig.Zookeepers,
		Config:     cgConfig,
		BufferSize: DEFAULT_BUFFER_SIZE,
	}
}

// ConsumerGroup
type KafkaConsumerGroup struct {
	consumer    *consumergroup.ConsumerGroup
	messageChan chan *KafkaConsumerMessage
}

func NewKafkaConsumerGroup(config *KafkaConsumerGroupConfig) (*KafkaConsumerGroup, error) {
	consumer, err := consumergroup.JoinConsumerGroup(
		config.GroupName,
		config.Topics,
		config.Zookeepers,
		config.Config,
	)
	if err != nil {
		return nil, err
	}

	if config.BufferSize < 0 {
		config.BufferSize = 0
	}
	messageChan := make(chan *KafkaConsumerMessage, config.BufferSize)
	go func() {
		defer close(messageChan)
		for m := range consumer.Messages() {
			messageChan <- fromSaramaConsumerMessage(m)
		}
	}()

	return &KafkaConsumerGroup{
		consumer:    consumer,
		messageChan: messageChan,
	}, nil
}

// you need to read Messages() channel until no data in the channel.
// for message := range Messages() will quit loop until no data in channel
func (c *KafkaConsumerGroup) Close() error {
	if c.consumer.Closed() {
		return nil
	}
	return c.consumer.Close()
}

func (c *KafkaConsumerGroup) Messages() <-chan *KafkaConsumerMessage {
	return c.messageChan
}

func (c *KafkaConsumerGroup) Errors() <-chan error {
	return c.consumer.Errors()
}

// CommitUpto : if you consume message successfully, commit manually
// Offset.CommitInterval of configuration will commit offset automatically,
// if you quit program unexpectedly,  there are some data in channel and they have been commited to zookeepers,
// means they are lost
func (c *KafkaConsumerGroup) CommitUpto(msg *KafkaConsumerMessage) {
	c.consumer.CommitUpto(toSaramaConsumerMessage(msg))
}
