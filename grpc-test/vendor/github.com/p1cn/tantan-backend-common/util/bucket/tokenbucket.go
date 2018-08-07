package bucket

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity int
	freq     int
	tokens   chan struct{}
	mu       sync.RWMutex
}

// NewTokenBucket new a bucket with freq times per seconds within capacity seconds
func NewTokenBucket(capacity, freq int) *TokenBucket {

	tb := &TokenBucket{
		capacity: capacity,
		freq:     freq,
		tokens:   make(chan struct{}, capacity*1000/freq),
	}

	go func() {
		t := time.NewTicker(time.Duration(1000000/freq) * time.Microsecond)
		for range t.C {
			tb.tokens <- struct{}{}
		}
	}()
	return tb
}

func (tb *TokenBucket) Take(n int) {
	if n <= 0 {
		n = 1
	}
	for i := 0; i < n; i++ {
		<-tb.tokens
	}
}

func (tb *TokenBucket) TakeSingle() {
	<-tb.tokens
}
