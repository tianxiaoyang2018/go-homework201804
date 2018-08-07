package dcl

// import (
// 	"hash/fnv"
// 	"sync"

// 	"github.com/p1cn/tantan-backend-common/dcl/processor"
// 	"github.com/p1cn/tantan-backend-common/dcl/producer"
// 	"github.com/p1cn/tantan-domain-schema/golang/event"
// 	"github.com/p1cn/tantan-backend-common/dcl/eventmeta"
// )

// func NewEventWorkerPool(
// 	concurrency int,
// 	maintainKeyOrder bool,
// 	processor processor.EventSyncProcessor) *EventWorkerPool {

// 	p := &EventWorkerPool{
// 		concurrency:      concurrency,
// 		processor:        processor,
// 		maintainKeyOrder: maintainKeyOrder,
// 	}

// 	for i := 0; i < concurrency; i++ {
// 		jobQueue := make(chan *job, 10)
// 		p.jobQueues = append(p.jobQueues, jobQueue)
// 		go p.runWorker(jobQueue)
// 	}

// 	return p
// }

// type EventWorkerPool struct {
// 	concurrency      int
// 	jobQueues        []chan *job
// 	processor        processor.EventSyncProcessor
// 	maintainKeyOrder bool
// 	iter             int
// 	iterMutex        sync.Mutex
// }

// func (p *EventWorkerPool) runWorker(jobQueue chan *job) {
// 	for job := range jobQueue {
// 		job.done <- p.processor.Process(job.event, job.meta)
// 	}
// }

// func (p *EventWorkerPool) Process(event *event.Event, meta *eventmeta.EventMetaData, done chan error) {
// 	if p.maintainKeyOrder {
// 		p.processOrder(event, meta, done)
// 	} else {
// 		p.processDisorder(event, meta, done)
// 	}
// }

// func (p *EventWorkerPool) processDisorder(event *event.Event, meta *eventmeta.EventMetaData, done chan error) {
// 	p.iterMutex.Lock()
// 	p.iter++
// 	nr := p.iter
// 	p.iterMutex.Unlock()

// 	index := nr % p.concurrency
// 	p.jobQueues[index] <- &job{event, meta, done}
// }

// func (p *EventWorkerPool) processOrder(event *event.Event, meta *eventmeta.EventMetaData, done chan error) {
// 	index := int(hash(producer.GetEventKey(event))) % p.concurrency
// 	p.jobQueues[index] <- &job{event, meta, done}
// }

// type job struct {
// 	event *event.Event
// 	meta  *eventmeta.EventMetaData
// 	done  chan error
// }

// func hash(s string) uint32 {
// 	h := fnv.New32a()
// 	h.Write([]byte(s))
// 	return h.Sum32()
// }
