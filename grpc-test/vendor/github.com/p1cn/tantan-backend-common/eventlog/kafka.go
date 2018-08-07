package eventlog

import (
	slog "github.com/p1cn/tantan-backend-common/log"

	"github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/mq/kafka/v1/producer"
)

type kafkaClient struct {
	producer *producer.KafkaASyncProducer
}

func newKafkaClient(defaultConfig config.KafkaConfig) (*kafkaClient, error) {
	pp, err := producer.NewKafkaASyncProducer(&defaultConfig)
	if err != nil {
		return nil, err
	}
	pp.RegisterPrometheus("eventlog")
	return &kafkaClient{
		producer: pp,
	}, nil
}

func (this *kafkaClient) Start() error {
	go func() {
		for range this.producer.Successes() {
		}
	}()
	go func() {
		for ee := range this.producer.Errors() {
			slog.Err("%s", ee)
		}
	}()
	return nil
}

func (this *kafkaClient) Close() error {
	return this.producer.Close()
}

func (this *kafkaClient) Send(message *RpcMessage) error {
	mm := &producer.KafkaProducerMessage{
		Key:       message.Key,
		Topic:     message.Topic,
		Value:     message.Data,
		Partition: message.ID,
	}
	err := this.producer.PushMessage(mm)
	if err != nil {
		return err
	}
	return nil
}
