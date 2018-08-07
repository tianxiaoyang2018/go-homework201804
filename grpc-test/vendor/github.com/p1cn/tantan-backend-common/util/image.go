package util

import (
	"fmt"
	"strings"
)

var specialMediaTypes map[string]string = map[string]string{
	"jpg": "image/jpeg",
}

func MediaType2Extension(m string) string {
	for ext, t := range specialMediaTypes {
		if t == m {
			return ext
		}
	}
	index := strings.Index(m, "/")
	return m[index+1:]
}

func Extension2MediaType(ext string, t string) string {
	for extension, mediaType := range specialMediaTypes {
		if ext == extension {
			return mediaType
		}
	}
	return fmt.Sprintf("%s/%s", t, ext)
}
