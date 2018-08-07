package external

type IdType struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type IdTypes []IdType

func (self IdTypes) Ids() []string {
	ids := make([]string, 0)
	for _, i := range self {
		ids = append(ids, i.Id)
	}
	return ids
}
