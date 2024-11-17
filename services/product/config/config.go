package config

import "gomall/common/config"

var (
	ServerName = "product"
	MID        = int64(1)
	EtcdAddr   = config.EtcdAddr
	conf       Config
)

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Mysql struct {
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
}

type ElasticSearch struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Service       Service       `yaml:"service"`
	Mysql         Mysql         `yaml:"mysql"`
	ElasticSearch ElasticSearch `yaml:"elasticSearch"`
	Jaeger        Jaeger        `yaml:"jaeger"`
}

func GetConf() *Config {
	return &conf
}
