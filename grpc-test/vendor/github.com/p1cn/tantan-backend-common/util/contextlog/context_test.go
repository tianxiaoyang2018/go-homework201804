package contextlog

import (
	"context"
	"errors"
	"testing"

	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/util/tracing"
)

func initLog() {
	slog.Init(slog.Config{
		Output: []string{"stderr", "syslog"},
		Flags:  []string{"level", "file", "date"},
	})
}

func BenchmarkErr(b *testing.B) {
	initLog()
	err := errors.New("ff")

	ctx := tracing.NewServiceContextToContext(context.Background())
	for i := 0; i < b.N; i++ {
		LogErr(ctx, err)
	}
}

func TestErr(t *testing.T) {
	initLog()
	err := errors.New("ff")

	ctx := tracing.NewServiceContextToContext(context.Background())

	LogErr(ctx, err)
	LogWarning(ctx, err)
	LogErrf(ctx, "test logerrf : %v", err)

	slog.Info("info")
	slog.Alert("alert")
	slog.Crit("crit")
	slog.Debug("debug")
	slog.Err("err")

}

func TestKvPair(t *testing.T) {
	initLog()

	err := errors.New("ff")

	ctx := tracing.NewServiceContextToContext(context.Background())

	LogAlertKv(ctx, KvPairs{
		"A":   "aaa",
		"B":   12,
		"err": err,
	})
}
