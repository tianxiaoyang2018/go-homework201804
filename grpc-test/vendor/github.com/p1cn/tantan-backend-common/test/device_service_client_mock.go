package test

import (
	"context"

	domain "github.com/p1cn/tantan-domain-schema/golang/common"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	device "github.com/p1cn/tantan-domain-schema/golang/device"
)

type MockDeviceServiceClient struct {
	mock.Mock
}

func (m *MockDeviceServiceClient) FindDevicesByIds(ctx context.Context, in *device.FindDevicesByIdsRequest, opts ...grpc.CallOption) (*device.DevicesReply, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*device.DevicesReply), args.Error(1)
}
func (m *MockDeviceServiceClient) FindDevicesByUserId(ctx context.Context, in *device.FindDevicesByUserIdRequest, opts ...grpc.CallOption) (*device.DevicesReply, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*device.DevicesReply), args.Error(1)
}
func (m *MockDeviceServiceClient) FindDevicesByUserIds(ctx context.Context, in *device.FindDevicesByUserIdsRequest, opts ...grpc.CallOption) (*device.DevicesReply, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*device.DevicesReply), args.Error(1)
}
func (m *MockDeviceServiceClient) FindDevicesByDeviceIdentifierTokens(ctx context.Context, in *device.FindDevicesByDeviceIdentifierTokensRequest, opts ...grpc.CallOption) (*device.DevicesReply, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*device.DevicesReply), args.Error(1)
}
func (m *MockDeviceServiceClient) FindExistsByUserId(ctx context.Context, in *device.FindExistsByUserIdRequest, opts ...grpc.CallOption) (*domain.BoolValue, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*domain.BoolValue), args.Error(1)
}
func (m *MockDeviceServiceClient) FindDeviceIdentifiersByDeviceIdentifierTokens(ctx context.Context, in *device.FindDeviceIdentifiersByDeviceIdentifierTokensRequest, opts ...grpc.CallOption) (*device.DeviceIdentifiersReply, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*device.DeviceIdentifiersReply), args.Error(1)
}
func (m *MockDeviceServiceClient) InsertDevice(ctx context.Context, in *device.InsertDeviceRequest, opts ...grpc.CallOption) (*device.DevicesReply, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*device.DevicesReply), args.Error(1)
}
func (m *MockDeviceServiceClient) UpdateDevice(ctx context.Context, in *device.UpdateDeviceRequest, opts ...grpc.CallOption) (*device.DevicesReply, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*device.DevicesReply), args.Error(1)
}
func (m *MockDeviceServiceClient) RemoveDevicesByIds(ctx context.Context, in *device.RemoveDevicesByIdsRequest, opts ...grpc.CallOption) (*device.DevicesReply, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*device.DevicesReply), args.Error(1)
}
func (m *MockDeviceServiceClient) RemoveDevicesByUserIds(ctx context.Context, in *device.RemoveDevicesByUserIdsRequest, opts ...grpc.CallOption) (*device.DevicesReply, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*device.DevicesReply), args.Error(1)
}
func (m *MockDeviceServiceClient) InvalidateToken(ctx context.Context, in *device.InvalidateTokenRequest, opts ...grpc.CallOption) (*device.InvalidateTokenReply, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*device.InvalidateTokenReply), args.Error(1)
}
