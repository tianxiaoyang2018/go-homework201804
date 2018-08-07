package external

const (
	AnswerType = "answer"
)

const (
	AnswerAttitudePositive = "positive"
	AnswerAttitudeNegative = "negative"
	AnswerAttitudeNeutral  = "neutral"
)

type Answer struct {
	Id       string `json:"id"`
	Attitude string `json:"-"`
	Text     string `json:"value"`
	Type     string `json:"-"`
}
