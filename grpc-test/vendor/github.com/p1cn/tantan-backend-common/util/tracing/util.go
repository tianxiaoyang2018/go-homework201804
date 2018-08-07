package tracing

import (
	"bytes"
	"fmt"
	"strings"
)

func TracingMapToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=%s,", key, value)
	}
	return b.String()
}

func TracingStringToMap(s string) map[string]string {
	mm := make(map[string]string)
	pairs := strings.Split(s, ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			mm[kv[0]] = kv[1]
		}
	}
	return mm
}
