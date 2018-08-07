package external

// import (
// 	"encoding/json"

// 	slog "github.com/p1cn/tantan-backend-common/log"
// 	"github.com/p1cn/tantan-backend-common/util"
// 	domain_external "github.com/p1cn/tantan-domain-schema/external/golang"
// )

// type Relationship struct {
// 	Id                string                      `json:"id"`
// 	Owner             domain_external.IdType      `json:"-"`
// 	OtherUser         domain_external.IdType      `json:"-"`
// 	State             string                      `json:"state"`
// 	OtherState        string                      `json:"-"`
// 	Category          string                      `json:"-"`
// 	Scenarios         []domain_external.IdType    `json:"scenarios,omitempty"`
// 	Status            []string                    `json:"status"`
// 	CreatedTime       domain_external.Iso8601Time `json:"createdTime"`
// 	UpdatedTime       domain_external.Iso8601Time `json:"updateTime"`
// 	ClientCreatedTime domain_external.Iso8601Time `json:"-"`
// 	Type              string                      `json:"type"`
// 	Additional        RelationshipAdditional      `json:"-"`
// }

// type RelationshipAdditional struct {
// 	Type            string
// 	SuperLikeQuota  int
// 	UndoQuota       int
// 	OldRelationship *Relationship
// }

// type Friendship struct {
// 	Id          string                      `json:"id"`
// 	Owner       domain_external.IdType      `json:"-"`
// 	OtherUser   domain_external.IdType      `json:"-"`
// 	State       string                      `json:"state"`
// 	Friendship  bool                        `json:"friendship"`
// 	OtherState  string                      `json:"-"`
// 	CreatedTime domain_external.Iso8601Time `json:"createdTime"`
// 	UpdatedTime domain_external.Iso8601Time `json:"updateTime"`
// 	Type        string                      `json:"type"`
// }

// type Md5PhoneNumber struct {
// 	Hash8       string
// 	Hash11      string
// 	PhoneNumber string
// }

// type Contact struct {
// 	Id           string
// 	UserId       string
// 	Name         string
// 	PhoneNumbers []Md5PhoneNumber
// 	Relationship *Relationship
// 	SecretCrush  bool
// 	Type         string
// }

// // BoostSettings boost状态设置
// // active: GET时表示，用户是否处于boost状态，PATCH时表示是否开启boost
// // Note boost GET和PATCH时所需字段不一致，详情参考：
// // https://github.com/p1cn/backend/wiki/vip-redesign#user-object
// type BoostSettings struct {
// 	Active      bool                        `json:"active"`
// 	Multiplier  int16                       `json:"multiplier"`
// 	Duration    int32                       `json:"duration"`
// 	ExpiresTime domain_external.Iso8601Time `json:"expiresTime"`
// 	Identifier  string                      `json:"identifier,omitempty"`
// }

// const (
// 	FollowshipType = "followship"

// 	FollowshipStateDefault   = "default"
// 	FollowshipStateFollowing = "following"
// 	FollowshipStatefollowed  = "followed"
// 	FollowshipStateMatched   = "matched"
// 	FollowshipStateUnfollow  = "unfollow"
// )

// type Followship struct {
// 	Id            string                      `json:"id"`
// 	Owner         domain_external.IdType      `json:"owner"`
// 	OtherUser     domain_external.IdType      `json:"otherUser"`
// 	State         string                      `json:"state"`
// 	Status        []string                    `json:"-"`
// 	OtherState    string                      `json:"-"`
// 	UserTime      domain_external.Iso8601Time `json:"userTime"`
// 	OtherUserTime domain_external.Iso8601Time `json:"otherUserTime"`
// 	Type          string                      `json:"type"`
// }

// func (self Followship) String() string {
// 	jsonSrc, err := json.Marshal(self)
// 	if err != nil {
// 		slog.Notice("%+v", err)
// 		return ""
// 	}
// 	return string(jsonSrc)
// }

// func (self Followship) HasValidState() bool {
// 	return util.InSlice(self.State, []string{FollowshipStateDefault, FollowshipStateFollowing, FollowshipStateUnfollow})
// }

// type Followships []Followship

// func (self Followships) GetReferencedUserIds() []string {
// 	ids := make([]string, 0)
// 	foundUsersMap := make(map[string]bool)
// 	for _, followship := range self {
// 		if !foundUsersMap[followship.Owner.Id] {
// 			ids = append(ids, followship.Owner.Id)
// 			foundUsersMap[followship.Owner.Id] = true
// 		}
// 		if !foundUsersMap[followship.OtherUser.Id] {
// 			ids = append(ids, followship.OtherUser.Id)
// 			foundUsersMap[followship.OtherUser.Id] = true
// 		}
// 	}
// 	return ids
// }
