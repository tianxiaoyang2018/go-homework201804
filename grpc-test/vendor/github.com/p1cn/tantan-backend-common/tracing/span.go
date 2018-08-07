package tracing

import (
	"time"
)

type ServiceSpanLog struct {
	TraceID        string
	ServiceName    string `json:"ServiceName,omitempty"`
	ServiceVersion string `json:"ServiceVer,omitempty"`
	ServiceHost    string `json:"Host,omitempty"`

	StartTime     time.Time
	Duration      time.Duration
	OperationName string `json:"OpName"`
	Context       SpanContext
	Tags          SpanTags    `json:"Tags,omitempty"`
	Logs          []LogRecord `json:"Logs,omitempty"`
	References    []Reference `json:"Refs,omitempty"`
}

type SpanContext struct {
	Debug    bool              `json:"debug,omitempty"`
	Sampled  float32           `json:"sampled,omitempty"`
	Baggage  map[string]string `json:"Baggage,omitempty"`
	SpanID   string            `json:"spanId,omitempty"`
	ParentID string            `json:"parentId,omitempty"`
}

type SpanTags map[string]interface{}

type LogRecord struct {
	Timestamp time.Time
	Fields    []Field `json:"Fields,omitempty"`
}

type Field struct {
	Key          string
	FieldType    FieldType
	NumericVal   *int64  `json:"int64,omitempty"`
	StringVal    *string `json:"str,omitempty"`
	InterfaceVal interface{}
}

type Reference struct {
	Type    int
	Context SpanContext
}

type FieldType int

const (
	StringType FieldType = iota
	BoolType
	IntType
	Int32Type
	Uint32Type
	Int64Type
	Uint64Type
	Float32Type
	Float64Type
	ErrorType
	ObjectType
	LazyLoggerType
	NoopType
)
