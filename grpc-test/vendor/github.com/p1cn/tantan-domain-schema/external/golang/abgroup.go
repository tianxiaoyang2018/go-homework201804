package external

import (
	"encoding/json"
)

type ABGroup struct {
	ThirdPartyValidation string `json:"third_party_validation,omitempty"`
	FacebookEntry        bool   `json:"facebook_entry,omitempty"`
}

func (g ABGroup) Bytes() []byte {
	data, err := json.Marshal(g)
	if err != nil {
		return nil
	}
	return data
}
