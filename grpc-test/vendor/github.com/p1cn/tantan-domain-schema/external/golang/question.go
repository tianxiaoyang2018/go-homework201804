package external

const (
	QuestionType = "question"
)

// question categories
const (
	QuestionFilterClassic  = "classic"
	QuestionFilterIntimate = "intimate"
	QuestionFilterProfile  = "profile"
	QuestionFilterTbh      = "tbh"
)

// question usage
const (
	QuestionUsageDefault = "default"
	QuestionUsageRequest = "request"
)

type Question struct {
	Id       string   `json:"id"`
	Text     string   `json:"text"`
	Answers  []Answer `json:"answers"`
	Category string   `json:"category"`
	Type     string   `json:"type"`
}
