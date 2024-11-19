package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"gomall/common/database"
	"os"
	"testing"
)

type TestEs struct {
	Host                   string `json:"host"`
	Port                   int    `json:"port"`
	Username               string `json:"username"`
	Password               string `json:"password"`
	CertificateFingerprint string `json:"certificate_fingerprint"`
}

var (
	url = TestEs{
		Host:                   "124.222.151.35",
		Port:                   9200,
		Username:               "elastic",
		Password:               "=Bq7JViaHHw_hdCabKVv",
		CertificateFingerprint: "039e813e2fc4847f3863ab90b05fbfa76adaaffe236ba9b26dac41b2136723fe",
	}
)

func TestESIndex(t *testing.T) {
	es := database.NewElasticSearch(url)

	// 打开 JSON 文件
	bytes, err := os.ReadFile("./data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}

	req := create.NewRequest()
	req, err = req.FromJSON(string(bytes))
	if err != nil {
		t.Error(err)
	}

	res, err := es.Indices.Create("product").
		Request(req).Do(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res)
}

func TestESAddDoc(t *testing.T) {
	//es := database.NewElasticSearch(url)

	// 打开 JSON 文件
	bytes, err := os.ReadFile("./doc.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(bytes))
	bytes, err = json.Marshal(bytes)

	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(bytes))
	//req := create.NewRequest()
	//
	//_, err = es.Create("product", "1").Request(bytes).Do(context.Background())
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
}

func TestSearch(t *testing.T) {
	client := database.NewElasticSearch(url)
	data, err := client.Search().
		Index("product").
		Request(&search.Request{
			Query: &types.Query{
				Term: map[string]types.TermQuery{
					"categories": {Value: "手机"},
				},
			},
		}).Do(context.Background())

	if err != nil {
		t.Error(err)
		return
	}

	for _, v := range data.Hits.Hits {
		res, err := v.Source_.MarshalJSON()
		if err != nil {
			panic(err)
		}

		fmt.Println(string(res))
	}
}
