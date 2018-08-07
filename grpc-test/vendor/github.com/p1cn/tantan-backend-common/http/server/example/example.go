package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/p1cn/tantan-backend-common/runner"
	"github.com/p1cn/tantan-backend-common/util/tracing"
	"github.com/p1cn/tantan-backend-common/version"

	"github.com/gin-gonic/gin"
	"github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/health"
	"github.com/p1cn/tantan-backend-common/http/server"
)

func NewServer() *Server {
	return &Server{}
}

type Server struct{}

func (s *Server) panic(c *gin.Context) {
	panic("panic")
}

func testServiceContext(c *gin.Context) {

	ctx := c.Request.Context()
	sctx := tracing.GetServiceContext(ctx)

	data, _ := json.Marshal(sctx)

	c.String(200, string(data))
}

func testAsyncContext(c *gin.Context) {

	go func() {
		pctx := tracing.PropagateContextWithServiceContext(c.Request.Context())
		select {
		case <-pctx.Done():
			fmt.Println("done")
		case <-time.After(2 * time.Second):
			fmt.Println("timeout")
		}
	}()

	reader, err := c.Request.GetBody()
	if err == nil {
		defer reader.Close()
		bb := bytes.NewBuffer([]byte{})
		fmt.Println(bb.String())
	}
}

func hello(c *gin.Context) {
	c.Writer.Write([]byte("hello"))
}

func usersFunc(c *gin.Context) {

	// sctx := tracing.GetServiceContext(c.Request.Context())
	// fmt.Printf("%+v\n", sctx)
	c.Writer.Write([]byte("hello"))
}

func (s *Server) Route(g *gin.Engine) {

	f1 := func(c *gin.Context) {
		wrapper(hello)(c)
	}

	f2 := func(c *gin.Context) {
		wrapper(hello)(c)
	}

	g.GET("/hello", f1)
	g.GET("/hello2", f2)
	g.GET("/panic", s.panic)
	gg := g.Group("/users/:uid")
	gg.GET("", usersFunc)
	gg.GET("/links/:id", hello)
}

func wrapper(ff gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ff(c)
	}
}

func middleware1(c *gin.Context) {
	fmt.Println("middleware 1")
	c.Next()
}

func middleware2(c *gin.Context) {
	fmt.Println("middleware 2")
	c.Next()
}

func NewService(name, listenAddress string) (*Service, error) {

	svr, err := server.NewHttpServer(server.HttpConfig{
		Name:        name,
		Listen:      listenAddress,
		Router:      &Server{},
		Mode:        server.ModeDebug,
		Middlewares: []gin.HandlerFunc{},
	})
	if err != nil {
		return nil, err
	}

	return &Service{svr}, nil
}

type Service struct {
	server.HttpServer
}

func (s *Service) GetHealthChecks() []health.HealthCheck {
	return nil
}

var configDir = flag.String("config", "", "")

func main() {
	flag.Parse()

	// err := config.InitCommon(
	// 	config.Config{
	// 		Log: slog.Config{
	// 			Output: []string{"stderr", "syslog"},
	// 			Level:  "debug",
	// 			Flags:  []string{"file", "level", "date"},
	// 		},
	// 		Trace: config.TraceConfig{
	// 			Log: &slog.Config{
	// 				Output: []string{"stderr", "syslog", "rotatefile"},
	// 				Level:  "info",
	// 				Flags:  []string{},
	// 				RotateFile: &slog.RotateFileConfig{
	// 					BaseName: "/tmp/tftest/trace-%s.log",
	// 					When:     "second",
	// 					Interval: 3,
	// 				},
	// 			},
	// 			Enable:      true,
	// 			EnableParts: []string{"logs", "tags"},
	// 			Format:      "text",
	// 		},
	// 	},
	// )
	// if err != nil {
	// 	panic(err)
	// }

	commonCfg := &config.Config{}
	err := config.ParseConfig(*configDir, map[int]interface{}{
		config.CommonConfig: commonCfg,
	})
	if err != nil {
		panic(err)
	}
	err = config.InitCommon(*commonCfg)
	if err != nil {
		panic(err)
	}

	version.Init("http-server-demo")

	service, err := NewService("http1", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	health.Init(":9090")

	runner.RunService(service).Wait()
}
