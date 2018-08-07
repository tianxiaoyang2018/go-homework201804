{{.CopyRight}}
// package handler
package handler


import (
	"github.com/stretchr/testify/mock"

	dcl "github.com/p1cn/tantan-domain-schema/golang/event"
	"github.com/p1cn/tantan-backend-common/dcl/eventmeta"

)


type MockEventHandler struct {
    mock.Mock
}

{{range .Events}}
func (self *MockEventHandler) {{.Handle}}(event *event.Event, meta *metadata.EventMetaData) error {
	args := self.Called(event, meta)
    return args.Error(0)
}
{{end}}

func (self *MockEventHandler) HealthCheck() error {
	return nil
}

func (self *MockEventHandler) Close() error {
	return nil
}
