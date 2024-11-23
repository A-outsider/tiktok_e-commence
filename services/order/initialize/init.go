package initialize

import (
	"gorm.io/gorm"
)

type ServiceContext struct {
	DB *gorm.DB
}

var svcContext *ServiceContext = new(ServiceContext)

func Init() {
	initMysql()
}

func GetServiceContext() *ServiceContext {
	return svcContext
}

func GetMysql() *gorm.DB {
	return GetServiceContext().DB
}
