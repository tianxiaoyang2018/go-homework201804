package bucket

import "testing"

func TestTokenBucketSingle(t *testing.T) {
	capacity := 2
	freq := 1
	tb := NewTokenBucket(capacity, freq)
	for i := 0; i < 10; i++ {
		tb.TakeSingle()
	}
}
