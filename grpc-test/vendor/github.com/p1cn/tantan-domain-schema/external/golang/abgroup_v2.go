package external

import (
	"time"
)

type ABV2Group struct {
	ID             string
	Name           string
	Type           string
	Status         string
	Services       string
	APIs           []string
	VersionLimit   []string
	ReleaseVersion string
	Business       string
	Desc           string
	StartTime      time.Time
	EndTime        time.Time
}

type ABInfo struct {
	Names    []string `json:"names"`
	ABHeader string   `json:"abHeader"`
}
