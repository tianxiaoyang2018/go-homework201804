package producer

import (
	"time"

	slog "github.com/p1cn/tantan-backend-common/log"

	"github.com/Shopify/sarama"
)

type KafkaProducerMessage struct {
	Topic     string
	Key       []byte
	Value     []byte
	Metadata  interface{}
	Offset    int64
	Partition int64
	Timestamp time.Time
}

func toSaramaProducerMessage(m *KafkaProducerMessage, partitions []int32) *sarama.ProducerMessage {
	p := m.Partition
	if p < 0 {
		p *= -1
	}
	return &sarama.ProducerMessage{
		Topic:     m.Topic,
		Key:       sarama.ByteEncoder(m.Key),
		Value:     sarama.ByteEncoder(m.Value),
		Partition: int32(p % int64(len(partitions))),
		Metadata:  m.Metadata,
	}
}

func fromSaramaProducerMessage(m *sarama.ProducerMessage) *KafkaProducerMessage {
	if m == nil {
		return nil
	}
	k, err := m.Key.Encode()
	if err != nil { // not good
		slog.Err("%v", err)
	}
	v, err := m.Value.Encode()
	if err != nil {
		slog.Err("%v", err)
	}
	return &KafkaProducerMessage{
		Topic:     m.Topic,
		Key:       k,
		Value:     v,
		Metadata:  m.Metadata,
		Offset:    m.Offset,
		Partition: int64(m.Partition),
		Timestamp: m.Timestamp,
	}
}

func fromSaramaProducerError(m *sarama.ProducerError) *KafkaProducerError {
	return &KafkaProducerError{
		Msg: fromSaramaProducerMessage(m.Msg),
		Err: m.Err,
	}
}
