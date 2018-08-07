package tracing

import (
	"fmt"
	"testing"
)

func TestNewSpanID(t *testing.T) {
	ss := NewSpanID()
	if len(ss) != 16 {
		t.Fatal("error length")
	}
	fmt.Println(ss)
	t.Log(ss)
}

func TestNewTraceID(t *testing.T) {
	ss := NewTraceID()
	if len(ss) != 32 {
		t.Log(ss)
		t.Fatal("error length")
	}
	fmt.Println(ss)
	t.Log(ss)
}
