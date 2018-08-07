package test

import (
	"github.com/stretchr/testify/mock"
)

type MockEventHandler struct {
	mock.Mock
}

func (m *MockEventHandler) Start() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockEventHandler) Stop() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockEventHandler) HealthCheck() error {
	args := m.Called()
	return args.Error(0)
}
