package util

import "testing"

func TestPrometheusFixLabel(t *testing.T) {
	strs := map[string]string{
		"abc":  "abc",
		"a.bc": "a_bc",
		"*abc": "_abc",
		"0abc": "_abc",
	}

	for k, v := range strs {
		if v != PrometheusFixLabel(k) {
			t.Error("illegal names")
		}
	}

}
