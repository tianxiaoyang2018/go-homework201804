package util

func CalSearchRadius(searchRadius, suggestSearchRadius int, searchRadiusLevels []int) int {
	oldVal := suggestSearchRadius
	if oldVal == 0 {
		oldVal = searchRadius
	}

	if searchRadiusLevels == nil {
		return oldVal
	}

	index := -1
	for i, v := range searchRadiusLevels {
		if v > oldVal {
			index = i
			break
		}
	}
	if index != -1 {
		return searchRadiusLevels[index]
	}
	return oldVal
}
