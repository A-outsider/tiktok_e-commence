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
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type ElasticSearch struct {
	Host                   string `json:"host"`
	Port                   int    `json:"port"`
	Username               string `json:"username"`
	PassWord               string `json:"password"`
	CertificateFingerprint string `json:"certificate_fingerprint"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Static struct {
	ProductPath string `yaml:"product_path"`
}

type Config struct {
	Service       Service       `yaml:"service"`
	Mysql         Mysql         `yaml:"mysql"`
	ElasticSearch ElasticSearch `yaml:"elasticSearch"`
	Jaeger        Jaeger        `yaml:"jaeger"`
	Static        Static        `yaml:"static"`
}

func GetConf() *Config {
	return &conf
}
