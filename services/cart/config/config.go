package config

import "gomall/common/config"

var (
	ServerName = "cart"
	MID        = int64(1)
	EtcdAddr   = config.EtcdAddr
	conf       Config
)

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Service Service `yaml:"service"`
	Redis   Redis   `yaml:"redis"`
	Jaeger  Jaeger  `yaml:"jaeger"`
}

func GetConf() *Config {
	return &conf
}
