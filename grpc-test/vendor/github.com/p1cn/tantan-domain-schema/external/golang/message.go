package external

const MessageType = "message"

const (
	ReferenceActionComment = "comment"
	ReferenceActionLike    = "like"
)

type Message struct {
	Id          string            `json:"id"`
	Owner       IdType            `json:"owner"`
	OtherUser   IdType            `json:"otherUser"`
	Accessory   *Accessory        `json:"accessory"`
	Reference   *MessageReference `json:"reference"`
	Location    *MessageLocation  `json:"location"`
	Media       []UserMedia       `json:"media"`
	Recalled    bool              `json:"recalled"`
	Value       string            `json:"value"`
	CreatedTime Iso8601Time       `json:"createdTime"`
	SentFrom    string            `json:"sentFrom"`
	Type        string            `json:"type"`
}

func (self Message) IsText() bool {
	return self.Value != "" &&
		!self.IsSticker() &&
		!self.IsQuestion() &&
		!self.IsAnswer() &&
		!self.IsMomentComment() &&
		!self.IsMomentLike() &&
		!self.IsMedia() &&
		!self.IsLocation()
}

func (self Message) IsSticker() bool {
	return self.Accessory != nil &&
		self.Accessory.Id != "" &&
		self.Accessory.Type == StickerType
}

func (self Message) IsQuestion() bool {
	return self.Accessory != nil &&
		self.Accessory.Type == QuestionType &&
		self.Accessory.Category != nil
}

func (self Message) IsAnswer() bool {
	return self.Reference != nil &&
		self.Reference.Id != "" &&
		self.Reference.Type == MessageType &&
		(self.Value != "" || self.IsSticker() || self.IsMedia())
}

func (self Message) IsMomentComment() bool {
	return self.Reference != nil &&
		self.Reference.Id != "" &&
		self.Reference.Type == MomentType &&
		self.Value != ""
}

func (self Message) IsMomentLike() bool {
	return self.Reference != nil &&
		self.Reference.Id != "" &&
		self.Reference.Type == MomentType &&
		self.Value == ""
}

func (self Message) IsMedia() bool {
	return len(self.Media) > 0
}

func (self Message) IsLocation() bool {
	return self.Location != nil &&
		len(self.Location.Coordinates) >= 2
}

func (self Message) ResetAccessory() {
	self.Accessory = nil
}

func (self Message) ResetReference() {
	self.Reference = nil
}

type Messages []Message

func (self Messages) GetReferencedUserIds() []string {
	userIds := make([]string, 0)
	foundUsersMap := make(map[string]bool)
	for _, m := range self {
		if !foundUsersMap[m.Owner.Id] {
			userIds = append(userIds, m.Owner.Id)
			foundUsersMap[m.Owner.Id] = true
		}
		if !foundUsersMap[m.OtherUser.Id] {
			userIds = append(userIds, m.OtherUser.Id)
			foundUsersMap[m.Owner.Id] = true
		}
	}
	return userIds
}

func (self Messages) GetReferencedStickerIds() []string {
	ids := make([]string, 0)
	for _, m := range self {
		if m.Accessory != nil && m.Accessory.Type == StickerType && len(m.Accessory.Id) > 0 {
			ids = append(ids, m.Accessory.Id)
		}
	}
	return ids
}

func (self Messages) GetReferencedQuestionIds() []string {
	ids := make([]string, 0)
	for _, m := range self {
		if m.Accessory != nil && m.Accessory.Type == QuestionType && len(m.Accessory.Id) > 0 {
			ids = append(ids, m.Accessory.Id)
		}
	}
	return ids
}
func IsBrandAccount(userId string) bool {
	var TeamAccountUserId string = "-1"
	// If need add some other brand accout, add here

	return userId == TeamAccountUserId
}
func (self Messages) GetAllReferencedMomentIds() []string {
	ids := make([]string, 0)
	idsFound := make(map[string]bool)
	for _, m := range self {
		if m.Reference != nil && m.Reference.Type == MomentType && len(m.Reference.Id) > 0 && !idsFound[m.Reference.Id] {
			ids = append(ids, m.Reference.Id)
			idsFound[m.Reference.Id] = true
		}
	}
	return ids
}

//func (self Messages) GetReferencedBrandAccountMomentIds() []string {
//	ids := make([]string, 0)
//	idsFound := make(map[string]bool)
//	for _, m := range self {
//		if IsBrandAccount(m.OtherUser.Id) && m.Reference != nil && m.Reference.Type == MomentType && len(m.Reference.Id) > 0 && !idsFound[m.Reference.Id] {
//			ids = append(ids, m.Reference.Id)
//			idsFound[m.Reference.Id] = true
//		}
//	}
//	return ids
//}
