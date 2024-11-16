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

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type Jaeger struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type Static struct {
	AvatarPath string `yaml:"avatar_path"`
}

type Config struct {
	Service Service `yaml:"service"`
	Mysql   Mysql   `yaml:"mysql"`
	Redis   Redis   `yaml:"redis"`
	Jaeger  Jaeger  `yaml:"jaeger"`
	Static  Static  `yaml:"static"`
}

func GetConf() *Config {
	return &conf
}
