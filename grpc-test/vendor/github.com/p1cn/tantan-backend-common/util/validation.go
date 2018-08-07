package util

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	CNCountryCode = 86
	KRCountryCode = 82
	JACountryCode = 81
)

const (
	allowedMinAge = 16
	allowedMaxAge = 100
)

var (
	validNameRegexp          = regexp.MustCompile(`^[\p{L}-.\040]{1,50}$`)
	validVersionRegexp       = regexp.MustCompile(`^\d+(\.\d+)+$`)
	validChineseNameRegexp   = regexp.MustCompile("^[\u4e00-\u9fa5]+$")
	validEmailRegexp         = regexp.MustCompile(`^(?i)\b[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}\b$`)
	validDateRegexp          = regexp.MustCompile(`^(19|20)\d\d-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$`)
	validMobileRegexp        = regexp.MustCompile(`^0?[1-9][0-9]{1,30}$`)
	validChineseMobileRegexp = regexp.MustCompile(`^1[3456789][0-9]{9}$`)
	zhRegexp                 = regexp.MustCompile("^[\u4e00-\u9fa5]+$")
	numRegexp                = regexp.MustCompile(`^[0-9]+$`)
	likeMobileNumRegexp      = regexp.MustCompile(`^[0-9\040]+$`)
	validDevicePushToken     = regexp.MustCompile(`^[a-f0-9]+$`)

	// AppsflyerID regexp
	validAppsflyerIDRegexp = regexp.MustCompile(`^\d{1,50}-\d{1,50}$`)
)

func ValidVersion(version string) bool {
	return validVersionRegexp.MatchString(version)
}

// ValidName validates user names
// @link http://www.regular-expressions.info/unicode.html
func ValidName(name string) bool {
	matched := validNameRegexp.MatchString(name)
	if matched {
		if len(SanitizeName(name)) == 0 {
			return false
		}
	}
	return matched
}

func ValidUserID(uid string) bool {
	id, err := strconv.Atoi(uid)
	if err != nil {
		return false
	}
	if id > math.MaxInt32 {
		return false
	}
	return true
}

// ValidChine
func ValidChineseName(name string) bool {
	return validChineseNameRegexp.MatchString(name)
}

// SanitizeName will do the following simple sanitization
// 1) remove leading and trailing spaces, dashes and dots
// 2) replace consecutive spaces, dashes and dots within the string
// @todo improve & add normalize
func SanitizeName(name string) string {
	name = strings.Trim(strings.TrimSpace(name), "-. ")
	rs := []rune(name)
	l := len(rs)
	fields := make([]rune, 0)
	if l < 3 {
		fields = rs
	} else {
		fields = append(fields, rs[0])
		for i := 0; i < l-2; i++ {
			valid := false
			switch rs[i+1] {
			case '-':
				valid = unicode.IsLetter(rs[i]) && unicode.IsLetter(rs[i+2])
			case '.':
				valid = unicode.IsLetter(rs[i]) && unicode.IsSpace(rs[i+2])
			default:
				valid = true
			}
			if valid {
				fields = append(fields, rs[i+1])
			}
			if i+3 == l {
				fields = append(fields, rs[i+2])
			}
		}
	}
	name = string(fields)
	names := strings.Fields(name)

	return strings.Join(names, " ")
}

// ValidEmail reports true if an email address matches the specified regular expression
// @link http://www.regular-expressions.info/email.html
func ValidEmail(email string) bool {
	return validEmailRegexp.MatchString(email)
}

func ValidEmailLength(email string) bool {
	return len(email) <= 50
}

// ValidPasswordLength reports true if the password lenght is not within specified range
func ValidPasswordLength(pwd string) bool {
	rs := []rune(pwd)
	return len(rs) >= 6 && len(rs) <= 256
}

// ValidGender validates the signup user gender
func ValidGender(gender string) bool {
	return InSlice(gender, []string{"male", "female"})
}

// ValidDate reports true if the provided date is between 1900-01-01 ~ 2099-12-31
// @link http://www.regular-expressions.info/dates.html
func ValidDate(date string) bool {
	return validDateRegexp.MatchString(date)
}

// ValidBirthdate validates the provided birthdate according to age range and
// the date must be in yyyy-mm-dd format
// Note: invalid date such as 1987-02-30 should be converted into
// correct date using time.Parse
func ValidBirthdate(birthdate time.Time) bool {
	age := CalculateAge(birthdate)
	return age >= allowedMinAge && age <= allowedMaxAge
}

func ValidateMobile(mobile string) bool {
	return validMobileRegexp.MatchString(mobile)
}

func ValidateChineseMobile(mobile string) bool {
	return validChineseMobileRegexp.MatchString(mobile)
}

func ValidIOSPushNotificationToken(token string) bool {
	return validDevicePushToken.MatchString(strings.ToLower(token))
}

// IsChinese return true if the provide string only
// contains Chinese characters
func IsChinese(s string) bool {
	return zhRegexp.MatchString(s)
}

func IsNumeric(str string) bool {
	return numRegexp.MatchString(str)
}

func LikeMobileNumber(str string) bool {
	return likeMobileNumRegexp.MatchString(str)
}

func LikeEmailAddress(str string) bool {
	return strings.Contains(str, "@")
}

func LikeUsername(str string) bool {
	return !(LikeEmailAddress(str) || LikeMobileNumber(str))
}

// judgeName judge if a user name is illegal.
func JudgeName(name string) bool {
	// '^[\u4e00-\u9fa5]+$']
	if !IsChinese(name) {
		return false
	}

	// length [2,3]
	count := utf8.RuneCountInString(name)
	if count < 2 || count > 3 {
		return false
	}

	// surname is valid
	bingo := false
	for _, r := range ChineseSurname {
		if strings.HasPrefix(name, r) {
			bingo = true
			break
		}
	}
	if !bingo {
		return false
	}

	// name include black lists word.
	for _, b := range SmsPromotionBlacklists {
		if strings.Contains(name, b) {
			return false
		}
	}
	return true
}

// ValidAppsflyerID reports whether s is valid AppsflyerID.
// AppsflyerID format(from db regexp conclusion): '^\d+-\d+$'
func ValidAppsflyerID(s string) bool {
	return validAppsflyerIDRegexp.MatchString(s)
}
