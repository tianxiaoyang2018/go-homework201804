package external

const LikeType = "like"

type Like struct {
	Id          string      `json:"id"`
	Owner       IdType      `json:"owner"`
	Content     IdType      `json:"content"`
	ContentUser IdType      `json:"-"`
	Recalled    bool        `json:"-"`
	CreatedTime Iso8601Time `json:"createdTime"`
	UpdatedTime Iso8601Time `json:"updatedTime"`
	Type        string      `json:"type"`
}

type Likes []Like

func (self Likes) GetReferencedUserIds() []string {
	userIds := make([]string, 0)
	userIdsFound := make(map[string]bool)
	for _, m := range self {
		if m.Owner.Id != "" && !userIdsFound[m.Owner.Id] {
			userIds = append(userIds, m.Owner.Id)
			userIdsFound[m.Owner.Id] = true
		}
	}
	return userIds
}

func (self Likes) GetReferencedMomentIds() []string {
	momentIds := make([]string, 0)
	momentIdsFound := make(map[string]bool)
	for _, m := range self {
		if m.Content.Type == MomentType && m.Content.Id != "" && !momentIdsFound[m.Content.Id] {
			momentIds = append(momentIds, m.Content.Id)
			momentIdsFound[m.Content.Id] = true
		}
	}
	return momentIds
}
