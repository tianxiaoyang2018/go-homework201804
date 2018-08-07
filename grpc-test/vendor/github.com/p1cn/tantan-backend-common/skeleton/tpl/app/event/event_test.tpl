{{.CopyRight}}
// package event
package event

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/p1cn/tantan-domain-schema/golang/event"
	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/test"
	"github.com/p1cn/tantan-backend-common/dcl/eventmeta"

	"github.com/p1cn/{{.RepoName}}/app/{{.AppName}}/handler"
)

{{range .Events}}
func Test{{.Handle}}(t *testing.T) {
	mh := &handler.MockEventHandler{}

	eh := &eventHandler{mHandler: mh}

	event := &event.Event{}
	meta := &eventmeta.EventMetaData{
		Timestamp: time.Now(),
	}

	mh.On("{{.Handle}}", event).Return(nil)
	err := eh.{{.Handle}}(event, meta)

	assert.Equal(t, nil, err, err)

	mh.AssertExpectations(t)
}
{{end}}

func TestStart(t *testing.T) {
	mc := &test.MockDCLToolsConsumer{}
	mh := &handler.MockEventHandler{}

	eh := &eventHandler{Consumer: mc, mHandler : mh,}

	defaltError := fmt.Errorf("error")
	mc.On("Start").Return(defaltError)

	err := eh.Start()

	assert.Equal(t, defaltError, err, err)

	mc.AssertExpectations(t)
}

func TestStop(t *testing.T) {
	mc := &test.MockDCLToolsConsumer{}
	mh := &handler.MockEventHandler{}

	eh := &eventHandler{Consumer: mc, mHander : mh,}

	defaltError := fmt.Errorf("error")
	mc.On("Stop").Return(defaltError)
	err := eh.Stop()

	assert.Equal(t, defaltError, err, err)

	mc.AssertExpectations(t)
}

func TestHealthCheck(t *testing.T) {
	slog.Init(slog.Config{})

	eh := &eventHandler{}

	err := eh.HealthCheck()
	assert.Equal(t, nil, err, err)

	defaultError := fmt.Errorf("error")
	eh.handleError(defaultError)
	err = eh.HealthCheck()

	assert.Equal(t, defaultError, err, err)
}
