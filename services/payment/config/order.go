package config

import "gomall/common/config"

var (
	ServerName = "payment"
	MID        = int64(1)
	EtcdAddr   = config.EtcdAddr
	conf       Config
)

type Mysql struct {
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
}

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
	Mysql   Mysql   `yaml:"mysql"`
	Service Service `yaml:"service"`
	Redis   Redis   `yaml:"redis"`
	Jaeger  Jaeger  `yaml:"jaeger"`
}

func GetConf() *Config {
	return &conf
}
