package endpoints

import "github.com/FoxComm/libs/configs"

var UserAPI *Endpoint

func init() {
	UserAPI = &Endpoint{
		Name:        "user",
		Domain:      configs.Get("UserHost"),
		DefaultPort: configs.Get("UserPort"),
		APIPrefix:   configs.Get("UserAPIPrefix"),
		IsFeature:   true,
	}

	Add(UserAPI)
}
