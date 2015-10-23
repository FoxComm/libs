package spree

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/FoxComm/libs/logger"
	"github.com/FoxComm/libs/utils"
)

type Cause struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Featured    bool        `json:"featured"`
	Banner      interface{} `json:"banner"`
	Icon        interface{} `json:"icon"`
	HasBanner   bool        `json:"has_banner"`
	HasIcon     bool        `json:"has_icon"`
}

func (ep *SpreeEndpoint) RequestCause(spreeToken string, userId int, storeHostName string) (*Cause, error) {
	httpReq, _ := http.NewRequest("GET", storeHostName+ep.APIPrefix+"/causes/"+strconv.Itoa(userId), nil)
	httpReq.Header["X-Spree-Token"] = []string{spreeToken}

	logger.Debug("Request: %+v", httpReq)
	if resp, err := utils.GetHttpSslFlexibleClient().Do(httpReq); err == nil {
		defer resp.Body.Close()
		logger.Debug("Spree response: %+v", resp)
		if resp.StatusCode == 200 {
			var cause Cause
			decoder := json.NewDecoder(resp.Body)
			err = decoder.Decode(&cause)
			return &cause, nil
		} else {
			return nil, errors.New("unexpected status: " + resp.Status)
		}
	} else {
		logger.Debug("There was an error verifying cause%s", err)
		return nil, err
	}
}

func (ep *SpreeEndpoint) RequestCauses(spreeToken, storeHost, query string) (*[]Cause, error) {
	query = "q[name_cont]=" + query
	httpReq, _ := http.NewRequest("GET", storeHost+ep.APIPrefix+"/causes?"+query, nil)
	httpReq.Header["X-Spree-Token"] = []string{spreeToken}

	if resp, err := utils.GetHttpSslFlexibleClient().Do(httpReq); err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			var causes struct {
				Results []Cause `json:results`
			}
			decoder := json.NewDecoder(resp.Body)
			err = decoder.Decode(&causes)
			return &causes.Results, nil
		} else {
			return nil, errors.New("unexpected status: " + resp.Status)
		}
	} else {
		logger.Debug("There was an error verifying user %s", err)
		return nil, err
	}
}
