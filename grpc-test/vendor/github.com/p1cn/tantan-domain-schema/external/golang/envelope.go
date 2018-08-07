package external

type Error struct {
	Message string  `json:"message"`
	Field   *string `json:"field,omitempty"`
	Code    *string `json:"code,omitempty"`
}

type Meta struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Errors  []Error `json:"errors,omitempty"`
}

type Counters struct {
	Conversations  ConversationCounters `json:"conversations"`
	Messages       MessageCounters      `json:"messages"`
	LikeLimit      LikeLimit            `json:"likeLimit"`
	SuperLikeLimit SuperLikeLimit       `json:"superLikeLimit"`
	UndoLimit      UndoLikeLimit        `json:"undoLimit"`
	Activities     ActivityCounters     `json:"activities"`
	SecretCrush    SecretCrushCounters  `json:"secretCrushLimit"`
	Scenarios      ScenarioCounters     `json:"scenarios"`
	PopularInfo    *PopularInfo         `json:"-"`
}

type PopularInfo struct {
	ReceivedLikes    int     `json:"-"`
	ReceivedDislikes int     `json:"-"`
	Popularity       float64 `json:"-"`
}

type LikeLimit struct {
	Remaining int `json:"remaining"`
	Total     int `json:"total"`
	Reset     int `json:"reset"`
}

type SuperLikeLimit struct {
	Remaining       int  `json:"remaining"`
	RewardForInvite *int `json:"rewardForInvite,omitempty"`
	RewardForShare  *int `json:"rewardForShare,omitempty"`
	Quota           int  `json:"quota"`
	Reset           int  `json:"reset"`
	ResetShare      int  `json:"resetShare"`
	Count           int  `json:"count"`
	Limit           int  `json:"limit"`
}

type UndoLikeLimit struct {
	Remaining       int  `json:"remaining"`
	RewardForInvite *int `json:"rewardForInvite,omitempty"`
	RewardForShare  *int `json:"rewardForShare,omitempty"`
	Quota           int  `json:"quota"`
	Reset           int  `json:"reset"`
	ResetShare      int  `json:"resetShare"`
	Count           int  `json:"count"`
	Limit           int  `json:"limit"`
}

type ConversationCounters struct {
	Total  int64 `json:"total"`
	Unread int64 `json:"unread"`
	Unseen int64 `json:"unseen"`
}

type MessageCounters struct {
	Unread int64 `json:"unread"`
}

type ActivityCounters struct {
	Unread int `json:"unread"`
}

type SecretCrushCounters struct {
	Total     int `json:"total"`
	Remaining int `json:"remaining"`
	Received  int `json:"received"`
}

type ScenarioCounters struct {
	Food  ScenarioCounter `json:"food"`
	Sport ScenarioCounter `json:"sport"`
	Movie ScenarioCounter `json:"movie"`
	Misc  ScenarioCounter `json:"misc"`
}

type ScenarioCounter struct {
	ReceivedLikes int                `json:"receivedLikes"`
	Duration      int                `json:"duration"`
	Reset         int                `json:"reset"`
	MatchLimit    ScenarioMatchLimit `json:"matchLimit"`
}

type ScenarioMatchLimit struct {
	Remaining int `json:"remaining"`
	Total     int `json:"total"`
	Reset     int `json:"reset"`
}

type Objects struct {
	Users         []User         `json:"users"`
	Relationships []Relationship `json:"relationships"`
	Devices       []Device       `json:"devices"`
	Conversations []Conversation `json:"conversations"`
	Messages      []Message      `json:"messages"`
	Contacts      []Contact      `json:"contacts"`
	Reports       []Report       `json:"reports"`
	Questions     []Question     `json:"questions"`
	Media         []Media        `json:"media"`
	Shops         []Shop         `json:"shops"`
	Stickers      []Sticker      `json:"stickers"`
	Packages      []Package      `json:"packages"`
	Bundles       []Bundle       `json:"bundles"`
	Campaigns     []Campaign     `json:"campaigns"`
	Pokes         []Poke         `json:"pokes"`
	UserLinks     []UserLink     `json:"links"`
	Moments       []Moment       `json:"moments"`
	Activities    []Activity     `json:"activities"`
	PassBys       []PassByCount  `json:"passbys"`
	Scenarios     []Scenario     `json:"scenarios"`
	Items         []Item         `json:"items"`
	Products      []Product      `json:"products"`
	Schools       []School       `json:"schools"`
	Friendships   []Friendship   `json:"friendships"`
	Polls         []Poll         `json:"polls"`
	Gems          []Gem          `json:"gems"`
}

type Links struct {
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
}

type Pagination struct {
	Total int64 `json:"total"`
	Limit int   `json:"limit"`
	Links Links `json:"links"`
}

type Envelope struct {
	Meta       Meta       `json:"meta"`
	Counters   *Counters  `json:"counters"`
	Data       Objects    `json:"data"`
	Pagination Pagination `json:"pagination"`
}
