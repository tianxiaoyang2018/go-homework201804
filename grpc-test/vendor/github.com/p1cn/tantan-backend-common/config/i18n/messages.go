package i18n

const (
	MsgPushMessageAppName Msg = iota

	MsgPushMessageSecretLikeTicker
	MsgPushMessageSecretLikeTickerShort
	MsgPushMessageSecretLikeTitle
	MsgPushMessageSecretLikeValue

	MsgPushMessageSecretMatchTicker
	MsgPushMessageSecretMatchTickerShort
	MsgPushMessageSecretMatchTitle
	MsgPushMessageSecretMatchValue

	MsgPushMessageMessageTicker
	MsgPushMessageMessageTickerShort
	MsgPushMessageImageMessageTicker
	MsgPushMessageAudioMessageTicker
	MsgPushMessageVideoMessageTicker
	MsgPushMessageStickerMessageTicker
	MsgPushMessageQuestionMessageTicker
	MsgPushMessageLocationMessageTicker
	MsgPushMessageMomentCommentTicker
	MsgPushMessageAggregatedMessageValue
	MsgPushMessageImageMessageValue
	MsgPushMessageAudioMessageValue
	MsgPushMessageVideoMessageValue
	MsgPushMessageStickerMessageValue
	MsgPushMessageQuestionMessageValue
	MsgPushMessageLocationMessageValue
	MsgPushMessageTbhFriendRequestTitle
	MsgPushMessageTbhFriendRequestTicker
	MsgPushMessageTbhVoteReceivedTitle
	MsgPushMessageTbhVoteReceivedTicker

	MsgPushMessageMatchTicker
	MsgPushMessageMatchTickerShort
	MsgPushMessageSingleMatchTitle
	MsgPushMessageSingleMatchValue
	MsgPushMessageAggregatedMatchValue

	MsgPushMessageReminderWithLikesTicker
	MsgPushMessageReminderWithLikesTitle
	MsgPushMessageReminderWithLikesValue
	MsgPushMessageReminderWithoutLikesTicker
	MsgPushMessageReminderWithoutLikesTitle
	MsgPushMessageReminderWithoutLikesValue
	MsgPushMessageReminderBeforeCNYMaleTicker
	MsgPushMessageReminderBeforeCNYFemaleTicker
	MsgPushMessageReminderBeforeCNYFemale2017Ticker
	MsgPushMessageReminderAfterCNYMaleTicker
	MsgPushMessageReminderAfterCNYFemaleTicker
	MsgPushMessageReminderAfterCNYFemaleNoCityTicker

	MsgPushMessageScenarioFoodLikeTicker
	MsgPushMessageScenarioFoodLikeTitle
	MsgPushMessageScenarioFoodLikeValue
	MsgPushMessageScenarioSportLikeTicker
	MsgPushMessageScenarioSportLikeTitle
	MsgPushMessageScenarioSportLikeValue
	MsgPushMessageScenarioMovieLikeTicker
	MsgPushMessageScenarioMovieLikeTitle
	MsgPushMessageScenarioMovieLikeValue

	MsgPushMessageMassMessageMaleTicker
	MsgPushMessageMassMessageFemaleTicker

	MsgPushMessageSuperLikeReceiveTicker
	MsgPushMessageSuperLikeReceiveTitle
	MsgPushMessageSuperLikeInviteTicker
	MsgPushMessageSuperLikeInviteTitle

	MsgQuestionMessageHeaderClassic
	MsgQuestionMessageHeaderIntimate
	MsgQuestionMessageFooterClassic
	MsgQuestionMessageFooterIntimate

	MsgRecalledMessage

	MsgMediaMessageImagePreview
	MsgMediaMessageImageUpgrade
	MsgMeidaMessageAudioPreview
	MsgMediaMessageAudioUpgrade
	MsgMediaMessageVideoPreview
	MsgMediaMessageVideoUpgrade
	MsgMediaMessageStickerPreview
	MsgMediaMessageStickerUpgrade
	MsgMediaMessageQuestionUpgrade
	MsgMediaMessageLocationPreview
	MsgMediaMessageLocationUpgrade
	MsgMeidaMessageMomentLikeUpgrade

	MsgSmsAccountSignUpVerification
	MsgSmsAccountSignInVerification
	MsgSmsAccountResetPassword
	MsgSmsAccountChangePhone
	MsgSmsTantanPromotion
	MsgSmsSecretCrushWithName
	MsgSmsSecretCrushWithoutName
	MsgSmsSecretCrushPromotionWithName
	MsgSmsSecretCrushPromotionWithoutName
	MsgSmsSecretCrushPromotionWithoutNameV3
	MsgSmsSecretCrushPromotionWithoutNameV4
	MsgSmsSecretCrushPromotionWithoutNameV5
	MsgSmsSecretCrushPromotionWithoutNameV6
	MsgSmsSecretCrushPromotionWithoutNameV7
	MsgSmsSecretCrushPromotionWithoutNameV8
	MsgSmsSecretCrushPromotionWithoutNameV9
	MsgSmsSecretCrushPromotionWithoutNameV10
	MsgSmsSecretCrushPromotionWithoutNameV11
	MsgSmsSecretCrushPromotionWithoutNameV12
	MsgSmsSecretCrushPromotionWithoutNameV13
	MsgSmsSecretCrushPromotionWithoutNameV14
	MsgSmsSecretCrushPromotionWithoutNameV20

	MsgTeamAccountName
	MsgTeamAccountDescription

	MsgOnBoardingUnpopularMale
	MsgOnBoardingUnpopularFemale
	MsgOnBoardingFewswipesMale
	MsgOnBoardingFewswipesFemale

	MsgPushMessageMomentMatchPostTicker
	MsgPushMessageMomentMatchPostValue
)

var msgLists = map[Msg]msg{
	MsgPushMessageAppName: msg{id: "PUSH_MESSAGE_APPNAME_TEXT"},

	MsgPushMessageSecretLikeTicker:      msg{id: "PUSH_MESSAGE_SECRET_LIKE_TICKER_TEXT"},
	MsgPushMessageSecretLikeTickerShort: msg{id: "PUSH_MESSAGE_SECRET_LIKE_TICKER_SHORT_TEXT"},
	MsgPushMessageSecretLikeTitle:       msg{id: "PUSH_MESSAGE_SECRET_LIKE_TITLE_TEXT"},
	MsgPushMessageSecretLikeValue:       msg{id: "PUSH_MESSAGE_SECRET_LIKE_VALUE_TEXT"},

	MsgPushMessageSecretMatchTicker:      msg{id: "PUSH_MESSAGE_SECRET_MATCH_TICKER_TEXT"},
	MsgPushMessageSecretMatchTickerShort: msg{id: "PUSH_MESSAGE_SECRET_MATCH_TICKER_SHORT_TEXT"},
	MsgPushMessageSecretMatchTitle:       msg{id: "PUSH_MESSAGE_SECRET_MATCH_TITLE_TEXT"},
	MsgPushMessageSecretMatchValue:       msg{id: "PUSH_MESSAGE_SECRET_MATCH_VALUE_TEXT"},

	MsgPushMessageSuperLikeReceiveTicker: msg{id: "PUSH_MESSAGE_SUPERLIKE_RECEIVE_TICKER_TEXT"},
	MsgPushMessageSuperLikeReceiveTitle:  msg{id: "PUSH_MESSAGE_SUPERLIKE_RECEIVE_TITLE_TEXT"},
	MsgPushMessageSuperLikeInviteTicker:  msg{id: "PUSH_MESSAGE_SUPERLIKE_INVITE_TICKER_TEXT"},
	MsgPushMessageSuperLikeInviteTitle:   msg{id: "PUSH_MESSAGE_SUPERLIKE_INVITE_TITLE_TEXT"},

	MsgPushMessageMessageTicker:          msg{id: "PUSH_MESSAGE_MESSAGE_TICKER_TEXT", placeholders: 1},
	MsgPushMessageMessageTickerShort:     msg{id: "PUSH_MESSAGE_MESSAGE_TICKER_SHORT_TEXT"},
	MsgPushMessageImageMessageTicker:     msg{id: "PUSH_MESSAGE_IMAGE_MESSAGE_TICKER_TEXT", placeholders: 1},
	MsgPushMessageAudioMessageTicker:     msg{id: "PUSH_MESSAGE_AUDIO_MESSAGE_TICKER_TEXT", placeholders: 1},
	MsgPushMessageVideoMessageTicker:     msg{id: "PUSH_MESSAGE_VIDEO_MESSAGE_TICKER_TEXT", placeholders: 1},
	MsgPushMessageStickerMessageTicker:   msg{id: "PUSH_MESSAGE_STICKER_MESSAGE_TICKER_TEXT", placeholders: 1},
	MsgPushMessageQuestionMessageTicker:  msg{id: "PUSH_MESSAGE_QUESTION_MESSAGE_TICKER_TEXT", placeholders: 1},
	MsgPushMessageLocationMessageTicker:  msg{id: "PUSH_MESSAGE_LOCATION_MESSAGE_TICKER_TEXT", placeholders: 1},
	MsgPushMessageMomentCommentTicker:    msg{id: "PUSH_MESSAGE_MOMENT_COMMENT_TICKER_TEXT", placeholders: 1},
	MsgPushMessageAggregatedMessageValue: msg{id: "PUSH_MESSAGE_AGGREGATED_MESSAGE_VALUE_TEXT", placeholders: 1},
	MsgPushMessageImageMessageValue:      msg{id: "PUSH_MESSAGE_IMAGE_MESSAGE_VALUE_TEXT"},
	MsgPushMessageAudioMessageValue:      msg{id: "PUSH_MESSAGE_AUDIO_MESSAGE_VALUE_TEXT"},
	MsgPushMessageVideoMessageValue:      msg{id: "PUSH_MESSAGE_VIDEO_MESSAGE_VALUE_TEXT"},
	MsgPushMessageStickerMessageValue:    msg{id: "PUSH_MESSAGE_STICKER_MESSAGE_VALUE_TEXT"},
	MsgPushMessageQuestionMessageValue:   msg{id: "PUSH_MESSAGE_QUESTION_MESSAGE_VALUE_TEXT"},
	MsgPushMessageLocationMessageValue:   msg{id: "PUSH_MESSAGE_LOCATION_MESSAGE_VALUE_TEXT"},
	MsgPushMessageTbhFriendRequestTitle:  msg{id: "PUSH_MESSAGE_TBH_FRIEND_REQUEST_TITLE_TEXT"},
	MsgPushMessageTbhFriendRequestTicker: msg{id: "PUSH_MESSAGE_TBH_FRIEND_REQUEST_TICKER_TEXT"},
	MsgPushMessageTbhVoteReceivedTitle:   msg{id: "PUSH_MESSAGE_TBH_VOTE_RECEIVED_TITLE_TEXT"},
	MsgPushMessageTbhVoteReceivedTicker:  msg{id: "PUSH_MESSAGE_TBH_VOTE_RECEIVED_TICKER_TEXT"},

	MsgPushMessageMatchTicker:          msg{id: "PUSH_MESSAGE_MATCH_TICKER_TEXT"},
	MsgPushMessageMatchTickerShort:     msg{id: "PUSH_MESSAGE_MATCH_TICKER_SHORT_TEXT"},
	MsgPushMessageSingleMatchTitle:     msg{id: "PUSH_MESSAGE_SINGLE_MATCH_TITLE_TEXT"},
	MsgPushMessageSingleMatchValue:     msg{id: "PUSH_MESSAGE_SINGLE_MATCH_VALUE_TEXT"},
	MsgPushMessageAggregatedMatchValue: msg{id: "PUSH_MESSAGE_AGGREGATED_MATCH_VALUE_TEXT", placeholders: 1},

	MsgPushMessageReminderWithLikesTicker:            msg{id: "PUSH_MESSAGE_REMINDER_WITH_LIKES_TICKER_TEXT"},
	MsgPushMessageReminderWithLikesTitle:             msg{id: "PUSH_MESSAGE_REMINDER_WITH_LIKES_TITLE_TEXT"},
	MsgPushMessageReminderWithLikesValue:             msg{id: "PUSH_MESSAGE_REMINDER_WITH_LIKES_VALUE_TEXT"},
	MsgPushMessageReminderWithoutLikesTicker:         msg{id: "PUSH_MESSAGE_REMINDER_WITHOUT_LIKES_TICKER_TEXT"},
	MsgPushMessageReminderWithoutLikesTitle:          msg{id: "PUSH_MESSAGE_REMINDER_WITHOUT_LIKES_TITLE_TEXT"},
	MsgPushMessageReminderWithoutLikesValue:          msg{id: "PUSH_MESSAGE_REMINDER_WITHOUT_LIKES_VALUE_TEXT"},
	MsgPushMessageReminderBeforeCNYMaleTicker:        msg{id: "PUSH_MESSAGE_REMINDER_BEFORE_CNY_MALE_TICKER_TEXT"},
	MsgPushMessageReminderBeforeCNYFemaleTicker:      msg{id: "PUSH_MESSAGE_REMINDER_BEFORE_CNY_FEMALE_TICKER_TEXT"},
	MsgPushMessageReminderBeforeCNYFemale2017Ticker:  msg{id: "PUSH_MESSAGE_REMINDER_BEFORE_CNY_FEMALE_2017_TICKER_TEXT"},
	MsgPushMessageReminderAfterCNYMaleTicker:         msg{id: "PUSH_MESSAGE_REMINDER_AFTER_CNY_MALE_TICKER_TEXT"},
	MsgPushMessageReminderAfterCNYFemaleTicker:       msg{id: "PUSH_MESSAGE_REMINDER_AFTER_CNY_FEMALE_TICKER_TEXT", placeholders: 1},
	MsgPushMessageReminderAfterCNYFemaleNoCityTicker: msg{id: "PUSH_MESSAGE_REMINDER_AFTER_CNY_FEMALE_NO_CITY_TICKER_TEXT"},

	MsgPushMessageScenarioFoodLikeTicker:  msg{id: "PUSH_MESSAGE_SCENARIO_FOOD_LIKE_TICKER_TEXT"},
	MsgPushMessageScenarioFoodLikeTitle:   msg{id: "PUSH_MESSAGE_SCENARIO_FOOD_LIKE_TITLE_TEXT"},
	MsgPushMessageScenarioFoodLikeValue:   msg{id: "PUSH_MESSAGE_SCENARIO_FOOD_LIKE_VALUE_TEXT"},
	MsgPushMessageScenarioSportLikeTicker: msg{id: "PUSH_MESSAGE_SCENARIO_SPORT_LIKE_TICKER_TEXT"},
	MsgPushMessageScenarioSportLikeTitle:  msg{id: "PUSH_MESSAGE_SCENARIO_SPORT_LIKE_TITLE_TEXT"},
	MsgPushMessageScenarioSportLikeValue:  msg{id: "PUSH_MESSAGE_SCENARIO_SPORT_LIKE_VALUE_TEXT"},
	MsgPushMessageScenarioMovieLikeTicker: msg{id: "PUSH_MESSAGE_SCENARIO_MOVIE_LIKE_TICKER_TEXT"},
	MsgPushMessageScenarioMovieLikeTitle:  msg{id: "PUSH_MESSAGE_SCENARIO_MOVIE_LIKE_TITLE_TEXT"},
	MsgPushMessageScenarioMovieLikeValue:  msg{id: "PUSH_MESSAGE_SCENARIO_MOVIE_LIKE_VALUE_TEXT"},

	MsgPushMessageMassMessageMaleTicker:   msg{id: "PUSH_MESSAGE_MASS_MESSAGE_MALE_TICKER_TEXT"},
	MsgPushMessageMassMessageFemaleTicker: msg{id: "PUSH_MESSAGE_MASS_MESSAGE_FEMALE_TICKER_TEXT"},

	MsgQuestionMessageHeaderClassic:  msg{id: "QUESTION_MESSAGE_HEADER_CLASSIC_TEXT"},
	MsgQuestionMessageHeaderIntimate: msg{id: "QUESTION_MESSAGE_HEADER_INTIMATE_TEXT"},
	MsgQuestionMessageFooterClassic:  msg{id: "QUESTION_MESSAGE_FOOTER_CLASSIC_TEXT"},
	MsgQuestionMessageFooterIntimate: msg{id: "QUESTION_MESSAGE_FOOTER_INTIMATE_TEXT"},

	MsgRecalledMessage: msg{id: "RECALLED_MESSAGE_TEXT"},

	MsgMediaMessageImagePreview:      msg{id: "MEDIA_MESSAGE_IMAGE_PREVIEW_TEXT"},
	MsgMediaMessageImageUpgrade:      msg{id: "MEDIA_MESSAGE_IMAGE_UPGRADE_TEXT"},
	MsgMeidaMessageAudioPreview:      msg{id: "MEDIA_MESSAGE_AUDIO_PREVIEW_TEXT"},
	MsgMediaMessageAudioUpgrade:      msg{id: "MEDIA_MESSAGE_AUDIO_UPGRADE_TEXT"},
	MsgMediaMessageVideoPreview:      msg{id: "MEDIA_MESSAGE_VIDEO_PREVIEW_TEXT"},
	MsgMediaMessageVideoUpgrade:      msg{id: "MEDIA_MESSAGE_VIDEO_UPGRADE_TEXT"},
	MsgMediaMessageStickerPreview:    msg{id: "MEDIA_MESSAGE_STICKER_PREVIEW_TEXT"},
	MsgMediaMessageStickerUpgrade:    msg{id: "MEDIA_MESSAGE_STICKER_UPGRADE_TEXT"},
	MsgMediaMessageQuestionUpgrade:   msg{id: "MEDIA_MESSAGE_QUESTION_UPGRADE_TEXT"},
	MsgMediaMessageLocationPreview:   msg{id: "MEDIA_MESSAGE_LOCATION_PREVIEW_TEXT"},
	MsgMediaMessageLocationUpgrade:   msg{id: "MEDIA_MESSAGE_LOCATION_UPGRADE_TEXT"},
	MsgMeidaMessageMomentLikeUpgrade: msg{id: "MEDIA_MESSAGE_MOMENTLIKE_UPGRADE_TEXT"},

	MsgSmsAccountSignUpVerification:          msg{id: "SMS_ACCOUNT_SIGN_UP_VERIFICATION"},
	MsgSmsAccountSignInVerification:          msg{id: "SMS_ACCOUNT_SIGN_IN_VERIFICATION"},
	MsgSmsAccountResetPassword:               msg{id: "SMS_ACCOUNT_RESET_PASSWORD"},
	MsgSmsAccountChangePhone:                 msg{id: "SMS_ACCOUNT_CHANGE_PHONE"},
	MsgSmsTantanPromotion:                    msg{id: "SMS_TANTAN_PROMOTION"},
	MsgSmsSecretCrushWithName:                msg{id: "SMS_SECRET_CRUSH_WITH_NAME"},
	MsgSmsSecretCrushWithoutName:             msg{id: "SMS_SECRET_CRUSH_WITHOUT_NAME"},
	MsgSmsSecretCrushPromotionWithName:       msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITH_NAME"},
	MsgSmsSecretCrushPromotionWithoutName:    msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITHOUT_NAME"},
	MsgSmsSecretCrushPromotionWithoutNameV3:  msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITH_NAME_V3"},
	MsgSmsSecretCrushPromotionWithoutNameV4:  msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITH_NAME_V4"},
	MsgSmsSecretCrushPromotionWithoutNameV5:  msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITH_NAME_V5"},
	MsgSmsSecretCrushPromotionWithoutNameV6:  msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITH_NAME_V6"},
	MsgSmsSecretCrushPromotionWithoutNameV7:  msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITH_NAME_V7"},
	MsgSmsSecretCrushPromotionWithoutNameV8:  msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITH_NAME_V8"},
	MsgSmsSecretCrushPromotionWithoutNameV9:  msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITH_NAME_V9"},
	MsgSmsSecretCrushPromotionWithoutNameV10: msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITH_NAME_V10"},
	MsgSmsSecretCrushPromotionWithoutNameV11: msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITH_NAME_V11"},
	MsgSmsSecretCrushPromotionWithoutNameV12: msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITH_NAME_V12"},
	MsgSmsSecretCrushPromotionWithoutNameV13: msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITH_NAME_V13"},
	MsgSmsSecretCrushPromotionWithoutNameV14: msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITH_NAME_V14"},
	MsgSmsSecretCrushPromotionWithoutNameV20: msg{id: "SMS_SECRET_CRUSH_PROMOTION_WITH_NAME_V20"},

	MsgTeamAccountName:        msg{id: "TEAM_ACCOUNT_NAME"},
	MsgTeamAccountDescription: msg{id: "TEAM_ACCOUNT_DESCRIPTION"},

	// onboarding message id
	MsgOnBoardingUnpopularMale: msg{id: "PUSH_MESSAGE_ONBOARDING_UNPOPULAR_MALE_MESSAGE", placeholders: 3},
	MsgOnBoardingFewswipesMale: msg{id: "PUSH_MESSAGE_ONBOARDING_FEWSWIPES_MALE_MESSAGE", placeholders: 3},

	MsgOnBoardingUnpopularFemale: msg{id: "PUSH_MESSAGE_ONBOARDING_UNPOPULAR_FEMALE_MESSAGE", placeholders: 3},
	MsgOnBoardingFewswipesFemale: msg{id: "PUSH_MESSAGE_ONBOARDING_FEWSWIPES_FEMALE_MESSAGE", placeholders: 3},

	MsgPushMessageMomentMatchPostTicker: msg{id: "PUSH_MESSAGE_MOMENT_MATCH_POST_TICKER_TEXT"},
	MsgPushMessageMomentMatchPostValue:  msg{id: "PUSH_MESSAGE_MOMENT_MATCH_POST_VALUE_TEXT", placeholders: 1},
}

type Msg uint

type msg struct {
	id           string
	placeholders uint8
}

func (m Msg) Id() string {
	return msgLists[m].id
}

func (m Msg) Text(lang string, vars ...interface{}) string {
	msg := msgLists[m]
	argc := int(msg.placeholders)
	if argc > 0 && len(vars) < argc {
		for i := 0; i < argc-len(vars); {
			vars = append(vars, "%v")
		}
		vars = vars[0:argc]
	}
	po, err := locales.Get(lang)
	if err != nil || po == nil {
		return ""
	}
	if text := po.Get(msg.id, vars...); text != "" && text != msg.id {
		return text
	}
	return ""
}
