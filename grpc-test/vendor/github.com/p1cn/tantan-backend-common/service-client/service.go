package client

import (
	"google.golang.org/grpc"

	grpcclient "github.com/p1cn/tantan-backend-common/grpc/client"
)

func NewConnection(cfg Config) (*grpc.ClientConn, error) {

	var opts []grpcclient.NewOption
	if len(cfg.Middleware) > 0 {
		opts = append(opts, grpcclient.WithMiddlewares(cfg.Middleware))
	}

	conn, err := grpcclient.NewClient(cfg.ServiceName, cfg.Peer, opts...)
	if err != nil {
		return nil, err
	}

	err = conn.Dial()
	if err != nil {
		return nil, err
	}

	return conn.Connection(), nil
}
