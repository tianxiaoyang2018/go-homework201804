package external

const CampaignType = "campaign"

type Campaign struct {
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	CampaignName string      `json:"campaignName"`
	CampaignCode string      `json:"campaignCode"`
	Status       string      `json:"status"`
	CreatedTime  Iso8601Time `json:"createdTime"`
	Objects      []IdType    `json:"objects"`
	Type         string      `json:"type"`
}

type Campaigns []Campaign

func (self Campaigns) GetReferencedBundleIds() []string {
	bundleIds := make([]string, 0)
	for _, c := range self {
		for _, obj := range c.Objects {
			if obj.Type == BundleType {
				bundleIds = append(bundleIds, obj.Id)
			}
		}
	}
	return bundleIds
}
