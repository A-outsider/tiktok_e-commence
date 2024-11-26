package handler

import (
	"context"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gomall/gateway/types/resp/common"
	product "gomall/kitex_gen/product"
	"gomall/services/product/dal/db"
	"gomall/services/product/dal/es"
	"gomall/services/product/dal/model"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

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

	*product.Categories = req.Product.Categories

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
