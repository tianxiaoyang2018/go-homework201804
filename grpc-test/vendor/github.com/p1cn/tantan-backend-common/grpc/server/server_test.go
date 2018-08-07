package server

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/p1cn/tantan-backend-common/test"
)

var defaultError error
var defaultGrpcError error

func init() {
	defaultError = fmt.Errorf("Error")
	defaultGrpcError = grpc.Errorf(codes.Internal, "%s", defaultError.Error())
}

func TestStart(t *testing.T) {
	mGrpcServer := &test.MockGrpcServer{}
	grpcServer := &grpcServer{
		server: mGrpcServer,
	}

	var netListener net.Listener
	mGrpcServer.On("Serve", netListener).Return(defaultError)

	err := grpcServer.Start()

	assert.Equal(t, defaultError, err, err)
	mGrpcServer.AssertExpectations(t)

}

func TestStop(t *testing.T) {
	mGrpcServer := &test.MockGrpcServer{}
	grpcServer := &grpcServer{
		server: mGrpcServer,
	}

	mGrpcServer.On("GracefulStop").Return()

	err := grpcServer.Stop()

	assert.Equal(t, nil, err, err)
	mGrpcServer.AssertExpectations(t)

}

func TestGrpcHealthCheck(t *testing.T) {
	s := grpcServer{}
	err := s.HealthCheck()
	assert.Equal(t, nil, err, "wrong error")
}
