package inner

import "time"

type DeviceOs struct {
	Name    string
	Version string
}

type IdType struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Md5PhoneNumber struct {
	Hash8       string
	Hash11      string
	PhoneNumber string
}

type Contact struct {
	Id           string
	UserId       string
	Name         string
	PhoneNumbers []Md5PhoneNumber
	Relationship *Relationship
	SecretCrush  bool
	Type         string
}

type Relationship struct {
	Id                string
	UserId            string
	OtherUserId       string
	State             string
	OtherState        string
	Category          string
	ScenarioIds       []string
	OtherScenarioIds  []string
	ScenarioMatched   bool
	CreatedTime       time.Time
	UpdatedTime       time.Time
	ClientCreatedTime time.Time
	Type              string
	Additional        RelationshipAdditional
}

type RelationshipAdditional struct {
	Type            string
	SuperLikeQuota  int
	UndoQuota       int
	OldRelationship *Relationship
}
