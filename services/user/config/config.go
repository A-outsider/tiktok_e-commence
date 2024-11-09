package config

import (
	"fmt"
	"gomall/common/config"
)

var (
	ServerName = "user"
	MID        = int64(1)
	EtcdAddr   = fmt.Sprintf("%s:%d", config.Common.Etcd.Host, config.Common.Etcd.Port)
	ServerAddr = "192.168.40.1:8081"
	conf       Config
)

type Config struct {
}

func GetConf() *Config {
	return &conf
}
