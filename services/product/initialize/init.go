package initialize

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"gomall/services/product/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	DB     *gorm.DB
	ES     *elasticsearch.TypedClient
	Client *redis.Client
}

var svcContext *ServiceContext = new(ServiceContext)

func Init() {
	fmt.Println(config.GetConf())
	initMysql()
	initElasticSearch()
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

func GetElasticSearchClient() *elasticsearch.TypedClient {
	return GetServiceContext().ES
}
