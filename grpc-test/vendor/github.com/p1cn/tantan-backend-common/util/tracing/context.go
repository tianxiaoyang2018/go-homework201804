package tracing

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/p1cn/tantan-backend-common/util/constant"
	"github.com/p1cn/tantan-backend-common/version"
	service "github.com/p1cn/tantan-domain-schema/golang/common"
)

// 繁衍一个context：提取出老context的service.Context 且放入新的context中
// 对于做一些异步操作，且不希望老context中的超时或cancel机制影响操作
// e.g.
// gin的context里面都包含cancel机制的。当一个http请求结束后，cancel就会被调用，
// 如果要做一些异步操作，且希望http的handle函数结束后继续执行异步操作，
// 则应该繁衍一个新的context；否则如果将gin的context传递到一个异步的grpc请求中，http请求结束会导致异步grpc请求也结束
func PropagateContextWithServiceContext(ctx context.Context) context.Context {
	sctx := GetServiceContext(ctx)
	return SetServiceContext(context.Background(), sctx)
}

// 新建一个service.Context 且放入到 参数context中
// 只适合没有上下文信息的job 类程序中。
func NewServiceContextToContext(ctx context.Context) context.Context {
	sctx := NewServiceContext()
	return SetServiceContext(ctx, sctx)
}

// 从http头中提取出service.Context , 且生成新的span id，
func NewServiceContextFromHttpHeader(h http.Header) *service.Context {
	s := version.ServiceName()
	tr := ServiceTraceFromHttpHeader(h)
	tr = append(tr, s)

	sctx := &service.Context{
		TraceID:      TraceIDFromHttpHeader(h),
		ParentSpanID: SpanIDFromHttpHeader(h),
		SpanID:       NewSpanID(),
		Service:      s,
		ServiceTrace: tr,
		DeviceId:     DeviceIdFromHttpHeader(h),
		Debug:        DebugFromHttpHeader(h),
		Ext:          ExtFromHttpHeader(h),
		Sampled:      SampledFromHttpHeader(h),
	}

	return sctx
}

// 从http头中提取出service.Context 不生成新的span id，建议不要轻易使用。
// 一般情况下，都要生成新的span id
func GetServiceContextFromHttpHeader(h http.Header) *service.Context {

	tr := ServiceTraceFromHttpHeader(h)
	var sn string
	if len(tr) > 0 {
		sn = tr[len(tr)-1]
	}

	sctx := &service.Context{
		TraceID:      TraceIDFromHttpHeader(h),
		ParentSpanID: ParentSpanIDFromHttpHeader(h),
		SpanID:       SpanIDFromHttpHeader(h),
		Service:      sn,
		ServiceTrace: tr,
		DeviceId:     DeviceIdFromHttpHeader(h),
		Debug:        DebugFromHttpHeader(h),
		Ext:          ExtFromHttpHeader(h),
		Sampled:      SampledFromHttpHeader(h),
	}

	return sctx
}

// 从context.Context中获取出service.Context的信息加入到http的header里面
// TraceID, ParentSpanID, SpanID
// 当调用一个http服务的时候，需要调用此函数把trace信息放入到http头中
func TraceFromContextToHttpHeader(ctx context.Context, h http.Header) {
	sctx := GetServiceContext(ctx)
	h.Add(constant.KeyTracingTrace, sctx.GetTraceID())
	h.Add(constant.KeyTracingParentSpan, sctx.GetParentSpanID())
	h.Add(constant.KeyTracingSpan, sctx.GetSpanID())
	h.Add(constant.KeyTracingSampled, strconv.FormatFloat(float64(sctx.GetSampled()), 'f', 6, 32))
	h.Add(constant.KeyTracingDeviceId, sctx.GetDeviceId())
	h.Add(constant.KeyTracingServiceTrace, strings.Join(sctx.GetServiceTrace(), ";"))
	h.Add(constant.KeyTracingDebug, strconv.FormatBool(sctx.GetDebug()))
	h.Add(constant.KeyTracingExt, TracingMapToString(sctx.GetExt()))
}

func GetServiceContext(ctx context.Context) *service.Context {
	ret, _ := ctx.Value(constant.KeyServiceContext).(*service.Context)
	return ret
}

func SetServiceContext(ctx context.Context, sctx *service.Context) context.Context {
	return context.WithValue(ctx, constant.KeyServiceContext, sctx)
}

func AppendServiceNameToServiceContext(serviceContext *service.Context, serviceName string) *service.Context {
	if serviceContext == nil {
		serviceContext = NewServiceContext()
	}

	serviceContext.ServiceTrace = append(serviceContext.ServiceTrace, serviceName)
	serviceContext.Service = serviceName
	return serviceContext
}
