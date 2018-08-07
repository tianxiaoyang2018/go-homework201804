package eventlog

import (
	"github.com/p1cn/tantan-backend-common/config"
)

func Init(cfg config.EventLog, mq config.MessageQueue) error {
	kafkaConfig := mq.Kafka[cfg.MqName]

	return InitEventLogWithKafka(
		kafkaConfig,
	)
}

func InitEventLogWithKafka(kafkaConfig config.KafkaConfig) (err error) {
	EventLog = &EventLogClient{}

	EventLog.client, err = newKafkaClient(kafkaConfig)
	if err != nil {
		return err
	}
	return EventLog.start()
}
