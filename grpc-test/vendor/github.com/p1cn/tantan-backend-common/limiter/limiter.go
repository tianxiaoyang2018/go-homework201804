package limiter

import "time"

func NewRequestLimiter(maxConcurrency int, timeout time.Duration) *RequestLimiter {
	l := &RequestLimiter{cc: make(chan struct{}, maxConcurrency), timeout: timeout}
	for i := 0; i < maxConcurrency; i++ {
		l.cc <- struct{}{}
	}
	return l
}

type RequestLimiter struct {
	cc      chan struct{}
	timeout time.Duration
}

func (l *RequestLimiter) Wait() (done func()) {
	select {
	case <-l.cc:
		return func() { l.cc <- struct{}{} }
	case <-time.After(l.timeout):
		return nil
	}
}
