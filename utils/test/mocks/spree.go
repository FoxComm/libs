package mocks

import (
	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/FoxComm/FoxComm/spree"
	"github.com/gin-gonic/gin"
)

type orders []*spree.Order

// FakeSpree is a mock Spree endpoint for use during testing
type FakeSpree struct {
	Server *httptest.Server
	users  []*spree.User
	orders map[int]orders
}

// NewFakeSpree creates a new FakeSpree object.
func NewFakeSpree() *FakeSpree {
	fss := &FakeSpree{
		users:  make([]*spree.User, 0),
		orders: make(map[int]orders),
	}

	r := gin.New()
	r.GET("/app/api/users/:user_id", fss.HandleGetUser)
	r.GET("/app/api/orders", fss.HandleUserOrders)

	fss.Server = httptest.NewServer(r)
	return fss
}

// AddUser adds a Spree::User to the mock.
func (fss *FakeSpree) AddUser(user *spree.User) {
	fss.users = append(fss.users, user)
}

// AddOrder adds a Spree::Order that's associated with a User.
func (fss *FakeSpree) AddOrder(userID int, order *spree.Order) {
	userOrders := fss.orders[userID]
	userOrders = append(userOrders, order)
	fss.orders[userID] = userOrders
}

// HandleGetUser handles an HTTP request that returns a Spree::User based on ID.
func (fss *FakeSpree) HandleGetUser(c *gin.Context) {
	id := c.Params.ByName("user_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	for _, user := range fss.users {
		if uid, err := strconv.Atoi(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		} else if uid == user.Id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{})
	return
}

// HandleUserOrders handles an HTTP request for a user's orders.
func (fss *FakeSpree) HandleUserOrders(c *gin.Context) {
	id := c.Request.URL.Query().Get("q[user_id_eq]")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	uid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	orders := fss.orders[uid]
	result := map[string]interface{}{
		"count":       len(orders),
		"total_count": len(orders),
		"per_page":    len(orders),
		"pages":       1,
		"results":     orders,
	}
	c.JSON(http.StatusOK, result)
	return
}
