/*
Package slog implements log to stderr and syslog.

Log Levels
----------
Emerg (panic)	System is unusable.					A "panic" condition usually affecting multiple apps/servers/sites. At this level it would usually notify all tech staff on call.
Alert			Action must be taken immediately.	Should be corrected immediately, therefore notify staff who can fix the problem. An example would be the loss of a primary ISP connection.
Crit			Critical conditions.				Should be corrected immediately, but indicates failure in a secondary system, an example is a loss of a backup ISP connection.
Err (error)		Error conditions.					Non-urgent failures, these should be relayed to developers or admins; each item must be resolved within a given time.
Warning (warn)	Warning conditions.					Warning messages, not an error, but indication that an error will occur if action is not taken, e.g. file system 85% full - each item must be resolved within a given time.
Notice			Normal but significant condition.	Events that are unusual but not error conditions - might be summarized in an email to developers or admins to spot potential problems - no immediate action required.
Info			Informational messages.				Normal operational messages - may be harvested for reporting, measuring throughput, etc. - no action required.
Debug			Debug-level messages.				Info useful to developers for debugging the application, not useful during operations.

*/

package log

import (
	"errors"
	"fmt"
	"sync/atomic"
)

type defaultLoggerImpl struct {
	closed   int32
	level    LogLevel
	handlers []LoggerHandler
}

func (self *defaultLoggerImpl) clone() (*defaultLoggerImpl, error) {

	var handlers []LoggerHandler
	for _, l := range self.handlers {
		ll, err := l.Clone()
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, ll)
	}
	return &defaultLoggerImpl{
		level:    self.level,
		handlers: handlers,
	}, nil
}

var (
	defaultLogger *defaultLoggerImpl
)

func (self *defaultLoggerImpl) GetLevel() LogLevel {
	return self.level
}

func (self *defaultLoggerImpl) Logf(level LogLevel, format string, v ...interface{}) error {
	if self.IsClosed() {
		return nil
	}
	if !self.level.Contains(level) {
		return ErrInvalidLevel
	}

	for _, l := range self.handlers {
		l.Logf(level, format, v...)
	}
	return nil
}

func (self *defaultLoggerImpl) Alertf(format string, v ...interface{}) {
	self.Logf(LevelAlert, format, v...)
}

func (self *defaultLoggerImpl) Critf(format string, v ...interface{}) {
	self.Logf(LevelCrit, format, v...)

}

func (self *defaultLoggerImpl) Debugf(format string, v ...interface{}) {
	self.Logf(LevelDebug, format, v...)

}

func (self *defaultLoggerImpl) Errf(format string, v ...interface{}) {
	self.Logf(LevelErr, format, v...)

}

func (self *defaultLoggerImpl) Infof(format string, v ...interface{}) {
	self.Logf(LevelInfo, format, v...)

}

func (self *defaultLoggerImpl) Noticef(format string, v ...interface{}) {
	self.Logf(LevelNotice, format, v...)

}

func (self *defaultLoggerImpl) Warningf(format string, v ...interface{}) {
	self.Logf(LevelWarning, format, v...)
}

func (self *defaultLoggerImpl) Flush() error {
	var errmsg string
	for _, l := range self.handlers {
		err := l.Flush()
		if err != nil {
			errmsg = fmt.Sprintf("%s  ,  err(%+)", errmsg, err)
		}
	}
	if len(errmsg) == 0 {
		return nil
	}

	return errors.New(errmsg)
}

func (self *defaultLoggerImpl) Close() error {
	atomic.CompareAndSwapInt32(&self.closed, 0, 1)
	if self.IsClosed() {
		return nil
	}
	var errmsg string
	for _, l := range self.handlers {
		err := l.Close()
		if err != nil {
			errmsg = fmt.Sprintf("%s  ,  err(%+)", errmsg, err)
		}
	}
	if len(errmsg) == 0 {
		return nil
	}

	return errors.New(errmsg)
}

func (self *defaultLoggerImpl) IsClosed() bool {
	return atomic.LoadInt32(&self.closed) == 1
}

func (self *defaultLoggerImpl) SetFileHeader(fh FileHeaderFunc) {
	for _, l := range self.handlers {
		l.SetFileHeader(fh)
	}
}
