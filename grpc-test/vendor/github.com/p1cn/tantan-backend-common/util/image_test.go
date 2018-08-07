package util

import "testing"

func TestMediaType2ExtensionAndExtension2MediaType(t *testing.T) {
	examples := map[string]string{
		"jpg": "image/jpeg",
		"png": "image/png",
	}
	for ext, mt := range examples {
		if res := MediaType2Extension(mt); res != ext {
			t.Errorf("mediatype:%s, extension:%s to extension: %s", mt, ext, res)
		}
		if res := Extension2MediaType(ext, "image"); res != mt {
			t.Errorf("mediatype:%s, extension:%s to extension: %s", mt, ext, res)
		}
	}
}
