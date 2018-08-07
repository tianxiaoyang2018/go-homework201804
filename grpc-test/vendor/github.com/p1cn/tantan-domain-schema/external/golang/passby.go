package external

const (
	PassByType = "passby"
)

type PassByCount struct {
	Id         string      `json:"id"`
	User       IdType      `json:"user"`
	Count      int         `json:"count"`
	LatestTime Iso8601Time `json:"latestTime"`
	Type       string      `json:"type"`
}
