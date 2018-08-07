package util

import (
	"errors"
	"strconv"
	"strings"
)

func VersionGreaterThan(v1, v2 string) bool {
	return VersionCompare(v1, v2) == 1
}

func VersionGreaterThanOrEqualTo(v1, v2 string) bool {
	c := VersionCompare(v1, v2)
	return c == 1 || c == 0
}

// greater than or equal to by comparing first 3 digits
// https://wiki.p1staff.com/wiki/Client_Versioning
func VersionGEByComparingFirst3Digits(v1, v2 string) (bool, error) {
	va1 := strings.Split(v1, ".")
	va2 := strings.Split(v2, ".")
	if len(va1) < 3 || len(va2) < 3 {
		return false, errors.New("wrong digit number")
	}
	for i := 0; i < 3; i++ {
		li, _ := strconv.Atoi(va1[i])
		ri, _ := strconv.Atoi(va2[i])
		if li > ri {
			return true, nil
		} else if li < ri {
			return false, nil
		}
	}
	return true, nil
}

func VersionCompare(v1, v2 string) (ret int) {
	v1s := strings.Split(v1, ".")
	v2s := strings.Split(v2, ".")
	length := len(v2s)
	if len(v1s) > len(v2s) {
		length = len(v1s)
	}
	for i := 0; i < length; i++ {
		var x, y string
		if len(v1s) > i {
			x = v1s[i]
		}
		if len(v2s) > i {
			y = v2s[i]
		}
		xi, _ := strconv.Atoi(x)
		yi, _ := strconv.Atoi(y)
		if xi > yi {
			ret = 1
		} else if xi < yi {
			ret = -1
		}
		if ret != 0 {
			break
		}
	}
	return
}
