package ip

import "testing"

func TestIsLocalAddress(t *testing.T) {
	data := map[string]bool{
		"127.0.0.0":     true,
		"10.0.0.0":      true,
		"10.191.160.35": true,
		"169.254.0.0":   true,
		"172.15.0.0":    false,
		"172.16.0.0":    true,
		"172.31.0.0":    true,
		"172.32.0.0":    false,
		"192.0.0.0":     true,
		"192.168.0.0":   true,
		"147.12.56.11":  false,
		"123.125.169.2": false,
	}
	for addr, isLocal := range data {
		if IsLocalAddress(addr) != isLocal {
			format := "%s should "
			if !isLocal {
				format += "not "
			}
			format += "be local address"
			t.Errorf(format, addr)
		}
	}
}
