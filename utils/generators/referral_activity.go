package generators

import (
	"github.com/FoxComm/FoxComm/models"
	"github.com/FoxComm/FoxComm/utils/test"
)

// ReferralSiteActivity generates a SiteActivity that shows an inbound referral
// action.
func ReferralSiteActivity(sharerToken, sessionToken string, fss *test.FakeSocialShopping) models.SiteActivity {
	return models.SiteActivity{
		Action:             "referrer",
		SharerToken:        sharerToken,
		SharingActivityTag: 4,
		RefererURL:         "",
		LandingURL:         "/signup",
		ApiRequestURL:      "",
		SessionToken:       sessionToken,
		Type:               "inbound",
		StoreURL:           fss.Server.URL,
	}
}
