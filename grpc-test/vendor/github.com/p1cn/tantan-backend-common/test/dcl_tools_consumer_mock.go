package test

import (
	"github.com/stretchr/testify/mock"

	dcl_tools "github.com/p1cn/tantan-backend-common/dcl"
)

type MockDCLToolsConsumer struct {
	mock.Mock
}

func (m *MockDCLToolsConsumer) AddConsumer(topic string, group string, p dcl_tools.Processor, eh dcl_tools.ErrorHandler) {
	m.Called(topic, group, p, eh)
}

func (m *MockDCLToolsConsumer) Start() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockDCLToolsConsumer) Stop() error {
	args := m.Called()
	return args.Error(0)
}
