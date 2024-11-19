package database

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/mitchellh/mapstructure"
	"log"
)

type ElasticSearch struct {
	Host                   string `json:"host"`
	Port                   int    `json:"port"`
	Username               string `json:"username"`
	PassWord               string `json:"password"`
	CertificateFingerprint string `json:"certificate_fingerprint"`
}

func NewElasticSearch(val any) *elasticsearch.TypedClient {

	var err error

	conf := new(ElasticSearch)

	err = mapstructure.Decode(val, conf) // 为结构体赋值
	if err != nil {
		log.Panicf("error decoding config to elasticSearch struct: %v", err)
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("https://%s:%d", conf.Host, conf.Port),
		},
		Username:               conf.Username,
		Password:               conf.PassWord,
		CertificateFingerprint: conf.CertificateFingerprint,
	}

	es, err := elasticsearch.NewTypedClient(cfg)

	if err != nil {
		panic(err)
	}

	return es

}
