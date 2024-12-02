package config

import "gomall/common/config"

var (
	ServerName = "product"
	MID        = int64(1)
	EtcdAddr   = config.EtcdAddr
	conf       Config
)

type Redis struct {
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type Static struct {
	ProductPath string `yaml:"product_path"`
}

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
	Host                   string `yaml:"host"`
	Port                   int    `yaml:"port"`
	Username               string `yaml:"username"`
	Password               string `yaml:"password"`
	CertificateFingerprint string `yaml:"certificate_fingerprint"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Redis         Redis         `yaml:"redis"`
	Static        Static        `yaml:"static"`
	Service       Service       `yaml:"service"`
	Mysql         Mysql         `yaml:"mysql"`
	ElasticSearch ElasticSearch `yaml:"elasticSearch"`
	Jaeger        Jaeger        `yaml:"jaeger"`
}

func GetConf() *Config {
	return &conf
}
