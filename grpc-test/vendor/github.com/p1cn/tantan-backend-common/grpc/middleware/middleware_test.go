package middleware

import (
	"context"
	"errors"
	"testing"

	"github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-domain-schema/golang/service"
	"google.golang.org/grpc"
)

func TestRequestLogUnaryInterceptorMiddleware(t *testing.T) {
	log.Init(log.Config{
		Output: []string{"stderr", "syslog"},
		Syslog: log.SyslogConfig{
			Facility: "local5",
			Address:  "localhost:514",
			Protocol: "udp",
		},
		Flags: log.LFlagDate | log.LFlagFile,
	})
	inter := RequestLogUnaryInterceptorMiddleware("testService", "testv.1.0.0")

	inter(context.Background(), &mockRpcRequest{}, &grpc.UnaryServerInfo{
		FullMethod: "testFunc",
	}, testhandlerWithErr)
}

type mockRpcRequest struct {
}

func (self *mockRpcRequest) GetContext() *service.Context {

	return &service.Context{
		TraceID:      "traceID",
		ParentSpanID: "parentSpanID",
		SpanID:       "spanID",
		DeviceId:     "deviceID",
		Service:      "testService2",
	}
}

func testhandlerWithErr(ctx context.Context, req interface{}) (interface{}, error) {
	return "resp", errors.New("error")
}
func testhandler(ctx context.Context, req interface{}) (interface{}, error) {
	return "resp", nil
}
