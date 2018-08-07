{{.CopyRight}}
// package httpserver
package httpserver

import (
	"github.com/stretchr/testify/mock"
)

type MockHttpServer struct {
	mock.Mock
}

func (self MockHttpServer) Start() error {
	args := self.Called()
	return args.Error(0)
}

func (self MockHttpServer) Stop() error {
	args := self.Called()
	return args.Error(0)
}

func (self MockHttpServer) HealthCheck() error {
	args := self.Called()
	return args.Error(0)
}
