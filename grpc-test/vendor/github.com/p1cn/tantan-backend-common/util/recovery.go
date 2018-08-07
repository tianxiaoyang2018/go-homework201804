package util

import (
	"runtime"

	slog "github.com/p1cn/tantan-backend-common/log"
)

func Recovery() {
	if r := recover(); r != nil {
		buf := make([]byte, 1<<18)
		n := runtime.Stack(buf, false)
		slog.Crit("%v, STACK: %s", r, buf[0:n])
	}
}
