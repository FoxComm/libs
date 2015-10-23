package endpoints

import (
	"github.com/FoxComm/libs/configs"
)

var SocialAnalyticsAPI *Endpoint
var SocialAnalyticsLoyaltyEngineAPI *Endpoint

func init() {
	SocialAnalyticsAPI = &Endpoint{
		Name:        "social_analytics",
		Description: "Social Analytics",
		Domain:      configs.Get("SocialAnalyticsHost"),
		DefaultPort: configs.Get("SocialAnalyticsPort"),
		APIPrefix:   configs.Get("SocialAnalyticsAPIPrefix"),
		IsFeature:   true,
	}

	SocialAnalyticsLoyaltyEngineAPI = &Endpoint{
		Name:        "loyalty_engine",
		Description: "Loyalty engine",
		Domain:      configs.Get("SocialAnalyticsHost"),
		DefaultPort: configs.Get("SocialAnalyticsPort"),
		APIPrefix:   configs.Get("LoyaltyEngineAPIPrefix"),
		IsFeature:   true,
	}

	Add(SocialAnalyticsAPI)
	Add(SocialAnalyticsLoyaltyEngineAPI)
}
