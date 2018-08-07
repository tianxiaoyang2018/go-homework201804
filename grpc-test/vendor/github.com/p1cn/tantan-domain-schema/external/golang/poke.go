package external

const (
	PokeType = "poke"
)

type Poke struct {
	Id          string      `json:"id"`
	Owner       IdType      `json:"-"`
	OtherUser   IdType      `json:"-"`
	State       string      `json:"state"`
	OtherState  string      `json:"-"`
	CreatedTime Iso8601Time `json:"createdTime"`
	UpdatedTime Iso8601Time `json:"updateTime"`
	Type        string      `json:"type"`
}
