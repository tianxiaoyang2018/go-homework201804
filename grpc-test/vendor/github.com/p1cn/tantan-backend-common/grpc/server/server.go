package server

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/grpc/middleware"
	"github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/tracing"
	"github.com/p1cn/tantan-backend-common/version"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

// GrpcServer
// grpc 服务器接口
type GrpcServer interface {
	Start() error
	Stop() error
	HealthCheck() error
}

// GrpcConfig
// grpc服务器的配置
type GrpcConfig struct {
	// 服务名字, 如果没有填写则用 version.ServiceName()
	ServiceName string
	// 服务的版本，如果没有填写则调用 version.ServiceVersion()
	ServiceVersion string

	// grpc server 名字，监控的name，不能重复； 可能一个服务有两个grpc服务，例如：tantan-backend-test服务有两个grpc服务(user 和 device)，
	// 这个地方应该配置 user-grpc 或者 device-grpc
	Name string
	// grpc配置
	Grpc config.Grpc
	// 中间件
	Middlewares []grpc.UnaryServerInterceptor
	// 注册回调
	Register func(*grpc.Server)

	// trace 配置，
	// 如果没有设置，则用common里的配置(拆分基础库后应该移除)
	Trace *TraceConfig

	// 日志
	// Logger log.Logger
}

//
type TraceConfig struct {
	// 开启trace
	Enable bool
	// 收集logs信息
	EnableLogs bool
	// 收集tags信息
	EnableTags bool
	// 收集器：以日志形式收集还是其他
	Collector tracing.SpanCollector
}

// 初始化一个grpc服务器
func NewGrpcServer(cfg GrpcConfig) (GrpcServer, error) {
	if len(cfg.ServiceName) == 0 {
		cfg.ServiceName = version.ServiceName()
	}
	if len(cfg.Name) == 0 {
		cfg.Name = version.ServiceName()
	}
	if len(cfg.ServiceVersion) == 0 {
		cfg.ServiceVersion = version.Version()
	}

	if cfg.Trace == nil {
		commonCfg := config.GetCommonConfig()

		cfg.Trace = &TraceConfig{
			Enable: commonCfg.Trace.Enable,
		}

		if cfg.Trace.Enable {
			for _, nn := range commonCfg.Trace.EnableParts {
				switch nn {
				case "tags":
					cfg.Trace.EnableTags = true
				case "logs":
					cfg.Trace.EnableLogs = true
				}
			}

			traceLogger := log.GetLogger(log.TraceLoggerName)
			if traceLogger == nil {
				traceLogger = log.GetLogger(log.DefaultLoggerName)
			}

			cc := &tracing.LoggerSpanCollector{
				Format: commonCfg.Trace.Format,
				Logger: traceLogger,
			}
			cfg.Trace.Collector = cc
		}
	}

	mws := []grpc.UnaryServerInterceptor{
		middleware.RecoveryUnaryInterceptorMiddleware(cfg.Name),
	}

	mws = append(mws, middleware.RequestContextUnaryInterceptorMiddleware(cfg.ServiceName, cfg.Name))

	mws = append(mws, middleware.PrometheusUnaryInterceptorMiddleware(cfg.Name))

	if cfg.Trace.Enable {
		mws = append(mws, middleware.RequestLogUnaryInterceptorMiddleware(cfg.ServiceName, cfg.ServiceVersion, cfg.Name, cfg.Trace.Collector, cfg.Trace.EnableLogs, cfg.Trace.EnableTags))
	}

	chain := grpc_middleware.ChainUnaryServer(append(mws, cfg.Middlewares...)...)
	ServerOption := grpc.UnaryInterceptor(chain)

	// all ServerOption must be defined in configuration
	srv := grpc.NewServer(ServerOption)
	cfg.Register(srv)
	reflection.Register(srv)

	listener, err := net.Listen("tcp", cfg.Grpc.Listen)
	if err != nil {
		log.Err("new grpc server failed : err(%+v)", err)
		return nil, err
	}
	log.Info("new grpc server successfully.")

	return &grpcServer{listener: listener, server: srv}, nil
}

type grpcServer struct {
	listener net.Listener
	server   interface {
		Serve(net.Listener) error
		GracefulStop()
	}
}

// 开始运行服务
func (s *grpcServer) Start() error {
	log.Info("grpc server starting...")
	return s.server.Serve(s.listener)
}

// 结束运行服务
func (s *grpcServer) Stop() error {
	log.Info("grpc server stopping...")
	s.server.GracefulStop()
	return nil
}

// 健康检查
func (s *grpcServer) HealthCheck() error {
	return nil
}
