package util

import "strings"

const (
	LanguageZh   = "zh-CN"
	LanguageZhTW = "zh-TW"
	LanguageEn   = "en-US"
	LanguageTh   = "th-TH"
	LanguageKo   = "ko-KR"
	LanguageJa   = "ja-JP"
)

var TranslatedLanguages = []string{
	LanguageZh,
	LanguageZhTW,
	LanguageEn,
	LanguageTh,
	LanguageKo,
	LanguageJa,
}

var AcceptLanguages = []string{
	LanguageZh,
	LanguageZhTW,
	LanguageTh,
	LanguageEn,
	LanguageKo,
	LanguageJa,
}

func ParseLanguage(text string) string {
	text = strings.ToLower(text)
	switch {
	case strings.Contains(text, "zh-hant") || strings.Contains(text, "zh-tw") || strings.Contains(text, "zh-hk"):
		return LanguageZhTW
	case strings.Contains(text, "zh"):
		return LanguageZh
	case strings.Contains(text, "th"):
		return LanguageTh
	case strings.Contains(text, "ko"):
		return LanguageKo
	case strings.Contains(text, "ja"):
		return LanguageJa
	}
	return LanguageEn
}

func ValidAcceptLanguage(language string) bool {
	return InSlice(language, AcceptLanguages)
}
