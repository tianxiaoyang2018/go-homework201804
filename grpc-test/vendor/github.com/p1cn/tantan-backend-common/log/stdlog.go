package log

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 非线程安全，仅调试用
type stdLogger struct {
	name  string
	flags LFlag
	fh    FileHeaderFunc
}

func newStdLogger(name string, flags LFlag) (LoggerHandler, error) {
	ll := &stdLogger{
		name:  name,
		flags: flags,
		fh:    fileHeader(5),
	}

	return ll, nil
}

func (self *stdLogger) Clone() (LoggerHandler, error) {
	return &stdLogger{
		name:  self.name,
		flags: self.flags,
		fh:    self.fh,
	}, nil
}

func (self *stdLogger) SetFileHeader(fh FileHeaderFunc) {
	self.fh = fh
}

func (self *stdLogger) Close() error {
	return nil
}

func (self *stdLogger) Flush() error {
	return nil
}

func (self *stdLogger) Logf(level LogLevel, format string, v ...interface{}) error {
	s := fmt.Sprintf(format, v...)
	if self.name == "stderr" {
		os.Stderr.WriteString(self.stdHeader(level) + s + "\n")
	} else if self.name == "stdout" {
		os.Stdout.WriteString(self.stdHeader(level) + s + "\n")
	}
	return nil
}

func (self *stdLogger) stdHeader(level LogLevel) string {
	var header string
	if self.flags&LFlagDate > 0 {
		header += time.Now().Format(LogTimeFormat)
	}

	if self.flags&LFlagLevel > 0 {
		if len(header) > 0 && header[len(header)-1] != ' ' {
			header += " "
		}
		header += strings.ToUpper(levelMap[level])
	}

	if len(header) > 0 && header[len(header)-1] != ' ' {
		header += " "
	}
	// header += path.Base(os.Args[0])

	if self.flags&LFlagFile > 0 {
		if len(header) > 0 && header[len(header)-1] != ' ' {
			header += " "
		}
		header += self.fh()
	}

	if len(header) > 0 && header[len(header)-1] != ' ' {
		header += " "
	}

	return header
}
