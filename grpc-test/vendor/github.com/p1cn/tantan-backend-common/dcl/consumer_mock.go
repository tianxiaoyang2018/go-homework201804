package dcl

import (
	"github.com/stretchr/testify/mock"
)

type MockConsumer struct {
	mock.Mock
}

// @todo unimplemented
func (self *MockConsumer) AddConsumer(topic string, group string, p Processor, eh ErrorHandler) {
	//args := self.Called(topic, group, p, eh)
	return
}

func (self *MockConsumer) Start() error {
	return nil
}

func (self *MockConsumer) Stop() error {
	return nil
}
