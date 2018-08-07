package external

const ShopType = "shop"

type Shop struct {
	Id         string   `json:"id"`
	Identifier string   `json:"identifier"`
	Bundles    []IdType `json:"bundles"`
	Campaigns  []IdType `json:"campaigns"`
	Type       string   `json:"type"`
}

type Shops []Shop

func (self Shops) GetReferencedBundleIds() []string {
	ids := make([]string, 0)
	for _, s := range self {
		for _, obj := range s.Bundles {
			ids = append(ids, obj.Id)
		}
	}
	return ids
}

func (self Shops) GetReferencedCampaignIds() []string {
	ids := make([]string, 0)
	for _, s := range self {
		for _, obj := range s.Campaigns {
			ids = append(ids, obj.Id)
		}
	}
	return ids
}
