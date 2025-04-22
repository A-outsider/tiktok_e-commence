package db

import (
	"context"
	"errors"
	"strconv"
	"time"

	"gomall/services/seckill/dal/model"
	"gomall/services/seckill/initialize"

	"gorm.io/gorm"
)

// StringToInt64 将字符串转为int64
func StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// CreateSeckillProduct 创建秒杀商品
func CreateSeckillProduct(ctx context.Context, product *model.SeckillProduct) (int64, error) {
	result := initialize.GetMysql().WithContext(ctx).Create(product)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.ID, nil
}

// GetSeckillProductBySPID 根据秒杀商品ID获取商品
func GetSeckillProductBySPID(ctx context.Context, spid string) (*model.SeckillProduct, error) {
	var product model.SeckillProduct
	spidInt, err := StringToInt64(spid)
	if err != nil {
		return nil, err
	}
	err = initialize.GetMysql().WithContext(ctx).Where("id = ?", spidInt).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetSeckillProduct 根据ID获取秒杀商品
func GetSeckillProduct(ctx context.Context, spid int64) (*model.SeckillProduct, error) {
	var product model.SeckillProduct
	result := initialize.GetMysql().WithContext(ctx).Where("id = ?", spid).First(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &product, nil
}

// GetActiveSeckillProducts 获取活动中的秒杀商品并带分页
func GetActiveSeckillProducts(ctx context.Context, page int, pageSize int) ([]*model.SeckillProduct, int64, error) {
	var products []*model.SeckillProduct
	var total int64
	now := time.Now()

	db := initialize.GetMysql().WithContext(ctx).
		Where("start_time <= ? AND end_time >= ?", now, now)

	// 计算总数
	if err := db.Model(&model.SeckillProduct{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询带分页的记录
	offset := (page - 1) * pageSize
	result := db.Offset(offset).Limit(pageSize).Find(&products)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return products, total, nil
}

// ListActiveSeckillProducts 获取当前处于活动时间内的秒杀商品列表
func ListActiveSeckillProducts(ctx context.Context) ([]*model.SeckillProduct, error) {
	var products []*model.SeckillProduct
	now := time.Now()
	err := initialize.GetMysql().WithContext(ctx).Where("start_time <= ? AND end_time >= ?", now, now).Find(&products).Error
	return products, err
}

// GetAllSeckillProducts 获取所有秒杀商品
func GetAllSeckillProducts(ctx context.Context) ([]*model.SeckillProduct, error) {
	var products []*model.SeckillProduct
	result := initialize.GetMysql().WithContext(ctx).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

// UpdateSeckillProductStock 更新秒杀商品库存
func UpdateSeckillProductStock(ctx context.Context, spid int64, quantity int32) error {
	result := initialize.GetMysql().WithContext(ctx).
		Model(&model.SeckillProduct{}).
		Where("id = ? AND stock >= ?", spid, quantity).
		Update("stock", gorm.Expr("stock - ?", quantity))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("库存不足或商品不存在")
	}

	return nil
}

// CreateInventoryFlow 创建库存流水记录
func CreateInventoryFlow(ctx context.Context, flow *model.InventoryFlow) error {
	return initialize.GetMysql().WithContext(ctx).Create(flow).Error
}

// GetInventoryFlow 获取库存流水
func GetInventoryFlow(ctx context.Context, flowID string) (*model.InventoryFlow, error) {
	var flow model.InventoryFlow
	result := initialize.GetMysql().WithContext(ctx).Where("flow_id = ?", flowID).First(&flow)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &flow, nil
}

// UpdateInventoryFlowStatus 更新库存流水状态
func UpdateInventoryFlowStatus(ctx context.Context, flowID string, status int, orderID string) error {
	updates := map[string]interface{}{
		"status": status,
	}
	if orderID != "" {
		updates["order_id"] = orderID
	}

	result := initialize.GetMysql().WithContext(ctx).
		Model(&model.InventoryFlow{}).
		Where("flow_id = ?", flowID).
		Updates(updates)

	return result.Error
}

// GetPendingInventoryFlows 获取待处理的库存流水
func GetPendingInventoryFlows(ctx context.Context, timeout time.Time) ([]*model.InventoryFlow, error) {
	var flows []*model.InventoryFlow

	result := initialize.GetMysql().WithContext(ctx).
		Where("status = ? AND created_at < ? AND retry_count < ?", 0, timeout, 3).
		Find(&flows)

	if result.Error != nil {
		return nil, result.Error
	}

	return flows, nil
}

// RollbackInventory 回滚库存
func RollbackInventory(ctx context.Context, flowID string) error {
	return initialize.GetMysql().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 查询流水
		var flow model.InventoryFlow
		if err := tx.Where("flow_id = ? AND status = ?", flowID, 0).First(&flow).Error; err != nil {
			return err
		}

		// 2. 更新流水状态
		if err := tx.Model(&flow).Updates(map[string]interface{}{
			"status":  2, // 失败
			"op_type": 3, // 回滚
		}).Error; err != nil {
			return err
		}

		// 3. 恢复库存
		if err := tx.Model(&model.SeckillProduct{}).
			Where("id = ?", flow.SPID).
			Update("stock", gorm.Expr("stock + ?", flow.Quantity)).Error; err != nil {
			return err
		}

		return nil
	})
}
