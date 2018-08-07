package util

func ValidCoordinates(coord []float64) bool {
	if len(coord) < 2 {
		return false
	}

	return ValidLatitude(coord[0]) && ValidLongitude(coord[1])
}

func ValidLatitude(lat float64) bool {
	return !(lat < -90 || lat > 90)
}

func ValidLongitude(lon float64) bool {
	return !(lon < -180 || lon > 180)
}
