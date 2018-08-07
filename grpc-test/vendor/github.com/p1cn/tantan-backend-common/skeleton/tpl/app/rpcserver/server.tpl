{{.CopyRight}}
// package rpcserver
package rpcserver

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	service_{{.ConstServiceName}} "github.com/p1cn/tantan-domain-schema/golang/{{.ConstServiceName}}"
  
    {{range .Models}}
	model_{{.PackageName}} "github.com/p1cn/{{$.RepoName}}/app/{{$.AppName}}/model/{{.PackageName}}"
	{{end}}
)

type {{.RPC.RpcInterface}} interface {
	service_{{.ConstServiceName}}.{{.RPC.RpcConstName}}ServiceServer
	HealthCheck() error
}

func New{{.RPC.RpcInterface}}({{range $i, $m := .Models}}{{$m.PackageName}} model_{{$m.PackageName}}.Model, {{end}}) {{.RPC.RpcInterface}} {
	return &{{.RPC.RpcServerClassName}}{
		{{range $i, $m := .Models}}
        m{{$m.DbName}}: model_{{$m.PackageName}},
		{{end}}
		demoModel : model_demo.Model,
    }
}

type {{.RPC.RpcServerClassName}} struct {
    {{range .Models}}m{{.DbName}}   model_{{.PackageName}}.Model
	{{end}}
	demoModel model_demo.Model
}

func (self *{{.RPC.RpcServerClassName}}) FindDemoById(ctx context.Context, in *service_{{.ConstServiceName}}.FindDemoByIdRequest) (*service_{{.ConstServiceName}}.DemosReply, error) {
	demos, err := self.demoModel.FindDemoByID(ctx, in.GetParams().GetId())
	if err != nil {
		return nil, grpc.Errorf(codeselfInternal, "%s", err.Error())
	}
	return &service_{{.ConstServiceName}}.DemosReply{
		Demos : demos,
	}, nil
}

func (self *{{.RPC.RpcServerClassName}}) HealthCheck() error {
	return nil
}
