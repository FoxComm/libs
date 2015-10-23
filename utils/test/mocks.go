package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/FoxComm/FoxComm/social_analytics/models"
	"github.com/FoxComm/FoxComm/spree"
	"github.com/gin-gonic/gin"
)

type FakeSocialShopping struct {
	Server *httptest.Server
	users  map[string]*spree.User
}

// NewFakeSocialShopping creates a new fake social_shopping instance.
func NewFakeSocialShopping() *FakeSocialShopping {
	fss := &FakeSocialShopping{
		users: make(map[string]*spree.User),
	}
	r := gin.New()
	r.GET("/foxcomm/social_shopping/api/entity_details/:token", fss.HandleEntityDetails)

	fss.Server = httptest.NewServer(r)
	return fss
}

// AddUser adds a user mapping for use in the mock.
func (fss *FakeSocialShopping) AddUser(token string, user *spree.User) {
	fss.users[token] = user
}

// HandleEntityDetails handles an HTTP request that returns an entity based
// on a referral token.
func (fss *FakeSocialShopping) HandleEntityDetails(c *gin.Context) {
	token := c.Params.ByName("token")
	if token == "" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	user, ok := fss.users[token]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"type": "user", "entity": user})
}

func (fss *FakeSocialShopping) ReferralToken(entity models.Entity) (string, error) {
	for k, v := range fss.users {
		if entity.Id == v.Id {
			return k, nil
		}
	}
	return "", fmt.Errorf("User with ID=%d not found", entity.Id)
}
