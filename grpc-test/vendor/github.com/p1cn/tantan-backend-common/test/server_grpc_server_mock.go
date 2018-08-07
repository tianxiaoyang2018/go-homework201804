package test

import "github.com/stretchr/testify/mock"

type MockServerGrpcServer struct {
	mock.Mock
}

func (m *MockServerGrpcServer) Start() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockServerGrpcServer) Stop() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockServerGrpcServer) HealthCheck() error {
	args := m.Called()
	return args.Error(0)
}
