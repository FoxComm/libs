package etcd_client

import (
	"fmt"
	"os"

	"github.com/coreos/go-etcd/etcd"
	"github.com/FoxComm/libs/logger"
)

var EtcdDefaults = map[string]string{
	"Host": "127.0.0.1",
	"Port": "4001",
}
var EtcdClient *etcd.Client

func init() {
	var EtcdHost, EtcdPort, EtcdFullHost string
	// Check the OS first, then let's default to what is set above.
	if host := os.Getenv("EtcdHost"); host != "" {
		EtcdHost = host
	} else {
		EtcdHost = EtcdDefaults["Host"]
	}

	if port := os.Getenv("EtcdPort"); port != "" {
		EtcdPort = port
	} else {
		EtcdPort = EtcdDefaults["Port"]
	}
	EtcdFullHost = fmt.Sprintf("http://%s:%s", EtcdHost, EtcdPort)
	logger.Debug("Mounting etcd watcher on: %s", EtcdFullHost)
	EtcdClient = etcd.NewClient([]string{EtcdFullHost})
}
