package external

type IdsOrder string

const (
	IdsOrderAsc     = "ascending"
	IdsOrderDesc    = "descending"
	IdsOrderLatest  = "latest"
	IdsOrderPopular = "popularity"
)

type IdsFilter string

const (
	IdsFilterNone      = "none"
	IdsFilterOngoing   = "ongoing"
	IdsFilterCompleted = "completed"
)

type IdsType struct {
	Ids    []string  `json:"ids"`
	Count  int64     `json:"count"`
	Order  IdsOrder  `json:"order"`
	Filter IdsFilter `json:"filter"`
	Type   string    `json:"type"`
}
