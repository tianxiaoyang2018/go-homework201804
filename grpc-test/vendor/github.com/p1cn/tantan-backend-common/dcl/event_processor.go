package dcl

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/p1cn/tantan-backend-common/dcl/eventmeta"
	"github.com/p1cn/tantan-backend-common/dcl/processor"
	"github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/metrics"
	"github.com/p1cn/tantan-domain-schema/golang/event"
)

type eventProcessorManager struct {
	name     string
	dclTimer *metrics.Timer
	dclLag   *metrics.Timer
}

var (
	dclTimer       *metrics.Timer
	dclLag         *metrics.Timer
	prometheusOnce sync.Once
)

func newEventProcessorManager(name string) (*eventProcessorManager, error) {

	prometheusOnce.Do(func() {
		dclTimer = metrics.NewTimer(metrics.NameSpaceTantan, "dcl_consume_process", "DCL Consume event", []string{"name", "topic", "ret"})
		dclLag = metrics.NewTimer(metrics.NameSpaceTantan, "dcl_consume_lag", "DCL Consume event lag", []string{"name", "topic"})
	})

	return &eventProcessorManager{
		name:     name,
		dclLag:   dclLag,
		dclTimer: dclTimer,
	}, nil
}

func (self eventProcessorManager) CreateEventProcessor(process Processor) processor.EventProcessor {
	return &eventProcessor{
		name:        self.name,
		dclTimer:    self.dclTimer,
		dclLag:      self.dclLag,
		processFunc: process,
	}
}

type eventProcessor struct {
	name        string
	dclTimer    *metrics.Timer
	dclLag      *metrics.Timer
	processFunc Processor
}

func (ep *eventProcessor) Process(ctx context.Context, e *event.Event, meta *eventmeta.EventMetaData, done chan error) {
	done <- ep.process(ctx, e, meta)
}

func (ep *eventProcessor) process(ctx context.Context, event *event.Event, meta *eventmeta.EventMetaData) (err error) {
	ep.dclLag.Observe(time.Since(meta.Timestamp), ep.name, event.GetTopic())

	record := ep.dclTimer.Timer()
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 1<<18)
			n := runtime.Stack(buf, false)
			err = fmt.Errorf("%v, STACK: %s", r, buf[0:n])
			log.Crit(err.Error())
			record(ep.name, event.GetTopic(), "panic")
		}
	}()

	err = ep.processFunc(ctx, event, meta)
	if err == nil {
		record(ep.name, event.GetTopic(), "OK")
	} else {
		record(ep.name, event.GetTopic(), "error")
	}

	return err
}
