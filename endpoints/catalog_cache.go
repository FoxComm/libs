package endpoints

import "github.com/FoxComm/libs/configs"

var CatalogCacheAPI *Endpoint

func init() {
	CatalogCacheAPI = &Endpoint{
		Name:        "catalog_cache",
		Domain:      configs.Get("CatalogCacheHost"),
		DefaultPort: configs.Get("CatalogCachePort"),
		APIPrefix:   configs.Get("CatalogCacheAPIPrefix"),
	}

	Add(CatalogCacheAPI)
}
