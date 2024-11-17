package db

import (
	"context"
	"gomall/services/product/dal/model"
	initialize "gomall/services/product/initialize"
)

func AddProduct(ctx context.Context, data *model.Product) error {
	return initialize.GetMysql().WithContext(ctx).Create(data).Error
}

func DeleteProduct(ctx context.Context, pid string) error {
	return initialize.GetMysql().WithContext(ctx).Delete(&model.Product{}, pid).Error
}

func GetProductByPid(ctx context.Context, pid string) (*model.Product, error) {
	data := &model.Product{}
	err := initialize.GetMysql().WithContext(ctx).Where("pid = ?", pid).First(data).Error
	return data, err
}
