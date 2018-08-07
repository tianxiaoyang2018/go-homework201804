package external

const (
	SchoolType = "school"

	SchoolLevelJunior     = "junior-middle-school"
	SchoolLevelSenior     = "senior-middle-school"
	SchoolLevelTechnical  = "technical-school"
	SchoolLevelUniversity = "university"

	SchoolIdOthers = "1"
)

type School struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	NameZh     string     `json:"-"`
	NameEn     string     `json:"-"`
	Type       string     `json:"type"`
	Level      string     `json:"level"`
	CountryId  string     `json:"-"`
	ProvinceId string     `json:"-"`
	CityId     string     `json:"-"`
	DistrictId string     `json:"-"`
	Region     UserRegion `json:"region"`
}
