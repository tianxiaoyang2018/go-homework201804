package dislocker

import (
	"errors"
	"time"

	"github.com/hashicorp/consul/api"
)

const (
	consulSessionTTL = "10s"
)

type consulXLocker struct {
	client *api.Client
	lock   *api.Lock
	key    string
}

func newConsulXLocker(key string) (XLocker, error) {
	client, err := api.NewClient(api.DefaultNonPooledConfig())
	if err != nil {
		return nil, err
	}

	return &consulXLocker{
		key:    key,
		client: client,
	}, nil
}

func (ll *consulXLocker) TryLock(waitTime time.Duration) error {

	lock, err := ll.client.LockOpts(&api.LockOptions{
		Key:          ll.key,
		SessionTTL:   consulSessionTTL,
		LockTryOnce:  true,
		LockWaitTime: waitTime,
	})
	if err != nil {
		return ErrConnectionFailed
	}

	ll.lock = lock

	return ll._lock()
}

func (ll *consulXLocker) Lock() error {
	lock, err := ll.client.LockOpts(&api.LockOptions{
		Key:        ll.key,
		SessionTTL: consulSessionTTL,
	})
	if err != nil {
		return ErrConnectionFailed
	}
	if lock == nil {
		return ErrLockFailed
	}

	ll.lock = lock

	return ll._lock()
}

func (ll *consulXLocker) UnLock() error {
	err := ll.lock.Unlock()
	if err != nil {
		return err
	}
	return ll.lock.Destroy()
}

func (ll *consulXLocker) _lock() error {
	stop := make(chan struct{})

	lost, err := ll.lock.Lock(stop)
	if err != nil {
		if err != api.ErrLockHeld {
			return ErrConnectionFailed
		}
		return err
	}
	if lost == nil {
		return errors.New("get lock failed")
	}

	go func() {
		<-lost
		ll.lock.Unlock()
		return
	}()

	return nil
}
