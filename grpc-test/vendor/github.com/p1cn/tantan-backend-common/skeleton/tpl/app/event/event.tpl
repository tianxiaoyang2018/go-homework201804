{{.CopyRight}}
// package event
package event

import (
	"time"

	"github.com/p1cn/tantan-domain-schema/golang/event"

	slog "github.com/p1cn/tantan-backend-common/log"
	common "github.com/p1cn/tantan-backend-common/config"
    "github.com/p1cn/tantan-backend-common/dcl/eventmeta"

	"github.com/p1cn/tantan-backend-common/dcl"
	"github.com/p1cn/{{.RepoName}}/app/{{.AppName}}/config"
	"github.com/p1cn/{{.RepoName}}/app/{{.AppName}}/handler"
)

const (
	{{range .Events}}{{.GroupVar}} = "{{.GroupName}}"
	{{end}}
)

type EventHandler interface {
	Start() error
	Stop() error
	HealthCheck() error
}

func NewEventHandler(dclh handler.EventHandler) (EventHandler, error) {
	consumer, err := dcl_tools.NewConsumer(common.ServiceNamePush,
		config.Get().DCL, config.Get().MQ)
	if err != nil {
		return nil, err
	}

	eh := &eventHandler{Consumer: consumer, mHandler: dclh}
    {{range .Events}}
	eh.AddConsumer("{{.Topic}}", {{.GroupVar}}, eh.{{.Handle}}, eh.handleError){{end}}
	return eh, nil
}

type eventHandler struct {
	dcl_tools.Consumer
	mHandler  handler.EventHandler
	err     error
	errTime time.Time
}

func (self *eventHandler) HealthCheck() error {
	return self.mHandler.HealthCheck()
}

func (self *eventHandler) Stop() error {
	err := self.Stop()
	if err != nil {
		return err
	}
	return self.mHandler.Close()
}

{{range .Events}}
func (self *eventHandler) {{.Handle}}(event *event.Event, meta *metadata.EventMetaData) error {
	return self.mHandler.{{.Handle}}(event, meta)
}
{{end}}

func (self *eventHandler) handleError(err error) {
	self.err = err
	self.errTime = time.Now()
	slog.Err("%v", self.err)
}
