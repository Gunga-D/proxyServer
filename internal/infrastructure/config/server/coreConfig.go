package server

import "github.com/Gunga-D/proxy-server/internal/infrastructure/config"

type CoreServerConfigEntity struct {
	Ip             string `yaml:"ip"`
	Port           string `yaml:"port"`
	MaxHeaderBytes int    `yaml:"maxHeaderBytes"`
	ReadTimeout    int    `yaml:"readTimeoutInSeconds"`
	WriteTimeout   int    `yaml:"writeTimeoutInSeconds"`
}

func GetServerCoreConfig() (*CoreServerConfigEntity, error) {
	entity := new(CoreServerConfigEntity)

	config.ReadConfig("../configs/server/core.yaml", entity)

	return entity, nil
}
