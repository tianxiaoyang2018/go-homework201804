package dcl

import (
	"context"
	"errors"
	"fmt"

	"github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/version"
	"github.com/p1cn/tantan-domain-schema/golang/event"

	dcl_consumer "github.com/p1cn/tantan-backend-common/dcl/consumer"
	"github.com/p1cn/tantan-backend-common/dcl/eventmeta"
)

// Processor : 处理器，消费者回调函数
type Processor func(context.Context, *event.Event, *eventmeta.EventMetaData) error

// 错误处理回调函数
type ErrorHandler func(error)

// DCL消费者接口
type Consumer interface {
	AddConsumer(topic string, group string, p Processor, eh ErrorHandler)
	Start() error
	Stop() error
}

// 初始化dcl消费者
// 后面将废弃，请使用NewKafkaConsumer
// name :消费者名字, 用来区分监控等
func NewConsumer(name string, dcl config.Dcl, mq config.MessageQueue) (Consumer, error) {
	if dcl.MqType != "kafka" {
		err := errors.New("new DCL consumer failed : only support kafka")
		log.Err("%+v", err)
		return nil, err
	}
	kafka, exist := mq.Kafka[dcl.MqName]
	if !exist {
		err := fmt.Errorf("new DCL consumer failed : mqName(%s)", dcl.MqName)
		log.Err("%+v", err)
		return nil, err
	}

	epm, err := newEventProcessorManager(name)
	if err != nil {
		err := fmt.Errorf("new DCL consumer failed : mqName(%s)", dcl.MqName)
		log.Err("%+v", err)
		return nil, err
	}

	log.Info("new DCL consumer successfully.")

	return &consumer{
		epm: epm,
		cfg: &consumerConfig{
			ServiceName: name,
			Kafka:       kafka,
		}}, nil
}

// new consumer of kafka
type KafkaConsumerConfig struct {
	// 服务名字
	ServiceName string
	// 服务版本
	ServiceVersion string
	// 消费者名字 : 用来区分监控等
	Name string
	// kafka 配置
	Kafka config.KafkaConfig
}

// 初始化kafka DCL 消费者
func NewKafkaConsumer(cfg KafkaConsumerConfig) (Consumer, error) {

	if len(cfg.ServiceName) == 0 {
		cfg.ServiceName = version.ServiceName()
	}
	name := cfg.Name
	if len(name) == 0 {
		name = cfg.ServiceName
	}

	epm, err := newEventProcessorManager(name)
	if err != nil {
		log.Err("new kakfa DCL consumer failed : err(%+v)", err)
		return nil, err
	}

	log.Info("new Kafka DCL consumer successfully.")

	return &consumer{
		epm: epm,
		cfg: &consumerConfig{
			ServiceName: cfg.ServiceName,
			Kafka:       cfg.Kafka,
		}}, nil
}

type consumerConfig struct {
	ServiceName string
	Kafka       config.KafkaConfig
}

type consumer struct {
	cfg             *consumerConfig
	epm             *eventProcessorManager
	consumerConfigs []dcl_consumer.ConsumerConfig
	mc              dcl_consumer.MultiConsumer
}

func (c *consumer) AddConsumer(topic string, group string, p Processor, eh ErrorHandler) {
	errChan := make(chan error, 100)
	c.consumerConfigs = append(c.consumerConfigs, dcl_consumer.ConsumerConfig{
		ServiceName: c.cfg.ServiceName,
		Topic:       topic,
		Group:       group,
		KafkaConfig: c.cfg.Kafka,
		Processor:   c.epm.CreateEventProcessor(p),
		Errors:      errChan,
	})
	go func() {
		for {
			eh(<-errChan)
		}
	}()
}

func (c *consumer) Start() (err error) {
	c.mc, err = dcl_consumer.NewMultiConsumer(c.consumerConfigs)
	if err != nil {
		log.Crit("consumer start failed : err(%+v)", err)
		return err
	}

	log.Info("consumer starting...")
	return c.mc.Start()
}

func (c *consumer) Stop() error {
	err := c.mc.Close()
	if err != nil {
		log.Err("consumer stop error : err(%v) .", err)
	} else {
		log.Info("consumer stop successfully .")
	}
	return err
}
