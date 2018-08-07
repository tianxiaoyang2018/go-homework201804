package external

import ()

const (
	ReportType = "report"
)

type Report struct {
	UserId            string      `json:"-"`
	ContentId         string      `json:"-"`
	ContentType       string      `json:"-"`
	ReportedBy        string      `json:"-"`
	Category          string      `json:"category"`
	Value             string      `json:"value"`
	CreatedTime       Iso8601Time `json:"createdTime"`
	Type              string      `json:"type"`
	Pictures          []UserMedia `json:"pictures"`
	UserPicturesCount int         `json:"user_pictures_count"`
}
