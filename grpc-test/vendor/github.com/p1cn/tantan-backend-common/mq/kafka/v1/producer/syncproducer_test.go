package producer

import (
	"flag"
	"fmt"
	"testing"

	"github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/metrics"
)

//var brokers = flag.String("brokers", "192.168.4.11:9092", "brokers' addresses")

func TestSyncProducer(t *testing.T) {
	flag.Parse()
	conf := new(config.KafkaConfig)
	conf.Brokers = []string{*brokers}

	conf.Producer = &config.KafkaProducerConfig{
		Return: &config.KafkaProducerReturnConfig{
			Errors:    true,
			Successes: true,
		},
	}

	kafka, err := NewKafkaSyncProducer(conf)
	if err != nil {
		t.Fatal(err)
	}
	kafka.RegisterPrometheus("test")
	partition, _, err := kafka.SendMessage(&KafkaProducerMessage{
		Topic:     "test",
		Value:     []byte("test"),
		Partition: 0,
	})

	if err != nil {
		t.Fatal(err)
	}

	if partition != 0 {
		t.Fatal("parition is not 0")
	}
	if *metric {
		fmt.Println(metrics.GetPromethuesAsFmtText())
	}
}
