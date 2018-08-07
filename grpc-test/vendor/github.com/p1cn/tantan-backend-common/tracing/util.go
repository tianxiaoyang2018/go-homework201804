package tracing

import (
	service "github.com/p1cn/tantan-domain-schema/golang/common"
)

const (
	MinSampled = 0.0001
)

func SpanContextFromServiceContext(ctx *service.Context) SpanContext {

	bag := map[string]string{}
	for _, v := range ctx.Baggages {
		bag[v.Key] = v.Value
	}

	return SpanContext{
		Sampled:  ctx.GetSampled(),
		Debug:    ctx.GetDebug(),
		SpanID:   ctx.GetSpanID(),
		ParentID: ctx.GetParentSpanID(),
		Baggage:  bag,
	}
}
