package external

const (
	LocationType = "location"
)

type LocationIntent string

const (
	LocationIntentBackgroundChange LocationIntent = "backgroundChange"
)

type Location struct {
	Coordinates []float64      `json:"coordinates"`
	Uncertainty int            `json:"uncertainty"`
	Intent      LocationIntent `json:"intent"`
}
