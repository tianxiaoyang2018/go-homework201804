package external

import "time"

type ScenarioRegionCounter struct {
	Id          string
	RegionType  string
	RegionId    string
	Gender      string
	ScenarioId  string
	ActiveUsers int64
	CreatedTime time.Time
	UpdatedTime time.Time
}
