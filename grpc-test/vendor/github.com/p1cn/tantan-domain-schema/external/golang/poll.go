package external

const (
	PollType            = "poll"
	TbhQuestionCategory = "tbh"
)

type PollQuestion struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Category string `json:"category"`
}

type Poll struct {
	Id          string       `json:"id"`
	Owner       IdType       `json:"owner"`
	Question    PollQuestion `json:"question"`
	Media       []UserMedia  `json:"media"`
	Candidates  []IdType     `json:"candidates"`
	Vote        *IdType      `json:"vote,omitempty"`
	Results     []float32    `json:"results,omitempty"`
	CreatedTime Iso8601Time  `json:"createdTime"`
	UpdatedTime Iso8601Time  `json:"updatedTime"`
	VotedTime   Iso8601Time  `json:"votedTime"`
	Type        string       `json:"type"`
}

type Polls []Poll

func (self Polls) GetReferencedUserIds() []string {
	userIds := make([]string, 0)
	foundUsersMap := make(map[string]bool)
	for _, p := range self {
		if !foundUsersMap[p.Owner.Id] {
			userIds = append(userIds, p.Owner.Id)
			foundUsersMap[p.Owner.Id] = true
		}

		for _, c := range p.Candidates {
			if !foundUsersMap[c.Id] {
				userIds = append(userIds, c.Id)
				foundUsersMap[c.Id] = true
			}
		}

		if p.Vote != nil {
			if !foundUsersMap[p.Vote.Id] {
				userIds = append(userIds, p.Vote.Id)
				foundUsersMap[p.Vote.Id] = true
			}
		}
	}

	return userIds
}

func (self Polls) GetReferencedQuestionIds() []string {
	qids := make([]string, 0)

	for _, p := range self {
		qids = append(qids, p.Question.Id)
	}

	return qids
}
