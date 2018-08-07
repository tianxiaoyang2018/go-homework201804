package server

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/http/middleware"
	"github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/tracing"
	"github.com/p1cn/tantan-backend-common/version"
)

const (
	defaultGracefulTimeout = 5 * time.Second
)

const (
	ModeDebug   = "debug"
	ModeRelease = "release"
)

type HttpServer interface {
	Start() error
	Stop() error

	HealthCheck() error
}

type HttpConfig struct {
	// 服务名字，不填写则默认用version.ServiceName()
	ServiceName string
	// 服务版本，不填写则默认用version.ServiceVersion()
	ServiceVersion string

	// http服务名字， 监控名字，不能重复； 可能一个服务有两个rest服务，例如：tantan-backend-test服务有两个rest服务(user 和 device)，
	// 这个地方应该配置 user-rest 或者 device-rest
	Name string
	// 监听网络接口和端口 :
	// interface:port
	Listen string

	TLSConfig         *tls.Config
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int

	// 路由接口
	Router Router
	// 额外需要的中间件
	Middlewares []gin.HandlerFunc

	// 状态变更回调
	ConnState func(net.Conn, http.ConnState)

	// 优雅退出的超时时间, 默认5秒
	GracefulTimeout time.Duration

	// 服务模式：
	// ModeDebug ： debug模式
	// ModeRelease : release模式
	Mode string

	// trace 配置，
	// 如果没有设置，则用common里的配置(拆分基础库后应该移除)
	Trace *TraceConfig

	// 日志
	// Logger log.Logger
}

func routeInfoKey(method, handler string) string {
	return method + "_" + handler
}

// 路由接口
// 请确保一个路由对应一个处理函数
// e.g.
//  g.GET("/hello", f1)
//	g.GET("/hello2", f2)
type Router interface {
	Route(*gin.Engine)
}

//@todo
// trace 应该分为是否trace上下文传递和是否收集
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

func NewHttpServer(cfg HttpConfig) (HttpServer, error) {

	switch cfg.Mode {
	case ModeDebug:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

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
			traceLogger := log.GetLogger(log.TraceLoggerName)
			if traceLogger == nil {
				traceLogger = log.GetLogger(log.DefaultLoggerName)
			}

			for _, nn := range commonCfg.Trace.EnableParts {
				switch nn {
				case "tags":
					cfg.Trace.EnableTags = true
				case "logs":
					cfg.Trace.EnableLogs = true
				}
			}
			cc := &tracing.LoggerSpanCollector{
				Format: commonCfg.Trace.Format,
				Logger: traceLogger,
			}
			cfg.Trace.Collector = cc
		}
	}

	engine := gin.New()

	// route info
	routesInfo := make(map[string]gin.RouteInfo)
	handlerRouteFunc := func(method string, handler string) string {
		return routesInfo[routeInfoKey(method, handler)].Path
	}

	mws := []gin.HandlerFunc{
		// recovery panic
		middleware.Recovery(cfg.Name),
	}

	// trace 上下文传递
	mws = append(mws, middleware.RequestContext(version.ServiceName(), cfg.Name))

	// promtheus监控
	// @todo 做成接口
	mws = append(mws, middleware.PrometheusHttpRequest(version.ServiceName(), cfg.Name, handlerRouteFunc))

	// trace 日志
	if cfg.Trace.Enable {
		mws = append(mws, middleware.Logging(cfg.ServiceName, cfg.ServiceVersion, cfg.Name, handlerRouteFunc, cfg.Trace.Collector, cfg.Trace.EnableLogs, cfg.Trace.EnableTags))
	}

	// custom middleware
	engine.Use(append(mws, cfg.Middlewares...)...)

	cfg.Router.Route(engine)
	for _, routeInfo := range engine.Routes() {
		routesInfo[routeInfoKey(routeInfo.Method, routeInfo.Handler)] = routeInfo
	}

	gto := defaultGracefulTimeout
	if cfg.GracefulTimeout != 0 {
		gto = cfg.GracefulTimeout
	}

	log.Info("new http server successfully : listen(%s)", cfg.Listen)

	return &httpServer{
		srv: &http.Server{
			Addr:    cfg.Listen,
			Handler: engine,

			TLSConfig:         cfg.TLSConfig,
			ReadTimeout:       cfg.ReadTimeout,
			ReadHeaderTimeout: cfg.ReadHeaderTimeout,
			WriteTimeout:      cfg.WriteTimeout,
			IdleTimeout:       cfg.IdleTimeout,
			MaxHeaderBytes:    cfg.MaxHeaderBytes,

			ConnState: cfg.ConnState,
		},
		gracefulTimeout: gto,
	}, nil
}

type httpServer struct {
	srv             *http.Server
	gracefulTimeout time.Duration
}

func (s httpServer) Start() error {
	log.Info("http server starting...")
	return s.srv.ListenAndServe()
}

func (s *httpServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.gracefulTimeout)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Err("http server stop failed : err(%+v)", err)
		return err
	}
	log.Info("http server stopped")
	return nil
}

func (s *httpServer) HealthCheck() error {
	return nil
}
