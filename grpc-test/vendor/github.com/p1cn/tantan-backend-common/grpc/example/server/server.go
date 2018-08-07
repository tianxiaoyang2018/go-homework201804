package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/p1cn/tantan-backend-common/health"
	"github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/metrics"

	"github.com/p1cn/tantan-backend-common/config"
	pb "github.com/p1cn/tantan-backend-common/grpc/example/proto"
	"google.golang.org/grpc"

	grpcServer "github.com/p1cn/tantan-backend-common/grpc/server"
)

var metric = flag.Bool("metric", false, "print metrics")

type RpcServer struct {
}

func (self *RpcServer) FindUserById(ctx context.Context, req *pb.FindUserByIdRequest) (*pb.FindUserByIdResponse, error) {
	// md, _ := metadata.FromIncomingContext(ctx)
	// sctx := tracing.GetServiceContext(ctx)
	// log.Debug("%+v", sctx)
	//time.Sleep(1 * time.Second)
	return &pb.FindUserByIdResponse{Name: "testusername"}, nil
}

func TestMiddleware() grpc.UnaryServerInterceptor {
	log.Info("created Test Middleware")
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (r interface{}, err error) {

		r, err = handler(ctx, req)
		//log.Info("test middleware")
		return r, err
	}
}

func server() {

	err := config.InitCommon(
		config.Config{
			Log: log.Config{
				Output: []string{"stderr", "syslog"},
				Level:  "debug",
				Flags:  []string{"file", "level", "date"},
			},
			Trace: config.TraceConfig{
				Log: &log.Config{
					Output: []string{"rotatefile"},
					Level:  "info",
					Flags:  []string{},
					RotateFile: &log.RotateFileConfig{
						BaseName:   "/tmp/tftest/trace-%s.log",
						When:       "minute",
						Interval:   1,
						BufferSize: 1,
					},
				},
				Enable:      true,
				EnableParts: []string{"logs"},
				Format:      "text",
			},
		},
	)
	if err != nil {
		panic(err)
	}

	ggg, err := grpcServer.NewGrpcServer(grpcServer.GrpcConfig{
		Name:        "test-grpc-server",
		ServiceName: "user",
		Grpc: config.Grpc{
			Listen: ":9090",
		},
		Register: func(srv *grpc.Server) {
			pb.RegisterUserServer(srv, &RpcServer{})
		},
		Middlewares: []grpc.UnaryServerInterceptor{TestMiddleware()},
	})
	if err != nil {
		log.Alert("%v", err)
		return
	}
	health.Init(":9091")

	ggg.Start()

}

func main() {
	go signalHandler()
	server()
}

func signalHandler() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGPIPE, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)
	for {
		select {
		case sig := <-sigCh:
			switch sig {
			case syscall.SIGPIPE:
			default:
				if *metric {
					fmt.Println(metrics.GetPromethuesAsFmtText())
				}
				os.Exit(0)
			}
		}
	}
}
