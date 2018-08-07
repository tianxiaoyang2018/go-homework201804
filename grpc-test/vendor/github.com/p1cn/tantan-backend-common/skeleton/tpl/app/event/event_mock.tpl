{{.CopyRight}}
// package event
package event

import (
	"github.com/stretchr/testify/mock"

	"github.com/p1cn/tantan-domain-schema/golang/event"
	"github.com/p1cn/tantan-backend-common/dcl/eventmeta"
	"github.com/p1cn/{{.RepoName}}/app/{{.AppName}}/handler"
)

type MockEventHandler struct {
    mock.Mock
    mConsumer dcl_tools.MockConsumer
    mHandler handler.MockHandler
}

func (self *MockEventHandler) HealthCheck() error {
	return nil
}

func (self *MockEventHandler) Stop() error {
	err := self.mConsumer.Stop()
	if err != nil {
		return err
	}
	return self.mHandler.Close()
}

{{range .Events}}
func (self *MockEventHandler) {{.Handle}}(event *event.Event, meta *metadata.EventMetaData) error {
	return self.mHandler.{{.Handle}}(event)
}
{{end}}

func (self *MockEventHandler) handleError(err error) {
	return nil
}

