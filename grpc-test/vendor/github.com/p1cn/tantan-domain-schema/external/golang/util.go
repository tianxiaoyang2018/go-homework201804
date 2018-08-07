package external

func InSlice(str string, list []string) bool {
	for _, s := range list {
		if str == s {
			return true
		}
	}

	return false
}
