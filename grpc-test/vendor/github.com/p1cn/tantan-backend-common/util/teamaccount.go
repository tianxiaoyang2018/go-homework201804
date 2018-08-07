package util

import (
	"strconv"
)

var (
	TeamAccountUserId = "-1"

	SwipePromotionUserIdMin = -1999
	SwipePromotionUserIdMax = -1000
)

func IsAdminUserId(userId string) bool {
	return userId == TeamAccountUserId
}

func IsSwipePromotionUserId(userId string) bool {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return false
	}
	return userIdInt >= SwipePromotionUserIdMin && userIdInt <= SwipePromotionUserIdMax
}

func IsSpecialUserId(userId string) bool {
	if len(userId) > 0 {
		return userId[0] == '-'
	}
	return false
}

func IsBrandAccount(userId string) bool {
	// If need add some other brand accout, add here

	return userId == TeamAccountUserId
}
