package external

type UserRegion struct {
	Country  *string `json:"country"`
	Province *string `json:"-"`
	City     *string `json:"city"`
	District *string `json:"district"`
}
