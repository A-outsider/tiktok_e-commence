package es

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"gomall/services/product/dal/model"
	"gomall/services/product/initialize"
)

func AddProduct(ctx context.Context, data *model.Product) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}

	initialize.GetElasticSearchClient().Create("product", data.Pid).
		Request(marshal).Do(ctx)

	return nil
}

func SearchProductByCategory(ctx context.Context, categoryName string) ([]model.Product, error) {
	client := initialize.GetElasticSearchClient()
	data, err := client.Search().
		Index("test").
		Request(&search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"categories": {Query: categoryName},
				},
			},
		}).Do(ctx)

	if err != nil {
		return nil, err
	}

	res := make([]model.Product, len(data.Hits.Hits))

	for i, v := range data.Hits.Hits {
		val, err := v.Source_.MarshalJSON()
		if err != nil {
			return nil, err
		}

		json.Unmarshal(val, &res[i])
	}

	return res, nil
}
