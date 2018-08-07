package util

import (
	"encoding/json"
	"runtime"
	"time"

	slog "github.com/p1cn/tantan-backend-common/log"
)

// Tracer is mainly used for tracing elapsed time during a func call.
type Tracer struct {
	startTime time.Time
	elapsed   time.Duration
	funcName  string
	msg       map[string]interface{}
}

func (t Tracer) renderJSON() string {
	b, _ := json.Marshal(map[string]interface{}{
		"elapsedTime": t.elapsed / time.Millisecond,
		"funcName":    t.funcName,
		"message":     t.msg,
	})
	return string(b)
}

// Stop prints log.
func (t *Tracer) Stop() {
	t.elapsed = time.Since(t.startTime)
	slog.Info("JSON %s", t.renderJSON())
}

// StopWithWarn will log WARN if func call costs time
// more than dur.
func (t *Tracer) StopWithWarn(dur time.Duration) {
	t.elapsed = time.Since(t.startTime)
	if t.elapsed > dur {
		slog.Warning("LOW FUNC CALL >(%s) JSON %s", dur.String(), t.renderJSON())
		return
	}
	slog.Info("JSON %s", t.renderJSON())
}

// NewTracer creates a new tracer.
// Usage: defer NewTracer(msg).Stop()
func NewTracer(msg map[string]interface{}) *Tracer {
	pc, _, _, _ := runtime.Caller(1)
	return &Tracer{
		startTime: time.Now(),
		funcName:  runtime.FuncForPC(pc).Name(),
		msg:       msg,
	}
}
