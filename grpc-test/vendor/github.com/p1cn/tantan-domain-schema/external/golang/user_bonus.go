package external

import "time"

type UserBonus struct {
	ID          string
	UserID      string
	ContentID   string
	ContentType string
	Status      string
	UpdatedTime time.Time
	CreatedTime time.Time
}
