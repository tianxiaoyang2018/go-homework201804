package external

type MessageReference struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	AnswerId string `json:"answerId,omitempty"`
	Action   string `json:"action,omitempty"`
}
