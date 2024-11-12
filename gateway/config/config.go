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

type Password struct {
	ErrorLockTime string `yaml:"ErrorLockTime"`
	ErrorLimit    int    `yaml:"ErrorLimit"`
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

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type Config struct {
	Jaeger     Jaeger     `yaml:"jaeger"`
	Role       Role       `yaml:"role"`
	Jwt        Jwt        `yaml:"jwt"`
	VisitLimit VisitLimit `yaml:"visitLimit"`
	Password   Password   `yaml:"password"`
	Service    Service    `yaml:"service"`
	Mysql      Mysql      `yaml:"mysql"`
	Redis      Redis      `yaml:"redis"`
}

func GetConf() *Config {
	return &conf
}
