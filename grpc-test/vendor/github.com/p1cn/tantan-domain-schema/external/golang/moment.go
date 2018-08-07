package external

const MomentType = "moment"
const (
	MOMENT_SHARE_TYPE_SHARE_TO     = "share_to"
	MOMENT_SHARE_TYPE_NOT_SHARE_TO = "not_share_to"
)

type MomentMessages struct {
	Count int      `json:"count"`
	Limit int      `json:"limit"`
	Ids   []string `json:"ids"`
	Links Links    `json:"links"`
}

type MomentLikes struct {
	Count int      `json:"count"`
	Limit int      `json:"limit"`
	Ids   []string `json:"ids"`
	Links Links    `json:"links"`
}

type MomentLocation struct {
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	Coordinates []float64 `json:"coordinates"`
}

type MomentShareRules struct {
	ShareType    string // "share_to" | "not_share_to"
	Regions      *ShareRuleRegions
	Gender       string // "male" | "female" | "both"
	Age          *ShareRuleAgeRange
	Mod          *ShareRuleMod
	Verification *ShareRuleVerification
	UserStatus   []string
}

type ShareRuleRegions struct {
	Countries []string
	Provinces []string
	Cities    []string
	Districts []string
}

type ShareRuleAgeRange struct {
	Min int
	Max int
}

type ShareRuleMod struct {
	Divisor    int
	Remainders []int
}

type ShareRuleVerification struct {
	Verified  bool
	SchoolIds []string
}

type Moment struct {
	Id                 string            `json:"id"`
	Owner              IdType            `json:"owner"`
	Value              string            `json:"value"`
	Location           *MomentLocation   `json:"location"`
	Messages           MomentMessages    `json:"messages"`
	Likes              MomentLikes       `json:"likes"`
	HaveLiked          bool              `json:"haveLiked"`
	Media              []UserMedia       `json:"media"`
	Type               string            `json:"type"`
	CreatedTime        Iso8601Time       `json:"createdTime"`
	LandingPage        string            `json:"landingPage"`
	PrivateMessageIds  []string          `json:"-"`
	InspiredMessageIds []string          `json:"-"`
	MomentLikedUserIds []string          `json:"-"`
	MomentLikeCount    int               `json:"-"`
	ShareRules         *MomentShareRules `json:"-"`
}

func (self *Moment) HasLocation() bool {
	return self.Location != nil && len(self.Location.Coordinates) >= 2
}

func (self *Moment) HasValidLocation() bool {
	if self.Location != nil && len(self.Location.Coordinates) >= 2 && len(self.Location.Name) > 0 && len(self.Location.Address) > 0 {
		return true
	}
	return false
}

type Moments []Moment

func (self Moments) GetPrivateMessageIdsMap() map[string][]string {
	messageIdsMap := make(map[string][]string)
	for _, m := range self {
		if !IsBrandAccount(m.Owner.Id) && m.PrivateMessageIds != nil && len(m.PrivateMessageIds) > 0 {
			messageIdsMap[m.Id] = m.PrivateMessageIds
		}
	}
	return messageIdsMap
}

func (self Moments) GetBrandAccountPrivateMessageIdsMap() map[string][]string {
	messageIdsMap := make(map[string][]string)
	for _, m := range self {
		if IsBrandAccount(m.Owner.Id) && m.PrivateMessageIds != nil && len(m.PrivateMessageIds) > 0 {
			messageIdsMap[m.Id] = m.PrivateMessageIds
		}
	}
	return messageIdsMap
}

func (self Moments) GetReferencedUserIds() []string {
	userIds := []string{}
	userIdsFound := make(map[string]bool)
	for _, m := range self {
		if m.Owner.Id != "" && !userIdsFound[m.Owner.Id] {
			userIds = append(userIds, m.Owner.Id)
			userIdsFound[m.Owner.Id] = true
		}
		if !IsBrandAccount(m.Owner.Id) {
			for _, uid := range m.Likes.Ids {
				if !userIdsFound[uid] {
					userIds = append(userIds, uid)
					userIdsFound[uid] = true
				}
			}
		}
	}
	return userIds
}

func (self Moments) GetBrandAccountLikeUserIds() []string {
	userIds := []string{}
	userIdsFound := make(map[string]bool)
	for _, m := range self {
		if IsBrandAccount(m.Owner.Id) {
			for _, uid := range m.Likes.Ids {
				if !userIdsFound[uid] {
					userIds = append(userIds, uid)
					userIdsFound[uid] = true
				}
			}
		}

	}
	return userIds
}

func (self Moments) GetBrandInspiredMessageIds() map[string][]string {
	messageIdsMap := make(map[string][]string)
	for _, m := range self {

		if IsBrandAccount(m.Owner.Id) && m.InspiredMessageIds != nil && len(m.InspiredMessageIds) > 0 {

			if val, ok := messageIdsMap[m.Id]; !ok {
				messageIdsMap[m.Id] = m.InspiredMessageIds
			} else {
				val = append(val, m.InspiredMessageIds...)
				messageIdsMap[m.Id] = val
			}

		}
	}
	return messageIdsMap
}
