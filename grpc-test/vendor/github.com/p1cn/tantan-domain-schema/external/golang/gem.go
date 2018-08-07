package external

const (
	GemType = "gem"
)

type Gem struct {
	Id        string      `json:"id"`
	Type      string      `json:"type"`
	UserId    string      `json:"-"`
	Count     int         `json:"count"`
	Reference IdType      `json:"reference"`
	Media     []UserMedia `json:"media"`
}

type Gems []Gem

func (self Gems) GetReferencedQuestionIds() []string {
	qids := make([]string, 0)

	for _, g := range self {
		qids = append(qids, g.Reference.Id)
	}

	return qids
}
