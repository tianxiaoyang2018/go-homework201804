package contextlog

import (
	"errors"
	"testing"

	"github.com/p1cn/tantan-backend-common/log"
)

func TestNewLog(t *testing.T) {

	log.Init(log.Config{
		Output: []string{"stderr", "syslog"},
		Flags:  []string{"level", "file", "date"},
	})

	clog := NewLogWithType("", "123")
	clog.SetDebug("debug")
	clog.SetError(errors.New("error"))
	clog.SetInfo(errors.New("info"))
	clog.SetAlert("alert")
	clog.SetMessage("mesg")

	log.Err("%s", clog.ToJson())
}
