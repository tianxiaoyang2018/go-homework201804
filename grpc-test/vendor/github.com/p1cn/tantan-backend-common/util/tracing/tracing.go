package tracing

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/p1cn/tantan-backend-common/contextlog"
	"github.com/p1cn/tantan-backend-common/util/constant"

	"github.com/p1cn/tantan-backend-common/version"
	service "github.com/p1cn/tantan-domain-schema/golang/common"
)

func GetDebug(ctx context.Context) bool {
	b, _ := ctx.Value(constant.KeyTracingDebug).(bool)
	return b
}

func DebugFromHttpHeader(h http.Header) bool {
	b := h.Get(constant.KeyTracingDebug)
	x, _ := strconv.ParseBool(b)
	return x
}

func ServiceTraceFromHttpHeader(h http.Header) []string {
	t := h.Get(constant.KeyTracingServiceTrace)
	return strings.Split(t, ";")
}

func TraceIDFromHttpHeader(h http.Header) string {
	return h.Get(constant.KeyTracingTrace)
}

func SpanIDFromHttpHeader(h http.Header) string {
	return h.Get(constant.KeyTracingSpan)
}

func ParentSpanIDFromHttpHeader(h http.Header) string {
	return h.Get(constant.KeyTracingParentSpan)
}

func UserIdFromHttpHeader(h http.Header) string {
	return h.Get(constant.HeaderUserID)
}

func ExtFromHttpHeader(h http.Header) map[string]string {
	exth := h.Get(constant.KeyTracingExt)
	return TracingStringToMap(exth)
}

func DeviceIdFromHttpHeader(h http.Header) string {
	return h.Get(constant.KeyTracingDeviceId)
}

func SampledFromHttpHeader(h http.Header) float32 {
	sampledh := h.Get(constant.KeyTracingSampled)
	if len(sampledh) > 0 {
		ff, _ := strconv.ParseFloat(sampledh, 32)
		return float32(ff)
	}
	return 0.0
}

func init() {
	rand.Seed(time.Now().Unix())
}

func NewTraceID() string {
	return fmt.Sprintf("%016x%016x", random(), random())
}

func NewSpanID() string {
	return fmt.Sprintf("%016x", random())
}

func random() int64 {
	return rand.Int63()
}

// create a new service.Context
func NewServiceContext() *service.Context {
	return &service.Context{
		TraceID:      NewTraceID(),
		ParentSpanID: "",
		SpanID:       NewSpanID(),
		DeviceId:     "",
		Service:      version.ServiceName(),
		ServiceTrace: []string{version.ServiceName()},
		Ext:          make(map[string]string),
	}
}

func NewLogFromContext(ctx context.Context) *log.Log {
	sctx := GetServiceContext(ctx)
	if sctx == nil {
		return nil
	}

	return &log.Log{
		TraceID: sctx.TraceID,
	}
}
