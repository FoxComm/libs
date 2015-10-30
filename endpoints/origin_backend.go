package endpoints

import (
	"time"

	"github.com/FoxComm/libs/Godeps/_workspace/src/github.com/mailgun/vulcan/location/httploc"
	"github.com/FoxComm/libs/configs"
)

var OriginBackendAPI *Endpoint

func init() {
	OriginBackendAPI = &Endpoint{
		Name:      "origin_backend",
		Domain:    configs.Get("OriginHost"),
		APIPrefix: configs.Get("OriginBackendPrefix"),
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

	Add(OriginBackendAPI)
}
