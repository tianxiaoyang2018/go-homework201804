package dcl

import (
	"github.com/p1cn/tantan-domain-schema/golang/event"
	"github.com/stretchr/testify/mock"
)

type MockProducer struct {
	mock.Mock
}

func (m *MockProducer) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockProducer) Commit(event *event.Event) error {
	args := m.Called(event)
	return args.Error(0)
}
