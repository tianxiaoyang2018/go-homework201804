package external

const ScenarioType = "scenario"

// scenario categories
const (
	ScenarioFilterFood  = "food"
	ScenarioFilterSport = "sport"
	ScenarioFilterMovie = "movie"
	ScenarioFilterMisc  = "misc"
)

type Scenario struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Activity struct {
		Users       int64       `json:"users"`
		UpdatedTime Iso8601Time `json:"updatedTime"`
	} `json:"activity"`
	Tags        []string    `json:"tags"`
	Media       []UserMedia `json:"media"`
	CreatedTime Iso8601Time `json:"createdTime"`
	Type        string      `json:"type"`
}

type Scenarios []Scenario

func IsScenarioFilterValid(filter string) bool {
	return filter == ScenarioFilterFood ||
		filter == ScenarioFilterSport ||
		filter == ScenarioFilterMovie || filter == ScenarioFilterMisc
}
