package util

import (
	"errors"
	"regexp"
	"strconv"
)

var RangeRegexp *regexp.Regexp = regexp.MustCompile("^bytes=([0-9]*)-([0-9]*)$")

func ParseRange(rangeHeader string) (*int, *int, error) {
	if !RangeRegexp.MatchString(rangeHeader) {
		return nil, nil, errors.New("invalid range header")
	}

	s := RangeRegexp.FindStringSubmatch(rangeHeader)
	if len(s) != 3 {
		return nil, nil, errors.New("invalid range header")
	}

	var start *int
	var end *int

	if s[1] != "" {
		x, err := strconv.Atoi(s[1])
		if err != nil {
			return nil, nil, err
		}
		start = &x
	}

	if s[2] != "" {
		x, err := strconv.Atoi(s[2])
		if err != nil {
			return nil, nil, err
		}
		end = &x
	}

	if start == nil && end == nil {
		return nil, nil, errors.New("invalid range header")
	}

	if start != nil && end != nil && *start > *end {
		return nil, nil, errors.New("invalid range header")
	}

	return start, end, nil
}
