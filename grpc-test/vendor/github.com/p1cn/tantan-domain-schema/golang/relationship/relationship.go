package relationship

func (r Relationship) GetMatched() bool {
	if (r.State == RelationshipStateEnum_LIKED || r.State == RelationshipStateEnum_SUPERLIKED) &&
		(r.OtherState == RelationshipStateEnum_LIKED || r.OtherState == RelationshipStateEnum_SUPERLIKED) {
		return true
	}
	return false
}

func (r Relationship) EitherSecrushCrushed() bool {
	for _, tag := range r.OtherTags {
		if tag.Category == "source" && tag.Name == "secretcrush" {
			return true
		}
	}
	for _, tag := range r.Tags {
		if tag.Category == "source" && tag.Name == "secretcrush" {
			return true
		}
	}
	return false
}

func (r Relationship) OtherVipBoostLiked() bool {
	for _, tag := range r.OtherTags {
		if tag.Category == "source" && tag.Name == "vipboost" {
			return true
		}
	}
	return false
}

func (r Relationship) OtherVipBadgeLiked() bool {
	for _, tag := range r.OtherTags {
		if tag.Category == "source" && tag.Name == "vipbadge" {
			return true
		}
	}
	return false
}

func (r Relationship) EitherSuperLiked() bool {
	return r.State == RelationshipStateEnum_SUPERLIKED || r.OtherState == RelationshipStateEnum_SUPERLIKED
}

func NewUpsertParam(user string, otherUser string, state RelationshipStateEnum) UpsertParam {
	r := RelationshipUpdate{
		UserId:      user,
		OtherUserId: otherUser,
		State:       state,
	}
	return UpsertParam{Relationship: &r}
}

func NewSecretCrushUpsertParam(user string, otherUser string) UpsertParam {
	r := RelationshipUpdate{
		UserId:      user,
		OtherUserId: otherUser,
		State:       RelationshipStateEnum_LIKED,
		Tags:        []*Tag{&Tag{Category: "source", Name: "secretcrush"}},
	}
	return UpsertParam{Relationship: &r}
}
