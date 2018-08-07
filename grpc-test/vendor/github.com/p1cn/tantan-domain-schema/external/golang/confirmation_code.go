package external

import (
	"encoding/json"
)

type ConfirmationCode struct {
	Id           string      `json:"id"`
	UserId       string      `json:"user_id"`
	Action       string      `json:"action"`
	CountryCode  int         `json:"country_code"`
	MobileNumber string      `json:"mobile_number"`
	Code         int         `json:"code"`
	CreatedTime  Iso8601Time `json:"created_time"`
	ExpiresTime  Iso8601Time `json:"expires_time"`
}

func (self *ConfirmationCode) Bytes() []byte {
	data, err := json.Marshal(self)
	if err != nil {
		return nil
	}

	return data
}
