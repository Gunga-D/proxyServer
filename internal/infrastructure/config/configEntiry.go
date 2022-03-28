package config

type serverConfigEntity struct {
	Ip             string `yaml:"ip"`
	Port           string `yaml:"port"`
	MaxHeaderBytes int    `yaml:"maxHeaderBytes"`
	ReadTimeout    int    `yaml:"readTimeoutInSeconds"`
	WriteTimeout   int    `yaml:"writeTimeoutInSeconds"`
}
