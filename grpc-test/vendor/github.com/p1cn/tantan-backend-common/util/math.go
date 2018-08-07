package util

import (
	"fmt"
	"strconv"
)

func AbsInt64(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}

// truncate float64 into specified precision
func SetFloatPrecision(n float64, precision uint) float64 {
	fmts := fmt.Sprintf("%%.%vf", precision)
	sn := fmt.Sprintf(fmts, n)
	no, _ := strconv.ParseFloat(sn, 64)

	return no
}
