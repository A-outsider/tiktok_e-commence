package initialize

import (
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	DB     *gorm.DB
	RDB    *RedisClient
	SMS    *dysmsapi20170525.Client
	Client *redis.Client
}

var svcContext *ServiceContext = new(ServiceContext)

func Init() {
	initMysql()
	initSms()
	InitRedis()
	initRedisWithNoContext()
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

func GetRedisWithNoContext() *redis.Client {
	return GetServiceContext().Client
}

func GetSms() *dysmsapi20170525.Client {
	return GetServiceContext().SMS
}
