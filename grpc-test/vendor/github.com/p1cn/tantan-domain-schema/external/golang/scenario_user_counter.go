package external

import "time"

type ScenarioUserCounter struct {
	UserId                string
	Category              string
	ScenarioIds           []string
	ScenarioExpiresTime   time.Time
	ScenarioReceivedLikes int
	ReceivedLikes         int
	ReceivedDislikes      int
	GivenLikes            int
	GivenDislikes         int
	Matches               int
	MatchesWithinLimit    int
	RemainingMatches      int
	MatchesLastResetTime  time.Time
	Type                  string
}

func (c ScenarioUserCounter) IsCategoryActive() bool {
	return len(c.ScenarioIds) > 0 &&
		time.Now().UTC().Before(c.ScenarioExpiresTime)
}
