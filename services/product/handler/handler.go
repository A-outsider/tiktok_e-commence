package handler

import (
	"context"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gomall/gateway/types/resp/common"
	product "gomall/kitex_gen/product"
	"gomall/services/product/config"
	"gomall/services/product/dal/cache"
	"gomall/services/product/dal/db"
	"gomall/services/product/dal/es"
	"gomall/services/product/dal/model"
	"os"
	"path/filepath"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

func (s *ProductCatalogServiceImpl) GetRankings(ctx context.Context, req *product.GetRankingsReq) (resp *product.GetRankingsResp, _ error) {
	//TODO implement me
	resp = new(product.GetRankingsResp)
	resp.StatusCode = common.CodeServerBusy

	details, err := cache.GetRankingsWithDetails(ctx, 100)
	if err != nil {
		zap.L().Error("GetRankings failed", zap.Error(err))
		return
	}

	resp.ProductItems = make([]*product.ProductItem, len(details))
	err = copier.Copy(&resp.ProductItems, details)
	if err != nil {
		zap.L().Error("GetRankings failed", zap.Error(err))
		return
	}

	resp.StatusCode = common.CodeSuccess
	return
}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, _ error) {
	// TODO: Your code here...
	resp = new(product.ListProductsResp)
	resp.StatusCode = common.CodeServerBusy
	data, err := es.SearchProductByCategory(ctx, req.CategoryName)
	if err != nil {
		return
	}

	resp.Products = make([]*product.Product, len(data))
	err = copier.Copy(&resp.Products, data)
	if err != nil {
		zap.L().Error("copier.Copy products failed", zap.Error(err))
		return
	}
	resp.StatusCode = common.CodeSuccess
	return
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, _ error) {
	// TODO: Your code here...
	resp = new(product.GetProductResp)
	resp.StatusCode = common.CodeServerBusy
	resp.Product = new(product.Product)

	data, err := db.GetProductByPid(ctx, req.GetId())

	err = copier.Copy(&resp.Product, data)
	if err != nil {
		zap.L().Error("copier.Copy product failed", zap.Error(err))
		return
	}

	// 增加商品热度, 因为错误不会影响到商品的正常使用, 所以忽略该错误
	err = cache.IncrProductHotness(ctx, req.GetId())
	if err != nil {
		zap.L().Error("cache.IncrProductHotness failed", zap.Error(err))
	}

	resp.Product.Categories = *data.Categories
	resp.StatusCode = common.CodeSuccess

	return
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, _ error) {
	// TODO: Your code here...
	resp = new(product.SearchProductsResp)
	resp.StatusCode = common.CodeServerBusy

	data, err := es.SearchProduct(ctx, req.Query)
	if err != nil {
		return
	}

	resp.Results = make([]*product.Product, len(data))
	err = copier.Copy(&resp.Results, data)
	if err != nil {
		zap.L().Error("copier.Copy products failed", zap.Error(err))
		return
	}
	resp.StatusCode = common.CodeSuccess
	return
}

func (s *ProductCatalogServiceImpl) AddProduct(ctx context.Context, req *product.AddProductReq) (resp *product.AddProductResp, _ error) {
	resp = new(product.AddProductResp)
	resp.StatusCode = common.CodeServerBusy

	product := &model.Product{
		Pid:         uuid.New().String(),
		Bid:         req.Product.Bid,
		Uid:         req.Product.Uid,
		Picture:     req.Product.Picture,
		Price:       req.Product.Price,
		Description: req.Product.Description,
		Name:        req.Product.Name,
		Categories:  new(model.Strings),
	}

	// 保存商品图片
	fileName := product.Pid + req.Ext

	os.MkdirAll(config.GetConf().Static.ProductPath, 0755)
	if err := os.WriteFile(filepath.Join(config.GetConf().Static.ProductPath, fileName), req.Body, 0755); err != nil {
		zap.L().Error("os.WriteFile failed", zap.Error(err))
		return
	}

	*product.Categories = req.Product.Categories
	product.Picture = fileName

	err := db.AddProduct(ctx, product)
	if err != nil {
		zap.L().Error("Failed to create product", zap.Error(err))
		return
	}

	// 添加商品到es, 该代码仅为测试时使用
	err = es.AddProduct(ctx, product)
	if err != nil {
		zap.L().Error("Failed to create es product", zap.Error(err))
	}

	// 添加商品到热度排行榜, 因为不影响实际商品的用图, 所以这里忽略它的报错
	err = cache.AddProductToRanking(ctx, product)
	if err != nil {
		zap.L().Error("Failed to add product to ranking", zap.Error(err))
	}

	resp.StatusCode = common.CodeSuccess

	return
}

func (s *ProductCatalogServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq) (resp *product.DeleteProductResp, _ error) {
	resp = new(product.DeleteProductResp)
	resp.StatusCode = common.CodeServerBusy

	err := db.DeleteProduct(ctx, req.Pid)
	if err != nil {
		zap.L().Error("Failed to delete product", zap.Error(err))
		return
	}

	// 从es删除该商品, 该代码仅为测试时使用
	err = es.DeleteProduct(ctx, req.Pid)
	if err != nil {
		zap.L().Error("Failed to create es product", zap.Error(err))
	}

	resp.StatusCode = common.CodeSuccess

	return
}

func NewProductCatalogServiceImpl() *ProductCatalogServiceImpl {
	return &ProductCatalogServiceImpl{}
}
