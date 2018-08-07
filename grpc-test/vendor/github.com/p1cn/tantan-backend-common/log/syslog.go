package log

import (
	"errors"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path"
	"strings"
	"time"
)

const maxLogMsgSize = 18000

type SyslogConfig struct {
	Protocol      string
	Address       string
	Facility      string
	Tag           string
	NewLineEscape string
	Level         string

	Priority syslog.Priority `json:"-"`
}

type sysLogger struct {
	writer      syslogWriter
	conf        SyslogConfig
	flags       LFlag
	level       LogLevel
	levelWriter map[LogLevel]func(m string) error
	fh          FileHeaderFunc
}

func newSysLogger(cfg SyslogConfig, flags LFlag) (LoggerHandler, error) {
	cfg2, err := adjustConfig(cfg)
	if err != nil {
		return nil, err
	}

	level := parseLevel(cfg2.Level)
	if !level.IsValid() {
		return nil, errors.New("invalid level configuration")
	}

	syslogWriter, err := syslog.Dial(cfg2.Protocol, cfg2.Address, cfg2.Priority, cfg2.Tag)
	if err != nil {
		return nil, err
	}

	lw := map[LogLevel]func(m string) error{
		LevelAlert:   syslogWriter.Alert,
		LevelCrit:    syslogWriter.Crit,
		LevelDebug:   syslogWriter.Debug,
		LevelErr:     syslogWriter.Err,
		LevelInfo:    syslogWriter.Info,
		LevelNotice:  syslogWriter.Notice,
		LevelWarning: syslogWriter.Warning,
	}

	return &sysLogger{
		writer:      syslogWriter,
		conf:        cfg2,
		flags:       flags,
		level:       level,
		levelWriter: lw,
		fh:          fileHeader(5),
	}, nil
}

func (self *sysLogger) Clone() (LoggerHandler, error) {
	syslogWriter, err := syslog.Dial(self.conf.Protocol, self.conf.Address, self.conf.Priority, self.conf.Tag)
	if err != nil {
		return nil, err
	}

	return &sysLogger{
		writer:      syslogWriter,
		conf:        self.conf,
		flags:       self.flags,
		level:       self.level,
		levelWriter: self.levelWriter,
		fh:          fileHeader(5),
	}, nil
}

func (self *sysLogger) SetFileHeader(fh FileHeaderFunc) {
	self.fh = fh
}

func (self *sysLogger) Close() error {
	return self.writer.Close()
}

func (self *sysLogger) Flush() error {
	return nil
}

func adjustConfig(cfg SyslogConfig) (SyslogConfig, error) {

	m := map[string]syslog.Priority{
		"kern":     syslog.LOG_KERN,
		"user":     syslog.LOG_USER,
		"mail":     syslog.LOG_MAIL,
		"daemon":   syslog.LOG_DAEMON,
		"auth":     syslog.LOG_AUTH,
		"syslog":   syslog.LOG_SYSLOG,
		"lpr":      syslog.LOG_LPR,
		"news":     syslog.LOG_NEWS,
		"uucp":     syslog.LOG_UUCP,
		"authpriv": syslog.LOG_AUTHPRIV,
		"ftp":      syslog.LOG_FTP,
		"cron":     syslog.LOG_CRON,
		"local0":   syslog.LOG_LOCAL0,
		"local1":   syslog.LOG_LOCAL1,
		"local2":   syslog.LOG_LOCAL2,
		"local3":   syslog.LOG_LOCAL3,
		"local4":   syslog.LOG_LOCAL4,
		"local5":   syslog.LOG_LOCAL5,
		"local6":   syslog.LOG_LOCAL6,
		"local7":   syslog.LOG_LOCAL7,
	}

	if cfg.Facility == "" {
		cfg.Facility = "local5"
	}
	if cfg.NewLineEscape == "" {
		cfg.NewLineEscape = "\\n"
	}

	if cfg.Tag == "" {
		cfg.Tag = path.Base(os.Args[0])
	}

	priority, ok := m[cfg.Facility]
	if !ok {
		return cfg, fmt.Errorf("Unsupported facility: %v", cfg.Facility)
	}
	priority |= syslog.LOG_DEBUG

	cfg.Priority = priority
	return cfg, nil
}

func (self *sysLogger) Logf(level LogLevel, format string, v ...interface{}) error {
	if self.writer == nil {
		return nil
	}

	if !self.level.Contains(level) {
		return ErrInvalidLevel
	}

	f, ok := self.levelWriter[level]
	if ok {
		self.logWithFunc(level, f, format, v...)
	}
	return nil
}

func (self *sysLogger) logWithFunc(level LogLevel, f func(m string) error, format string, v ...interface{}) {
	if self.writer == nil {
		return
	}

	var m string

	if self.flags&LFlagDate > 0 {
		m += time.Now().Format(LogTimeFormat)
	}

	if self.flags&LFlagLevel > 0 {
		m += fmt.Sprintf(" %v ", strings.ToUpper(levelMap[level]))
	}

	if self.flags&LFlagFile > 0 {
		fh := self.fh()
		m += fh
	}

	if len(m) > 0 && m[len(m)-1] != ' ' {
		m += " "
	}

	m = m + strings.Replace(fmt.Sprintf(format, v...), "\n", self.conf.NewLineEscape, -1)
	lm := len(m)
	var mlsw string

	if lm > maxLogMsgSize {
		m = m[:maxLogMsgSize] + " This log line is truncated due to bigger than maxLogMsgSize"
		mlsw = fmt.Sprintf("This log line contains %d bytes", lm)
	}

	err := f(m)
	if err != nil {
		log.Printf("ERROR: Log error %v on message %s", err, m)
	}

	if mlsw == "" {
		return
	}

	err = self.writer.Warning(mlsw)
	if err != nil {
		log.Printf("ERROR: Log error %v on message %s", err, mlsw)
	}
}

type syslogWriter interface {
	Alert(m string) error
	Crit(m string) error
	Debug(m string) error
	Emerg(m string) error
	Err(m string) error
	Info(m string) error
	Notice(m string) error
	Warning(m string) error
	Close() error
}
