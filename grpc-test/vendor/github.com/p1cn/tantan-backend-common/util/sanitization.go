package util

import (
	"regexp"
	"strings"
)

// regular expression to replace standard hyperlinks from text
var (
	hyperLinkTagsRegexp  = regexp.MustCompile(`<a\s+.*?>(.*?)</a>`)
	chineseRegexp        = regexp.MustCompile("[\u4e00-\u9fa5]")
	specialLettersRegexp = regexp.MustCompile("[^(?!\u4e00-\u9fa5a-zA-Z0-9-_)$]")
)

func SanitizeHyperLinkTags(str string) string {
	if strings.Contains(str, `</a>`) {
		return hyperLinkTagsRegexp.ReplaceAllString(str, "$1")
	}
	return str
}

func SanitizeDBCtoSBC(s string) string {
	retstr := ""
	for _, i := range s {
		insideCode := i
		if insideCode == 12288 {
			insideCode = 32
		} else {
			insideCode -= 65248
		}
		if insideCode < 32 || insideCode > 126 {
			retstr += string(i)
		} else {
			retstr += string(insideCode)
		}
	}
	return retstr
}

func SanitizeRemoveChinese(s string) string {
	if s != "" {
		return chineseRegexp.ReplaceAllString(s, "")
	}
	return s
}

func SanitizeRemoveSpaces(s string) string {
	return strings.Replace(s, " ", "", -1)
}

func SanitizeRemoveSpecialLetters(s string) string {
	if s != "" {
		return specialLettersRegexp.ReplaceAllString(s, "")
	}
	return s
}

func SanitizeDigital(s string) string {
	//add digital map here
	rot13 := func(r rune) rune {
		switch {
		case r >= 9312 && r <= 9320:
			return r - 9263
		case r == 9450:
			return r - 9402
		case r >= 10102 && r <= 10110:
			return r - 10053
		case r >= 9332 && r <= 9340:
			return r - 9283
		case r >= 9352 && r <= 9360:
			return r - 9303
		case r >= 8467 && r <= 8476:
			return r - 8419
		}
		return r
	}
	return strings.Map(rot13, s)
}
