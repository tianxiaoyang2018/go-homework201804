{{.CopyRight}}
// package rpcserver
package rpcserver

import (
	"context"

	"github.com/stretchr/testify/mock"

    service_{{.ConstServiceName}} "github.com/p1cn/tantan-domain-schema/golang/{{.ConstServiceName}}"
)

type Mock{{.RPC.RpcInterface}} struct {
	mock.Mock
}

func (self Mock{{.RPC.RpcInterface}}) FindDemoById(ctx context.Context, in *service_{{.ConstServiceName}}.FindDemoByIdRequest) (*service_{{.ConstServiceName}}.DemosReply, error) {
	args := self.Called(ctx, in)
	return args.Get(0).(*service_{{.ConstServiceName}}.DemosReply), args.Error(1)
}

func (self Mock{{.RPC.RpcInterface}}) HealthCheck() error {
	args := self.Called()
	return args.Error(0)
}
