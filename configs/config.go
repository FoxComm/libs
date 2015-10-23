package configs

import (
	"github.com/FoxComm/libs/etcd_client"
	"github.com/FoxComm/libs/logger"
	ct "github.com/daviddengcn/go-colortext"

	"fmt"
	"os"
	"strings"
)

var ENV = "development"

func init() {
	osenv := os.Getenv("FC_ENV")
	if osenv != "" {
		ENV = strings.ToUpper(osenv)
	}
}

var DefaultSettings = map[string]string{
	"Host":       "http://localhost",
	"RouterPort": "8000",

	// The idea here is that we'll have 999 ports for app-servers for each
	// service
	"CatalogCacheHost":      "http://localhost",
	"CatalogCachePort":      "2000",
	"CatalogCacheAPIPrefix": "/catalog",

	"CausesCacheHost":      "http://localhost",
	"CausesCachePort":      "9000",
	"CausesCacheAPIPrefix": "/app/api/causes",

	"SocialShoppingHost":      "http://localhost",
	"SocialShoppingPort":      "4000",
	"SocialShoppingAPIPrefix": "/foxcomm/social_shopping",

	"SocialAnalyticsHost":      "http://localhost",
	"SocialAnalyticsPort":      "5000",
	"SocialAnalyticsAPIPrefix": "/foxcomm/social_analytics",
	"LoyaltyEngineAPIPrefix":   "/foxcomm/loyalty_engine",

	"CoreHost":      "http://localhost",
	"CorePort":      "5500",
	"CoreAPIPrefix": "/foxcomm/core",

	"UIHost":      "http://localhost",
	"UIPort":      "7000",
	"UIAPIPrefix": "/foxcomm/ui",

	"UserHost":      "http://localhost",
	"UserPort":      "7500",
	"UserAPIPrefix": "/session",

	"OriginHost":           "http://localhost:8080",
	"OriginFrontendPrefix": "/",
	"OriginBackendPrefix":  "/admin/",

	"BackupsPrefix": "/foxcomm/backups",
	"BackupsPort":   "4500",

	"TomlServersDir": "/tmp/foxcomm_servers/",

	// FC specific configs
	"FC_ENV":         "development",
	"FC_CORE_DB_URL": "",

	//Debugging control
	"DEBUG_MONGO":     "disabled",
	"DB_AUTO_MIGRATE": "disabled",
	"PRIVATE_IPV4":    "localhost",

	// Etcd
	"EtcdConfigKeyPrefix":     "foxcommerce.com/core/config/",
	"EtcdTestConfigKeyPrefix": "foxcommerce.com/test/core/config/",

	"RedisHost": "localhost:6379",
}

var EtcdCache = make(map[string]string)

// Get (aka MustGet), will panic if not found, it should be only invoked in initialize process
func Get(name string) string {

	defer ct.ResetColor()

	//We first check the Env variable.  Then we check etcd.
	//Finally, we default to the settings that are defined above.
	if value := GetSettingFromEnv(name); value != "" {
		logger.Debug("Got variable %s from OS!", name)
		return value
	} else if value, _ := GetSettingFromEtcd(name); value != "" {
		logger.Debug("Got variable %s from ETCD!", name)
		return value
	} else if value := getEtcdCache(name); value != "" {
		logger.Debug("Got variable %s from ETCD CACHE!", name)
		return value
	} else if value, _ := DefaultSettings[name]; value != "" {
		logger.Debug("Got variable %s from Defaults!", name)
		return value
	} else {
		ct.ChangeColor(ct.Red, true, ct.Black, false)
		fmt.Println("The following are not available via OS Env or the default configuration:")
		fmt.Println("Setting: ", name)
		panic(fmt.Sprintf("%v is not set!", name))
	}
}

func GetSettingFromEnv(name string) string {
	val := os.Getenv(name + "_" + ENV)
	if val != "" {
		return val
	}
	return os.Getenv(name)
}

func GetSettingFromEtcd(name string) (string, error) {
	keyString := DefaultSettings["EtcdConfigKeyPrefix"] + name
	if os.Getenv("FC_ENV") == "test" {
		keyString = DefaultSettings["EtcdTestConfigKeyPrefix"] + name
	}

	resp, err := etcd_client.EtcdClient.Get(keyString, true, false)

	//If we get a value from etcd, let's cache it in case etcd goes down in the future
	if resp != nil {
		if resp.Node != nil {
			if resp.Node.Value != "" {
				setEtcdCache(keyString, resp.Node.Value)
				return resp.Node.Value, nil
			}
		}
	}

	return "", err
}

func setEtcdCache(keyStr, value string) {
	EtcdCache[keyStr] = value
}

func getEtcdCache(keyStr string) string {
	value, _ := EtcdCache[keyStr]
	return value
}

func Set(name, value string) {
	DefaultSettings[name] = value
}
