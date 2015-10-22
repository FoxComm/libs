package generators

import (
	"fmt"

	"github.com/FoxComm/FoxComm/models"
	"github.com/FoxComm/FoxComm/repositories"
	"github.com/FoxComm/FoxComm/spree"
	"github.com/FoxComm/FoxComm/utils/test"
	"github.com/gin-gonic/gin"
)

var totalUsers = 0

// SignupSiteActivity generates a SiteActivity that shows a signup action.
func SignupSiteActivity(referrerURL, sessionToken string, fss *test.FakeSocialShopping) models.SiteActivity {
	totalUsers = totalUsers + 1
	user := spree.User{
		Id:        totalUsers,
		Email:     fake.Email(),
		FirstName: fake.FirstName(),
		LastName:  fake.LastName(),
	}

	return models.SiteActivity{
		Action:             "signup",
		SharingActivityTag: 0,
		RefererURL:         referrerURL,
		LandingURL:         "",
		ApiRequestURL:      "/app/signup.js",
		SessionToken:       sessionToken,
		Entity:             models.NewUserEntity(user),
		SignupDetails: models.SignupActivityDetails{
			User:     user,
			IsSocial: false,
		},
		Type:     "onsite",
		StoreURL: fss.Server.URL,
	}
}

// UserSignup creates the signup SiteActivity for a user that has not been
// referred. It then submits the SiteActivity.
func UserSignup(c *gin.Context, fss *test.FakeSocialShopping) (signup models.SiteActivity, rt string, err error) {
	rootSessionToken := fake.Characters(8)
	rt = fake.Characters(16)
	signup = SignupSiteActivity("http://foxcommerce.com/signup", rootSessionToken, fss)

	signupUser := signup.SignupDetails.User
	fss.AddUser(rt, &signupUser)
	return
}

// ReferredUserSignup creates the inbound and signup SiteActivities for a user
// that has been referrer. It then submits the SiteActivities.
func ReferredUserSignup(
	c *gin.Context,
	referredToken string,
	fss *test.FakeSocialShopping) (signup models.SiteActivity, referralToken string, err error) {

	sessionToken := fake.Characters(8)
	referralToken = fake.Characters(16)

	inbound := ReferralSiteActivity(referredToken, sessionToken, fss)

	signupURL := "http://localhost:8000/signup"
	if referredToken != "" {
		signupURL = fmt.Sprintf("http://localhost:8000/signup?fc_at=4&fc_ut=%s", referredToken)
	}
	signup = SignupSiteActivity(signupURL, sessionToken, fss)

	user := signup.SignupDetails.User

	fss.AddUser(referralToken, &user)

	repo, err := repositories.NewSiteActivityRepo(c.Request)
	if err != nil {
		return
	}

	if err = repo.Create(&inbound); err != nil {
		return
	}

	err = repo.Create(&signup)
	return
}
