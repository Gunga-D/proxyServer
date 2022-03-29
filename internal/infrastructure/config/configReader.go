package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ReadConfig(path string, entity interface{}) error {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configFile, entity)
	if err != nil {
		return err
	}

	return nil
}
