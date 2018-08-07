package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/p1cn/tantan-backend-common/config"
	pb "github.com/p1cn/tantan-backend-common/grpc/example/proto"
	slog "github.com/p1cn/tantan-backend-common/log"
	grpcclient "github.com/p1cn/tantan-backend-common/service-client"
	"github.com/p1cn/tantan-backend-common/util/tracing"
	"github.com/p1cn/tantan-backend-common/version"
)

func rawClient() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	client, err := grpc.Dial("localhost:9090", opts...)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("state : ", client.GetState())
	cc := pb.NewUserClient(client)

	ctx := tracing.NewServiceContextToContext(context.Background())
	tracing.GetServiceContext(ctx).GetExt()["xxx"] = "yyy"

	sctx := tracing.GetServiceContext(ctx)
	slog.Debug("%+v", sctx)

	resp, err := cc.FindUserById(ctx, &pb.FindUserByIdRequest{Id: "1"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}

func main() {

	err := config.InitCommon(
		config.Config{
			Log: slog.Config{
				Output: []string{"stderr", "syslog"},
				Level:  "debug",
				Flags:  []string{"file", "level", "date"},
			},
		},
	)
	if err != nil {
		panic(err)
	}

	client, err := grpcclient.NewConnection(grpcclient.Config{
		ServiceName: "test",
		Peer: config.PeerService{
			Naming: config.PeerServiceNaming{
				Type:   "file",
				Target: "localhost:9090",
			},
			Grpc: &config.GrpcClient{
				Dial: &config.GrpcDialOptions{
					WithBlock:    true,
					WithInsecure: true,
					WithTimeout:  config.Duration(1 * time.Second),
				},
				Call: &config.GrpcCallOptions{
					Timeout: config.Duration(1 * time.Second),
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	version.Init("grpc-demo-client")

	cc := pb.NewUserClient(client)

	ctx := tracing.NewServiceContextToContext(context.Background())
	sctx := tracing.GetServiceContext(ctx)
	sctx.GetExt()["xxx"] = "yyy"
	sctx.Sampled = 0.56
	sctx.Debug = true
	sctx.DeviceId = "123"

	time.Sleep(1 * time.Second)
	slog.Info("calling")

	st := time.Now()
	count := 1000000
	for i := 0; i < count; i++ {
		_, err := cc.FindUserById(ctx, &pb.FindUserByIdRequest{Id: "1"})
		if err != nil {
			fmt.Println("error : ", i)
		}
	}
	et := time.Now()
	fmt.Println(float32(et.Sub(st).Nanoseconds()) / float32(count))
	fmt.Println(float32(count) / float32(et.Sub(st).Seconds()))

}
