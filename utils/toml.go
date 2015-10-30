package utils

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/serenize/snaker"
	"io/ioutil"

	"github.com/FoxComm/libs/configs"
)

// Probably it's temporary solution to make usage of local FoxComm easily

var ServicesWithEndpoints = [...]string{"origin_backend", "health_check",
	"causes_cache", "catalog_cache", "user",
	"core", "loyalty_engine", "backups", "ui", "social_shopping", "social_analytics", "origin_frontend"}

type ServersConfig struct {
	Servers map[string][]Server
}

type Server struct {
	URL string
}

func NewServersConfig() *ServersConfig {
	cfg := ServersConfig{}
	cfg.Servers = make(map[string][]Server)
	return &cfg
}

func WriteDefaultServersConfig(filename string) {
	cfg := NewServersConfig()
	for _, service := range ServicesWithEndpoints {
		var url string

		switch service {
		case "health_check":
			url = "http://localhost:54354" // not used but required for now by vulcand Engine
		case "backups":
			url = fmt.Sprintf("http://localhost:%s", configs.Get("BackupsPort"))
		case "origin":
			fallthrough
		case "origin_frontend":
			fallthrough
		case "origin_backend":
			url = configs.Get("OriginHost")
		case "loyalty_engine":
			url = fmt.Sprintf("%s:%s", configs.Get("SocialAnalyticsHost"), configs.Get("SocialAnalyticsPort"))
		default:
			cfgHostName := snaker.SnakeToCamel(service) + "Host"
			cfgPortName := snaker.SnakeToCamel(service) + "Port"
			url = fmt.Sprintf("%s:%s", configs.Get(cfgHostName), configs.Get(cfgPortName))
		}

		server := Server{URL: url}
		cfg.Servers[service] = []Server{server}
	}

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(cfg); err != nil {
		panic(err)
	}

	ioutil.WriteFile(filename, buf.Bytes(), 0644)
}
