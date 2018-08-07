package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/dcl"
	"github.com/p1cn/tantan-backend-common/dcl/eventmeta"
	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/metrics"
	"github.com/p1cn/tantan-domain-schema/golang/event"
)

func Process(ctx context.Context, event *event.Event, meta *eventmeta.EventMetaData) error {

	data, err := json.Marshal(event)
	if err != nil {
		slog.Info("PROCESS ERROR:", err)
		return err
	}
	data2, err := json.Marshal(meta)
	if err != nil {
		slog.Info("PROCESS ERROR:", err)
		return err
	}

	if !*dryRun {
		slog.Info("test : %s, %s", string(data), string(data2))
	}

	return nil
}

var topic = flag.String("topic", "dcl.test", "topic name")
var kafka = flag.String("kafka", "192.168.4.5:9092", "kafka brokers")
var consumerGroup = flag.String("group", "test-consumer", "specify consumer group name")
var dryRun = flag.Bool("dry", false, "dry run:dont' print message")
var metric = flag.Bool("metric", false, "print metrics")

func signalHandler() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGPIPE, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)
	for {
		select {
		case sig := <-sigCh:
			switch sig {
			case syscall.SIGPIPE:
			default:
				if *metric {
					fmt.Println(metrics.GetPromethuesAsFmtText())
				}
				os.Exit(0)
			}
		}
	}
}

func main() {
	flag.Parse()

	slog.Init(slog.Config{
		Output: []string{"stderr", "syslog"},
		Level:  "info",
		Flags:  []string{"file", "level", "date"},
	})

	go signalHandler()

	group := "test-service-" + *topic
	if len(*consumerGroup) > 0 {
		group = *consumerGroup
	}

	kconfig := dcl.KafkaConsumerConfig{

		ServiceName: "test-service",
		Name:        "test-dcl-consumer",
		Kafka: config.KafkaConfig{

			Brokers: strings.Split(*kafka, ","),
		},
	}

	c, err := dcl.NewKafkaConsumer(kconfig)
	if err != nil {
		log.Fatal(err)
	}

	c.AddConsumer(*topic, group, Process, errHanlder)

	c.Start()
}

func errHanlder(err error) {
	fmt.Println(err)
}
