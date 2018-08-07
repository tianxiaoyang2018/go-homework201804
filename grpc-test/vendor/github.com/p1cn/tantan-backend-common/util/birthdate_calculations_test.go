package util

import (
	"testing"
	"time"
)

func TestCalculateZodiac(t *testing.T) {

	type testCase struct {
		date   string
		zodiac string
	}

	testCases := []testCase{
		//utc cases
		{"1940-12-21T00:00:00+00:00", "sagittarius"},
		{"1940-12-21T00:00:00+00:00", "sagittarius"},
		{"1940-12-21T08:00:00+00:00", "sagittarius"},
		{"1940-12-21T23:59:59+00:00", "sagittarius"},
		{"1940-12-22T00:00:00+00:00", "capricorn"},
		{"1940-12-22T08:00:00+00:00", "capricorn"},
		{"1940-12-22T23:59:59+00:00", "capricorn"},
		//prc case
		{"1940-12-21T00:00:00+08:00", "sagittarius"},
		{"1940-12-21T00:00:00+08:00", "sagittarius"},
		{"1940-12-21T08:00:00+08:00", "sagittarius"},
		{"1940-12-21T23:59:59+08:00", "sagittarius"},
		{"1940-12-22T00:00:00+08:00", "capricorn"},
		{"1940-12-22T08:00:00+08:00", "capricorn"},
		{"1940-12-22T23:59:59+08:00", "capricorn"},
	}

	for _, testCase := range testCases {
		birthdate, err := time.Parse(time.RFC3339, testCase.date)
		if err != nil {
			t.Errorf("%+v", err)
		}

		zodiac := CalculateZodiac(birthdate)

		if zodiac != testCase.zodiac {
			t.Errorf("Expected zodiac %v but got %v for birthdate %v", testCase.zodiac, zodiac, birthdate)
		}
	}
}
