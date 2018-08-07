package util

import (
	"testing"
)

func TestVersionGreaterThan(t *testing.T) {
	vers := map[string]string{
		"1.0.1":   "1.0",
		"1.2.1":   "1.1.0",
		"1.4.1":   "1.3.2",
		"1.4.5":   "1.4.1",
		"1.4.15":  "1.4.5",
		"1.5.0.1": "1.4.1",
		"1.5.0.2": "1.5.0.0",
		"1.5.0.3": "1.4.50.20",
	}
	for v1, v2 := range vers {
		if !VersionGreaterThan(v1, v2) {
			t.Errorf("version %s is great than version %s", v1, v2)
		}
	}
}

func TestVersionVersionGEByComparingFirst3Digits(t *testing.T) {
	type input struct {
		l, r     string
		res, err bool
	}
	inputs := []input{
		{"1.5.4", "1.5.4.1", true, false},
		{"1.5.10", "1.5.9", true, false},
		{"1.5.4.0", "1.5.4.1", true, false},
		{"1.5.4.2", "1.5.4.1", true, false},
		{"1.5.4", "1.5.4", true, false},
		{"1.5.3", "1.5.4", false, false},
		{"1.5.3.4", "1.5.4", false, false},
		{"1.5", "1.5.4", false, true},
		{"1.5.3.4", "1", false, true},
	}

	for _, in := range inputs {
		res, err := VersionGEByComparingFirst3Digits(in.l, in.r)
		if err != nil && !in.err {
			t.Errorf("got err: %v, expect no error", err)
		} else if err == nil && in.err {
			t.Errorf("got err: %v, expect error wrong digit number", err)
		}
		if res != in.res {
			t.Errorf("%+v got res: %v, expect %v", in, res, in.res)
		}
	}
}

func TestVersionGreaterThanOrEqualTo(t *testing.T) {
	vers := map[string]string{
		"1":        "1.0",
		"1.0":      "1.0",
		"1.0.1":    "1.0",
		"1.1.0":    "1.1",
		"1.1.1":    "1.1.0",
		"1.4.1":    "1.3.2",
		"1.4.5":    "1.4.1",
		"1.4.15":   "1.4.5",
		"1.5.0.1":  "1.4.1",
		"1.5.0.01": "1.5.0.1",
		"1.5.0.2":  "1.5.0.0",
		"1.5.0.3":  "1.4.50.20",
	}
	for v1, v2 := range vers {
		if !VersionGreaterThanOrEqualTo(v1, v2) {
			t.Errorf("version %s is great than or equal version %s", v1, v2)
		}
	}
}

func TestVersionCompare(t *testing.T) {
	vers := map[string]string{
		"1.1":      "1.0",
		"1.3":      "1.1.5",
		"1.4":      "1.3",
		"1.5.0":    "1.4.6",
		"1.5.0.0":  "1.4.0.1",
		"1.5.0.1":  "1.4.20.20",
		"1.5.1.10": "1.5.1.9",
	}
	for v1, v2 := range vers {
		if VersionCompare(v1, v2) != 1 {
			t.Errorf("version %s is great than version %s", v1, v2)
		}
	}
	for v1, v2 := range vers {
		if VersionCompare(v2, v1) != -1 {
			t.Errorf("version %s is less than version %s", v2, v1)
		}
	}
	for v1, _ := range vers {
		if VersionCompare(v1, v1) != 0 {
			t.Errorf("version %s is equal to version %s", v1, v1)
		}
	}
}
