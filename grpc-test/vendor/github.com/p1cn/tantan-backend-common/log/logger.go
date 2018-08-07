package log

import (
	"errors"
)

const (
	DefaultLoggerName = "common-default-log"
	TraceLoggerName   = "common-trace-log"
)

type Logger interface {
	Close() error
	IsClosed() bool
	Flush() error
	SetFileHeader(fh FileHeaderFunc)
	GetLevel() LogLevel

	Alertf(format string, v ...interface{})
	Critf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Errf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Noticef(format string, v ...interface{})
	Warningf(format string, v ...interface{})
}

var (
	ErrInvalidLevel = errors.New("invalid level")
	ErrLoggerExist  = errors.New("register failed : logger name does exist")
)

var (
	globalLoggerM = map[string]Logger{}
)

// 克隆一个Logger，然后以name为key添加到LoggerManager中
func GetLogger(name string) Logger {
	l, ok := globalLoggerM[name]
	if !ok {
		ll, err := defaultLogger.clone()
		if err != nil {
			Crit("defaultLogger.clone failed : err(%+v)", err)
			return defaultLogger
		}
		globalLoggerM[name] = ll
		l = ll
	}
	return l
}

func NewLogger(conf Config) (Logger, error) {
	return newDefaultLoggerImpl(conf)
}

func RegisterLogger(name string, logger Logger) error {
	if _, ok := globalLoggerM[name]; ok {
		return ErrLoggerExist
	}
	globalLoggerM[name] = logger
	return nil
}

func newDefaultLoggerImpl(conf Config) (*defaultLoggerImpl, error) {
	if len(conf.Output) == 0 {
		return nil, errors.New("output is null")
	}

	logger := new(defaultLoggerImpl)

	flags := parseFlags(conf.Flags)
	level := parseLevel(conf.Level)

	for _, c := range conf.Output {
		switch c {
		case "syslog":
			l, err := newSysLogger(conf.Syslog, flags)
			if err != nil {
				return nil, err
			}
			logger.handlers = append(logger.handlers, l)
		case "stderr":
			l, err := newStdLogger(c, flags)
			if err != nil {
				return nil, err
			}
			logger.handlers = append(logger.handlers, l)
		case "rotatefile":
			if conf.RotateFile != nil {
				tflags := parseFlags(conf.RotateFile.Flags)
				tlevel := parseLevel(conf.RotateFile.Level)
				if tflags == 0 {
					tflags = flags
				}
				if tlevel > level {
					tlevel = level
				}
				ll, err := NewRotateFileHandler(RotateFileHandlerConfig{
					BaseName:   conf.RotateFile.BaseName,
					When:       conf.RotateFile.When,
					Interval:   conf.RotateFile.Interval,
					Flags:      tflags,
					Level:      tlevel,
					Fh:         fileHeader(5),
					BufferSize: conf.RotateFile.BufferSize,
					// 由于一直使用legacy接口，所以
					// @todo 后面去掉，不应该默认加上
					LineBreak: true,
				})
				if err != nil {
					return nil, err
				}
				logger.handlers = append(logger.handlers, ll)
			}
		}
	}

	logger.level = level
	return logger, nil
}
