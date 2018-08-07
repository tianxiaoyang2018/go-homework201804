package util

import (
	"math"
	"strconv"
	"time"
)

func getLikes(user_id int, seconds float64) int {

	// vary the speed of likes by user_id
	speed := float64(user_id % 199)

	// offset starting point by the user_id (0 to about 5 minutes)
	offset := float64(user_id % 293)
	seconds = seconds - offset

	// ever increasing number of likes in a slow linear fashion
	y1 := seconds / (60 * 60 * 24) * 10

	// add a "random" looking factor so that its not a totally smooth curve
	r2 := seconds / 60 / 60 / 2 * 2 * math.Pi
	y2 := 5 * math.Cos(r2+math.Pi) * math.Pow(2, (-seconds/10000))

	// increase likes a lot in the first ~8 hours and then reduce speed
	r3 := seconds / 60 / 60 / 24 / 2 * 2 * math.Pi
	y3 := (speed + 200) * math.Tanh(r3)

	// add four to make it start at 0 at time 0.
	likes := 4 + int(y1+y2+y3)

	return likes
}

func getLikesOffset(user_id int, seconds, startSeconds float64) int {
	likes := getLikes(user_id, seconds+startSeconds) - getLikes(user_id, startSeconds)

	if likes < 0 {
		return 0
	}
	return likes
}

func LikesForBannedUser(userId string, createdTime, bannedTime time.Time, startSwipeTime *time.Time) int {
	userIdInt, _ := strconv.Atoi(userId)

	seconds := time.Since(bannedTime).Seconds()

	startSeconds := seconds - time.Since(createdTime).Seconds()

	if startSwipeTime != nil {
		delta := startSwipeTime.Sub(bannedTime).Seconds()
		seconds = math.Min(seconds, delta)
	}

	likes := getLikesOffset(userIdInt, seconds, startSeconds)
	return likes
}
