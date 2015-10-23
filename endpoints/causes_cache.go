package endpoints

import (
	"fmt"

	"github.com/FoxComm/libs/configs"
)

var CausesCacheAPI *Endpoint

func init() {
	CausesCacheAPI = &Endpoint{
		Name:        "causes_cache",
		Domain:      configs.Get("CausesCacheHost"),
		DefaultPort: configs.Get("CausesCachePort"),
		APIPrefix:   configs.Get("CausesCacheAPIPrefix"),
		routePrefix: fmt.Sprintf(`TrieRoute("GET", "%v")`, configs.Get("CausesCacheAPIPrefix")),
	}

	Add(CausesCacheAPI)
}
