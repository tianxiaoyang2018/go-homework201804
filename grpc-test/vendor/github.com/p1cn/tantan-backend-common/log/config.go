package log

// 日志配置
type Config struct {
	Syslog     SyslogConfig
	RotateFile *RotateFileConfig
	Output     []string
	Flags      []string
	Level      string
}

type RotateFileConfig struct {
	BaseName   string
	When       string
	Interval   int
	BufferSize uint16

	Flags []string
	Level string
}

type LFlag int64

// 日志Flag
const (
	// 是否打印文件名
	LFlagFile LFlag = 1 << iota
	// 是否打印日期
	LFlagDate LFlag = 1 << iota
	// 是否打印日志Level
	LFlagLevel LFlag = 1 << iota
)

type LogLevel int

// if configuration doesnot have level definition, LevelAll is default
const (
	LevelOff     LogLevel = iota // "off"
	LevelAlert                   // "alert"
	LevelCrit                    // "crit"
	LevelErr                     // "err"
	LevelWarning                 // "warning"
	LevelNotice                  // "notice"
	LevelInfo                    // "info"
	LevelDebug                   // "debug"
	LevelTrace                   // "trace"
	LevelAll                     // "all"
)

var (
	levelMap = map[LogLevel]string{
		LevelOff:     "off",
		LevelAlert:   "alert",
		LevelCrit:    "crit",
		LevelErr:     "error",
		LevelWarning: "warning",
		LevelNotice:  "notice",
		LevelInfo:    "info",
		LevelDebug:   "debug",
		LevelAll:     "all",
	}
)

func (self LogLevel) IsValid() bool {
	return self >= LevelOff && self <= LevelAll
}

func (self LogLevel) Contains(level LogLevel) bool {
	return self >= level
}
