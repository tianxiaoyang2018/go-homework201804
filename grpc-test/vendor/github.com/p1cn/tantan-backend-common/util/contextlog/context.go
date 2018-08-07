package contextlog

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	log "github.com/p1cn/tantan-backend-common/contextlog"

	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/util/tracing"
)

type KvPairs map[string]interface{}

var (
	contextLogger slog.Logger
	once          sync.Once
)

//
/*
e.g. : gin handler
	func Handle(ctx *gin.Context) {
		NewLog(ctx.Request.Context)
	}
*/
func NewLog(ctx context.Context) *log.Log {
	sctx := tracing.GetServiceContext(ctx)
	clog := log.NewLogWithType(log.LogTypeContext, sctx.GetTraceID())
	return clog
}

func getLogger() slog.Logger {
	once.Do(func() {
		contextLogger = slog.GetLogger("contextlog")
		contextLogger.SetFileHeader(contextFileHeader)
	})
	return contextLogger
}

func contextFileHeader() string {
	var (
		ok   bool
		file string
		line int
	)
	_, file, line, ok = runtime.Caller(6)
	if !ok {
		file = "???"
		line = 0
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	return fmt.Sprintf("%v:%v: ", file, line)
}

func getFile() string {
	var (
		ok   bool
		file string
		line int
	)
	_, file, line, ok = runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	return fmt.Sprintf("%v:%v: ", file, line)
}

func setFileContextLog(clog *log.Log) {

	clog.SetExt("file", getFile())
}
