package spree

import (
	"net/http"

	"github.com/FoxComm/libs/Godeps/_workspace/src/github.com/jmcvetta/napping"
	"github.com/FoxComm/libs/logger"
	"github.com/FoxComm/libs/utils"
	_ "github.com/FoxComm/libs/utils/ssl"
)

func (ep *SpreeEndpoint) CreatePromotion(spreeToken, spreeUrl string, params interface{}) (*http.Response, error) {
	url := spreeUrl + ep.APIPrefix + "/promotions"
	logger.Debug("Creating promotion url=%s", url)
	session := &napping.Session{Client: utils.GetHttpSslFlexibleClient()}
	session.Header = &http.Header{}
	session.Header.Set("X-Spree-Token", spreeToken)
	res, err := session.Post(url, params, nil, nil)
	if err != nil {
		return nil, err
	}
	return res.HttpResponse(), err
}
