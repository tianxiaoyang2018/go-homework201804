{{.CopyRight}}
// package handler
package handler

import (
	"time"

	"github.com/p1cn/tantan-domain-schema/golang/event"


	slog "github.com/p1cn/tantan-backend-common/log"
	common "github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/dcl/eventmeta"
)

type EventHandler interface {
    {{range .Events}}
    {{.Handle}}(event *event.Event, meta *metadata.EventMetaData) error
    {{end}}
	HealthCheck() error
	Close() error
}

type EventHandlerImpl struct {

}

func NewEventHandlerImpl() (EventHandler, error) {
    return &EventHandlerImpl{}, nil
}

{{range .Events}}
func (self *EventHandlerImpl) {{.Handle}}(event *event.Event, meta *metadata.EventMetaData) error {
	return nil
}
{{end}}

func (self *EventHandlerImpl) HealthCheck() error {
	return nil
}

func (self *EventHandlerImpl) Close() error {
	return nil
}
