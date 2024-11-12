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
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Redis struct {
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type Jaeger struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type Config struct {
	Service Service `yaml:"service"`
	Mysql   Mysql   `yaml:"mysql"`
	Redis   Redis   `yaml:"redis"`
	Jaeger  Jaeger  `yaml:"jaeger"`
}

func GetConf() *Config {
	return &conf
}
