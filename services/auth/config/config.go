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

type PhotoCaptcha struct {
	Height   int     `yaml:"height"`
	Width    int     `yaml:"width"`
	Length   int     `yaml:"length"`
	MaxSkew  float64 `yaml:"maxSkew"`
	DotCount int     `yaml:"dotCount"`
	Expire   string  `yaml:"expire"`
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

type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Phone struct {
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	RegionId        string `yaml:"regionId"`
	ExpirationTime  string `yaml:"expiration_time"`
	SendInterval    string `yaml:"sendInterval"`
}

type Redis struct {
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
}

type Jwt struct {
	Issuer            string `yaml:"issuer"`
	AccessExpireTime  string `yaml:"accessExpireTime"`
	RefreshExpireTime string `yaml:"refreshExpireTime"`
	AccessSecret      string `yaml:"accessSecret"`
	RefreshSecret     string `yaml:"refreshSecret"`
}

type Password struct {
	ErrorLimit    int    `yaml:"ErrorLimit"`
	ErrorLockTime string `yaml:"ErrorLockTime"`
}

type Config struct {
	Email        Email        `yaml:"email"`
	PhotoCaptcha PhotoCaptcha `yaml:"photoCaptcha"`
	Service      Service      `yaml:"service"`
	Mysql        Mysql        `yaml:"mysql"`
	Jaeger       Jaeger       `yaml:"jaeger"`
	Phone        Phone        `yaml:"phone"`
	Redis        Redis        `yaml:"redis"`
	Jwt          Jwt          `yaml:"jwt"`
	Password     Password     `yaml:"password"`
}

func GetConf() *Config {
	return &conf
}
