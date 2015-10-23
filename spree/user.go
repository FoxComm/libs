package spree

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/FoxComm/libs/logger"
	"github.com/FoxComm/libs/utils"
)

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id           int       `json:"id"`
	Email        string    `json:"email"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CurrentOrder string    `json:"current_order"`
	ActiveCause  Cause     `json:"active_cause"`
	Roles        []Role    `json:"spree_roles"`
}

func (user *User) IsAdmin() bool {
	for _, role := range user.Roles {
		if role.Name == "admin" {
			return true
		}
	}
	return false
}

func (user *User) Name() string {
	if user.FirstName == "" {
		return user.Email
	} else {
		return user.FirstName + " " + user.LastName
	}
}

func (ep *SpreeEndpoint) RequestUser(spreeToken string, userId int, storeHostName string) (*User, error) {
	url := fmt.Sprintf("%s%s/users/%s", storeHostName, ep.APIPrefix, strconv.Itoa(userId))
	httpReq, _ := http.NewRequest("GET", url, nil)
	httpReq.Header["X-Spree-Token"] = []string{spreeToken}

	logger.Debug("Request: %+v", httpReq)
	if resp, err := utils.GetHttpSslFlexibleClient().Do(httpReq); err == nil {
		logger.Debug("Spree response: %+v", resp)
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			var user User
			decoder := json.NewDecoder(resp.Body)
			err = decoder.Decode(&user)
			return &user, err
		} else {
			return nil, fmt.Errorf(
				"Status %s when requesting Spree user %d from endpoint %s",
				resp.Status,
				userId,
				url,
			)
		}
	} else {
		logger.Debug("There was an error verifying user %s", err)
		return nil, err
	}
}

func (ep *SpreeEndpoint) RequestUsers(spreeToken, storeHost, query string) (*[]User, error) {
	query = "q[first_name_or_last_name_cont]=" + query
	httpReq, _ := http.NewRequest("GET", storeHost+ep.APIPrefix+"/users?"+query, nil)
	httpReq.Header["X-Spree-Token"] = []string{spreeToken}

	if resp, err := utils.GetHttpSslFlexibleClient().Do(httpReq); err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			var users struct {
				Results []User `json:results`
			}
			decoder := json.NewDecoder(resp.Body)
			err = decoder.Decode(&users)
			return &users.Results, err
		} else {
			return nil, errors.New("unexpected status: " + resp.Status)
		}
	} else {
		logger.Debug("There was an error verifying user %s", err)
		return nil, err
	}
}
