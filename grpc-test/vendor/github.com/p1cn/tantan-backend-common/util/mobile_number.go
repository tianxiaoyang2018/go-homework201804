package util

import (
	"errors"
	"fmt"

	"github.com/ttacon/libphonenumber"
)

var (
	ErrInvalidNumber = errors.New("invalid mobile number")
)

func ParseMobileNumber(countryCode int, mobileNumber string) (string, error) {
	return parseMobileNumber(countryCode, mobileNumber, "")
}

// region like: CN which can be found here: https://countrycode.org/china
func ParseMobileNumberWithDefaultRegion(countryCode int, mobileNumber string, region string) (string, error) {
	return parseMobileNumber(countryCode, mobileNumber, region)
}

func parseMobileNumber(countryCode int, mobileNumber string, region string) (string, error) {
	num, err := libphonenumber.Parse(fmt.Sprintf("+%d %s", countryCode, mobileNumber), region)
	if err != nil {
		return "", err
	}
	if !libphonenumber.IsValidNumber(num) {
		return "", ErrInvalidNumber
	}
	return fmt.Sprintf("%d", num.GetNationalNumber()), nil
}
