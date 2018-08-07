package producer

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/p1cn/tantan-backend-common/config"
)

const (
	COUNTER_SUCCESS = "OK"
	COUNTER_FAILURE = "failure"
)

var (
	ErrTimeout      = errors.New("timeout")
	ErrInvalidTopic = errors.New("invalid topic")
)

type KafkaProducerError struct {
	Msg *KafkaProducerMessage
	Err error
}

type BaseKafkaProducer struct {
	prometheusCounter *prometheus.CounterVec
	client            *KafkaClient
}

func (k *BaseKafkaProducer) Init(conf *config.KafkaConfig) error {
	client, err := NewKafkaClient(conf)
	if err != nil {
		return err
	}
	k.client = client
	return nil
}

func (k *BaseKafkaProducer) Close() error {
	if k.client != nil && !k.client.Closed() {
		return k.client.Close()
	}
	return nil
}

func (k *BaseKafkaProducer) TopicPartitions(topic string) []int32 {
	if k.client != nil {
		return k.client.TopicPartitions(topic)
	}
	return nil
}

// RegisterPrometheus
// name is subsystem
func (k *BaseKafkaProducer) RegisterPrometheus(name string) error {

	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "tantan",
		Subsystem: name,
		Name:      "kafka_producer_c",
		Help:      "kafka producer counter",
	}, []string{"topic", "ret"},
	)

	err := prometheus.Register(counter)
	if err != nil {
		return err
	}
	k.prometheusCounter = counter
	return nil
}

func (k *BaseKafkaProducer) CounterInc(label ...string) {
	if k.prometheusCounter != nil {
		k.prometheusCounter.WithLabelValues(label...).Inc()
	}
}

func (k *BaseKafkaProducer) CounterAdd(count int, label ...string) {
	if k.prometheusCounter != nil {
		k.prometheusCounter.WithLabelValues(label...).Add(float64(count))
	}
}
