package config

import "gomall/common/config"

var (
	ServerName = "gateway"
	MID        = int64(1)
	EtcdAddr   = config.EtcdAddr
	conf       Config
)

type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Static struct {
	AvatarPath  string `yaml:"avatarPath"`
	ProductPath string `yaml:"productPath"`
}

type Role struct {
	Model  string `yaml:"model"`
	Policy string `yaml:"policy"`
}

type Jwt struct {
	AccessExpireTime  string `yaml:"accessExpireTime"`
	RefreshExpireTime string `yaml:"refreshExpireTime"`
	AccessSecret      string `yaml:"accessSecret"`
	RefreshSecret     string `yaml:"refreshSecret"`
	Issuer            string `yaml:"issuer"`
}

type VisitLimit struct {
	RateLimitInterval string `yaml:"rateLimitInterval"`
	RateLimitCap      int    `yaml:"rateLimitCap"`
}

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Jaeger     Jaeger     `yaml:"jaeger"`
	Static     Static     `yaml:"static"`
	Role       Role       `yaml:"role"`
	Jwt        Jwt        `yaml:"jwt"`
	VisitLimit VisitLimit `yaml:"visitLimit"`
	Service    Service    `yaml:"service"`
}

func GetConf() *Config {
	return &conf
}
