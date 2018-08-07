package util

import "encoding/base64"

// EncodeBase64 encodes using encoding/base64.RawURLEncoding.
func EncodeBase64(bs []byte) string {
	return base64.RawURLEncoding.EncodeToString(bs)
}

// DecodeBase64 decodes using encoding/base64.RawURLEncoding.
func DecodeBase64(s string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(s)
}
