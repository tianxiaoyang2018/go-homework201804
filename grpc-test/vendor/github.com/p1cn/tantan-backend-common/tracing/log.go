package tracing

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/p1cn/tantan-backend-common/log"
)

type SpanCollector interface {
	Collect(span *ServiceSpanLog) error
}

type LoggerSpanCollector struct {
	Logger log.Logger
	Format string
}

func (self *LoggerSpanCollector) Collect(span *ServiceSpanLog) error {
	if self == nil {
		return nil
	}

	var str string
	if self.Format == "json" {
		dd, err := json.Marshal(span)
		if err != nil {
			return err
		}
		str = bytesToStringUnsafe(dd)
	} else {
		tags := "-"
		if len(span.Tags) > 0 {
			dd, _ := json.Marshal(&span.Tags)
			tags = bytesToStringUnsafe(dd)
		}
		logs := "-"
		if len(span.Logs) > 0 {
			dd, _ := json.Marshal(&span.Logs)
			logs = bytesToStringUnsafe(dd)
		}

		str = fmt.Sprintf("span\t%s\t%s\t-\t-\t%s\t%s\t%d\t%s\t%v\t%f\t%s\t%s\t%s\t%s",
			span.ServiceHost,
			span.ServiceName,
			span.TraceID,
			span.StartTime.Format("2006-01-02T15:04:05.000000"),
			span.Duration.Nanoseconds(),
			span.OperationName,
			span.Context.Debug,
			span.Context.Sampled,
			span.Context.ParentID,
			span.Context.SpanID,
			tags,
			logs,
		)
	}

	self.Logger.Infof(str)
	return nil
}

func bytesToStringUnsafe(b []byte) string {
	var str string
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	s := (*reflect.StringHeader)(unsafe.Pointer(&str))
	s.Data = h.Data
	s.Len = h.Len
	return str
}
