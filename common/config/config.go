package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var Common = new(Default)

// Default
type Default struct {
	Nacos  Nacos  `yaml:"nacos"`
	Log    Log    `yaml:"log"`
	Jaeger Jaeger `yaml:"jaeger"`
}

// Jaeger
type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Nacos
type Nacos struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Log
type Log struct {
	BasePath   string `yaml:"basePath"`
	MaxAge     int    `yaml:"maxAge"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxSize    int    `yaml:"maxSize"`
	Mode       string `yaml:"mode"`
	Level      string `yaml:"level"`
}

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	vip := viper.New()
	vip.AddConfigPath(path + "/common/config")
	vip.SetConfigName("common")
	vip.SetConfigType("yaml")

	if err := vip.ReadInConfig(); err != nil {
		log.Println(err)
	}

	if err = vip.Unmarshal(Common); err != nil {
		log.Println(err)
	}

}
