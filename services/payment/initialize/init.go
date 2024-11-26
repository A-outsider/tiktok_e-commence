package initialize

import (
	"github.com/redis/go-redis/v9"
	"github.com/smartwalle/alipay/v3"
	"gomall/kitex_gen/order/orderservice"
	"gorm.io/gorm"
)

type ServiceContext struct {
	DB        *gorm.DB
	Client    *redis.Client
	AlipayCli *alipay.Client
	OrderCli  orderservice.Client
}

var svcContext *ServiceContext = new(ServiceContext)

func Init() {
	initMysql()
	initRedis()
	InitAlipay()
	initOrderCli()
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

func GetAlipay() *alipay.Client { return GetServiceContext().AlipayCli }

func GetOrderClient() orderservice.Client {
	return GetServiceContext().OrderCli
}
