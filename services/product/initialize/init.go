package initialize

import (
	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	DB *gorm.DB
	ES *elasticsearch.TypedClient
}

var svcContext *ServiceContext = new(ServiceContext)

func Init() {
	initMysql()
	initElasticSearch()
}

func GetServiceContext() *ServiceContext {
	return svcContext
}

func GetMysql() *gorm.DB {
	return GetServiceContext().DB
}

func GetElasticSearchClient() *elasticsearch.TypedClient {
	return GetServiceContext().ES
}
