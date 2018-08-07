package test

import (
	"net"

	"github.com/stretchr/testify/mock"
)

type MockGrpcServer struct {
	mock.Mock
}

func (m *MockGrpcServer) Serve(listener net.Listener) error {
	args := m.Called(listener)
	return args.Error(0)
}

func (m *MockGrpcServer) GracefulStop() {
	m.Called()
	return
}
