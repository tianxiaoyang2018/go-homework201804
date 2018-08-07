package external

type CloudObjects struct {
	Media []Media `json:"media"`
}

type CloudEnvelope struct {
	Meta Meta         `json:"meta"`
	Data CloudObjects `json:"data"`
}
