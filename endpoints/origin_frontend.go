package endpoints

import (
	"time"

	"github.com/mailgun/vulcan/location/httploc"
	"github.com/FoxComm/libs/configs"
)

var OriginFrontendAPI *Endpoint

func init() {
	OriginFrontendAPI = &Endpoint{
		Name:      "origin_frontend",
		Domain:    configs.Get("OriginHost"),
		APIPrefix: configs.Get("OriginFrontendPrefix"),
		Options: httploc.Options{
			Limits: httploc.Limits{
				MaxMemBodyBytes: 20971520,
				MaxBodyBytes:    -1,
			},
			Timeouts: httploc.Timeouts{
				Read:         (time.Duration(90) * time.Second),
				Dial:         (time.Duration(90) * time.Second),
				TlsHandshake: (time.Duration(10) * time.Second),
			},
		},
	}

	Add(OriginFrontendAPI)
}
