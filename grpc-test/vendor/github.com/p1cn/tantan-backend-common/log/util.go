package log

import (
	"fmt"
	"runtime"
)

const (
	LogTimeFormat = "2006-01-02T15:04:05.000000"
)

type FileHeaderFunc func() string

func fileHeader(deep int) FileHeaderFunc {

	return func() string {
		var (
			ok   bool
			file string
			line int
		)
		_, file, line, ok = runtime.Caller(deep)
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
}

func parseFlags(flagStrs []string) LFlag {
	var flags LFlag
	for _, c := range flagStrs {
		switch c {
		case "level":
			flags |= LFlagLevel
		case "date":
			flags |= LFlagDate
		case "file":
			flags |= LFlagFile
		}
	}
	return flags
}

func parseLevel(s string) LogLevel {
	for k, v := range levelMap {
		if s == v {
			return k
		}
	}
	// default level
	return LevelAll
}
