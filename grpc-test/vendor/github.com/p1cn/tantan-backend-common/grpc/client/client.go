package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	common "github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/grpc/client/resolver"
	"github.com/p1cn/tantan-backend-common/grpc/middleware"
	"github.com/p1cn/tantan-backend-common/log"
)

type Client interface {
	Connection() *grpc.ClientConn
	Dial() error
}

type client struct {
	conn        *grpc.ClientConn
	peer        common.PeerService
	serviceName string
	middlwares  []grpc.UnaryClientInterceptor
}

type newOptions struct {
	middlwares []grpc.UnaryClientInterceptor
}

type NewOption func(*newOptions)

func WithMiddlewares(mws []grpc.UnaryClientInterceptor) NewOption {
	return func(o *newOptions) {
		o.middlwares = mws
	}
}

// NewClient : 初始化grpc client
// serviceName : grpc服务的名字
// peer : grpc服务的相关参数
// 后面将废弃，请使用 NewGrpcClient
func NewClient(serviceName string, peer common.PeerService, opts ...NewOption) (Client, error) {

	log.Info("new grpc client successfully.")

	nOpts := &newOptions{}
	for _, opt := range opts {
		opt(nOpts)
	}

	return &client{
		peer:        peer,
		serviceName: serviceName,
		middlwares:  nOpts.middlwares,
	}, nil
}

func (self *client) Connection() *grpc.ClientConn {
	return self.conn
}

func (self *client) Dial() error {

	// balance
	lb, err := balancer(self.serviceName, &self.peer)
	if err != nil {
		return err
	}

	opts := []grpc.DialOption{
		grpc.WithBalancer(lb),
	}

	dialOpts := self.dialConfig()
	callOpts := self.callConfig()

	// middlewares
	mws := []grpc.UnaryClientInterceptor{middleware.ContextClientUnaryInterceptorMiddleware()}
	for _, mw := range self.middlwares {
		mws = append(mws, mw)
	}

	if callOpts.Timeout.Duration() > 0 {
		mws = append(mws, middleware.TimeoutClientUnaryInterceptorMiddleware(callOpts.Timeout.Duration()))
	}

	if len(mws) > 0 {
		opts = append(opts, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(mws...)))
	}

	if dialOpts.WithInsecure {
		opts = append(opts, grpc.WithInsecure())
	}
	if dialOpts.WithBlock {
		opts = append(opts, grpc.WithBlock())
	}

	// timeout is useless while WithBlock is false
	ctx, cancel := context.WithTimeout(context.Background(), dialOpts.WithTimeout.Duration())
	defer cancel()

	conn, err := grpc.DialContext(ctx, self.peer.Naming.Target, opts...)
	if err != nil {
		log.Err("dial grpc server failed : err(%+v)", err)
		return err
	}

	log.Info("dialling grpc server.")

	self.conn = conn
	return nil
}

func balancer(serviceName string, cfg *common.PeerService) (grpc.Balancer, error) {
	re, err := discovery(serviceName, cfg)
	if err != nil {
		log.Err("%+v", err)
		return nil, err
	}

	balancer := grpc.RoundRobin(re)

	if cfg.Balance == nil {
		return balancer, nil
	}

	switch cfg.Balance.Type {
	default:
	}

	return balancer, err
}

func discovery(serviceName string, cfg *common.PeerService) (naming.Resolver, error) {
	if cfg.Naming.Type == common.NamingTypeFile {
		addrs := strings.Split(cfg.Naming.Target, ",")
		if len(addrs) == 0 {
			return nil, common.ErrInvalidConfig
		}
		return resolver.NewFakeResolver(addrs)
	} else if cfg.Naming.Type == common.NamingTypeDns {
		return naming.NewDNSResolver()
	}
	// else if cfg.Naming.Type == "consul" {
	// 	re, err = resolver.NewConsulResolver(serviceName, cfg.Tags)
	// }
	err := fmt.Errorf("service discovery does not support : type(%s)", cfg.Naming.Type)
	log.Err("%+v", err)
	return nil, err
}

func (self *client) dialConfig() common.GrpcDialOptions {
	if self.peer.Grpc != nil && self.peer.Grpc.Dial != nil {
		return *self.peer.Grpc.Dial
	}

	return common.GrpcDialOptions{
		WithBlock:    false,
		WithTimeout:  common.Duration(3 * time.Second),
		WithInsecure: true,
	}
}

func (self *client) callConfig() common.GrpcCallOptions {
	if self.peer.Grpc != nil && self.peer.Grpc.Call != nil {
		return *self.peer.Grpc.Call
	}

	return common.GrpcCallOptions{}
}
