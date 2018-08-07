package resolver

import (
	"google.golang.org/grpc/naming"
)

type FakeResolver struct {
	addresses []string
}

func NewFakeResolver(addresses []string) (naming.Resolver, error) {
	return &FakeResolver{addresses: addresses}, nil
}

func (self *FakeResolver) Resolve(target string) (naming.Watcher, error) {
	return newFakeWatcher(self.addresses)
}

func newFakeWatcher(addresses []string) (naming.Watcher, error) {
	return &fakeWatcher{
		addresses: addresses,
	}, nil
}

type fakeWatcher struct {
	addresses []string
	init      bool
}

func (self *fakeWatcher) Close() {
}

func (self *fakeWatcher) Next() ([]*naming.Update, error) {
	if self.init {
		dl := make(chan struct{})
		<-dl
	}

	var updates []*naming.Update
	for _, addr := range self.addresses {
		updates = append(updates, &naming.Update{
			Op:   naming.Add,
			Addr: addr,
		})
	}

	self.init = true
	return updates, nil
}
