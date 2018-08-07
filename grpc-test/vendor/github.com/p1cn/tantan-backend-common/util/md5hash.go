package util

import (
	"crypto/md5"
	"fmt"
)

func PhoneNumberToHashes(phone string) (hash8 string, hash11 string) {
	hash8 = fmt.Sprintf("%x", md5.Sum([]byte(phone[MaxInt(0, len(phone)-8):len(phone)])))
	hash11 = fmt.Sprintf("%x", md5.Sum([]byte(phone[MaxInt(0, len(phone)-11):len(phone)])))
	return hash8, hash11
}
