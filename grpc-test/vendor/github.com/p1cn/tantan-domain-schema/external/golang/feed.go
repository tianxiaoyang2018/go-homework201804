package external

const FeedType = "feed"

type Feed struct {
	Id          string      `json:"id"`
	Owner       IdType      `json:"owner"`
	Moment      IdType      `json:"moment"`
	MomentOwner IdType      `json:"momentOwner"`
	CreatedTime Iso8601Time `json:"createdTime"`
	Type        string      `json:"type"`
}

type Feeds []Feed

func (self Feeds) GetReferencedUserIds() []string {
	userIds := make([]string, 0)
	userIdsFound := make(map[string]bool)
	for _, f := range self {
		if f.Owner.Id != "" && !userIdsFound[f.Owner.Id] {
			userIds = append(userIds, f.Owner.Id)
			userIdsFound[f.Owner.Id] = true
		}
		if f.MomentOwner.Id != "" && !userIdsFound[f.MomentOwner.Id] {
			userIds = append(userIds, f.MomentOwner.Id)
			userIdsFound[f.MomentOwner.Id] = true
		}
	}
	return userIds
}

func (self Feeds) GetReferencedMomentIds() []string {
	momentIds := make([]string, 0)
	momentIdsFound := make(map[string]bool)
	for _, f := range self {
		if f.Moment.Id != "" && !momentIdsFound[f.Moment.Id] {
			momentIds = append(momentIds, f.Moment.Id)
			momentIdsFound[f.Moment.Id] = true
		}
	}
	return momentIds
}
