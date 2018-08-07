package external

import (
	"encoding/json"
)

const (
	FollowshipType = "followship"

	FollowshipStateDefault   = "default"
	FollowshipStateFollowing = "following"
	FollowshipStatefollowed  = "followed"
	FollowshipStateMatched   = "matched"
	FollowshipStateUnfollow  = "unfollow"
)

type Followship struct {
	Id            string      `json:"id"`
	Owner         IdType      `json:"owner"`
	OtherUser     IdType      `json:"otherUser"`
	State         string      `json:"state"`
	Status        []string    `json:"-"`
	OtherState    string      `json:"-"`
	UserTime      Iso8601Time `json:"userTime"`
	OtherUserTime Iso8601Time `json:"otherUserTime"`
	Type          string      `json:"type"`
}

func (self Followship) String() string {
	jsonSrc, err := json.Marshal(self)
	if err != nil {
		return ""
	}
	return string(jsonSrc)
}

func (self Followship) HasValidState() bool {
	return InSlice(self.State, []string{FollowshipStateDefault, FollowshipStateFollowing, FollowshipStateUnfollow})
}

type Followships []Followship

func (self Followships) GetReferencedUserIds() []string {
	ids := make([]string, 0)
	foundUsersMap := make(map[string]bool)
	for _, followship := range self {
		if !foundUsersMap[followship.Owner.Id] {
			ids = append(ids, followship.Owner.Id)
			foundUsersMap[followship.Owner.Id] = true
		}
		if !foundUsersMap[followship.OtherUser.Id] {
			ids = append(ids, followship.OtherUser.Id)
			foundUsersMap[followship.OtherUser.Id] = true
		}
	}
	return ids
}
