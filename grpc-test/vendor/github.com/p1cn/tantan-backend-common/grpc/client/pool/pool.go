package pool

import (
	"context"
	"errors"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var (
	ErrClosed        = errors.New("grpc pool: client pool is closed")
	ErrTimeout       = errors.New("grpc pool: client pool timed out")
	ErrAlreadyClosed = errors.New("grpc pool: the connection was already closed")
	ErrFullPool      = errors.New("grpc pool: closing a ClientConn into a full pool")
)

// 创建一个grpc client connection
type Factory func() (*grpc.ClientConn, error)

// grpc 连接池
type Pool struct {
	clients     chan ClientConn
	factory     Factory
	idleTimeout time.Duration
	mu          sync.RWMutex
}

type ClientConn struct {
	*grpc.ClientConn
	pool     *Pool
	timeUsed time.Time
}

// 初始化grpc 连接池
// init ： 初始化时候连接数量
// capacity ： 连接容量
//
func New(factory Factory, init, capacity int, idleTimeout time.Duration) (*Pool, error) {
	if capacity <= 0 {
		capacity = 1
	}
	if init < 0 {
		init = 0
	}
	if init > capacity {
		init = capacity
	}
	p := &Pool{
		clients:     make(chan ClientConn, capacity),
		factory:     factory,
		idleTimeout: idleTimeout,
	}
	for i := 0; i < init; i++ {
		c, err := factory()
		if err != nil {
			return nil, err
		}

		p.clients <- ClientConn{
			ClientConn: c,
			pool:       p,
			timeUsed:   time.Now(),
		}
	}
	// 填充空连接
	for i := 0; i < capacity-init; i++ {
		p.clients <- ClientConn{
			pool: p,
		}
	}
	return p, nil
}

func (p *Pool) getClients() chan ClientConn {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.clients
}

// 清空连接池且关闭所有连接
func (p *Pool) Close() {
	p.mu.Lock()
	clients := p.clients
	p.clients = nil
	p.mu.Unlock()

	if clients == nil {
		return
	}

	close(clients)
	for i := 0; i < p.Capacity(); i++ {
		client := <-clients
		if client.ClientConn == nil {
			continue
		}
		client.ClientConn.Close()
	}
}

// IsClosed returns true if the client pool is closed.
func (p *Pool) IsClosed() bool {
	return p == nil || p.getClients() == nil
}

// 获取连接
// 返回client或者超时错误，context设置超时时间
//
func (p *Pool) Get(ctx context.Context) (*ClientConn, error) {
	clients := p.getClients()
	if clients == nil {
		return nil, ErrClosed
	}

	wrapper := ClientConn{
		pool: p,
	}
	select {
	case wrapper = <-clients:
	case <-ctx.Done():
		return nil, ErrTimeout
	}

	idleTimeout := p.idleTimeout
	if wrapper.ClientConn != nil && idleTimeout > 0 &&
		wrapper.timeUsed.Add(idleTimeout).Before(time.Now()) {

		wrapper.ClientConn.Close()
		wrapper.ClientConn = nil
	}

	var err error
	if wrapper.ClientConn == nil {
		wrapper.ClientConn, err = p.factory()
		if err != nil {
			clients <- ClientConn{
				pool: p,
			}
		}
	}

	return &wrapper, err
}

// 将空闲连接放入连接池
func (c *Pool) Put(client *ClientConn) error {
	if client == nil {
		return nil
	}
	if client.ClientConn == nil {
		return ErrAlreadyClosed
	}
	if c.IsClosed() {
		return ErrClosed
	}

	wrapper := ClientConn{
		pool:       c,
		ClientConn: client.ClientConn,
		timeUsed:   time.Now(),
	}

	select {
	case c.clients <- wrapper:
	default:
		return ErrFullPool
	}

	return nil
}

func (p *Pool) Capacity() int {
	if p.IsClosed() {
		return 0
	}
	return cap(p.clients)
}

func (p *Pool) Available() int {
	if p.IsClosed() {
		return 0
	}
	return len(p.clients)
}
