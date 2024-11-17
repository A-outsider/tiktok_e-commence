package initialize

import (
	"gomall/common/database"
	"gomall/services/product/config"
)

func initElasticSearch() {
	svcContext.ES = database.NewElasticSearch(config.GetConf().ElasticSearch)
}
