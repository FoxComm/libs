package spree

import (
	"fmt"
	"net/url"

	"github.com/FoxComm/FoxComm/configs"
	"github.com/FoxComm/FoxComm/endpoints"
)

type SpreeEndpoint endpoints.Endpoint

func (ep *SpreeEndpoint) Url() string {
	url, _ := url.Parse(fmt.Sprintf("%v%v", ep.Domain, ep.APIPrefix))
	return url.String()
}

var SpreeAPI = &SpreeEndpoint{
	Name:      "spree",
	Domain:    configs.Get("OriginHost"),
	APIPrefix: "/app/api",
}
