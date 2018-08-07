package middleware

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/opentracing/opentracing-go/ext"

	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/metrics"
	tracingLog "github.com/p1cn/tantan-backend-common/tracing"
	"github.com/p1cn/tantan-backend-common/util/constant"
	utilContextLog "github.com/p1cn/tantan-backend-common/util/contextlog"
	"github.com/p1cn/tantan-backend-common/util/tracing"
	service "github.com/p1cn/tantan-domain-schema/golang/common"
)

var (
	initOnce sync.Once
	opTimer  *metrics.Timer
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func PrometheusHttpRequest(serviceName string, serverName string, handlerRoute func(method string, handler string) string) gin.HandlerFunc {
	slog.Info("installing prometheus middleware to http server(%s)", serverName)

	initOnce.Do(func() {
		opTimer = metrics.NewTimer(metrics.NameSpaceTantan, "http_request", "HTTP request", []string{"http_server", "method", "url", "status_code"})
	})
	return func(c *gin.Context) {
		record := opTimer.Timer()

		opName := handlerRoute(c.Request.Method, c.HandlerName())
		c.Next()

		record(serverName, c.Request.Method, opName, strconv.Itoa(c.Writer.Status()))
	}
}

func Recovery(serverName string) gin.HandlerFunc {
	slog.Info("installing recovery middleware to http server(%s)", serverName)

	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 1<<15)
				n := runtime.Stack(buf, false)
				err := fmt.Errorf("%v, STACK: %s", r, buf[0:n])

				utilContextLog.LogCrit(c.Request.Context(), err)

				c.AbortWithError(500, err)
			}
		}()
		c.Next()
	}
}

func Logging(serviceName, serviceVersion, serverName string, handlerRoute func(method string, handler string) string, collector tracingLog.SpanCollector, logs, tags bool) gin.HandlerFunc {
	slog.Info("installing trace collection middleware to http server(%s)", serverName)
	hostName, _ := os.Hostname()
	return func(c *gin.Context) {
		writeLog := httpRequestLogger(c, serviceName, serviceVersion, serverName, handlerRoute(c.Request.Method, c.HandlerName()), hostName, collector, logs, tags)
		c.Next()

		var e error
		if len(c.Errors) > 0 {
			e = c.Errors.Last()
		}
		writeLog(e)
	}
}

func getServiceContextFromGin(c *gin.Context, serviceName string) *service.Context {

	TraceID := tracing.TraceIDFromHttpHeader(c.Request.Header)
	SpanID := tracing.SpanIDFromHttpHeader(c.Request.Header)

	ServiceTrace := tracing.ServiceTraceFromHttpHeader(c.Request.Header)
	ServiceTrace = append(ServiceTrace, serviceName)

	if len(TraceID) == 0 {
		TraceID = tracing.NewTraceID()
		slog.Warning("trace id is empty : %v", serviceName)
	}

	// new span
	if len(SpanID) == 0 {
		SpanID = tracing.NewSpanID()
	}

	sctx := &service.Context{
		TraceID:      TraceID,
		ParentSpanID: SpanID,
		SpanID:       tracing.NewSpanID(),

		Service:      serviceName,
		ServiceTrace: ServiceTrace,
		Debug:        tracing.DebugFromHttpHeader(c.Request.Header),
		DeviceId:     tracing.DeviceIdFromHttpHeader(c.Request.Header),
		Ext:          tracing.ExtFromHttpHeader(c.Request.Header),
		Sampled:      tracing.SampledFromHttpHeader(c.Request.Header),
	}

	return sctx
}

func RequestContext(serviceName, serverName string) gin.HandlerFunc {
	slog.Info("installing trace context middleware to http server(%s)", serverName)

	return func(c *gin.Context) {

		if c.Request != nil {
			serviceContext := getServiceContextFromGin(c, serviceName)

			ctx := c.Request.Context()
			ctx = context.WithValue(ctx, constant.KeyServiceContext, serviceContext)
			ctx = context.WithValue(ctx, constant.HeaderUserID, c.Request.Header.Get(constant.HeaderUserID))
			ctx = context.WithValue(ctx, constant.KeyTracingReqStartTime, time.Now().UnixNano())

			c.Request = c.Request.WithContext(ctx)
		} else {
			slog.Warning("request is nil : %v", serviceName)
		}

		c.Next()
	}
}

func httpRequestLogger(c *gin.Context, serviceName, serviceVersion, serverName, route string, hostName string, collector tracingLog.SpanCollector, logs, tags bool) (logRequest func(error)) {
	now := time.Now()
	return func(err error) {

		if c.Request == nil {
			slog.Warning("request is nil : %v , %v", serviceName, err)
			return
		}

		sctx, _ := c.Request.Context().Value(constant.KeyServiceContext).(*service.Context)
		if sctx == nil {
			slog.Warning("request context is nil : %v, %v", serviceName, err)
			return
		}

		if sctx.Sampled < tracingLog.MinSampled {
			return
		}

		var spanLogs []tracingLog.LogRecord
		if err != nil && logs {
			spanLogs = append(spanLogs, tracingLog.LogRecord{
				Timestamp: time.Now(),
				Fields: []tracingLog.Field{
					tracingLog.Field{
						Key:          tracingLog.SpanFieldErrorObject,
						FieldType:    tracingLog.ObjectType,
						InterfaceVal: err.Error(),
					},
				},
			})
		}

		span := tracingLog.ServiceSpanLog{
			TraceID:        sctx.TraceID,
			ServiceName:    serviceName,
			ServiceVersion: serviceVersion,
			ServiceHost:    hostName,
			StartTime:      now,
			Duration:       time.Since(now),
			OperationName:  c.Request.Method + "_" + route,
			Context:        tracingLog.SpanContextFromServiceContext(sctx),
		}

		if tags {
			span.Tags = tracingLog.SpanTags{
				string(ext.SpanKind):       ext.SpanKindRPCServerEnum,
				string(ext.HTTPStatusCode): c.Writer.Status(),
				string(ext.HTTPMethod):     c.Request.Method,
				string(ext.HTTPUrl):        c.Request.RequestURI,
			}
		}
		if logs {
			span.Logs = spanLogs
		}

		collector.Collect(&span)
	}
}

func handlerName(c *gin.Context) string {
	caller := runtime.FuncForPC(reflect.ValueOf(c.Handler()).Pointer())

	split := strings.Split(caller.Name(), ".")
	fname := split[len(split)-1]
	fname = strings.Split(fname, ")")[0]

	return fname
}
