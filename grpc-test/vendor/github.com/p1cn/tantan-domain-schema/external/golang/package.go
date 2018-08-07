package external

const PackageType = "package"

type Package struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	CreatedTime Iso8601Time `json:"createdTime"`
	Pictures    []UserMedia `json:"pictures"`
	Objects     []IdType    `json:"objects"`
	PackageType string      `json:"packageType"`
	Activated   bool        `json:"activated"`
	Type        string      `json:"type"`
}

type Packages []Package

func (self Packages) GetReferencedStickerIds() []string {
	stickerIds := make([]string, 0)
	for _, p := range self {
		for _, obj := range p.Objects {
			if obj.Type == StickerType {
				stickerIds = append(stickerIds, obj.Id)
			}
		}
	}
	return stickerIds
}
