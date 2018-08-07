package util

import (
	"fmt"
	"time"
)

func CalculateAge(birthdate time.Time) int {
	now := time.Now()
	age := now.Year() - birthdate.Year()
	if now.Month() < birthdate.Month() {
		age--
	} else if now.Month() == birthdate.Month() {
		if now.Day() < birthdate.Day() {
			age--
		}
	}
	return age
}

type zd struct {
	StartMonth time.Month
	StartDay   int
	EndMonth   time.Month
	EndDay     int
}

// ref http://en.wikipedia.org/wiki/Zodiac "Tropical Zodiac 2011"
var dateToZodiac = map[zd]string{
	zd{StartMonth: 3, StartDay: 21, EndMonth: 4, EndDay: 19}:   "aries",
	zd{StartMonth: 4, StartDay: 20, EndMonth: 5, EndDay: 20}:   "taurus",
	zd{StartMonth: 5, StartDay: 21, EndMonth: 6, EndDay: 21}:   "gemini",
	zd{StartMonth: 6, StartDay: 22, EndMonth: 7, EndDay: 22}:   "cancer",
	zd{StartMonth: 7, StartDay: 23, EndMonth: 8, EndDay: 22}:   "leo",
	zd{StartMonth: 8, StartDay: 23, EndMonth: 9, EndDay: 22}:   "virgo",
	zd{StartMonth: 9, StartDay: 23, EndMonth: 10, EndDay: 23}:  "libra",
	zd{StartMonth: 10, StartDay: 24, EndMonth: 11, EndDay: 22}: "scorpio",
	zd{StartMonth: 11, StartDay: 23, EndMonth: 12, EndDay: 21}: "sagittarius",
	zd{StartMonth: 12, StartDay: 22, EndMonth: 12, EndDay: 31}: "capricorn",
	zd{StartMonth: 1, StartDay: 1, EndMonth: 1, EndDay: 19}:    "capricorn",
	zd{StartMonth: 1, StartDay: 20, EndMonth: 2, EndDay: 18}:   "aquarius",
	zd{StartMonth: 2, StartDay: 19, EndMonth: 3, EndDay: 20}:   "pisces",
}

func CalculateZodiac(birthdate time.Time) string {
	for k, v := range dateToZodiac {
		start := time.Date(birthdate.Year(), k.StartMonth, k.StartDay, 0, 0, 0, 0, time.UTC)
		end := time.Date(birthdate.Year(), k.EndMonth, k.EndDay, 0, 0, 0, 0, time.UTC).AddDate(0, 0, 1)

		utcBD := time.Date(birthdate.Year(), birthdate.Month(), birthdate.Day(), 0, 0, 0, 0, time.UTC)

		if !utcBD.Before(start) && !(utcBD.After(end) || utcBD.Equal(end)) {
			return v
		}
	}
	panic(fmt.Sprintf("Birthdate doesnt have zodiac! %v", birthdate))
}
