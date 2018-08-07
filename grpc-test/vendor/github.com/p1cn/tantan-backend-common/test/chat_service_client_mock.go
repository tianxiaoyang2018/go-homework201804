package test

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	chat "github.com/p1cn/tantan-domain-schema/golang/chat"
)

type MockChatServiceClient struct {
	mock.Mock
}

func (m *MockChatServiceClient) FindChatCountersByUserId(ctx context.Context, in *chat.FindChatCountersByUserIdRequest, opts ...grpc.CallOption) (*chat.ChatCountersReply, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*chat.ChatCountersReply), args.Error(1)
}

func (m *MockChatServiceClient) FindConversationsByOtherIds(ctx context.Context, in *chat.FindConversationsByOtherIdsRequest, opts ...grpc.CallOption) (*chat.ConversationsReply, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*chat.ConversationsReply), args.Error(1)
}

func (m *MockChatServiceClient) FindConversationOtherUserIds(ctx context.Context, in *chat.FindConversationOtherUserIdsRequest, opts ...grpc.CallOption) (*chat.FindConversationOtherUserIdsReply, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*chat.FindConversationOtherUserIdsReply), args.Error(1)
}
