package i18n

import "github.com/p1cn/tantan-backend-common/util"

type OnboardingUserTagType int

var (
	TeamAccounts = map[string]TeamAccount{}

	QuestionMessageHeaders = map[string]map[string]string{}
	QuestionMessageFooters = map[string]map[string]string{}

	MediaMessagePreviews = map[string]map[string]string{}
	MediaMessageUpgrades = map[string]map[string]string{}

	PushMessageFormats = map[string]PushMessageFormat{}

	MassMessages = map[string]MassMessage{}

	//i18n.OnBoardingMessages[lang].Male[tagType])
	OnBoardingMessages = map[string]OnBoardingMessage{}

	SmsTemplates = map[string]map[string]string{}
)

type TeamAccount struct {
	Name        string
	Description string
}

type PushMessageFormat struct {
	AppName  string
	Language string

	SecretLikeTicker      string
	SecretLikeTickerShort string
	SecretLikeTitle       string
	SecretLikeValue       string

	SecretMatchTicker      string
	SecretMatchTickerShort string
	SecretMatchTitle       string
	SecretMatchValue       string

	MessageTicker          string
	MessageTickerShort     string
	ImageMessageTicker     string
	AudioMessageTicker     string
	VideoMessageTicker     string
	StickerMessageTicker   string
	QuestionMessageTicker  string
	LocationMessageTicker  string
	MomentCommentTicker    string
	AggregatedMessageValue string
	ImageMessageValue      string
	AudioMessageValue      string
	VideoMessageValue      string
	StickerMessageValue    string
	QuestionMessageValue   string
	LocationMessageValue   string
	TbhFriendRequestTitle  string
	TbhFriendRequestTicker string
	TbhVoteReceivedTitle   string
	TbhVoteReceivedTicker  string

	MatchTicker          string
	MatchTickerShort     string
	SingleMatchTitle     string
	SingleMatchValue     string
	AggregatedMatchValue string

	ReminderWithLikesTicker string
	ReminderWithLikesTitle  string
	ReminderWithLikesValue  string

	ReminderWithoutLikesTicker string
	ReminderWithoutLikesTitle  string
	ReminderWithoutLikesValue  string

	ReminderBeforeCNYMaleTicker        string
	ReminderBeforeCNYFemaleTicker      string
	ReminderBeforeCNYFemale2017Ticker  string
	ReminderAfterCNYMaleTicker         string
	ReminderAfterCNYFemaleTicker       string
	ReminderAfterCNYFemaleNoCityTicker string

	ScenarioFoodLikeTicker string
	ScenarioFoodLikeTitle  string
	ScenarioFoodLikeValue  string

	ScenarioSportLikeTicker string
	ScenarioSportLikeTitle  string
	ScenarioSportLikeValue  string

	ScenarioMovieLikeTicker string
	ScenarioMovieLikeTitle  string
	ScenarioMovieLikeValue  string

	SuperLikeReceiveTicker string
	SuperLikeReceiveTitle  string

	SuperLikeInviteTicker string
	SuperLikeInviteTitle  string

	MomentMatchPostTicker string
	MomentMatchPostValue  string
}

type MassMessage struct {
	Male   string
	Female string
}

type OnBoardingMessage struct {
	UnpopularMale   string
	UnpopularFemale string
	FewswipesFemale string
	FewswipesMale   string
}

func InitTranslations() {
	for _, lang := range util.TranslatedLanguages {
		TeamAccounts[lang] = TeamAccount{
			Name:        MsgTeamAccountName.Text(lang),
			Description: MsgTeamAccountDescription.Text(lang),
		}

		QuestionMessageHeaders[lang] = map[string]string{
			"classic":  MsgQuestionMessageHeaderClassic.Text(lang),
			"intimate": MsgQuestionMessageHeaderIntimate.Text(lang),
		}
		QuestionMessageFooters[lang] = map[string]string{
			"classic":  MsgQuestionMessageFooterClassic.Text(lang),
			"intimate": MsgQuestionMessageFooterIntimate.Text(lang),
		}

		MediaMessagePreviews[lang] = map[string]string{
			"image":      MsgMediaMessageImagePreview.Text(lang),
			"audio":      MsgMeidaMessageAudioPreview.Text(lang),
			"video":      MsgMediaMessageVideoPreview.Text(lang),
			"sticker":    MsgMediaMessageStickerPreview.Text(lang),
			"question":   "",
			"location":   MsgMediaMessageLocationPreview.Text(lang),
			"momentLike": "",
		}
		MediaMessageUpgrades[lang] = map[string]string{
			"image":      MsgMediaMessageImageUpgrade.Text(lang),
			"audio":      MsgMediaMessageAudioUpgrade.Text(lang),
			"video":      MsgMediaMessageVideoUpgrade.Text(lang),
			"sticker":    MsgMediaMessageStickerUpgrade.Text(lang),
			"question":   MsgMediaMessageQuestionUpgrade.Text(lang),
			"location":   MsgMediaMessageLocationUpgrade.Text(lang),
			"momentLike": MsgMeidaMessageMomentLikeUpgrade.Text(lang),
		}

		PushMessageFormats[lang] = PushMessageFormat{
			AppName:  MsgPushMessageAppName.Text(lang),
			Language: lang,

			SecretLikeTicker:      MsgPushMessageSecretLikeTicker.Text(lang),
			SecretLikeTickerShort: MsgPushMessageSecretLikeTickerShort.Text(lang),
			SecretLikeTitle:       MsgPushMessageSecretLikeTitle.Text(lang),
			SecretLikeValue:       MsgPushMessageSecretLikeValue.Text(lang),

			SuperLikeReceiveTicker: MsgPushMessageSuperLikeReceiveTicker.Text(lang),
			SuperLikeReceiveTitle:  MsgPushMessageSuperLikeReceiveTitle.Text(lang),
			SuperLikeInviteTicker:  MsgPushMessageSuperLikeInviteTicker.Text(lang),
			SuperLikeInviteTitle:   MsgPushMessageSuperLikeInviteTitle.Text(lang),

			SecretMatchTicker:      MsgPushMessageSecretMatchTicker.Text(lang),
			SecretMatchTickerShort: MsgPushMessageSecretMatchTickerShort.Text(lang),
			SecretMatchTitle:       MsgPushMessageSecretMatchTitle.Text(lang),
			SecretMatchValue:       MsgPushMessageSecretMatchValue.Text(lang),

			MessageTicker:          MsgPushMessageMessageTicker.Text(lang),
			MessageTickerShort:     MsgPushMessageMessageTickerShort.Text(lang),
			ImageMessageTicker:     MsgPushMessageImageMessageTicker.Text(lang),
			AudioMessageTicker:     MsgPushMessageAudioMessageTicker.Text(lang),
			VideoMessageTicker:     MsgPushMessageVideoMessageTicker.Text(lang),
			StickerMessageTicker:   MsgPushMessageStickerMessageTicker.Text(lang),
			QuestionMessageTicker:  MsgPushMessageQuestionMessageTicker.Text(lang),
			LocationMessageTicker:  MsgPushMessageLocationMessageTicker.Text(lang),
			MomentCommentTicker:    MsgPushMessageMomentCommentTicker.Text(lang),
			AggregatedMessageValue: MsgPushMessageAggregatedMessageValue.Text(lang),
			ImageMessageValue:      MsgPushMessageImageMessageValue.Text(lang),
			AudioMessageValue:      MsgPushMessageAudioMessageValue.Text(lang),
			VideoMessageValue:      MsgPushMessageVideoMessageValue.Text(lang),
			StickerMessageValue:    MsgPushMessageStickerMessageValue.Text(lang),
			QuestionMessageValue:   MsgPushMessageQuestionMessageValue.Text(lang),
			LocationMessageValue:   MsgPushMessageLocationMessageValue.Text(lang),
			TbhFriendRequestTitle:  MsgPushMessageTbhFriendRequestTitle.Text(lang),
			TbhFriendRequestTicker: MsgPushMessageTbhFriendRequestTicker.Text(lang),
			TbhVoteReceivedTitle:   MsgPushMessageTbhVoteReceivedTitle.Text(lang),
			TbhVoteReceivedTicker:  MsgPushMessageTbhVoteReceivedTicker.Text(lang),

			MatchTicker:          MsgPushMessageMatchTicker.Text(lang),
			MatchTickerShort:     MsgPushMessageMatchTickerShort.Text(lang),
			SingleMatchTitle:     MsgPushMessageSingleMatchTitle.Text(lang),
			SingleMatchValue:     MsgPushMessageSingleMatchValue.Text(lang),
			AggregatedMatchValue: MsgPushMessageAggregatedMatchValue.Text(lang),

			ReminderWithLikesTicker:    MsgPushMessageReminderWithLikesTicker.Text(lang),
			ReminderWithLikesTitle:     MsgPushMessageReminderWithLikesTitle.Text(lang),
			ReminderWithLikesValue:     MsgPushMessageReminderWithLikesValue.Text(lang),
			ReminderWithoutLikesTicker: MsgPushMessageReminderWithoutLikesTicker.Text(lang),
			ReminderWithoutLikesTitle:  MsgPushMessageReminderWithoutLikesTitle.Text(lang),
			ReminderWithoutLikesValue:  MsgPushMessageReminderWithoutLikesValue.Text(lang),

			ReminderBeforeCNYMaleTicker:        MsgPushMessageReminderBeforeCNYMaleTicker.Text(lang),
			ReminderBeforeCNYFemaleTicker:      MsgPushMessageReminderBeforeCNYFemaleTicker.Text(lang),
			ReminderBeforeCNYFemale2017Ticker:  MsgPushMessageReminderBeforeCNYFemale2017Ticker.Text(lang),
			ReminderAfterCNYMaleTicker:         MsgPushMessageReminderAfterCNYMaleTicker.Text(lang),
			ReminderAfterCNYFemaleTicker:       MsgPushMessageReminderAfterCNYFemaleTicker.Text(lang),
			ReminderAfterCNYFemaleNoCityTicker: MsgPushMessageReminderAfterCNYFemaleNoCityTicker.Text(lang),

			ScenarioFoodLikeTicker:  MsgPushMessageScenarioFoodLikeTicker.Text(lang),
			ScenarioFoodLikeTitle:   MsgPushMessageScenarioFoodLikeTitle.Text(lang),
			ScenarioFoodLikeValue:   MsgPushMessageScenarioFoodLikeValue.Text(lang),
			ScenarioSportLikeTicker: MsgPushMessageScenarioSportLikeTicker.Text(lang),
			ScenarioSportLikeTitle:  MsgPushMessageScenarioSportLikeTitle.Text(lang),
			ScenarioSportLikeValue:  MsgPushMessageScenarioSportLikeValue.Text(lang),
			ScenarioMovieLikeTicker: MsgPushMessageScenarioMovieLikeTicker.Text(lang),
			ScenarioMovieLikeTitle:  MsgPushMessageScenarioMovieLikeTitle.Text(lang),
			ScenarioMovieLikeValue:  MsgPushMessageScenarioMovieLikeValue.Text(lang),

			MomentMatchPostTicker: MsgPushMessageMomentMatchPostTicker.Text(lang),
			MomentMatchPostValue:  MsgPushMessageMomentMatchPostValue.Text(lang),
		}

		MassMessages[lang] = MassMessage{
			Male:   MsgPushMessageMassMessageMaleTicker.Text(lang),
			Female: MsgPushMessageMassMessageFemaleTicker.Text(lang),
		}

		OnBoardingMessages[lang] = OnBoardingMessage{
			UnpopularMale:   MsgOnBoardingUnpopularMale.Text(lang),
			UnpopularFemale: MsgOnBoardingUnpopularFemale.Text(lang),
			FewswipesMale:   MsgOnBoardingFewswipesMale.Text(lang),
			FewswipesFemale: MsgOnBoardingFewswipesFemale.Text(lang),
		}

		SmsTemplates[lang] = map[string]string{
			"tantan_sign_up_verification":      MsgSmsAccountSignUpVerification.Text(lang),
			"tantan_sign_in_verification":      MsgSmsAccountSignInVerification.Text(lang),
			"tantan_reset_password":            MsgSmsAccountResetPassword.Text(lang),
			"tantan_change_phone":              MsgSmsAccountChangePhone.Text(lang),
			"tantan_promotion":                 MsgSmsTantanPromotion.Text(lang), // not used anymore?
			"tantan_secret_crush_with_name":    MsgSmsSecretCrushWithName.Text(lang),
			"tantan_secret_crush_without_name": MsgSmsSecretCrushWithoutName.Text(lang),
			"tantan_promotion_with_name":       MsgSmsSecretCrushPromotionWithName.Text(lang),
			"tantan_promotion_without_name":    MsgSmsSecretCrushPromotionWithoutName.Text(lang),
			"tantan_promotion_with_name_v3":    MsgSmsSecretCrushPromotionWithoutNameV3.Text(lang),
			"tantan_promotion_with_name_v4":    MsgSmsSecretCrushPromotionWithoutNameV4.Text(lang),
			"tantan_promotion_with_name_v5":    MsgSmsSecretCrushPromotionWithoutNameV5.Text(lang),
			"tantan_promotion_with_name_v6":    MsgSmsSecretCrushPromotionWithoutNameV6.Text(lang),
			"tantan_promotion_with_name_v7":    MsgSmsSecretCrushPromotionWithoutNameV7.Text(lang),
			"tantan_promotion_with_name_v8":    MsgSmsSecretCrushPromotionWithoutNameV8.Text(lang),
			"tantan_promotion_with_name_v9":    MsgSmsSecretCrushPromotionWithoutNameV9.Text(lang),
			"tantan_promotion_with_name_v10":   MsgSmsSecretCrushPromotionWithoutNameV10.Text(lang),
			"tantan_promotion_with_name_v11":   MsgSmsSecretCrushPromotionWithoutNameV11.Text(lang),
			"tantan_promotion_with_name_v12":   MsgSmsSecretCrushPromotionWithoutNameV12.Text(lang),
			"tantan_promotion_with_name_v13":   MsgSmsSecretCrushPromotionWithoutNameV13.Text(lang),
			"tantan_promotion_with_name_v14":   MsgSmsSecretCrushPromotionWithoutNameV14.Text(lang),
			"tantan_promotion_with_name_v20":   MsgSmsSecretCrushPromotionWithoutNameV20.Text(lang),
		}
	}
}
