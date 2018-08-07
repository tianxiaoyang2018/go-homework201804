package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

var topic = flag.String("topic", "dcl.test", "topic name")
var kafka = flag.String("kafka", "192.168.4.5:9092", "kafka brokers")
var consumerGroup = flag.String("group", "test-consumer", "specify consumer group name")
var old = flag.Bool("oldoffset", false, "old offset : true or false")

func signalHandler() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGPIPE, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)
	for {
		select {
		case sig := <-sigCh:
			switch sig {
			case syscall.SIGPIPE:
			default:

				os.Exit(0)
			}
		}
	}
}

func main() {
	flag.Parse()

	go signalHandler()

	group := "test-service-" + *topic
	if len(*consumerGroup) > 0 {
		group = *consumerGroup
	}

	brokers := strings.Split(*kafka, ",")
	config := cluster.NewConfig()
	config.Group.Mode = cluster.ConsumerModePartitions
	config.Version = sarama.V0_10_0_0

	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.MaxProcessingTime = 1 * time.Second
	config.Consumer.MaxWaitTime = 1 * time.Second
	config.Group.Offsets.Synchronization.DwellTime = 1 * time.Second
	if *old {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	fmt.Printf("topic(%s) , offset(%v) , group(%s)\n", *topic, config.Consumer.Offsets.Initial, group)

	topics := strings.Split(*topic, ",")
	clusterConsumer, err := cluster.NewConsumer(brokers, group, topics, config)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for msg := range clusterConsumer.Errors() {
			fmt.Println(msg)
		}
	}()

	go func() {
		for msg := range clusterConsumer.Notifications() {
			fmt.Println(msg)
		}
	}()

	for {
		fmt.Println("starting")
		select {
		case partition, ok := <-clusterConsumer.Partitions():
			if !ok {
				log.Fatal("partition error")
			}
			fmt.Println("starting consume 1")
			go consumePartition(partition, clusterConsumer)

		}
	}

}

func consumePartition(pc cluster.PartitionConsumer, consumer *cluster.Consumer) {
	fmt.Println("starting consume 2")
	for msg := range pc.Messages() {
		// var event dcl.Event
		// err := proto.Unmarshal(msg.Value, &event)
		// if err != nil {
		// 	fmt.Printf("error : %+v\n", err)
		// }
		// data, _ := json.Marshal(event)
		// fmt.Printf("dcl : %s\n", data)

		consumer.MarkOffset(msg, "")
	}
}
