package initialize

import "github.com/redis/go-redis/v9"

type ServiceContext struct {
	Client *redis.Client
}

var svcContext *ServiceContext = new(ServiceContext)

func Init() {
	initRedis()
}

func GetServiceContext() *ServiceContext {
	return svcContext
}

func GetRedis() *redis.Client {
	return GetServiceContext().Client
}
