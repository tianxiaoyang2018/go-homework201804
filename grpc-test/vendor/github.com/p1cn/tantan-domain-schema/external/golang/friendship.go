package external

const (
	FriendshipType = "friendship"

	FriendshipStateDefault  = "default"
	FriendshipStateOutgoing = "outgoing"
	FriendshipStateIncoming = "incoming"
	FriendshipStateAccepted = "accepted"
)

type Friendship struct {
	Id          string      `json:"id"`
	Owner       IdType      `json:"-"`
	OtherUser   IdType      `json:"-"`
	State       string      `json:"state"`
	Friendship  bool        `json:"friendship"`
	OtherState  string      `json:"-"`
	CreatedTime Iso8601Time `json:"createdTime"`
	UpdatedTime Iso8601Time `json:"updateTime"`
	Type        string      `json:"type"`
}

func (self Friendship) HasValidInputState() bool {
	return InSlice(self.State, []string{FriendshipStateDefault, FriendshipStateOutgoing, FriendshipStateAccepted})
}

func (self Friendship) HasValidState() bool {
	return InSlice(self.State, []string{FriendshipStateDefault, FriendshipStateOutgoing, FriendshipStateIncoming, FriendshipStateAccepted})
}

type Friendships []Friendship

func (self Friendships) GetReferencedUserIds() []string {
	ids := make([]string, 0)
	foundUsersMap := make(map[string]bool)
	for _, friendship := range self {
		if !foundUsersMap[friendship.Id] {
			ids = append(ids, friendship.Owner.Id)
			foundUsersMap[friendship.Owner.Id] = true
		}
		if !foundUsersMap[friendship.OtherUser.Id] {
			ids = append(ids, friendship.OtherUser.Id)
			foundUsersMap[friendship.OtherUser.Id] = true
		}
	}
	return ids
}
