package util

import (
	"testing"
	"time"
)

func TestLikesForBannedUserUserId(t *testing.T) {
	type testCase struct {
		userId string
		banned time.Time
		likes  int
	}

	now := time.Now()

	testCases := []testCase{
		{"000", now.Add(-240 * time.Hour), 305},
		{"000", now.Add(-24 * time.Hour), 214},
		{"000", now.Add(-time.Hour), 35},
		{"000", now.Add(-4 * time.Minute), 2},
		{"198", now.Add(-4 * time.Minute), 3},
		{"200", now.Add(-4 * time.Minute), 2},
		{"000", now.Add(0), 1},
		{"198", now.Add(0), 0},
		{"200", now.Add(0), 0},
		{"0", now.Add(3 * time.Minute), 0},
		{"0", now.Add(time.Hour), 0},
		{"0", now.Add(24 * time.Hour), 0},
		{"0", now.Add(240 * time.Hour), 0},
	}

	for _, v := range testCases {
		likes := LikesForBannedUser(v.userId, v.banned, v.banned, nil)
		if likes != v.likes {
			t.Errorf("test %v %v %v failed: gave: %v != %v",
				v.userId, v.banned, v.likes, likes, v.likes)
		}
	}
}

func TestLikesForBannedUserCreatedTime(t *testing.T) {
	type testCase struct {
		created time.Time
		banned  time.Time
		likes   int
	}

	now := time.Now()

	testCases := []testCase{
		{now.Add(2 * -time.Hour), now.Add(-time.Hour), 16},
		{now.Add(-time.Hour), now.Add(-time.Hour), 35},
		{now, now.Add(-time.Hour), 18},
		{now, now.Add(0), 1},
		{now.Add(time.Hour), now.Add(0), 0},
		{now.Add(-time.Hour), now.Add(0), 0},
		{now, now.Add(time.Hour), 0},
		{now.Add(time.Hour), now.Add(time.Hour), 0},
		{now.Add(-time.Hour), now.Add(time.Hour), 0},
	}

	for _, v := range testCases {
		likes := LikesForBannedUser("0", v.created, v.banned, nil)
		if likes != v.likes {
			t.Errorf("test %v %v %v failed: gave: %v != %v",
				v.created, v.banned, v.likes, likes, v.likes)
		}
	}
}

func TestLikesForBannedUserSwipeTime(t *testing.T) {
	type testCase struct {
		start time.Time
		end   time.Time
		likes int
	}

	now := time.Now()
	testCases := []testCase{
		{now.Add(-time.Hour), now, 35},
		{now.Add(-time.Hour), now.Add(10 * time.Minute), 35},
		{now.Add(-time.Hour), now.Add(-10 * time.Minute), 30},
		{now, now, 0},
		{now, now.Add(time.Hour), 1},
		{now, now.Add(-time.Hour), 0},
		{now.Add(time.Hour), now, 0},
		{now.Add(time.Hour), now.Add(time.Hour), 0},
		{now.Add(time.Hour), now.Add(-time.Hour), 0},
	}

	for _, v := range testCases {
		likes := LikesForBannedUser("0", v.start, v.start, &v.end)
		if likes != v.likes {
			t.Errorf("start: %v end: %v gave: %v != %v", v.start, v.end, likes, v.likes)
		}
	}
}
