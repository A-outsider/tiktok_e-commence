package config

import (
	"gomall/common/config"
)

var (
	ServerName = "auth"
	MID        = int64(1)
	EtcdAddr   = config.EtcdAddr
	conf       Config
)

type Email struct {
	Addresses      string `yaml:"addresses"`
	Email          string `yaml:"email"`
	Host           string `yaml:"host"`
	Name           string `yaml:"name"`
	Password       string `yaml:"password"`
	Port           int    `yaml:"port"`
	ExpirationTime string `yaml:"expiration_time"`
	SendInterval   string `yaml:"sendInterval"`
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

type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type PhotoCaptcha struct {
	Width    int     `yaml:"width"`
	Length   int     `yaml:"length"`
	MaxSkew  float64 `yaml:"maxSkew"`
	DotCount int     `yaml:"dotCount"`
	Expire   string  `yaml:"expire"`
	Height   int     `yaml:"height"`
}

type Jwt struct {
	AccessExpireTime  string `yaml:"accessExpireTime"`
	RefreshExpireTime string `yaml:"refreshExpireTime"`
	AccessSecret      string `yaml:"accessSecret"`
	RefreshSecret     string `yaml:"refreshSecret"`
	Issuer            string `yaml:"issuer"`
}

type Password struct {
	ErrorLimit    int    `yaml:"ErrorLimit"`
	ErrorLockTime string `yaml:"ErrorLockTime"`
}

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Email        Email        `yaml:"email"`
	Phone        interface{}  `yaml:"phone"`
	Mysql        Mysql        `yaml:"mysql"`
	Redis        Redis        `yaml:"redis"`
	Jaeger       Jaeger       `yaml:"jaeger"`
	PhotoCaptcha PhotoCaptcha `yaml:"photoCaptcha"`
	Jwt          Jwt          `yaml:"jwt"`
	Password     Password     `yaml:"password"`
	Service      Service      `yaml:"service"`
}

func GetConf() *Config {
	return &conf
}
