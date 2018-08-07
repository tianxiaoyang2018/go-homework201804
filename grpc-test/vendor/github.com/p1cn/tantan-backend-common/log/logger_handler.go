package log

// LoggerHanlder
type LoggerHandler interface {
	Logf(level LogLevel, format string, v ...interface{}) error
	Clone() (LoggerHandler, error)
	Close() error
	Flush() error
	SetFileHeader(fh FileHeaderFunc)
}
