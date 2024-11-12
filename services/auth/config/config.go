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

type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Email struct {
	Addresses      interface{} `yaml:"addresses"`
	Email          interface{} `yaml:"email"`
	Host           interface{} `yaml:"host"`
	Name           string      `yaml:"name"`
	Password       string      `yaml:"password"`
	Port           int         `yaml:"port"`
	ExpirationTime string      `yaml:"expiration_time"`
}

type PhotoCaptcha struct {
	Length   int     `yaml:"length"`
	MaxSkew  float64 `yaml:"maxSkew"`
	DotCount int     `yaml:"dotCount"`
	Expire   string  `yaml:"expire"`
	Height   int     `yaml:"height"`
	Width    int     `yaml:"width"`
}

type Config struct {
	Service      Service      `yaml:"service"`
	Mysql        Mysql        `yaml:"mysql"`
	Redis        Redis        `yaml:"redis"`
	Jaeger       Jaeger       `yaml:"jaeger"`
	Phone        interface{}  `yaml:"phone"`
	Email        Email        `yaml:"email"`
	PhotoCaptcha PhotoCaptcha `yaml:"photoCaptcha"`
}

func GetConf() *Config {
	return &conf
}
