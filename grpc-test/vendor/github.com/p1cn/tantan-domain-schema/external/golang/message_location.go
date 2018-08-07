package external

const (
	MessageLocationType = "location"
)

type MessageLocation struct {
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	Coordinates []float64 `json:"coordinates"`
}
