package external

import (
	"encoding/json"
)

const (
	TokenType = "Bearer"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	UserId      string `json:"user_id"`
}

func (self AccessToken) Bytes() []byte {
	data, err := json.Marshal(self)
	if err != nil {
		return nil
	}

	return data
}
