package external

const ConversationType = "conversation"

type ConversationMessages struct {
	Count  int      `json:"count"`
	Ids    []string `json:"ids"`
	Links  Links    `json:"links"`
	Unread int      `json:"unread"` // unread message counter
}

type Conversation struct {
	Id                     string               `json:"id"`
	Owner                  IdType               `json:"owner"`
	OtherUser              IdType               `json:"otherUser"`
	OtherUserLastActivity  Iso8601Time          `json:"-"`
	Messages               ConversationMessages `json:"messages"`
	Read                   bool                 `json:"read"`
	ReadUntil              string               `json:"readUntil"`
	ClearedUntil           string               `json:"clearedUntil"`
	AcceptIntimateQuestion bool                 `json:"-"`
	ClearedTime            Iso8601Time          `json:"-"`
	CreatedTime            Iso8601Time          `json:"createdTime"`
	LatestTime             Iso8601Time          `json:"latestTime"`
	Type                   string               `json:"type"`
	UnreadMessages         int                  `json:"unreadMessages"` // unread message counter
}

type Conversations []Conversation

func (self Conversations) GetReferencedUserIds() []string {
	userIds := make([]string, 0)
	foundUsersMap := make(map[string]bool)
	for _, c := range self {
		if !foundUsersMap[c.Owner.Id] {
			userIds = append(userIds, c.Owner.Id)
			foundUsersMap[c.Owner.Id] = true
		}
		if !foundUsersMap[c.OtherUser.Id] {
			userIds = append(userIds, c.OtherUser.Id)
			foundUsersMap[c.Owner.Id] = true
		}
	}
	return userIds
}

func (self Conversations) GetReferencedMessageIds() []string {
	messageIds := make([]string, 0)
	for _, c := range self {
		if len(c.Messages.Ids) > 0 {
			messageIds = append(messageIds, c.Messages.Ids...)
		}
	}
	return messageIds
}

func (self Conversations) GetReferencedRelationshipIds() []string {
	relationshipIds := make([]string, 0)
	for _, c := range self {
		relationshipIds = append(relationshipIds, c.OtherUser.Id)
	}
	return relationshipIds
}
