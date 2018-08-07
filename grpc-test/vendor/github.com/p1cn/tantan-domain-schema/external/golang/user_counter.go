package external

import (
	"time"
)

const UserCounterType = "user_counter"

type UserCounter struct {
	UserId              string
	Gender              string
	ReceivedLikes       int
	ReceivedDislikes    int
	ReceivedBlocks      int
	GivenLikes          int
	GivenDislikes       int
	GivenBlocks         int
	Matches             int
	UnreadConversations int
	LikeRating          int
	Popularity          float64
	Moments             int
	UnreadActivities    *int
	TotalCrushes        int
	GivenCrushes        int
	ReceivedCrushes     int
	Type                string
	GivenSuperLikes     int
	GivenUndos          int
	ReceivedSuperLikes  int

	TotalLikeLimit int
	RemainingLikes int

	LastResetTime time.Time
	LastShareTime time.Time

	SuperLikeReward     int
	SuperLikeBuy        int
	SuperLikeCount      int
	SuperLikeRemaining  int
	SuperLikeQuota      int
	SuperLikeDailyQuota *int

	UndoCount     int
	UndoReward    int
	UndoBuy       int
	UndoRemaining int
	UndoQuota     int
}

func (self UserCounter) GetSuperLikeRemaining() int {
	return self.SuperLikeRemaining
}

func (self UserCounter) GetUndoRemaining() int {
	return self.UndoRemaining
}

// func (self UserCounter) GetSuperLikeShare() int {
// 	if time.Now().Before(time.Time(self.LastShareTime).Add(time.Hour * 24 * time.Duration(config.Conf.SuperLikeLimits.ShareResetDays))) {
// 		return config.Conf.SuperLikeLimits.RewardForShare
// 	}
// 	return 0
// }

// func (self UserCounter) GetUndoShare() int {
// 	if time.Now().Before(time.Time(self.LastShareTime).Add(time.Hour * 24 * time.Duration(config.Conf.UndoLimits.ShareResetDays))) {
// 		return config.Conf.UndoLimits.RewardForShare
// 	}
// 	return 0
// }
