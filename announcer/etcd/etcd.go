package etcd

import (
	"fmt"
	"github.com/FoxComm/libs/etcd_client"
	"github.com/FoxComm/libs/logger"
)

type EtcdAnnouncer struct{}

func (e *EtcdAnnouncer) AnnounceStart(endpoint, host, port string) error {
	// Create endpoints dir if it don't exist
	etcd_client.EtcdClient.CreateDir("endpoints", 0)

	endpointDir := fmt.Sprintf("endpoints/%s", endpoint)
	etcd_client.EtcdClient.CreateDir(endpointDir, 0)

	key := fmt.Sprintf("%s/%s:%s", endpointDir, host, port)

	if resp, err := etcd_client.EtcdClient.Create(key, "", 0); err != nil {
		return err
	} else {
		logger.Debug("[Announcer:Start] Etcd response: %+v", *resp)
	}
	return nil
}

func (e *EtcdAnnouncer) AnnounceStop(endpoint, host, port string) error {
	key := fmt.Sprintf("endpoints/%s/%s:%s", endpoint, host, port)

	if resp, err := etcd_client.EtcdClient.Delete(key, true); err != nil {
		return err
	} else {
		logger.Debug("[Announcer:Stop] Etcd response: %+v", *resp)
	}
	return nil
}
