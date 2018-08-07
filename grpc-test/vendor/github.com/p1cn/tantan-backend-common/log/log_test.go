package log

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestInfo(t *testing.T) {
	err := Init(Config{
		Output: []string{"stderr", "syslog"},
		Syslog: SyslogConfig{
			Facility: "local6",
			Address:  "127.0.0.1:5140",
			Protocol: "udp",
		},
		Flags: []string{"file", "date", "level"},
		Level: "info",
	})
	if err != nil {
		t.Error(err)
	}

	defaultLogger.level = LevelInfo

	for i := 0; i < 1; i++ {
		Alert("%+v", i)
		//time.Sleep(300 * time.Millisecond)
	}

	Close()

}

func TestTimeRotateFile(t *testing.T) {
	// new logger
	logger, err := NewLogger(Config{
		Output: []string{"rotatefile", "syslog"},
		Syslog: SyslogConfig{
			Facility: "local5",
			Address:  "localhost:514",
			Protocol: "udp",
		},
		Flags: []string{"file", "date", "level"},
		Level: "info",
	})
	if err != nil {
		t.Error(err)
	}

	// register to manager
	err = RegisterLogger("file", logger)
	if err != nil {
		t.Error(err)
	}

	//后面都可以通过这个或者logger
	logger = GetLogger("file")

	logger.Warningf("test : %s", "warning")
	logger.Infof("test : %s", "info")
	logger.Debugf("test : %s", "debug")

	logger.Errf("test : %s", "err")

}

func BenchmarkTimeRotateFile(t *testing.B) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	baseName := fmt.Sprintf("%s/test-%s.log", tempDir, "%s")

	logger, err := NewLogger(Config{
		Output: []string{"rotatefile"},
		RotateFile: &RotateFileConfig{
			BaseName:   baseName,
			When:       "second",
			Interval:   60,
			Flags:      []string{"file", "date", "level"},
			Level:      "info",
			BufferSize: 0,
		},
		Flags: []string{"file", "date", "level"},
		Level: "info",
	})
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < t.N; i++ {
		logger.Warningf("test : %s", "warning")
	}
	err = logger.Close()
	if err != nil {
		t.Error(err)
	}
}

func testInit(t *testing.T) {
	err := Init(Config{
		Output: []string{"stderr", "syslog"},
		Syslog: SyslogConfig{
			Facility: "local5",
			Address:  "localhost:514",
			Protocol: "udp",
		},
		Flags: []string{"file", "date", "level"},
		Level: "info",
	})
	if err != nil {
		t.Error(err)
	}
}

func getSysLogger(t *testing.T) LoggerHandler {
	cfg := SyslogConfig{
		Facility: "local5",
	}
	cfg2, err := adjustConfig(cfg)
	if err != nil {
		t.Fatal(err)
	}

	syslogWriter := &mockSyslogWriter{}
	syslogWriter.On("Info", "test").Return(nil)

	return &sysLogger{
		writer: syslogWriter,
		conf:   cfg2,
	}
}
