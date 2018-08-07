package producer

import (
	"sync/atomic"
	"time"

	"github.com/Shopify/sarama"

	"github.com/p1cn/tantan-backend-common/config"
)

type KafkaASyncProducer struct {
	BaseKafkaProducer
	producer       sarama.AsyncProducer
	successChan    chan *KafkaProducerMessage
	errorChan      chan *KafkaProducerError
	producerStatus int32
}

func NewKafkaASyncProducer(conf *config.KafkaConfig) (*KafkaASyncProducer, error) {
	async, err := sarama.NewAsyncProducer(conf.Brokers, parseKafkaConfig(conf))
	if err != nil {
		return nil, err
	}

	pp := &KafkaASyncProducer{
		producer:    async,
		successChan: make(chan *KafkaProducerMessage, 1000),
		errorChan:   make(chan *KafkaProducerError, 1000),
	}

	err = pp.Init(conf)
	if err != nil {
		async.Close()
		return nil, err
	}

	if conf.Producer != nil && conf.Producer.Return != nil {
		if conf.Producer.Return.Successes {
			go func() {
				for m := range async.Successes() {
					if m != nil {
						pp.CounterInc(m.Topic, COUNTER_SUCCESS)
						pp.successChan <- fromSaramaProducerMessage(m)
					}
				}
			}()
		}

		if conf.Producer.Return.Errors {
			go func() {
				for m := range async.Errors() {
					if m.Msg != nil {
						pp.CounterInc(m.Msg.Topic, "error")
					} else {
						pp.CounterInc("null", "error_msg_nil")
					}

					pp.errorChan <- fromSaramaProducerError(m)
				}
			}()
		}
	}

	return pp, nil
}

func (k *KafkaASyncProducer) Close() error {
	if atomic.CompareAndSwapInt32(&k.producerStatus, 0, 1) {
		err := k.producer.Close()
		if err != nil {
			atomic.SwapInt32(&k.producerStatus, 0)
			return err
		}
	}

	return k.BaseKafkaProducer.Close()
}

func (k *KafkaASyncProducer) IsClosed() bool {
	return atomic.LoadInt32(&k.producerStatus) == 1
}

func (k *KafkaASyncProducer) PushMessage(m *KafkaProducerMessage) error {
	if atomic.LoadInt32(&k.producerStatus) == 1 {
		return ErrProducerClosed
	}

	p := k.TopicPartitions(m.Topic)
	if len(p) == 0 {
		return ErrInvalidTopic
	}
	msg := toSaramaProducerMessage(m, p)

	k.producer.Input() <- msg
	return nil
}

func (k *KafkaASyncProducer) PushMessageTimeout(m *KafkaProducerMessage, timeout time.Duration) error {
	if atomic.LoadInt32(&k.producerStatus) == 1 {
		return ErrProducerClosed
	}

	p := k.TopicPartitions(m.Topic)
	if len(p) == 0 {
		return ErrInvalidTopic
	}
	msg := toSaramaProducerMessage(m, p)
	select {
	case k.producer.Input() <- msg:
	case <-time.After(timeout):
		return ErrTimeout
	}
	return nil
}

// success don't guarantee message has been sent to kafka if RequiredAcks of configuration is false
func (k *KafkaASyncProducer) Successes() <-chan *KafkaProducerMessage {
	return k.successChan
}

func (k *KafkaASyncProducer) Errors() <-chan *KafkaProducerError {
	return k.errorChan
}

func parseKafkaConfig(conf *config.KafkaConfig) *sarama.Config {
	pconfig := sarama.NewConfig()
	pconfig.Producer.Partitioner = sarama.NewManualPartitioner
	if conf.Producer != nil {
		if conf.Producer.RequiredAcks != nil {
			pconfig.Producer.RequiredAcks = sarama.RequiredAcks(*conf.Producer.RequiredAcks)
		}
		if conf.Producer.Compression != nil {
			pconfig.Producer.Compression = sarama.CompressionCodec(*conf.Producer.Compression)
		}
		if conf.Producer.Flush != nil {
			pconfig.Producer.Flush.MaxMessages = conf.Producer.Flush.MaxMessages
			pconfig.Producer.Flush.Messages = conf.Producer.Flush.Messages
			pconfig.Producer.Flush.Frequency = conf.Producer.Flush.Frequency.Duration()
			pconfig.Producer.Flush.Bytes = conf.Producer.Flush.Bytes
		}
		if conf.Producer.MaxMessageBytes != nil {
			pconfig.Producer.MaxMessageBytes = *conf.Producer.MaxMessageBytes
		}
		if conf.Producer.Timeout != nil {
			pconfig.Producer.Timeout = conf.Producer.Timeout.Duration()
		}

		if conf.Producer.Retry != nil {
			pconfig.Producer.Retry.Max = conf.Producer.Retry.Max
			pconfig.Producer.Retry.Backoff = conf.Producer.Retry.Backoff.Duration()
		}
		if conf.Producer.Return != nil {
			pconfig.Producer.Return.Successes = conf.Producer.Return.Successes
			pconfig.Producer.Return.Errors = conf.Producer.Return.Errors
		}
	}

	return pconfig
}
