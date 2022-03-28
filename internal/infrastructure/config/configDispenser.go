package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func GetServerConfig() (*serverConfigEntity, error) {
	entity := new(serverConfigEntity)

	configFile, err := ioutil.ReadFile("../configs/serverConfig.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configFile, entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
