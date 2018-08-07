// References:
// https://en.wikipedia.org/wiki/Reserved_IP_addresses
// https://github.com/tomasen/realip
package ip

import "net"

var reservedCidrs []*net.IPNet

func init() {
	cidrBlocks := []string{
		"127.0.0.1/8",
		"10.0.0.0/8",
		"100.64.0.0/10",
		"169.254.0.0/16",
		"172.16.0.0/12",
		"192.0.0.0/24",
		"192.168.0.0/16",
		"198.18.0.0/15",
	}
	reservedCidrs = []*net.IPNet{}
	for _, v := range cidrBlocks {
		_, cidrnet, err := net.ParseCIDR(v)
		if err != nil {
			continue
		}
		reservedCidrs = append(reservedCidrs, cidrnet)
	}
}

func IsLocalAddress(addr string) bool {
	for idx := range reservedCidrs {
		if reservedCidrs[idx].Contains(net.ParseIP(addr)) {
			return true
		}
	}
	return false
}
