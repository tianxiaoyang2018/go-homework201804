package test

// import (
// 	"github.com/p1cn/tantan-domain-schema/golang/service"
// 	"github.com/stretchr/testify/mock"
// )

// type clientsMock struct {
// 	mock.Mock
// }

// func (m *clientsMock) InitUserServiceClients(addresses []string) error {
// 	args := m.Called(addresses)
// 	return args.Error(0)
// }
// func (m *clientsMock) InitDeviceServiceClients(addresses []string) error {
// 	args := m.Called(addresses)
// 	return args.Error(0)
// }
// func (m *clientsMock) InitChatServiceClients(addresses []string) error {
// 	args := m.Called(addresses)
// 	return args.Error(0)
// }
// func (m *clientsMock) GetUserServiceClient() (service.UserServiceClient, error) {
// 	args := m.Called()
// 	return args.Get(0).(service.UserServiceClient), args.Error(1)
// }
// func (m *clientsMock) GetDeviceServiceClient() (service.DeviceServiceClient, error) {
// 	args := m.Called()
// 	return args.Get(0).(service.DeviceServiceClient), args.Error(1)
// }
// func (m *clientsMock) GetChatServiceClient() (service.ChatServiceClient, error) {
// 	args := m.Called()
// 	return args.Get(0).(service.ChatServiceClient), args.Error(1)
// }

// type MockedClientsWrapper struct {
// 	mock.Mock

// 	UserCLient   *MockUserServiceClient
// 	DeviceCLient *MockDeviceServiceClient
// 	ChatCLient   *MockChatServiceClient
// }

// func (m *MockedClientsWrapper) InitUserServiceClients(addresses []string) error {
// 	args := m.Called(addresses)
// 	return args.Error(0)
// }
// func (m *MockedClientsWrapper) InitDeviceServiceClients(addresses []string) error {
// 	args := m.Called(addresses)
// 	return args.Error(0)
// }
// func (m *MockedClientsWrapper) InitChatServiceClients(addresses []string) error {
// 	args := m.Called(addresses)
// 	return args.Error(0)
// }
// func (m *MockedClientsWrapper) GetUserServiceClient() (service.UserServiceClient, error) {
// 	return m.UserCLient, nil
// }
// func (m *MockedClientsWrapper) GetDeviceServiceClient() (service.DeviceServiceClient, error) {
// 	return m.DeviceCLient, nil
// }
// func (m *MockedClientsWrapper) GetChatServiceClient() (service.ChatServiceClient, error) {
// 	return m.ChatCLient, nil
// }
