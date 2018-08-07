package util

import (
	"reflect"
	"unicode/utf8"
	"unsafe"
)

// Substr returns substring by bytes length without breaking characters,
// the length of the returned string is as close as possible to the
// `length` argument (less than or equal to)
func Substr(s string, start int, length int) string {
	s = MultiByteSubstr(s, start, length)
	l := len(s)
	if l > length {
		rs := []rune(s)
		rl := len(rs)
		rs2 := rs[0 : rl-1]
		for len(string(rs2)) > length {
			rl--
			rs = rs[0:rl]
			rs2 = rs[0 : rl-1]
			if len(string(rs)) <= length {
				return string(rs)
			}
		}
		return string(rs2)
	}

	return s
}

// MultiByteSubstr returns substring by character length (rune)
func MultiByteSubstr(s string, start int, length int) string {
	rs := []rune(s)
	rl := len(rs)
	if start == 0 && length > rl {
		return s
	}
	if start < 0 {
		start = rl - 1 + start
	}
	end := start + length

	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func TruncateStringOnSize(s string, length int) string {
	if len(s) <= length {
		return s
	}
	var runes []rune
	size := 0
	for _, r := range []rune(s) {
		size += utf8.RuneLen(r)
		if size > length {
			break
		}
		runes = append(runes, r)
	}
	return string(runes)
}

func StringToBytesUnsafe(s string) []byte {
	var b []byte
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	h.Len = len(s)
	h.Cap = len(s)
	return b
}

func BytesToStringUnsafe(b []byte) string {
	var str string
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	s := (*reflect.StringHeader)(unsafe.Pointer(&str))
	s.Data = h.Data
	s.Len = h.Len
	return str
}
