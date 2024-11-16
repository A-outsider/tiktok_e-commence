package initialize

import (
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"gomall/common/database"
	"gorm.io/gorm"
)

type ServiceContext struct {
	DB  *gorm.DB
	RDB *database.RedisClient
	SMS *dysmsapi20170525.Client
}

var svcContext *ServiceContext = new(ServiceContext)

func Init() {
	initMysql()
	initSms()
	InitRedis()
}

func GetServiceContext() *ServiceContext {
	return svcContext
}

func GetMysql() *gorm.DB {
	return GetServiceContext().DB
}

func GetRedis() *database.RedisClient {
	return GetServiceContext().RDB
}

func GetSms() *dysmsapi20170525.Client {
	return GetServiceContext().SMS
}
