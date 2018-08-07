package external

const StickerType = "sticker"

type Sticker struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Package     IdType      `json:"package"`
	Pictures    []UserMedia `json:"pictures"`
	CreatedTime Iso8601Time `json:"createdTime"`
	Source      string      `json:"source"`
	Type        string      `json:"type"`
}

type Stickers []Sticker

func (self Stickers) GetReferencedPackageIds() []string {
	packIds := make([]string, 0)
	for _, s := range self {
		packIds = append(packIds, s.Id)
	}
	return packIds
}
