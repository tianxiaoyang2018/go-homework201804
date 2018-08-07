{{.CopyRight}}
// package httpserver
package httpserver

import (
	"github.com/gin-gonic/gin"
	http_server "github.com/p1cn/tantan-backend-common/http/server"
	common "github.com/p1cn/tantan-backend-common/config"
)

func NewHttpServer(listenAddress string) (http_server.HttpServer, error) {
	return http_server.NewHttpServer(http_server.HttpConfig{
		ServiceName: common.ServiceName{{$.AbbServiceName}},
		Listen: listenAddress,
		Router: &server{},
	})
}

type server struct{}

func (s *server) handlerHelloWorld(c *gin.Context) {
	c.String(200, "Hello World")
}

func (s *server) Route(g *gin.Engine) {
	g.GET("hello", s.handlerHelloWorld)
}
