package toml

import (
	"bytes"
	"fmt"
	"github.com/FoxComm/libs/Godeps/_workspace/src/github.com/BurntSushi/toml"
	"github.com/FoxComm/libs/configs"
	"github.com/FoxComm/libs/utils"
	"io/ioutil"
	"os"
	"path"
)

type TomlAnnouncer struct{}

func getFileLocation(endpoint string) string {
	return fmt.Sprintf("%s.toml", path.Join(configs.Get("TomlServersDir"), endpoint))
}

func (e *TomlAnnouncer) AnnounceStart(endpoint, host, port string) error {
	// Create endpoints dir if it don't exist

	os.Mkdir(configs.Get("TomlServersDir"), 0755)
	filename := getFileLocation(endpoint)
	// TODO: https?
	url := fmt.Sprintf("http://%s:%s", host, port)
	if err := writeServerConfig(endpoint, url, filename); err != nil {
		return err
	}
	return nil
}

func (e *TomlAnnouncer) AnnounceStop(endpoint, host, port string) error {
	return os.Remove(getFileLocation(endpoint))
}

func NewAnnouncer() *TomlAnnouncer {
	return &TomlAnnouncer{}
}

func writeServerConfig(service string, url string, filename string) error {
	cfg := utils.NewServersConfig()
	cfg.Servers[service] = []utils.Server{utils.Server{URL: url}}
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(cfg); err != nil {
		return err
	}

	return ioutil.WriteFile(filename, buf.Bytes(), 0644)
}
