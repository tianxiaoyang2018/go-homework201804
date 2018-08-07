package contextlog

import (
	"context"
	"fmt"

	log "github.com/p1cn/tantan-backend-common/contextlog"
	"github.com/p1cn/tantan-backend-common/util/tracing"
)

// alert
func LogAlertf(ctx context.Context, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)

	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())

	clog.SetAlert(s)
	setFileContextLog(clog)
	getLogger().Alertf("%s", clog.ToJson())
}

func LogAlert(ctx context.Context, alert interface{}) {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	clog.SetAlert(alert)
	setFileContextLog(clog)
	getLogger().Alertf("%s", clog.ToJson())
}

/*
	LogAlertKV(ctx, KvPairs{
		"request" : request,
		"response" : response,
	})
*/
func LogAlertKv(ctx context.Context, kvs KvPairs) {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())

	for k, v := range kvs {
		if es, ok := v.(error); ok {
			clog.Ext[k] = es.Error()
		} else {
			clog.Ext[k] = v
		}
	}
	setFileContextLog(clog)
	getLogger().Alertf("%s", clog.ToJson())
}

// critical
func LogCritf(ctx context.Context, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	clog.SetError(s)
	setFileContextLog(clog)
	getLogger().Critf("%s", clog.ToJson())
}

func LogCrit(ctx context.Context, err interface{}) {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	clog.SetError(err)
	setFileContextLog(clog)
	getLogger().Critf("%s", clog.ToJson())
}

func LogCritKv(ctx context.Context, kvs KvPairs) {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	for k, v := range kvs {
		if es, ok := v.(error); ok {
			clog.Ext[k] = es.Error()
		} else {
			clog.Ext[k] = v
		}
	}
	setFileContextLog(clog)
	getLogger().Critf("%s", clog.ToJson())
}

// error
func LogErrf(ctx context.Context, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	clog.SetError(s)
	setFileContextLog(clog)
	getLogger().Errf("%s", clog.ToJson())
}

func LogErr(ctx context.Context, err interface{}) {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	clog.SetError(err)
	setFileContextLog(clog)
	getLogger().Errf("%s", clog.ToJson())
}

func LogErrKv(ctx context.Context, kvs KvPairs) {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	for k, v := range kvs {
		if es, ok := v.(error); ok {
			clog.Ext[k] = es.Error()
		} else {
			clog.Ext[k] = v
		}
	}
	setFileContextLog(clog)
	getLogger().Errf("%s", clog.ToJson())
}

// warning
func LogWarningf(ctx context.Context, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	clog.SetWarning(s)
	setFileContextLog(clog)
	getLogger().Warningf("%s", clog.ToJson())
}

func LogWarning(ctx context.Context, warning interface{}) {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	clog.SetWarning(warning)
	setFileContextLog(clog)
	getLogger().Warningf("%s", clog.ToJson())
}

func LogWarningKv(ctx context.Context, kvs KvPairs) {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	for k, v := range kvs {
		if es, ok := v.(error); ok {
			clog.Ext[k] = es.Error()
		} else {
			clog.Ext[k] = v
		}
	}
	setFileContextLog(clog)
	getLogger().Warningf("%s", clog.ToJson())
}

// Notice
func LogNoticef(ctx context.Context, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	clog.SetNotice(s)
	setFileContextLog(clog)
	getLogger().Noticef("%s", clog.ToJson())
}

func LogNotice(ctx context.Context, notice interface{}) {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	clog.SetNotice(notice)
	setFileContextLog(clog)
	getLogger().Noticef("%s", clog.ToJson())
}

func LogNoticeKv(ctx context.Context, kvs KvPairs) {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	for k, v := range kvs {
		if es, ok := v.(error); ok {
			clog.Ext[k] = es.Error()
		} else {
			clog.Ext[k] = v
		}
	}
	setFileContextLog(clog)
	getLogger().Noticef("%s", clog.ToJson())
}

// Info
func LogInfof(ctx context.Context, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	clog.SetInfo(s)
	setFileContextLog(clog)
	getLogger().Infof("%s", clog.ToJson())
}

/*
contextlog.LogInfo(ctx, response)
*/
func LogInfo(ctx context.Context, info interface{}) {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	clog.SetInfo(info)
	setFileContextLog(clog)
	getLogger().Infof("%s", clog.ToJson())
}

func LogInfoKv(ctx context.Context, kvs KvPairs) {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	for k, v := range kvs {
		if es, ok := v.(error); ok {
			clog.Ext[k] = es.Error()
		} else {
			clog.Ext[k] = v
		}
	}
	setFileContextLog(clog)
	getLogger().Infof("%s", clog.ToJson())
}

// Debug
func LogDebugf(ctx context.Context, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	clog.SetDebug(s)
	setFileContextLog(clog)
	getLogger().Debugf("%s", clog.ToJson())
}

func LogDebug(ctx context.Context, debug interface{}) {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	clog.SetDebug(debug)
	setFileContextLog(clog)
	getLogger().Debugf("%s", clog.ToJson())
}

func LogDebugKv(ctx context.Context, kvs KvPairs) {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	for k, v := range kvs {
		if es, ok := v.(error); ok {
			clog.Ext[k] = es.Error()
		} else {
			clog.Ext[k] = v
		}
	}
	setFileContextLog(clog)
	getLogger().Debugf("%s", clog.ToJson())
}
