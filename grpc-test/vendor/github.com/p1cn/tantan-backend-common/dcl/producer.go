package dcl

import (
	"context"
	"errors"

	"github.com/Shopify/sarama"

	"github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/dcl/producer"
	log "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-domain-schema/golang/event"
)

type Producer interface {
	Commit(context.Context, *event.Event) error
	Close() error
}

// 初始化DCL生产者
// 后面将废弃，请使用NewKafkaProducer
func NewProducer(name string, pro config.Dcl, mq config.MessageQueue) (Producer, error) {

	if pro.MqType != "kafka" {
		return nil, errors.New("only support kafka")
	}
	kafka, exist := mq.Kafka[pro.MqName]
	if !exist {
		return nil, errors.New("kafka cluster name does not exist")
	}

	prod, err := producer.NewProducer(&producer.ProducerConfig{
		Name:        name,
		KafkaConfig: kafka,
	})

	if err != nil {
		log.Err("new DCL producer error : name(%s) , err(%+v)", name, err)
		return nil, err
	}

	log.Info("new DCL producer successfully : name(%s)", name)

	return prod, nil
}

// New kafka dcl producer
type KafkaProducerConfig struct {
	Name                string
	KafkaConfig         config.KafkaConfig
	ProducerPartitioner sarama.PartitionerConstructor
}

// 初始化kakfa DCL 生产者
func NewKafkaProducer(cfg KafkaProducerConfig) (Producer, error) {

	pro, err := producer.NewProducer(&producer.ProducerConfig{
		Name:        cfg.Name,
		KafkaConfig: cfg.KafkaConfig,
		Partitioner: cfg.ProducerPartitioner,
	})
	if err != nil {
		log.Err("new kafka DCL producer error : err(%+v)", err)
		return nil, err
	}

	log.Info("new kafka DCL producer successfully.")
	return pro, nil
}
