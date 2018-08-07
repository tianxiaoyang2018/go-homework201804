package producer

import (
	"flag"
	"fmt"
	"testing"

	"sync"

	"github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/metrics"
)

var brokers = flag.String("brokers", "192.168.4.5:9092", "brokers' addresses")
var metric = flag.Bool("metric", true, "print metrics")

func TestASyncProducer(t *testing.T) {
	flag.Parse()
	conf := new(config.KafkaConfig)
	conf.Brokers = []string{*brokers}

	conf.Producer = &config.KafkaProducerConfig{
		Return: &config.KafkaProducerReturnConfig{
			Errors:    true,
			Successes: true,
		},
	}

	kafka, err := NewKafkaASyncProducer(conf)
	if err != nil {
		t.Fatal(err)
	}
	kafka.RegisterPrometheus("test")

	waitGroup := new(sync.WaitGroup)
	count := 0
	waitGroup.Add(1)
	testN := 10
	go func() {
		for {
			select {
			case e := <-kafka.Successes():
				t.Log(e)
			case e := <-kafka.Errors():
				t.Fatal(e)
			}
			count++
			if count == testN {
				waitGroup.Done()
			}
		}
	}()

	for i := 0; i < testN; i++ {
		err = kafka.PushMessage(&KafkaProducerMessage{
			Topic:     "test",
			Value:     []byte("test"),
			Partition: 0,
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	waitGroup.Wait()
	err = kafka.Close()
	if err != nil {
		t.Fatal(err)
	}

	if *metric {
		fmt.Println(metrics.GetPromethuesAsFmtText())
	}
}
