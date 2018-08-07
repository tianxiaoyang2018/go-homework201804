package middleware

import (
	"context"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/util/constant"
	"github.com/p1cn/tantan-backend-common/util/tracing"
)

func ContextClientUnaryInterceptorMiddleware() grpc.UnaryClientInterceptor {

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		sctx := GetServiceContext(ctx)
		if sctx != nil {

			data := map[string]string{
				constant.KeyTracingTrace:        sctx.TraceID,
				constant.KeyTracingParentSpan:   sctx.ParentSpanID,
				constant.KeyTracingSpan:         sctx.SpanID,
				constant.KeyTracingDeviceId:     sctx.DeviceId,
				constant.KeyTracingSampled:      strconv.FormatFloat(float64(sctx.Sampled), 'f', 6, 32),
				constant.KeyTracingServiceTrace: strings.Join(sctx.ServiceTrace, ";"),
				constant.KeyTracingDebug:        strconv.FormatBool(sctx.Debug),
				constant.KeyTracingExt:          tracing.TracingMapToString(sctx.Ext),
			}

			ctx = metadata.NewOutgoingContext(ctx, metadata.New(data))
		} else {
			slog.Warning("meta of service context does not exist : %v", method)
		}

		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}

func TimeoutClientUnaryInterceptorMiddleware(duration time.Duration) grpc.UnaryClientInterceptor {

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx2, cancel := context.WithTimeout(ctx, duration)
		defer cancel()
		err := invoker(ctx2, method, req, reply, cc, opts...)
		return err
	}
}
