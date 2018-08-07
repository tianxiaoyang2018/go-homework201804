package util

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	slog "github.com/p1cn/tantan-backend-common/log"
)

var ErrCircuitBreakerOpen = errors.New("circuit breaker is open")

type CircuitBreakerStatus uint8

const (
	CircuitBreakerClosed CircuitBreakerStatus = iota
	CircuitBreakerOpen
	CircuitBreakerHalfOpen
)

func NewCircuitBreaker(threshold int64, retryDur time.Duration, callerInfo fmt.Stringer) *CircuitBreaker {
	return &CircuitBreaker{threshold: threshold, retryDur: retryDur, s: CircuitBreakerClosed, callerInfo: callerInfo}
}

type CircuitBreaker struct {
	failedCount int64
	threshold   int64
	retryDur    time.Duration
	callerInfo  fmt.Stringer

	rwm   sync.RWMutex
	timer *time.Timer
	s     CircuitBreakerStatus
}

func (cb *CircuitBreaker) Run(f func() (err error, skip bool)) error {
	switch cb.Status() {
	case CircuitBreakerOpen:
		if cb.callerInfo == nil {
			return ErrCircuitBreakerOpen
		}
		return fmt.Errorf("circuit breaker is open %s", cb.callerInfo)
	case CircuitBreakerClosed, CircuitBreakerHalfOpen:
		err, skip := f()
		switch {
		case err == nil:
			cb.setStatus(CircuitBreakerClosed)
		case !skip:
			cb.errHappened()
		}
		return err
	}
	return nil
}

func (cb *CircuitBreaker) errHappened() {
	switch cb.Status() {
	case CircuitBreakerClosed:
		if atomic.AddInt64(&cb.failedCount, 1) > cb.threshold {
			cb.setStatus(CircuitBreakerOpen)
		}
	case CircuitBreakerHalfOpen:
		atomic.AddInt64(&cb.failedCount, 1)
		cb.setStatus(CircuitBreakerOpen)
	case CircuitBreakerOpen:
		atomic.AddInt64(&cb.failedCount, 1)
	}
}

func (cb *CircuitBreaker) setStatus(cbs CircuitBreakerStatus) CircuitBreakerStatus {
	cb.rwm.Lock()
	defer cb.rwm.Unlock()
	if cb.s != cbs {
		switch cbs {
		case CircuitBreakerOpen:
			slog.Err("circuit breaker is open %s", cb.callerInfo)
		case CircuitBreakerHalfOpen:
			slog.Info("circuit breaker is half open %s", cb.callerInfo)
		case CircuitBreakerClosed:
			slog.Info("circuit breaker is closed %s", cb.callerInfo)
		}
	}
	cb.s = cbs
	switch cb.s {
	case CircuitBreakerOpen:
		if cb.timer == nil {
			cb.timer = time.AfterFunc(cb.retryDur, func() {
				cb.setStatus(CircuitBreakerHalfOpen)
			})
		} else {
			cb.timer.Reset(cb.retryDur)
		}
	case CircuitBreakerClosed:
		atomic.StoreInt64(&cb.failedCount, 0)
	}
	return cb.s
}

func (cb *CircuitBreaker) Status() CircuitBreakerStatus {
	cb.rwm.RLock()
	defer cb.rwm.RUnlock()
	return cb.s
}
