package external

const TbhCounterType = "tbh_counter"

type TbhCounter struct {
	UserId              string
	GivenFriendships    int
	ReceivedFriendships int
	TotalFriends        int
	ReceivedVotedPolls  int
	ReceivedVotes       int
	GivenVotes          int
}
