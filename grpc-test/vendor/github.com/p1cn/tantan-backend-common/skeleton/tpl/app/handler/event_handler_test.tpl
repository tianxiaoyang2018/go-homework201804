{{.CopyRight}}
// package handler
package handler

import "testing"

{{range .Events}}
func Test{{.Handle}}(t *testing.T) {
}
{{end}}


