package util

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var errFT = errors.New("err for test")

func testCB(t *testing.T, f func() (error, bool)) *CircuitBreaker {
	cb := NewCircuitBreaker(10, time.Minute, nil)
	for i := 0; i < 11; i++ {
		assert.EqualError(t, cb.Run(f), errFT.Error())
	}
	return cb
}

func TestCircuitBreakerOpen(t *testing.T) {
	t.Skip()
	f := func() (error, bool) {
		return errFT, false
	}
	cb := testCB(t, f)
	assert.EqualError(t, cb.Run(f), "circuit breaker is open")
}

func TestCircuitBreakerHalfOpenClose(t *testing.T) {
	t.Skip()
	cb := testCB(t, func() (error, bool) {
		return errFT, false
	})
	time.Sleep(time.Minute + time.Second)
	assert.Equal(t, CircuitBreakerHalfOpen, cb.Status())
	cb.Run(func() (error, bool) { return nil, false })
	assert.Equal(t, CircuitBreakerClosed, cb.Status())
}

func TestCircuitBreakerSkip(t *testing.T) {
	t.Skip()
	cb := testCB(t, func() (error, bool) {
		return errFT, true
	})
	assert.Equal(t, CircuitBreakerClosed, cb.Status())
}
