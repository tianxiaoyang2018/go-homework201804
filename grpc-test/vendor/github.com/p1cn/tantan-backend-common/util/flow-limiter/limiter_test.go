package flowlimiter

import (
	"math"
	"testing"
	"time"
)

func TestRatioLimit(t *testing.T) {
	steps := uint(10)
	duration := time.Hour * 24

	interval := duration / time.Duration(steps)

	res := make([][2]int, steps+1)

	for j := 0; j <= int(steps); j++ {
		for i := 0; i < 10000; i++ {
			limit := RatioLimit(steps, time.Now().UnixNano(), duration, time.Now().Add(time.Duration(-j)*interval+time.Second))
			if limit {
				res[j][0]++
			} else {
				res[j][1]++
			}
		}
	}

	expected := make([]float64, steps+1)
	for i := 0; i <= int(steps); i++ {
		expected[i] = 1 - float64(i)/float64(steps)
	}
	for i, r := range res {
		if diff := float64(r[0])/float64(r[0]+r[1]) - expected[i]; math.Abs(diff) > 0.01 {
			t.Errorf("deviation is too large")
		}
	}
}
