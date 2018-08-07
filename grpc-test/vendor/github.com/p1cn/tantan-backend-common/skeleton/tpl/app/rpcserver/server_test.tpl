{{.CopyRight}}
// package rpcserver
package rpcserver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	domain_{{.ConstServiceName}} "github.com/p1cn/tantan-domain-schema/golang/{{.ConstServiceName}}"

    {{range .Models}}"github.com/p1cn/{{$.RepoName}}/app/{{$.AppName}}/model/{{.PackageName}}"{{end}}
	"github.com/p1cn/{{.RepoName}}/app/{{.AppName}}/model/demo"
	
	"github.com/p1cn/tantan-domain-schema/test"
)

func TestServiceFindDemoById(t *testing.T) {

	mDemoModel := &demo.MockDemoModel{}

	ctx := context.Background()
	retDemo := []*domain_{{.ConstServiceName}}.Demo{}
	mModel.On("FindDemoById", ctx, "1").Return(retDemo[0], nil)

	s := demoServer{
		demoModel: mModel,
	}

	c, err := s.FindDemoById(ctx, &domain_{{.ConstServiceName}}.FindDemoByIdRequest{
		Params: &domain_{{.ConstServiceName}}.FindDemoByIdParams{
			ID: "1",
		},
	})
	assert.Equal(t, err, nil, err)
	assert.Equal(t, c, &domain_{{.ConstServiceName}}.DemosReply{}, "wrong reply")

	mModel.AssertExpectations(t)
}

func TestHealthCheck(t *testing.T) {

	s := {{.RPC.RpcServerClassName}}{}
	err := s.HealthCheck()
	assert.Equal(t, nil, err, "wrong error")
}

