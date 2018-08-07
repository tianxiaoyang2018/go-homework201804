package consumer

import (
	"flag"
	"testing"

	"github.com/p1cn/tantan-backend-common/config"

	"sync"
)

var zookeepers = flag.String("zookeepers", "192.168.4.11:2181", "zookeepers' addresses")

func TestConsume(t *testing.T) {
	kafkaConfig := &config.KafkaConfig{
		Zookeepers: []string{*zookeepers},
	}
	kafkaConfig.Zookeepers = []string{*zookeepers}
	conf := NewKafkaConsumerGroupConfig("test1", []string{"test", "eventlog"}, kafkaConfig)
	conf.Config.Offsets.ResetOffsets = true
	c, err := NewKafkaConsumerGroup(conf)
	if err != nil {
		t.Fatal(err)
	}

	count := 0
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(1)
	go func() {
		for {
			select {
			case e, ok := <-c.Errors():
				if !ok {
					return
				}
				t.Fatal(e)
			case e, ok := <-c.Messages():
				if !ok {
					return
				}
				t.Log(e)
				c.CommitUpto(e)
			}
			count++
			if count == 1 {
				waitGroup.Done()
			}
		}
	}()

	waitGroup.Wait()
	err = c.Close()

	if err != nil {
		t.Fatal(err)
	}
}
