package spree

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/FoxComm/libs/logger"
	"github.com/FoxComm/libs/utils"
)

type Variant struct {
	Id        int
	ProductId int
}

type LineItem struct {
	Id      int
	Variant Variant `json:"variant"`
}

type Order struct {
	Id          int        `json:"id"`
	Number      string     `json:"number"`
	ItemTotal   float64    `json:"item_total,string"`
	CompletedAt time.Time  `json:"completed_at"`
	Total       float64    `json:"total,string"`
	LineItems   []LineItem `json:"line_items"`
}

type OrderResult struct {
	Count      int     `json:"count"`
	TotalCount int     `json:"total_count"`
	PerPage    int     `json:"per_page"`
	Pages      int     `json:"pages"`
	Orders     []Order `json:"orders"`
}

func (ep *SpreeEndpoint) MyOrders(spreeToken, spreeUrl string) (orders []Order, err error) {
	url := fmt.Sprintf("%v%v/orders/mine?q[state_eq]=complete", spreeUrl, ep.APIPrefix)
	return requestOrders(url, spreeToken)
}

func (ep *SpreeEndpoint) CompleteOrders(userId int, spreeToken, spreeUrl string) (orders []Order, err error) {
	url := fmt.Sprintf("%v%v/orders?q[state_eq]=complete&q[user_id_eq]=%v", spreeUrl, ep.APIPrefix, userId)
	return requestOrders(url, spreeToken)
}

func requestOrders(url string, spreeToken string) ([]Order, error) {
	logger.Debug("[Spree::Orders] Request for orders %s", url)

	httpClient := utils.GetHttpSslFlexibleClient()
	req, _ := http.NewRequest("GET", url, nil)
	req.Header["X-Spree-Token"] = []string{spreeToken}

	var orders []Order
	resp, err := httpClient.Do(req)

	if err != nil {
		logger.Error("There was an error fetching my orders %s", err)
		return orders, err
	}

	if resp.StatusCode != 200 {
		message := fmt.Sprintf("[Spree::Orders] Error loading data from Spree. Status: %d", resp.StatusCode)
		logger.Warn(message)
		return orders, errors.New(message)
	}

	defer resp.Body.Close()

	var orderResult OrderResult
	err = json.NewDecoder(resp.Body).Decode(&orderResult)

	if err != nil && err != io.EOF {
		logger.Warn("[Spree::Orders] Error decoding message %s", err.Error())
		return orders, err
	}

	return orderResult.Orders, nil
}

func (ep *SpreeEndpoint) Order(orderNumber, spreeToken, orderToken string) (Order, error) {
	url := "/orders/mine"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header["X-Spree-Token"] = []string{spreeToken}
	req.Header["X-Spree-Order-Token"] = []string{orderToken}

	if resp, err := utils.GetHttpSslFlexibleClient().Do(req); err == nil && resp.StatusCode == 200 {
		var order Order
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&order); err == nil {
			return order, nil
		}
	} else {
		logger.Error("There was an error fetching my orders %s", err)
	}
	return Order{}, errors.New("Order not found")
}
