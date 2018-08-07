package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/dcl"
	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/metrics"
	"github.com/p1cn/tantan-domain-schema/golang/event"
)

//var topic = flag.String("topic", "dcl.test", "topic name")
var kafka = flag.String("kafka", "192.168.4.5:9092", "kafka brokers")
var count = flag.Int("count", 1, "count")
var metric = flag.Bool("metric", false, "print metrics")
var data = flag.String("data", "{}", "data want be sent to kafka")

func main() {

	flag.Parse()

	slog.Init(slog.Config{
		Output: []string{"stderr", "syslog"},
		Level:  "debug",
		Flags:  []string{"file", "level", "date"},
	})

	kafkacluster := strings.Split(*kafka, ",")

	producer, err := dcl.NewKafkaProducer(dcl.KafkaProducerConfig{
		Name: "test_dcl_producer",
		KafkaConfig: config.KafkaConfig{
			Brokers: kafkacluster,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	ee := event.Event{}
	err = json.Unmarshal([]byte(*data), &ee)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < *count; i++ {
		if i%100 == 0 {
			fmt.Println("produce id : ", i)
		}

		err = producer.Commit(context.Background(), &ee)

		if err != nil {
			fmt.Println(err)
		}
	}

	producer.Close()

	if *metric {
		fmt.Println(metrics.GetPromethuesAsFmtText())
	}

}
