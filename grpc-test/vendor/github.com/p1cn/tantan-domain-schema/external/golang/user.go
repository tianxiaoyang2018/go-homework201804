package external

import (
	"fmt"
	"time"
)

const (
	UserType = "user"

	UserStatusHidden      = "hidden"
	UserStatusTeamAccount = "teamaccount"
	UserStatusBrand       = "brand"
	UserStatusBoosted     = "boosted"
	// used for remind low popularity male user ab test
	UserStatusMediumPopularity = "mediumPopularity"
	UserStatusLowPopularity    = "lowPopularity"
	THRLD_LOW                  = 0.03 // the low threshold of judge user popularity
	THRLD_MEDIUM               = 0.06 // the medium threshold of judge user popularity
	THRLD_SWIPED_TIMES         = 500  // the times user has been swiped that the user's popularity in user counters has meaning
)

type Date time.Time

func (self Date) IsZero() bool {
	return (time.Time)(self).IsZero()
}

func (self *Date) MarshalJSON() ([]byte, error) {
	return []byte(
		"\"" + (*time.Time)(self).Format("2006-01-02") + "\"",
	), nil
}

func (self *Date) UnmarshalJSON(data []byte) error {
	if string(data) == `""` {
		data = []byte(`"0001-01-01"`)
	}

	aLen := len(data)
	if aLen < 2 {
		return fmt.Errorf("Bad Time")
	}
	t, err := time.Parse("2006-01-02", string(data[1:aLen-1]))
	if err == nil {
		*self = Date(t)
	}
	return err
}

type User struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	RealName    *string `json:"realName,omitempty"`
	Description string  `json:"description"`
	Gender      string  `json:"gender"`
	Location struct {
		Distance    int         `json:"distance"`
		UpdatedTime Iso8601Time `json:"updatedTime"`
		Region      UserRegion  `json:"region"`
		RegionEn    UserRegion  `json:"-"`
		RegionKo    UserRegion  `json:"-"`
		RegionJa    UserRegion  `json:"-"`
		Passby      *UserPassby `json:"passby"`
	} `json:"location"`
	Pictures     []UserMedia      `json:"pictures"`
	Age          int              `json:"age"`
	Settings     *UserSettings    `json:"settings"`
	Membership   UserMembership   `json:"membership"`
	Memberships  []UserMembership `json:"memberships"`
	Profile      *UserProfile     `json:"profile"`
	CreatedTime  Iso8601Time      `json:"createdTime"`
	Status       []string         `json:"status"`
	InnerStatus  string           `json:"-"`
	VerifiedType *string          `json:"-"`
	Type         string           `json:"type"`
	Source       string           `json:"source"`
	TbhActive    bool             `json:"-"`
}

type UserPassby struct {
	Location   PassbyLocation `json:"location"`
	Count      int            `json:"count"`
	LatestTime Iso8601Time    `json:"latestTime"`
}

type PassbyLocation struct {
	Coordinates []float64 `json:"coordinates"`
}

type PhoneNumber struct {
	CountryCode int    `json:"countryCode"`
	Number      string `json:"number"`
}

type NotificationSettings struct {
	PreviewPushMessage bool `json:"previewPushMessage"`
}

type ConversationSettings struct {
	ShowMomentLikes bool `json:"showMomentLikes"`
}

// BoostSettings boost状态设置
// active: GET时表示，用户是否处于boost状态，PATCH时表示是否开启boost
// Note boost GET和PATCH时所需字段不一致，详情参考：
// https://github.com/p1cn/backend/wiki/vip-redesign#user-object
type BoostSettings struct {
	Active      bool        `json:"active"`
	Multiplier  int16       `json:"multiplier"`
	Duration    int32       `json:"duration"`
	ExpiresTime Iso8601Time `json:"expiresTime"`
	Identifier  string      `json:"identifier,omitempty"`
}

type TbhStudies struct {
	Id    string `json:"id"`
	Grade string `json:"grade"`
	Type  string `json:"type"`
}

type TbhSetting struct {
	Studies TbhStudies `json:"studies"`
}

type UserSettings struct {
	LookingFor         string `json:"lookingFor"`
	Intent             string `json:"-"`
	HideContacts       bool   `json:"hideContacts"`
	HideMutualContacts bool   `json:"hideMutualContacts"`
	SearchRadius struct {
		Value          int `json:"value"`
		AllowedMaximum int `json:"allowedMaximum"`
		AllowedMinimum int `json:"allowedMinimum"`
	} `json:"searchRadius"`
	SearchLocation struct {
		Coordinates []float64 `json:"coordinates"`
	} `json:"searchLocation"`
	SuggestSearchRadius     int     `json:"-"`
	AutoAdjustSuggestRadius *bool   `json:"autoAdjustSuggestRadius"`
	Longitude               float64 `json:"-"`
	Latitude                float64 `json:"-"`
	SearchAge struct {
		AllowedMinimum int `json:"allowedMinimum"`
		AllowedMaximum int `json:"allowedMaximum"`
		Minimum        int `json:"minimum"`
		Maximum        int `json:"maximum"`
	} `json:"searchAge"`
	Birthdate        Date                 `json:"birthdate"`
	Email            string               `json:"email"`
	PhoneNumber      PhoneNumber          `json:"phoneNumber"`
	Notifications    NotificationSettings `json:"notifications"`
	Conversations    ConversationSettings `json:"conversations"`
	Boost            BoostSettings        `json:"boost"`
	Verification     *Verification        `json:"verification"`
	SearchPriorities []string             `json:"searchPriorities"`
	Tbh              *TbhSetting          `json:"tbh,omitempty"`
	Moment struct {
		HidePublicMoments bool `json:"hidePublicMoments"`
	} `json:"moment"`
}

func (s *UserSettings) AdjustSearchAgeForYoung() {
	s.SearchAge.AllowedMaximum = 17
	if s.SearchAge.Maximum > s.SearchAge.AllowedMaximum {
		s.SearchAge.Maximum = s.SearchAge.AllowedMaximum
		if s.SearchAge.Minimum > s.SearchAge.Maximum {
			s.SearchAge.Minimum = s.SearchAge.Maximum
		}
	}
}

type UserMembership struct {
	Name        *string              `json:"name"`
	Duration    int64                `json:"duration"`
	Active      bool                 `json:"active"`
	ExpiresTime *Iso8601Time         `json:"expiresTime,omitempty"`
	Links       *UserMembershipLinks `json:"links,omitempty"`
}

type UserMembershipLinks struct {
	Renew   string `json:"renew,omitempty"`
	Current string `json:"current,omitempty"`
	Upgrade string `json:"upgrade,omitempty"`
}

type UserProfileTag struct {
	Value    string `json:"value"`
	Category string `json:"category"`
}

type UserProfileSocialAccount struct {
	Value   string `json:"value"`
	Network string `json:"network"`
}

type UserProfileAnswer struct {
	Answer   string `json:"value"`
	Question IdType `json:"question"`
}

type UserProfileMoments struct {
	//Count int      `json:"count"`
	Ids   []string `json:"ids"`
	Links Links    `json:"links"`
}

type MutualContacts struct {
	//Count int      `json:"count"`
	Ids   []string `json:"ids"`
	Links Links    `json:"links"`
}

type UserProfileWork struct {
	Industry   string `json:"industry"`
	Company    string `json:"company"`
	Department string `json:"department"`
	Active     bool   `json:"active"`
}

type UserProfileStudy struct {
	School   string  `json:"school,omitempty"`
	Major    string  `json:"major"`
	Active   bool    `json:"active"`
	Verified bool    `json:"verified"`
	Grade    *string `json:"grade,omitempty"`
}

type UserProfileTbh struct {
	Friends            int         `json:"friends"`
	ReceivedVotes      int         `json:"receivedVotes"`
	ReceivedVotedPolls int         `json:"receivedVotedPolls"`
	CreatedTime        Iso8601Time `json:"createdTime"`
}

type UserProfile struct {
	ReceivedLikes       *int                       `json:"receivedLikes"`
	ReceivedLikesRank   string                     `json:"receivedLikesRank"` //只有android 2.6.1~2.6.3版本使用,默认为空即可
	Occupation          string                     `json:"occupation"`
	School              string                     `json:"school,omitempty"`
	Job                 string                     `json:"job"`
	Zodiac              string                     `json:"zodiac"`
	Work                *UserProfileWork           `json:"work"`
	Studies             *UserProfileStudy          `json:"studies"`
	Hometown            string                     `json:"hometown"`
	Hangouts            string                     `json:"hangouts"`
	Tags                []UserProfileTag           `json:"tags"`
	SocialAccounts      []UserProfileSocialAccount `json:"social"`
	Answers             []UserProfileAnswer        `json:"answers"`
	Scenarios           []IdType                   `json:"scenarios"`
	ScenarioCategory    *string                    `json:"-"`
	ScenarioExpiresTime *Iso8601Time               `json:"-"`
	MutualContacts      MutualContacts             `json:"mutualContacts"`
	Moments             UserProfileMoments         `json:"publicMoments"`
	Tbh                 *UserProfileTbh            `json:"tbh,omitempty"`
}

type Users []User

func (self Users) Ids() []string {
	ids := make([]string, 0)
	for _, u := range self {
		ids = append(ids, u.Id)
	}
	return ids
}

func (self Users) IdsExcept(exceptId string) []string {
	ids := make([]string, 0)
	for _, u := range self {
		if u.Id != exceptId {
			ids = append(ids, u.Id)
		}
	}
	return ids
}

func (self Users) GetReferencedQuestionIds() []string {
	ids := make([]string, 0)
	for _, u := range self {
		ids = append(ids, UserProfileAnswers(u.Profile.Answers).QuestionIds()...)
	}
	return ids
}

func (self Users) GetReferencedContactIds() []string {
	ids := make([]string, 0)
	for _, u := range self {
		ids = append(ids, u.Profile.MutualContacts.Ids...)
	}
	return ids
}

func (self Users) GetReferencedScenarioIds() []string {
	ids := make([]string, 0)
	for _, u := range self {
		ids = append(ids, IdTypes(u.Profile.Scenarios).Ids()...)
	}
	return ids
}

func (self Users) FillReferencedMoments(moments Moments) {
	usersMap := map[string]*[]string{}
	for i := 0; i < len(self); i++ {
		if self[i].Profile == nil {
			self[i].Profile = &UserProfile{
				Moments: UserProfileMoments{
					Ids: []string{},
				},
			}
		}

		usersMap[self[i].Id] = &self[i].Profile.Moments.Ids
	}
	for _, moment := range moments {
		*usersMap[moment.Owner.Id] = append(*usersMap[moment.Owner.Id], moment.Id)
	}
}

func (self *User) IsStudent() bool {
	if self.Profile == nil {
		return false
	}
	if self.Profile.Studies == nil {
		return false
	}

	if self.Profile.Studies.Active {
		return true
	}
	return false
}

func (self *User) IsStudentAge(minAge int, maxAge int) bool {
	return self.Age >= minAge && self.Age <= maxAge
}

type UserProfileAnswers []UserProfileAnswer

func (self UserProfileAnswers) QuestionIds() []string {
	ids := make([]string, 0)
	for _, a := range self {
		ids = append(ids, a.Question.Id)
	}
	return ids
}
func (self Users) FetchUserIdsFromUsers() []string {
	result := []string{}
	for i, _ := range self {
		if !self[i].Settings.Moment.HidePublicMoments {
			result = append(result, self[i].Id)
		}
	}
	return result
}
