package consumer

import "time"
import "github.com/Shopify/sarama"

type KafkaConsumerMessage struct {
	Key, Value []byte
	Topic      string
	Partition  int32
	Offset     int64
	Timestamp  time.Time
}

func toSaramaConsumerMessage(m *KafkaConsumerMessage) *sarama.ConsumerMessage {
	return &sarama.ConsumerMessage{
		Key:       m.Key,
		Value:     m.Value,
		Topic:     m.Topic,
		Partition: m.Partition,
		Offset:    m.Offset,
		Timestamp: m.Timestamp,
	}
}

func fromSaramaConsumerMessage(m *sarama.ConsumerMessage) *KafkaConsumerMessage {
	return &KafkaConsumerMessage{
		Key:       m.Key,
		Value:     m.Value,
		Topic:     m.Topic,
		Partition: m.Partition,
		Offset:    m.Offset,
		Timestamp: m.Timestamp,
	}
}
