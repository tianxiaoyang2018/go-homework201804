package log

import "fmt"

// 初始化默认日志
func Init(conf Config) error {
	ll, err := newDefaultLoggerImpl(conf)
	if err != nil {
		return err
	}

	globalLoggerM[DefaultLoggerName] = ll
	defaultLogger = ll

	return nil
}

// Flush所有日志
func Flush() error {
	for k, v := range globalLoggerM {
		if k == DefaultLoggerName {
			continue
		}
		err := v.Flush()
		if err != nil {
			Crit("Flush Logger error : err(%+v)", err)
		}
	}
	if defaultLogger != nil {
		return defaultLogger.Flush()
	}
	return nil
}

// 关闭所有日志
func Close() error {

	for k, v := range globalLoggerM {
		if k == DefaultLoggerName {
			continue
		}
		err := v.Close()
		if err != nil {
			Crit("Close Logger error : err(%+v)", err)
		}
	}
	if defaultLogger != nil {
		return defaultLogger.Close()
	}
	return nil
}

// panic 打印的日志，
func Panic(format string, v ...interface{}) {
	Crit(format, v...)
	Flush()
	panic(fmt.Sprintf(format, v...))
}

// 重大错误报警用
func Fatal(format string, v ...interface{}) {
	Alert(format, v...)
	Flush()
	panic(fmt.Sprintf(format, v...))
}

// Legacy functions
// 历史遗留
func Alert(format string, v ...interface{}) {
	defaultLogger.Logf(LevelAlert, format, v...)
	Flush()
}

func Crit(format string, v ...interface{}) {
	defaultLogger.Logf(LevelCrit, format, v...)

}

func Debug(format string, v ...interface{}) {
	defaultLogger.Logf(LevelDebug, format, v...)

}

func Err(format string, v ...interface{}) {
	defaultLogger.Logf(LevelErr, format, v...)

}

func Info(format string, v ...interface{}) {
	defaultLogger.Logf(LevelInfo, format, v...)

}

func Notice(format string, v ...interface{}) {
	defaultLogger.Logf(LevelNotice, format, v...)
}

func Warning(format string, v ...interface{}) {
	defaultLogger.Logf(LevelWarning, format, v...)
}

func GetLevel() LogLevel {
	return defaultLogger.GetLevel()
}
