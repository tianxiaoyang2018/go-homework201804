package external

const BundleType = "bundle"

const (
	BundleStatusStock     = "stock"
	BundleStatusLocked    = "locked"
	BundleStatusPurchased = "purchased"
)

type Bundle struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Price       float64     `json:"price"`
	Currency    string      `json:"currency"`
	Description string      `json:"description"`
	Status      string      `json:"status"`
	CreatedTime Iso8601Time `json:"createdTime"`
	Pictures    []UserMedia `json:"pictures"`
	Objects     []IdType    `json:"objects"`
	Type        string      `json:"type"`
}

type Bundles []Bundle

func (self Bundles) GetReferencedPackageIds() []string {
	packIds := make([]string, 0)
	for _, c := range self {
		for _, obj := range c.Objects {
			if obj.Type == PackageType {
				packIds = append(packIds, obj.Id)
			}
		}
	}
	return packIds
}
