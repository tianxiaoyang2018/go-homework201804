package util

func PrometheusFixLabel(str string) string {
	bytes := make([]rune, len(str))
	for i, s := range str {
		if !((s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || s == '_' || (s >= '0' && s <= '9' && i > 0)) {
			bytes[i] = '_'
		} else {
			bytes[i] = s
		}
	}
	return string(bytes)
}
