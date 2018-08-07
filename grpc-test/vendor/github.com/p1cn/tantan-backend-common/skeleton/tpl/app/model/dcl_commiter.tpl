{{.CopyRight}}
// package {{.Data.PackageName}}
package {{.Data.PackageName}}


import (
	"golang.org/x/net/context"

	"github.com/p1cn/tantan-backend-common/dcl/eventmeta"
	"github.com/p1cn/tantan-backend-common/dcl"
	"github.com/p1cn/tantan-backend-common/util"
	"github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-domain-schema/golang/event"
)

type dclCommiter struct {
	Model
	dclProducer dcl.Producer
}

func (self *dclCommiter) Close() error {
	err := self.Close()
	if err != nil {
		return err
	}
	return self.dclProducer.Close()
}

func (self *dclCommiter) HealthCheck() error {
	return self.HealthCheck()
}

func (self *dclCommiter) FindDemoByIds(ctx context.Context, ids []string) ([]*domain.Demo, []*domain.Demo, error) {
	new, old, err := self.Model.FindDemoByIds(ctx, userIds)
	self.commitToDCL(ctx, new, old)
	return new, old, err
}

func (self *dclCommiter) commitToDCL(ctx context.Context, newDemos []*domain.Demo, oldDemos []*domain.Demo) {

	for i, newDemo := range newDemos {
		event := &event.Event{
			Topic:   eventmeta.DemoTopic,
			Demo: &event.DemoEvent{
				New: newDemo,
			},
		}
		if len(oldDemos) > i && oldDemos[i] != nil {
			event.Demo.Old = oldDemos[i]
		}
		err := self.dclProducer.Commit(ctx, event)
		if err != nil {
			log.Err("Could not commit demo event %s", event.String())
		}
	}
}
