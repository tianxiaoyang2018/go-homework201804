package external

const (
	BanEntryRelationship = "relationship"
	BanEntrySignUp       = "signup"
	BanEntryMessage      = "messages"
	BanEntryCoreRequest  = "core-request"
	BanEntrySignIn       = "signin"
)

type BanEntry struct {
	UserID  string
	Handler string
	Reason  string
}

type UserMobileCheckEntry struct {
	UserID string
}
