package es

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"gomall/services/product/dal/model"
	"gomall/services/product/initialize"
)

var (
	Analyzer = "ik_smart"
)

func AddProduct(ctx context.Context, data *model.Product) error {
	_, err := initialize.GetElasticSearchClient().Create("product", data.Pid).Request(data).Do(ctx)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(ctx context.Context, pid string) error {
	_, err := initialize.GetElasticSearchClient().Delete("product", pid).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func SearchProductByCategory(ctx context.Context, categoryName string) ([]model.Product, error) {
	client := initialize.GetElasticSearchClient()

	data, err := client.Search().
		Index("product").
		Request(&search.Request{
			Query: &types.Query{
				Term: map[string]types.TermQuery{
					"categories": {Value: categoryName},
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

		err = json.Unmarshal(val, &res[i])
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func SearchProduct(ctx context.Context, query string) ([]model.Product, error) {
	client := initialize.GetElasticSearchClient()
	data, err := client.Search().
		Index("product").
		Request(&search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"name": {Query: query, Analyzer: &Analyzer},
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
