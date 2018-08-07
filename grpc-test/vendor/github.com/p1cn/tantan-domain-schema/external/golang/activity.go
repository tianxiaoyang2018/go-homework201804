package external

const ActivityType = "activity"

const (
	ActivityActionLike    = "like"
	ActivityActionComment = "comment"
	ActivityActionFollow  = "follow"
)

type Activity struct {
	Id          string      `json:"id"`
	Value       string      `json:"value"`
	Owner       IdType      `json:"owner"`
	Actors      []IdType    `json:"actors"`
	Action      string      `json:"action"`
	Reference   IdType      `json:"reference"`
	Read        bool        `json:"read"`
	CreatedTime Iso8601Time `json:"createdTime"`
	Type        string      `json:"type"`
}

func (self Activity) IsLike() bool {
	return self.Action == ActivityActionLike
}

func (self Activity) IsComment() bool {
	return self.Action == ActivityActionComment
}

func (self Activity) IsFollow() bool {
	return self.Action == ActivityActionFollow
}

type Activities []Activity

func (self Activities) GetReferencedUserIds() []string {
	userIds := []string{}
	userIdsFound := make(map[string]bool)
	for _, a := range self {
		if a.Owner.Id != "" && !userIdsFound[a.Owner.Id] {
			userIds = append(userIds, a.Owner.Id)
			userIdsFound[a.Owner.Id] = true
		}
		for _, u := range a.Actors {
			if u.Id != "" && !userIdsFound[u.Id] {
				userIds = append(userIds, u.Id)
				userIdsFound[u.Id] = true
			}
		}
	}
	return userIds
}

func (self Activities) GetReferencedMomentIds() []string {
	momentIds := []string{}
	momentIdsFound := make(map[string]bool)
	for _, a := range self {
		if a.Reference.Type == MomentType && a.Reference.Id != "" && !momentIdsFound[a.Reference.Id] {
			momentIds = append(momentIds, a.Reference.Id)
			momentIdsFound[a.Reference.Id] = true
		}
	}
	return momentIds
}
