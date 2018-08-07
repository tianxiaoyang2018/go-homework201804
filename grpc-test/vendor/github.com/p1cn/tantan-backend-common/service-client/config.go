package client

import (
	"google.golang.org/grpc"

	"github.com/p1cn/tantan-backend-common/config"
)

type Config struct {
	ServiceName string
	Peer        config.PeerService
	Middleware  []grpc.UnaryClientInterceptor
}
