package endpoints

import (
	"github.com/FoxComm/libs/configs"
)

var UIAPI *Endpoint

func init() {
	UIAPI = &Endpoint{
		Name:      "ui",
		Domain:    configs.Get("UIHost"),
		APIPrefix: configs.Get("UIAPIPrefix"),
	}

	Add(UIAPI)
}
