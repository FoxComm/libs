package generators

import (
	"fmt"

	core_models "github.com/FoxComm/FoxComm/social_analytics/models"
	"github.com/FoxComm/FoxComm/social_analytics/repositories"
	"github.com/FoxComm/FoxComm/social_analytics/models"
	"github.com/FoxComm/FoxComm/utils/test"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// CreateReferralChain strings together a series of signup ActivityAttributions
// that represent users signing up through referrals.
func CreateReferralChain(
	c *gin.Context,
	depth int,
	referrer core_models.Entity,
	fss *test.FakeSocialShopping,
) ([]models.ActivityAttribution, error) {

	var chain []models.ActivityAttribution
	sar, err := repositories.NewSiteActivityRepo(c.Request)
	if err != nil {
		return chain, err
	}

	aar := repositories.NewActivityAttributionRepo(c)

	lastReferrer := referrer
	for i := 0; i < depth; i++ {
		// First, create an inbound SiteActivity so that we can establish the
		// relationship with the referring user.
		st := fake.Characters(8)
		rt, err := fss.ReferralToken(lastReferrer)
		if err != nil {
			return nil, err
		}

		inboundActivity := ReferralSiteActivity(rt, st, fss)
		err = sar.Create(&inboundActivity)
		if err != nil {
			return nil, err
		}

		// Second, create a signup SiteActivity to represent the user signup.
		url := fmt.Sprintf("http://foxcommerce.com/s/%s/4", rt)
		signupActivity := SignupSiteActivity(url, st, fss)
		err = sar.Create(&signupActivity)
		if err != nil {
			return nil, err
		}

		// Third, add the newly created user to the FakeSocialShopping instance so
		// that it can be used in future referral signups.
		user := signupActivity.SignupDetails.User
		fss.AddUser(fake.Characters(16), &user)

		// Fourth, create a signup ActivityAttribution. This will enable points to be
		// accumulated for the referrer.
		signupAttribution := models.NewActivityAttribution(signupActivity, inboundActivity, lastReferrer, 100)
		id, err := aar.CreateWithId(&signupAttribution)
		if err != nil {
			return nil, err
		}
		signupAttribution.Id = bson.ObjectIdHex(id)

		lastReferrer = signupAttribution.Referred
		chain = append(chain, signupAttribution)
	}

	return chain, nil
}
