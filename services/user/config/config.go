package config

import (
	"gomall/common/config"
)

var (
	ServerName = "user"
	MID        = int64(1)
	EtcdAddr   = config.EtcdAddr
	conf       Config
)

type Mysql struct {
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Redis struct {
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Static struct {
	AvatarPath string `yaml:"avatarPath"`
}

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Mysql   Mysql   `yaml:"mysql"`
	Redis   Redis   `yaml:"redis"`
	Jaeger  Jaeger  `yaml:"jaeger"`
	Static  Static  `yaml:"static"`
	Service Service `yaml:"service"`
}

func GetConf() *Config {
	return &conf
}
