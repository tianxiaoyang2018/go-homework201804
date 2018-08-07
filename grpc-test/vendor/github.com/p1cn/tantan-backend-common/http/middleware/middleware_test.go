package middleware

import (
	stdJson "encoding/json"
	"testing"
	"time"

	"github.com/p1cn/tantan-backend-common/contextlog"
	tracingLog "github.com/p1cn/tantan-backend-common/tracing"
	"github.com/pquerna/ffjson/ffjson"

	jsoniter "github.com/json-iterator/go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/p1cn/tantan-backend-common/util/tracing"
)

func testGetLog() contextlog.Log {
	sctx := tracing.NewServiceContext()
	serviceLog := contextlog.Log{
		Type:           contextlog.LogTypeSpan,
		ServiceName:    "test-json-demo",
		ServiceVersion: "v1.0.0",
		ServerName:     "http-server",
		TraceID:        "5822107fcfd52d29a0f5f3f164fd",

		Span: &tracingLog.ServiceSpanLog{
			StartTime:     time.Now(),
			Duration:      1 * time.Second,
			OperationName: "GET_/users/:id",
			Context:       tracingLog.SpanContextFromServiceContext(sctx),
		},
		Ext: make(map[string]interface{}),
	}

	serviceLog.Span.Tags = tracingLog.SpanTags{
		string(ext.SpanKind):       ext.SpanKindRPCServerEnum,
		string(ext.HTTPStatusCode): "200",
		string(ext.HTTPMethod):     "GET",
		string(ext.HTTPUrl):        "/users/:id",
	}
	return serviceLog
}

func BenchmarkJson(b *testing.B) {
	clog := testGetLog()
	for i := 0; i < b.N; i++ {
		_, err := stdJson.Marshal(clog)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDidiJson(b *testing.B) {
	didiJson := jsoniter.ConfigCompatibleWithStandardLibrary

	clog := testGetLog()
	for i := 0; i < b.N; i++ {
		_, err := didiJson.Marshal(clog)

		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkffJson(b *testing.B) {
	clog := testGetLog()
	for i := 0; i < b.N; i++ {
		_, err := ffjson.Marshal(clog)

		if err != nil {
			b.Fatal(err)
		}
	}
}

//BenchmarkJson-8   	  200000	      9385 ns/op	    3824 B/op	      32 allocs/op
//BenchmarkDidiJson-8     200000	      5996 ns/op	    2791 B/op	      29 allocs/op
