package external

const (
	MediaType = "media"
)

type Media struct {
	Name      string  `json:"name"`
	Url       string  `json:"url"`
	Size      []int   `json:"size"`
	Duration  float64 `json:"duration"`
	MediaType string  `json:"mediaType"`
}

type EmbeddedMediaInfo struct {
	Identifier string  `json:"id"`
	Width      int     `json:"w"`
	Height     int     `json:"h"`
	Duration   float64 `json:"d"`
	MediaType  string  `json:"mt"`
	DHash      uint64  `json:"dh"`
}
