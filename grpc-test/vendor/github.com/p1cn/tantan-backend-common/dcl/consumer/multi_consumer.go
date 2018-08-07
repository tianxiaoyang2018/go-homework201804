package consumer

import (
	slog "github.com/p1cn/tantan-backend-common/log"
)

type MultiConsumer interface {
	Start() error
	Close() error
}

// NewMultiConsumer ...
func NewMultiConsumer(configs []ConsumerConfig) (MultiConsumer, error) {
	consumers := []*Consumer{}
	for _, cfg := range configs {
		c, err := NewConsumer(cfg)
		if err != nil {
			return nil, err
		}
		consumers = append(consumers, c)
	}
	return &multiConsumer{consumers: consumers}, nil
}

// MultiConsumer ...
type multiConsumer struct {
	consumers []*Consumer
}

// Start ...
func (mc *multiConsumer) Start() error {
	startRet := make(chan error)
	for _, c := range mc.consumers {
		go func(consumer *Consumer) {
			startRet <- consumer.Start()
		}(c)
	}
	err := <-startRet
	retNum := 1
	go func() {
		for {
			<-startRet
			retNum++
			if retNum == len(mc.consumers) {
				return
			}
		}
	}()
	for _, c := range mc.consumers {
		c.Close()
	}
	return err
}

// Close ...
func (mc *multiConsumer) Close() error {
	closeRet := make(chan struct{})
	for _, c := range mc.consumers {
		go func(consumer *Consumer) {
			slog.Info("Closing topic: %v group: %v\n", consumer.c.cfg.Topic, consumer.c.cfg.Group)
			consumer.Close()
			closeRet <- struct{}{}
		}(c)
	}
	retNum := 0
	for {
		<-closeRet
		retNum++
		if retNum == len(mc.consumers) {
			return nil
		}
	}
}
