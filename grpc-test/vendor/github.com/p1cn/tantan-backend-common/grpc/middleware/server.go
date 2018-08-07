package middleware

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/metrics"
	"github.com/p1cn/tantan-backend-common/tracing"
	"github.com/p1cn/tantan-backend-common/util/constant"
	log "github.com/p1cn/tantan-backend-common/util/contextlog"
	utilTracing "github.com/p1cn/tantan-backend-common/util/tracing"
	service "github.com/p1cn/tantan-domain-schema/golang/common"
)

var (
	promtheusOnce  sync.Once
	promtheusTimer *metrics.Timer
)
var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// 普罗米修斯中间件
func PrometheusUnaryInterceptorMiddleware(name string) grpc.UnaryServerInterceptor {
	// return grpc_prometheus.UnaryServerInterceptor
	slog.Info("installing prometheus middleware to grpc server(%s)", name)

	promtheusOnce.Do(func() {
		promtheusTimer = metrics.NewTimer(metrics.NameSpaceTantan, "rpc_request", "RPC request", []string{"name", "caller", "op_name", "ret"})
	})

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (r interface{}, err error) {
		sctx := GetServiceContext(ctx)
		st := sctx.GetServiceTrace()
		var caller string
		if len(st) > 1 {
			caller = st[len(st)-2]
		}

		record := promtheusTimer.Timer()
		r, err = handler(ctx, req)

		record(name, caller, extractMethodFromFullMethod(info.FullMethod), grpc.Code(err).String())
		return r, err
	}
}

// recovery中间件
func RecoveryUnaryInterceptorMiddleware(name string) grpc.UnaryServerInterceptor {
	slog.Info("installing recovery middleware to grpc server(%s)", name)

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 1<<15)
				n := runtime.Stack(buf, false)

				err = fmt.Errorf("%v, STACK: %s", r, buf[0:n])
				log.LogCrit(ctx, err)
			}
		}()

		return handler(ctx, req)
	}
}

// 打印tracing日志
func RequestLogUnaryInterceptorMiddleware(serviceName, serviceVersion string, serverName string, collector tracing.SpanCollector, logs, tags bool) grpc.UnaryServerInterceptor {
	slog.Info("installing trace collection middleware to grpc server(%s)", serverName)

	hostName, _ := os.Hostname()
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (r interface{}, err error) {
		sctx := GetServiceContext(ctx)

		writeLog := grpcRequestLogger(sctx, serviceName, serviceVersion, info.FullMethod, hostName, collector, logs, tags)

		r, err = handler(ctx, req)

		writeLog(err)
		return r, err
	}
}

func grpcRequestLogger(serviceContext *service.Context, serviceName, serviceVersion, method, hostName string, collector tracing.SpanCollector, logs, tags bool) (logRequest func(error)) {
	now := time.Now()
	return func(err error) {

		if serviceContext == nil {
			slog.Warning("service context is nil: %v, %v , %v", serviceName, method, err)
			return
		}
		if serviceContext.Sampled < tracing.MinSampled {
			return
		}

		var spanLogs []tracing.LogRecord
		if err != nil && logs {
			spanLogs = append(spanLogs, tracing.LogRecord{
				Timestamp: time.Now(),
				Fields: []tracing.Field{
					tracing.Field{
						Key:          tracing.SpanFieldErrorObject,
						FieldType:    tracing.ObjectType,
						InterfaceVal: err.Error(),
					},
				},
			})
		}

		span := tracing.ServiceSpanLog{
			TraceID:        serviceContext.TraceID,
			ServiceName:    serviceName,
			ServiceVersion: serviceVersion,
			ServiceHost:    hostName,

			StartTime:     now,
			Duration:      time.Since(now),
			OperationName: method,
			Context:       tracing.SpanContextFromServiceContext(serviceContext),
		}

		if tags {
			span.Tags = tracing.SpanTags{
				string(ext.SpanKind): ext.SpanKindRPCServerEnum,
			}
		}

		collector.Collect(&span)
	}
}

func RequestContextUnaryInterceptorMiddleware(serviceName, serverName string) grpc.UnaryServerInterceptor {
	slog.Info("installing trace context middleware to grpc server(%s)", serverName)

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (r interface{}, err error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if ok && len(md) > 0 {
			sctx := utilTracing.NewServiceContext()
			// trace id
			if vv := md[strings.ToLower(constant.KeyTracingTrace)]; len(vv) > 0 {
				sctx.TraceID = vv[0]
			}
			// parent span id
			if vv := md[strings.ToLower(constant.KeyTracingSpan)]; len(vv) > 0 {
				sctx.ParentSpanID = vv[0]
			}
			// span id
			sctx.SpanID = utilTracing.NewSpanID()

			// debug
			if vv := md[strings.ToLower(constant.KeyTracingDebug)]; len(vv) > 0 {
				x, _ := strconv.ParseBool(vv[0])
				sctx.Debug = x
			}

			// sampled
			if vv := md[strings.ToLower(constant.KeyTracingSampled)]; len(vv) > 0 {
				x, _ := strconv.ParseFloat(vv[0], 32)
				sctx.Sampled = float32(x)
			}
			// device id
			if vv := md[strings.ToLower(constant.KeyTracingDeviceId)]; len(vv) > 0 {
				sctx.DeviceId = vv[0]
			}

			// service name
			sctx.Service = serviceName

			// service trace
			if vv := md[strings.ToLower(constant.KeyTracingServiceTrace)]; len(vv) > 0 {
				st := strings.Split(vv[0], ";")
				st = append(st, serviceName)
				sctx.ServiceTrace = st
			}

			// ext
			if vv := md[strings.ToLower(constant.KeyTracingExt)]; len(vv) > 0 {
				sctx.Ext = utilTracing.TracingStringToMap(vv[0])
			}

			ctx = SetServiceContext(ctx, sctx)

		} else {
			slog.Warning("invalid request context : %v", info.FullMethod)
		}

		r, err = handler(ctx, req)

		return r, err
	}
}

func SetServiceContext(ctx context.Context, sCtx *service.Context) context.Context {
	return context.WithValue(ctx, constant.KeyServiceContext, sCtx)
}

func GetServiceContext(ctx context.Context) *service.Context {
	ret, _ := ctx.Value(constant.KeyServiceContext).(*service.Context)
	return ret
}

func extractMethodFromFullMethod(fullMethod string) string {
	split := strings.Split(fullMethod, "/")
	return split[len(split)-1]
}
