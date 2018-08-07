package external

const (
	RelationshipType = "relationship"

	RelationshipStateDefault  = "default"
	RelationshipStateLiked    = "liked"
	RelationshipStateDisliked = "disliked"
	RelationshipStateBlocked  = "blocked"
	RelationshipStateMatched  = "matched"

	RelationshipStatusSecretCrush                  = "secretcrush"
	RelationshipStatusScenario                     = "scenario"
	RelationshipStatusSuperLike                    = "superLiked"
	RelationshipStatusReceivedLikeDuringBoost      = "boosted"
	RelationshipStatusOtherBoosting                = "otherBoosting"
	RelationshipStatusReceivedLikeDuringBoostBadge = "boostBadge"
	RelationshipStatusOtherBoostBadge              = "otherBoostBadge"
)

const (
	BoostStatusBoost      uint64 = 1 << 0
	BoostStatusBoostBadge uint64 = 1 << 1
)

const (
	AdditionalSuperLike     = "superlike"
	AdditionalUndo          = "undo"
	AdditionalUndoSuperLike = "undoSuperlike"
	AdditionalLike          = "like"
	AdditionalScenario      = "scenario"
)

type Relationship struct {
	Id                string                 `json:"id"`
	Owner             IdType                 `json:"-"`
	OtherUser         IdType                 `json:"-"`
	State             string                 `json:"state"`
	OtherState        string                 `json:"-"`
	Category          string                 `json:"-"`
	Scenarios         []IdType               `json:"scenarios,omitempty"`
	Status            []string               `json:"status"`
	CreatedTime       Iso8601Time            `json:"createdTime"`
	UpdatedTime       Iso8601Time            `json:"updateTime"`
	ClientCreatedTime Iso8601Time            `json:"-"`
	Type              string                 `json:"type"`
	Additional        RelationshipAdditional `json:"-"`
}

type RelationshipAdditional struct {
	Type            string
	SuperLikeQuota  int
	UndoQuota       int
	OldRelationship *Relationship
}

type SuperLikeRelationship struct {
	Id          string
	UserId      string
	OtherUserId string
	State       string
	OtherState  string
}

func (self Relationship) IsScenarioSwipe() bool {
	return len(self.Scenarios) > 0 &&
		len(self.Status) > 0 &&
		self.Status[0] == ScenarioType
}

func (self Relationship) IsSuperLikeSwipe() bool {
	return InSlice(RelationshipStatusSuperLike, self.Status)
}

func (self Relationship) IsUndo() bool {
	return self.State == RelationshipStateDefault
}

func (self Relationship) HasValidInputState() bool {
	return self.State == RelationshipStateLiked ||
		self.State == RelationshipStateDisliked ||
		self.State == RelationshipStateBlocked ||
		self.State == RelationshipStateDefault
}

func (self Relationship) HasValidState() bool {
	return self.State == RelationshipStateLiked ||
		self.State == RelationshipStateDisliked ||
		self.State == RelationshipStateMatched ||
		self.State == RelationshipStateBlocked
}

func (self Relationship) HasValidScenarioState() bool {
	return self.State == RelationshipStateLiked ||
		self.State == RelationshipStateDisliked ||
		self.State == RelationshipStateMatched
}

func (self Relationship) HasValidInputStatus() bool {
	for _, status := range self.Status {
		if status != RelationshipStatusSecretCrush &&
			status != RelationshipStatusScenario &&
			status != RelationshipStatusSuperLike {
			return false
		}
	}
	return true
}

type Relationships []Relationship

func (self Relationships) GetReferencedScenarioIds() []string {
	ids := make([]string, 0)
	for _, r := range self {
		ids = append(ids, IdTypes(r.Scenarios).Ids()...)
	}
	return ids
}

// GetReferenceUserIds Get relative user ids for the specific relationships
func (r Relationships) GetReferenceUserIds() []string {
	uids := make([]string, len(r))

	for i, record := range r {
		uids[i] = record.Id
	}
	return uids
}
