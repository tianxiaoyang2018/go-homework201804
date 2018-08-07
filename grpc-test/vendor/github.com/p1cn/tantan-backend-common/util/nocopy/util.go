package nocopy

import (
	"reflect"
	"unsafe"
)

func StringToBytes(s string) []byte {
	var b []byte
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	h.Len = len(s)
	h.Cap = len(s)
	return b
}

func BytesToString(b []byte) string {
	var str string
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	s := (*reflect.StringHeader)(unsafe.Pointer(&str))
	s.Data = h.Data
	s.Len = h.Len
	return str
}
