package seeds

import (
	"github.com/FoxComm/FoxComm/endpoints"
	"github.com/FoxComm/FoxComm/models"
	"github.com/FoxComm/FoxComm/repositories"
)

var Preferences = []models.PromotionPreferences{
	models.PromotionPreferences{
		Name:            "Welcome",
		Description:     "A gift from FoxComm",
		MatchPolicy:     "all",
		CouponRecipient: models.REFEREE,
		TriggerAction:   "signup",
		Rules: []models.PromotionRule{
			models.PromotionRule{
				Type: "FirstOrder",
			},
		},
		Actions: []models.PromotionAction{
			models.PromotionAction{
				Type:            "CreateAdjustment",
				Calculator:      "FlatPercentItemTotal",
				PreferenceName:  "preferred_flat_percent",
				PreferenceValue: "20",
			},
		},
		AutoApply: true,
		Active:    false,
	},
	models.PromotionPreferences{
		Name:            "Welcome",
		Description:     "A gift from FoxComm",
		MatchPolicy:     "all",
		CouponRecipient: models.REFEREE,
		TriggerAction:   "signup",
		Rules: []models.PromotionRule{
			models.PromotionRule{
				Type: "FirstOrder",
			},
		},
		Actions: []models.PromotionAction{
			models.PromotionAction{
				Type:            "CreateAdjustment",
				Calculator:      "FlatRate",
				PreferenceName:  "preferred_amount",
				PreferenceValue: "25",
			},
		},
		AutoApply: true,
		Active:    true,
	},
	models.PromotionPreferences{
		Name:            "Give/Get",
		Description:     "A gift from FoxComm",
		MatchPolicy:     "all",
		CouponRecipient: models.REFERRER,
		TriggerAction:   "checkout",
		Rules: []models.PromotionRule{
			models.PromotionRule{
				Type:            models.PromotionRuleSpreeMapping["ItemTotal"],
				PreferenceName:  models.PromotionPreferenceNameSpreeMapping["FlatRate"],
				PreferenceValue: "50.0",
				ComparisonType:  "gte",
			},
		},
		Actions: []models.PromotionAction{
			models.PromotionAction{
				Type:            "CreateAdjustment",
				Calculator:      "FlatRate",
				PreferenceName:  "preferred_amount",
				PreferenceValue: "25",
			},
		},
		AutoApply: false,
		Active:    true,
	},
}

func (sr SeedRunner) GeneratePromotionPreferences() {
	storeID := StoreID()

	promotionPrefRepo := repositories.PromotionPreferencesRepo{}
	promotionPrefRepo.InitializeWithStoreID(storeID, repositories.PromotionPreferencesCollection, endpoints.SocialAnalyticsAPI)
	promotionPrefRepo.DestroyAll()

	for _, pref := range Preferences {
		promotionPrefRepo.Create(&pref)
	}
}
