package external

import (
	"time"
)

const UserLinkType = "link"

type UserLink struct {
	Id          string     `json:"id"`
	Intent      string     `json:"intent"`
	Href        string     `json:"href"`
	Owner       *IdType    `json:"owner"`
	Resources   IdTypes    `json:"resources"`
	State       string     `json:"state"`
	Channel     string     `json:"channel"`
	CreatedTime time.Time  `json:"created_time"`
	Title       string     `json:"title"`
	TestCase    string     `json:"-"`
	ExpiresTime *time.Time `json:"-"`
	Type        string     `json:"type"`
}

func (self UserLink) IsExpired() bool {
	if self.ExpiresTime != nil {
		return self.ExpiresTime.Before(time.Now().UTC())
	} else {
		return false
	}
}
