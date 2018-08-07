package util

import (
	"math"
	"testing"
)

var (
	posInf = math.Inf(1)
	negInf = math.Inf(-1)
)

var setFloatPrecisionTests = []struct {
	in        float64
	precision uint
	out       float64
}{
	//positive
	{1.123456, 4, 1.1235},
	{1.1234564, 6, 1.123456},
	{1.123456, 6, 1.123456},
	{1.123456, 10, 1.123456},
	{1.123456, 0, 1.0},
	//negative
	{-1.123456, 4, -1.1235},
	{-1.1234564, 6, -1.123456},
	{-1.123456, 6, -1.123456},
	{-1.123456, 10, -1.123456},
	{-1.123456, 0, -1.0},
	// float infinites
	{posInf, 4, posInf},
	{negInf, 4, negInf},
}

func TestSetFloatPrecision(t *testing.T) {
	for _, tt := range setFloatPrecisionTests {
		s := SetFloatPrecision(tt.in, tt.precision)
		if s != tt.out {
			t.Errorf("SetFloatPrecision(%v, %v) = %v want %v", tt.in, tt.precision, s, tt.out)
		}
	}
}
