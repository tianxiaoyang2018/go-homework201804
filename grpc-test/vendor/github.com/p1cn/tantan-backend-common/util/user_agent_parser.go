package util

import (
	"strings"
)

const (
	OSAndroid = "android"
	OSIOS     = "ios"
)

func ParseUserAgent(userAgent string) (string, string) {
	userAgent = strings.ToLower(userAgent)
	osName := ""
	if strings.Contains(userAgent, OSAndroid) {
		osName = OSAndroid
	} else if strings.Contains(userAgent, OSIOS) ||
		strings.Contains(userAgent, "darwin") ||
		strings.Contains(userAgent, "iphone") {
		osName = OSIOS
	}
	pieces := strings.Split(userAgent, " ")
	if len(pieces) < 1 {
		return osName, ""
	}
	pieces = strings.Split(pieces[0], "/")
	if len(pieces) < 2 {
		return osName, ""
	}
	return osName, pieces[1]
}
