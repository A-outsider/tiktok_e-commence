package initialize

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	DB     *gorm.DB
	Client *redis.Client
}

var svcContext *ServiceContext = new(ServiceContext)

func Init() {
	initMysql()
	initRedis()
}

func GetServiceContext() *ServiceContext {
	return svcContext
}

func GetMysql() *gorm.DB {
	return GetServiceContext().DB
}

func GetRedis() *redis.Client {
	return GetServiceContext().Client
}
