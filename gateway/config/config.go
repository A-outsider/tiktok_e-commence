package config

import "gomall/common/config"

var (
	ServerName = "gateway"
	MID        = int64(1)
	EtcdAddr   = config.EtcdAddr
	conf       Config
)

type Server struct {
	Addr string `yaml:"addr"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Server Server `yaml:"server"`
	Jaeger Jaeger `yaml:"jaeger"`
}

func GetConf() *Config {
	return &conf
}
