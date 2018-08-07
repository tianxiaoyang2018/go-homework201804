package quit

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrClosed = errors.New("already closed")
)

type Manager struct {
	sync.WaitGroup
	closeFlag int32
}

func (m *Manager) Close() {
	atomic.CompareAndSwapInt32(&m.closeFlag, 0, 1)
}

func (m *Manager) Closed() bool {
	return atomic.LoadInt32(&m.closeFlag) == 1
}
