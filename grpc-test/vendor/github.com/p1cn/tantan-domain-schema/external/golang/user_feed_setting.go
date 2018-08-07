package external

const (
	UserFeedSettingType = "user_feed_setting"
)

type UserFeedSetting struct {
	Id          string
	UserId      string
	OtherUserId string
	Muted       bool
	OtherMuted  bool
	Hidden      bool
	OtherHidden bool
	Status      string
	Type        string
}

type UserFeedSettings []UserFeedSetting

func (self UserFeedSettings) GetMutedUserIds() []string {
	ids := []string{}
	for _, s := range self {
		if s.Muted {
			ids = append(ids, s.OtherUserId)
		}
	}
	return ids
}

func (self UserFeedSettings) GetHiddenUserIds() []string {
	ids := []string{}
	for _, s := range self {
		if s.Hidden {
			ids = append(ids, s.OtherUserId)
		}
	}
	return ids
}
