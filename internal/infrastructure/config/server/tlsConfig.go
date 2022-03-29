package server

import "github.com/Gunga-D/proxy-server/internal/infrastructure/config"

type TLSServerConfigEntity struct {
	KeyPath string `yaml:"keyPath"`
	PemPath string `yaml:"pemPath"`
}

func GetServerTLSConfig() (*TLSServerConfigEntity, error) {
	entity := new(TLSServerConfigEntity)

	config.ReadConfig("../../configs/server/tls.yaml", entity)

	return entity, nil
}
