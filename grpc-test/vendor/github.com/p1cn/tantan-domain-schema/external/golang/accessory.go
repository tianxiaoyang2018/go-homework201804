package external

type Accessory struct {
	Id       string  `json:"id"`
	Type     string  `json:"type"`
	Category *string `json:"category,omitempty"`
}
