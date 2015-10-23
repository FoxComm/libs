package endpoints

import (
	"github.com/FoxComm/libs/configs"
	"github.com/FoxComm/libs/etcd_client"
	"github.com/FoxComm/libs/logger"
	"github.com/FoxComm/core_services/feature_manager/core"
	"github.com/coreos/go-etcd/etcd"
)

var CoreAPI *Endpoint

func WatchFeatureManagerChanges() {
	go func() {
		resp := make(chan *etcd.Response)
		go etcd_client.EtcdClient.Watch("feature_manager_updated_at", 0, true, resp, nil)

		for res := range resp {
			logger.Debug("[FeatureManger] Clearing cached store feature map", res)
			core.ClearCacheMap()
		}
	}()
}

func init() {
	CoreAPI = &Endpoint{
		Name:        "core",
		Domain:      configs.Get("CoreHost"),
		DefaultPort: configs.Get("CorePort"),
		APIPrefix:   configs.Get("CoreAPIPrefix"),
	}
	WatchFeatureManagerChanges()

	Add(CoreAPI)
}
