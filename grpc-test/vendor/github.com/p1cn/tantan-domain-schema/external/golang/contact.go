package external

const ContactType = "contact"

type Md5PhoneNumber struct {
	Hash8       string `json:"hash8"`
	Hash11      string `json:"hash11"`
	PhoneNumber string `json:"number"`
}

type Contact struct {
	Id           string           `json:"id"`
	Owner        IdType           `json:"owner"`
	Name         string           `json:"name"`
	PhoneNumbers []Md5PhoneNumber `json:"phoneNumbers"`
	SecretCrush  bool             `json:"secretCrush"`
	Match        *IdType          `json:"match"`
	Type         string           `json:"type"`
}

type Contacts []Contact

func (self Contacts) GetReferencedUserIds() []string {
	userIds := make([]string, 0)
	foundUsersMap := make(map[string]bool)
	for _, c := range self {
		if !foundUsersMap[c.Owner.Id] {
			userIds = append(userIds, c.Owner.Id)
			foundUsersMap[c.Owner.Id] = true
		}
		if c.Match == nil {
			continue
		}
		if !foundUsersMap[c.Match.Id] {
			userIds = append(userIds, c.Match.Id)
			foundUsersMap[c.Match.Id] = true
		}
	}

	return userIds
}
