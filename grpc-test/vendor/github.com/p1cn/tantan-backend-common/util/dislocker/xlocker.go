package dislocker

import (
	"errors"
	"time"
)

const (
	LockTypeXLocker = 1
)

const (
	LockImpTypeConsul = 1
	LockImpTypeEtcd   = 2
	LockImpTypeZk     = 3
	LockImpTypeRedis  = 4
)

var (
	ErrLockFailed       = errors.New("lock failed")
	ErrConnectionFailed = errors.New("connection failed")
	ErrLockNoAcquired   = errors.New("no lock")
)

type XLocker interface {
	TryLock(waitTime time.Duration) error
	Lock() error
	UnLock() error
}

type XLockerConfig struct {
	Type int
	// host schema ... etc...  @TODO
}

func GetDefaultXLockerConfig() *XLockerConfig {
	return &XLockerConfig{
		Type: LockImpTypeConsul,
	}
}

func NewXLocker(key string, config *XLockerConfig) (XLocker, error) {
	if config == nil {
		config = GetDefaultXLockerConfig()
	}
	return newConsulXLocker(key)
}
