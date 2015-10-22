package fixtures

import (
	"io/ioutil"

	"gopkg.in/yaml.v1"
)

func Load(fixtures string, result interface{}) error {
	fixtureFile, err := ioutil.ReadFile("db/fixtures/" + fixtures + ".yml")
	if err == nil {
		yaml.Unmarshal(fixtureFile, result)
	}
	return err
}
