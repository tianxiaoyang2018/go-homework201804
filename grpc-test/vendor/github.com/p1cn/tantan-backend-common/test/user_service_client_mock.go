package test

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	user "github.com/p1cn/tantan-domain-schema/golang/user"
)

type MockUserServiceClient struct {
	mock.Mock
}

func (m *MockUserServiceClient) FindUsersByIds(ctx context.Context, in *user.FindUsersByIdsRequest, opts ...grpc.CallOption) (*user.UsersReply, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*user.UsersReply), args.Error(1)
}
