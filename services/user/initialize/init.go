package initialize

import (
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"gorm.io/gorm"
)

type ServiceContext struct {
	DB  *gorm.DB
	RDB *RedisClient
	SMS *dysmsapi20170525.Client
}

var svcContext *ServiceContext = new(ServiceContext)

func Init() {
	initMysql()
	InitRedis()
}

func GetServiceContext() *ServiceContext {
	return svcContext
}

func GetMysql() *gorm.DB {
	return GetServiceContext().DB
}

func GetRedis() *RedisClient {
	return GetServiceContext().RDB
}
