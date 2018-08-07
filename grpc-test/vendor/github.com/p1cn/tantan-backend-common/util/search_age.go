package util

import (
	"time"
)

func AllowedSearchAge(birthdate time.Time) (min, max, maxPlus int) {
	age := CalculateAge(birthdate)

	if age < 18 {
		min = 16
		max = 22
		maxPlus = 22
	} else if age <= 22 {
		min = 16
		max = 50
		maxPlus = 200
	} else {
		min = 18
		max = 50
		maxPlus = 200
	}
	return
}

func DefaultSearchAge(birthdate time.Time) (min, max int) {
	age := CalculateAge(birthdate)

	if age < 18 {
		min = 16
		max = 22
	} else {
		min = 18
		max = 200
	}
	return
}
