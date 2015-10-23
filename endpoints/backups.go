package endpoints

import "github.com/FoxComm/libs/configs"

var BackupsAPI *Endpoint

func init() {
	BackupsAPI = &Endpoint{
		Name:        "backups",
		DefaultPort: configs.Get("BackupsPort"),
		APIPrefix:   configs.Get("BackupsPrefix"),
	}

	Add(BackupsAPI)
}
