package util

import (
	"strings"
	"testing"
	"time"
)

func TestValidName(t *testing.T) {
	name := "Henry Ren"
	if ValidName(name) == false {
		t.Errorf("%s is a valid name", name)
	}
}

func BenchmarkValidName(b *testing.B) {
	name := "James Ma"
	for n := 0; n < b.N; n++ {
		ValidName(name)
	}
}

func TestValidEmail(t *testing.T) {
	emails := []string{
		"username@domain.cn",
		"username@domain.com",
		"username@domain.asia",
		"usernaMe@DOMAIN.com",
		"username@domain.COM",
		"user0name1@domain2.com",
	}
	for _, email := range emails {
		if ValidEmail(email) == false {
			t.Errorf("%s is a valid email", email)
		}
	}
	emails = []string{
		"username@domain.office",
		"@domain.com",
		"username@domain",
		".@domain.com",
	}
	for _, email := range emails {
		if ValidEmail(email) == true {
			t.Errorf("%s is not a valid email", email)
		}
	}
}

func TestValidPasswordLength(t *testing.T) {
	char := "p"
	counts := []int{1, 3, 257}
	for _, count := range counts {
		str := strings.Repeat(char, count)
		if ValidPasswordLength(str) == true {
			t.Errorf("%d is not a valid password length", len(str))
		}
	}
	counts = []int{6, 100, 256}
	for _, count := range counts {
		str := strings.Repeat(char, count)
		if ValidPasswordLength(str) == false {
			t.Errorf("%d is a valid password length", len(str))
		}
	}
}

func TestValidDate(t *testing.T) {
	dates := []string{
		"1900-01-01",
		"1987-02-03",
		"2050-04-05",
	}
	for _, date := range dates {
		if ValidDate(date) == false {
			t.Errorf("%s is a valid date", date)
		}
	}
	dates = []string{
		"1899-01-01",
		"2100-02-03",
	}
	for _, date := range dates {
		if ValidDate(date) == true {
			t.Errorf("%s is not a valid date", date)
		}
	}
}

func TestValidBirthdate(t *testing.T) {
	min, max := 16, 100
	leadDayAdjustment := 0
	if time.Now().Month() == time.February && time.Now().Day() == 29 {
		// If today is Feb 29 and we add one year to that we will end up on March 1. This compensates.
		leadDayAdjustment = -1
	}
	validDates := []time.Time{
		time.Now().AddDate(-min, 0, leadDayAdjustment),
		time.Now().AddDate(-(min + 1), 0, 0),
		time.Now().AddDate(-max, 0, 0),
		time.Now().AddDate(-(max - 1), 0, 0),
	}
	invalidDates := []time.Time{
		time.Now().AddDate(-(min - 1), 0, 0),
		time.Now().AddDate(-min, 1, 0),
		time.Now().AddDate(-min, 0, 1),
		time.Now().AddDate(-(max + 1), 0, leadDayAdjustment),
	}
	for _, date := range validDates {
		if ValidBirthdate(date) == false {
			t.Errorf("%s is a valid birthdate between the age range %d ~ %d", date, min, max)
		}
	}
	for _, date := range invalidDates {
		if ValidBirthdate(date) == true {
			t.Errorf("%s is not a valid birthdate between the age range %d ~ %d", date, min, max)
		}
	}
}

func TestValidMobileNumberFormat(t *testing.T) {
	mobiles := []string{
		"089123456",
		"89123456",
		"18612345678",
	}
	for _, mobile := range mobiles {
		if !ValidateMobile(mobile) {
			t.Errorf("%s is a valid mobile number format", mobile)
		}
	}
}

func TestInvalidMobileNumberFormat(t *testing.T) {
	mobiles := []string{
		"188-8888-8888",
		"188 8888 8888",
		"0",
		"00",
		"0089123456",
		"00089123456",
	}
	for _, mobile := range mobiles {
		if ValidateMobile(mobile) {
			t.Errorf("%s is not a valid mobile number format", mobile)
		}
	}
}

func TestValidateChineseMobileNumberFormat(t *testing.T) {
	mobiles := []string{
		"13112345678",
		"14712345678",
		"15812345678",
		"16812345678",
		"17012345678",
		"18612345678",
	}
	for _, mobile := range mobiles {
		if !ValidateChineseMobile(mobile) {
			t.Errorf("%s is a valid Chinese mobile number format", mobile)
		}
	}
}

func TestInvalidateChineseMobileNumberFormat(t *testing.T) {
	mobiles := []string{
		"12112345678",
		"11712345678",
		"1861234567",
		"186123456789",
		"89123456",
	}
	for _, mobile := range mobiles {
		if ValidateChineseMobile(mobile) {
			t.Errorf("%s is not a valid Chinese mobile number format", mobile)
		}
	}
}

func TestValidIOSPushNotificationToken(t *testing.T) {
	tokens := []string{
		"0123456789",
		"0abcdef0",
		"0abcdef",
		"abcdef0",
		"abcdef",
		"ab324345456cdef",
	}
	for _, token := range tokens {
		if !ValidIOSPushNotificationToken(token) {
			t.Errorf("%s is a valid ios push notification token format", token)
		}
	}
}

func TestInvalidIOSPushNotificationToken(t *testing.T) {
	tokens := []string{
		"0123456789h",
		"0abcdef0#",
		"0abcdef*",
		"abcd*ef0",
		"abcdefz",
		"zab324345456cdef",
		"@ab324345456cdef",
	}
	for _, token := range tokens {
		if ValidIOSPushNotificationToken(token) {
			t.Errorf("%s is not a valid ios push notification token format", token)
		}
	}
}

func TestJudgeName(t *testing.T) {
	vers := map[string]bool{
		//count not between 2 and 3
		"欧阳震华": false,
		"赵":    false,
		"赵匡胤":  true,

		// invalid surname
		"们": false,

		// in black lists
		"老师": false,
		"先生": false,
		"小姐": false,
	}

	for name, res := range vers {
		if judge := JudgeName(name); judge != res {
			t.Errorf("isChineseChar(%v) = %v; want %v", name, judge, res)
		}
	}
}

func TestValidAppsflyerID(t *testing.T) {
	vers := map[string]bool{
		"1-1":     true,
		"123-1":   true,
		"123-123": true,
		"12323333333333333333333333333333333-123333333333333333333333333333333333333333333": true,

		"":     false,
		"1":    false,
		"1-":   false,
		"-1":   false,
		"1a-1": false,
		"1😄-1": false,
		// more than 50 before -
		"012345678901234567890123456789012345678901234567891-1": false,
		// more than 50 after -
		"1-012345678901234567890123456789012345678901234567891": false,
	}

	for id, res := range vers {
		if r := ValidAppsflyerID(id); r != res {
			t.Errorf("ValidAppsflyerID(%v) = %v; want %v", id, r, res)
		}
	}
}
