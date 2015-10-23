package endpoints

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/FoxComm/libs/configs"
	"github.com/FoxComm/libs/logger"
	"github.com/FoxComm/libs/utils"
)

type SocialShoppingEndpoint Endpoint

var SocialShoppingAPI *SocialShoppingEndpoint

func init() {
	SocialShoppingAPI = &SocialShoppingEndpoint{
		Name:        "social_shopping",
		Description: "Social Shopping",
		Domain:      configs.Get("SocialShoppingHost"),
		DefaultPort: configs.Get("SocialShoppingPort"),
		APIPrefix:   configs.Get("SocialShoppingAPIPrefix"),
		IsFeature:   true,
	}

	Add((*Endpoint)(SocialShoppingAPI))
}

func (ep *SocialShoppingEndpoint) GetUserFromToken(storeUrl string, token string) int {
	url := storeUrl + (*Endpoint)(ep).APIPrefix + "/users/" + token + "/attributes"
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := utils.GetHttpSslFlexibleClient().Do(req)
	defer resp.Body.Close()
	if err == nil && resp.StatusCode == 200 {
		parsedResp := map[string]interface{}{}
		respbytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(respbytes, &parsedResp)
		logger.Debug("[SocialShoppingAPI] user attributes: %+v", parsedResp)
		return int(parsedResp["OriginUserId"].(float64))
	}
	return 0
}

func (ep *SocialShoppingEndpoint) GetEntityIdFromToken(storeUrl string, token string) (int, string) {
	logger.Debug("[Social Shopping] Asking SS for the owner of the token: %s", token)

	url := storeUrl + (*Endpoint)(ep).APIPrefix + "/api/entity/" + token
	req, _ := http.NewRequest("GET", url, nil)

	if resp, err := utils.GetHttpSslFlexibleClient().Do(req); err == nil && resp.StatusCode == 200 {
		parsedResp := map[string]interface{}{}
		defer resp.Body.Close()

		respBytes, _ := ioutil.ReadAll(resp.Body)

		json.Unmarshal(respBytes, &parsedResp)
		logger.Debug("[Social Shopping] Response encoded: %-v", parsedResp)

		return int(parsedResp["EntityId"].(float64)), parsedResp["EntityType"].(string)
	} else {
		logger.Error("[Social Shopping] Error trying to contact social shopping: %s", err)
	}
	return 0, ""
}

func (ep *SocialShoppingEndpoint) GetEntityFromToken(storeUrl string, token string) (map[string]interface{}, error) {

	url := storeUrl + (*Endpoint)(ep).APIPrefix + "/api/entity_details/" + token
	logger.Debug("[Social Shopping] GetEntityFromToken storeUrl=%s token=%s full_url=%s", storeUrl, token, url)
	req, err := http.NewRequest("GET", url, nil)

	parsedResp := map[string]interface{}{}

	if err != nil {
		logger.Error("[Social Shopping] error=%s", err)
		return parsedResp, err
	}

	resp, err := utils.GetHttpSslFlexibleClient().Do(req)

	if err != nil {
		logger.Error("[Social Shopping] error=%s", err)
		return parsedResp, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		logger.Warn("[Social Shopping] error=%s", "Entity Not Found")
		return parsedResp, errors.New("Entity Not Found")
	} else if resp.StatusCode != 200 {
		logger.Error("[Social Shopping] Unexpected error", logger.Map{"status": resp.StatusCode})
		return parsedResp, errors.New("Unexpected error")
	}

	err = json.NewDecoder(resp.Body).Decode(&parsedResp)

	if err != nil {
		logger.Error("[Social Shopping] error=%s", err)
		return parsedResp, err
	}

	logger.Debug("[Social Shopping] storeUrl=%s token=%s response=%+v", storeUrl, token, parsedResp)

	return parsedResp, err
}
