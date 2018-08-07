{{.CopyRight}}
// package service
package service

import (
	"google.golang.org/grpc"

    {{if .Config.RPC}}
	service_{{.ConstServiceName}} "github.com/p1cn/tantan-domain-schema/golang/{{.ConstServiceName}}"
    {{end}}
    {{if .Config.DB}}
	database "github.com/p1cn/tantan-backend-common/db/postgres"
    {{end}}
	"github.com/p1cn/tantan-backend-common/util"
	"github.com/p1cn/{{.RepoName}}/app/{{.AppName}}/config"
    {{range .Models}}
	"github.com/p1cn/{{$.RepoName}}/app/{{$.AppName}}/model/{{.PackageName}}"
	{{end}}
    {{if .Config.RPC}}
	"github.com/p1cn/{{.RepoName}}/app/{{.AppName}}/rpcserver"
    grpc_server "github.com/p1cn/tantan-backend-common/grpc/server"
    {{end}}
	common "github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/health"
	"github.com/p1cn/tantan-backend-common/service"
	dcl_tools "github.com/p1cn/tantan-backend-common/dcl"
	"github.com/p1cn/tantan-backend-common/util"

	{{if .Config.HTTP}}
	http_server "github.com/p1cn/tantan-backend-common/http/server"
	http_server_impl "github.com/p1cn/{{.RepoName}}/app/{{.AppName}}/httpserver"
	{{end}}

	{{if .Config.Dependencies}}
		"github.com/p1cn/{{.RepoName}}/app/{{.AppName}}/rpcclient"
	{{end}}
)

func New{{.ServiceName}}(cfg *config.Config) (service.IService, error) {
	service := {{.ServiceClassName}}{
		graceful: util.NewGraceful(),
	}

    {{if .Config.EventLog}}
	// init eventlog
	err := eventlog.Init(cfg.EventLog, cfg.MQ)
	if err != nil {
		return nil, err
	}
	{{end}}

    {{if .Config.Dependencies}}
    // init dependencies
    {{range .Deps}}
	// initialize  {{.ServiceName}}
	{{.LwConstServiceName}}Client, err := rpcclient.New{{.ConstServiceName}}Client(
		common.ServiceName{{.ConstServiceName}}Grpc, 
		cfg.Dependencies.Services[common.ServiceName{{.ConstServiceName}}Grpc])
	if err != nil {
		return nil, err
	}
	{{end}}
	{{end}}

    {{if .DclCommiter}}
	// new dcl producer
	producer, err := dcl_tools.NewProducer(common.ServiceName, cfg.DCL, cfg.MQ)
	if err != nil {
		return nil, err
	}
	service.producer = producer
	{{end}}

    {{if .Config.DB}}
	// init database

	// set model
    {{range .Models}}
	db{{.DbName}}, err := database.NewDB(database.Config{
		Name : "{{.DbName}}",
		Cluster: cfg.Database.Postgres[config.DbName{{.DbName}}], 
		Graceful:service.graceful),
	if err != nil {
		return nil, err
	}

	service.{{.PackageName}}Model = {{.PackageName}}.NewModel(db{{.DbName}}{{if $.DclCommiter}}, producer{{end}})
	{{end}}
    {{end}}

    {{if .Config.RPC}}
	// new grpc service
	service.server = server.New{{.RPC.RpcInterface}}({{range $i, $m := .Models}}service.{{$m.PackageName}}Model, {{end}})

	// new grpc server and register service
	service.grpcServer, err = grpcserver.NewGrpcServer(grpcserver.GrpcConfig{
		ServiceName: common.ServiceName{{.ConstServiceName}},
		Grpc:        *cfg.RPC.Grpc,
		Register: func(srv *grpc.Server) {
			service_{{.ConstServiceName}}.Register{{.RPC.RpcConstName}}ServiceServer(srv, service.server)
		},
	})
	if err != nil {
		return nil, err
	}
    {{end}}

	{{if .Config.HTTP}}
		service.httpServer, err = http_server_impl.NewHttpServer(cfg.HTTP.Listen)
		if err != nil {
			return nil, err
		}
	{{end}}


    {{if .Config.DCL}}
	// implement interface of event handler
	p, err := handler.NewHandler()
	if err != nil {
		return nil, err
	}

	// new event(DCL) handler
	service.eventHandler, err = event.NewEventHandler(p)
	if err != nil {
		return nil, err
	}
    {{end}}

	return service, nil
}

type {{.ServiceClassName}} struct {
    {{range .Models}}{{.PackageName}}Model	{{.PackageName}}.Model
    {{end}}
	{{if .Config.RPC}}
	server			server.{{.RPC.RpcInterface}}
	grpcServer		grpcserver.GrpcServer
	{{end}}
	{{if .Config.HTTP}}
	httpServer httpserver.HttpServer
	{{end}}
	graceful		util.GracefulMonitor
	{{if .DclCommiter}}producer	dcl_tools.Producer
	{{end}}
	{{if .Config.DCL}}
	eventHandler event.EventHandler
	{{end}}
}

func (self *{{.ServiceClassName}}) Start() (error) {
	starters := []util.Starter{}
	
	{{if .Config.RPC}}
	starters = append(starters, self.grpcServer)
	{{end}}
	{{if .Config.DCL}}
	starters = append(starters, self.eventHandler)
	{{end}}
	{{if .Config.HTTP}}
	starters = append(starters, self.httpServer)
	{{end}}
	return util.RunMultiStarter(starters)
}

func (self *{{.ServiceClassName}}) Stop() (err error) {
	{{if .Config.DCL}}
	err = self.eventHandler.Stop()
	if err != nil {
		return
	}
	{{end}}
	{{if .Config.RPC}}
	err = self.grpcServer.Stop()
	if err != nil { // if fail to close, graceful.Wait might be blocked forever.
		return 
	}
	{{end}}
	{{if .Config.HTTP}}
	err = self.httpServer.Stop()
	if err != nil { // if fail to close, graceful.Wait might be blocked forever.
		return 
	}
	{{end}}
	self.graceful.Wait()
	{{if .DclCommiter}}
	err = self.producer.Close()
	if err != nil {
		return
	}
	{{end}}
	return 
}

func (self *{{.ServiceClassName}}) GetHealthChecks() []health.HealthCheck {
	return []health.HealthCheck{
		{{range .Models}}health.NewHealthCheck(self.{{.PackageName}}Model.HealthCheck, true, "[Model]"),
		{{end}}
		{{if .Config.DCL}}
		health.NewHealthCheck(self.eventHandler.HealthCheck, true, "[event handler]"),
		{{end}}
		{{if .Config.RPC}}
		health.NewHealthCheck(self.server.HealthCheck, true, "[RPC Server]"),
		health.NewHealthCheck(self.grpcServer.HealthCheck, true, "[GRpc Server]"),
		{{end}}
		{{if .Config.HTTP}}
		health.NewHealthCheck(self.httpServer.HealthCheck, true, "[http server]"),
		{{end}}
	}
}


