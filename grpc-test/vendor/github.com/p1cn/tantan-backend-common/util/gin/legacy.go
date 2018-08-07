package gin

type ResourceType int

const (
	GlobalMaxLimit     = 1000
	GlobalDefaultLimit = 100

	ResourceTypeVoid ResourceType = 1 << iota
	ResourceTypeUser
	ResourceTypeRelationship
	ResourceTypeDevice
	ResourceTypeConversation
	ResourceTypeMessage
	ResourceTypeContact
	ResourceTypeReport
	ResourceTypeQuestion
	ResourceTypeShop
	ResourceTypePackage
	ResourceTypeBundle
	ResourceTypeSticker
	ResourceTypeCampaign
	ResourceTypePoke
	ResourceTypeUserLink
	ResourceTypeMoment
	ResourceTypeMomentComment
	ResourceTypeLike
	ResourceTypeActivity
	ResourceTypeUserFeedSetting
	ResourceTypeMuted
	ResourceTypeHidden
	ResourceTypeFeed
	ResourceTypePassBy
	ResourceTypeLocation
	ResourceTypeScenario
	ResourceTypeBuildInfo
	ResourceTypeItem
	ResourceTypeProduct
	ResourceTypeUserPublicMoments
	ResourceTypeSchool
	ResourceTypeFriendship
	ResourceTypePoll
	ResourceTypeGem
)

var (
	allowedResources = map[string]ResourceType{
		"void":               ResourceTypeVoid,
		"users":              ResourceTypeUser,
		"relationships":      ResourceTypeRelationship,
		"devices":            ResourceTypeDevice,
		"conversations":      ResourceTypeConversation,
		"messages":           ResourceTypeMessage,
		"contacts":           ResourceTypeContact,
		"reports":            ResourceTypeReport,
		"questions":          ResourceTypeQuestion,
		"shops":              ResourceTypeShop,
		"packages":           ResourceTypePackage,
		"bundles":            ResourceTypeBundle,
		"stickers":           ResourceTypeSticker,
		"campaigns":          ResourceTypeCampaign,
		"pokes":              ResourceTypePoke,
		"links":              ResourceTypeUserLink,
		"moments":            ResourceTypeMoment,
		"likes":              ResourceTypeLike,
		"activities":         ResourceTypeActivity,
		"muted":              ResourceTypeMuted,
		"hidden":             ResourceTypeHidden,
		"passbys":            ResourceTypePassBy,
		"locations":          ResourceTypeLocation,
		"scenarios":          ResourceTypeScenario,
		"ad-data":            ResourceTypeBuildInfo,
		"items":              ResourceTypeItem,
		"products":           ResourceTypeProduct,
		"user.publicMoments": ResourceTypeUserPublicMoments,
		"schools":            ResourceTypeSchool,
		"friendships":        ResourceTypeFriendship,
		"polls":              ResourceTypePoll,
		"gems":               ResourceTypeGem,
	}
	resourcesMaxLimit = map[ResourceType]int{
		ResourceTypeConversation: 10000,
	}
	resourcesRecursiveRequestMaxLimit = map[ResourceType]int{
		ResourceTypeConversation: 20,
	}
	resourcesSortMaxLimit = map[ResourceType]int{
		ResourceTypeConversation: 100,
	}
	resourcesDefaultLimit = map[ResourceType]int{
		ResourceTypeConversation: 20,
	}
)

type Resource struct {
	Id     *string
	Type   ResourceType
	Parent *Resource
}

func (self *Resource) RecursiveRequestMaxLimit() int {
	if self == nil {
		return GlobalMaxLimit
	}
	if limit, ok := resourcesRecursiveRequestMaxLimit[self.Type]; ok {
		return limit
	}
	return GlobalMaxLimit
}

func (self *Resource) SortMaxLimit() int {
	if self == nil {
		return GlobalMaxLimit
	}
	if limit, ok := resourcesSortMaxLimit[self.Type]; ok {
		return limit
	}
	return GlobalMaxLimit
}

func (self *Resource) MaxLimit() int {
	if self == nil {
		return GlobalMaxLimit
	}
	if limit, ok := resourcesMaxLimit[self.Type]; ok {
		return limit
	}
	return GlobalMaxLimit
}

func (self *Resource) DefaultLimit() int {
	if self == nil {
		return GlobalDefaultLimit
	}
	if limit, ok := resourcesDefaultLimit[self.Type]; ok {
		return limit
	}
	return GlobalDefaultLimit
}

func (self *Resource) String() string {
	if self == nil {
		return ""
	}
	suffix := ""
	if self.Id != nil {
		suffix = "/" + *self.Id
	}
	return self.Parent.String() + "/" + self.name() + suffix

}

func (self *Resource) name() string {
	for k, v := range allowedResources {
		if v == self.Type {
			return k
		}
	}
	return "unknown"
}
