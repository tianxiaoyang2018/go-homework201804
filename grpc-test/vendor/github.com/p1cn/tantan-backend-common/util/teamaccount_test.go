package util

import "testing"

func TestIsAdminUserId(t *testing.T) {
	adminUserIds := []string{"-1"}
	nonAdminUserIds := []string{"1", "0", "-0", "-2", "test"}

	for _, v := range adminUserIds {
		if !IsAdminUserId(v) {
			t.Errorf("%v should be admin user id", v)
		}
	}

	for _, v := range nonAdminUserIds {
		if IsAdminUserId(v) {
			t.Errorf("%v should not be admin user id", v)
		}
	}
}

func TestIsSwipePromotionUserId(t *testing.T) {
	swipePromotionUserIds := []string{"-1000", "-1001", "-1999"}
	nonSwipePromotionUserIds := []string{"1000", "1001", "1999", "-999", "-2000", "test"}

	for _, v := range swipePromotionUserIds {
		if !IsSwipePromotionUserId(v) {
			t.Errorf("%v should be a swipe promotion user id", v)
		}
	}

	for _, v := range nonSwipePromotionUserIds {
		if IsSwipePromotionUserId(v) {
			t.Errorf("%v should not be a swipe promotion user id", v)
		}
	}
}

func TestIsSpecialUserId(t *testing.T) {
	specialUserIds := []string{"-1", "-test", "-", "-0"}
	nonSpecialUserIds := []string{"1", "test", "", "12"}

	for _, v := range specialUserIds {
		if !IsSpecialUserId(v) {
			t.Errorf("%v should be a special user id", v)
		}
	}

	for _, v := range nonSpecialUserIds {
		if IsSpecialUserId(v) {
			t.Errorf("%v should not be a special user id", v)
		}
	}
}
