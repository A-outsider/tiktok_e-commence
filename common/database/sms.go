package database

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/mitchellh/mapstructure"
	"log"
)

type Phone struct {
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	RegionId        string `yaml:"regionId"`
}

func NewSms(val any) *dysmsapi20170525.Client {

	conf := new(Phone)

	var client *dysmsapi20170525.Client
	var err error
	err = mapstructure.Decode(val, conf) // 为结构体赋值
	if err != nil {
		log.Panicf("error decoding config to Pgsql struct: %v", err)
	}

	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(conf.AccessKeyId),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(conf.AccessKeySecret),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")

	client, err = dysmsapi20170525.NewClient(config)
	if err != nil {
		panic(err)
	}

	return client
}
