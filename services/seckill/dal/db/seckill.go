package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gomall/services/seckill/dal/model"
	"gomall/services/seckill/initialize"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreateSeckillProduct 创建秒杀商品
func CreateSeckillProduct(ctx context.Context, product *model.SeckillProduct) error {
	return initialize.GetMysql().WithContext(ctx).Create(product).Error
}

// GetSeckillProductBySPID 根据秒杀商品ID获取商品
func GetSeckillProductBySPID(ctx context.Context, spid string) (*model.SeckillProduct, error) {
	var product model.SeckillProduct
	err := initialize.GetMysql().WithContext(ctx).Where("spid = ?", spid).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// ListActiveSeckillProducts 获取当前处于活动时间内的秒杀商品列表
func ListActiveSeckillProducts(ctx context.Context) ([]*model.SeckillProduct, error) {
	var products []*model.SeckillProduct
	now := time.Now()
	err := initialize.GetMysql().WithContext(ctx).Where("is_active = ? AND start_time <= ? AND end_time >= ?", true, now, now).Find(&products).Error
	return products, err
}

// UpdateSeckillProductStock 更新秒杀商品库存（使用乐观锁）
func UpdateSeckillProductStock(ctx context.Context, spid string, quantity int64, version int64) error {
	result := initialize.GetMysql().WithContext(ctx).
		Model(&model.SeckillProduct{}).
		Where("spid = ? AND version = ? AND stock >= ?", spid, version, quantity).
		Updates(map[string]interface{}{
			"stock":   gorm.Expr("stock - ?", quantity),
			"version": gorm.Expr("version + 1"),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("库存不足或商品已被修改")
	}

	return nil
}

// CreateInventoryFlow 创建库存流水记录
func CreateInventoryFlow(ctx context.Context, flow *model.InventoryFlow) error {
	return initialize.GetMysql().WithContext(ctx).Create(flow).Error
}

// UpdateInventoryFlowStatus 更新库存流水状态
func UpdateInventoryFlowStatus(ctx context.Context, flowID string, status int) error {
	return initialize.GetMysql().WithContext(ctx).
		Model(&model.InventoryFlow{}).
		Where("flow_id = ?", flowID).
		Update("status", status).Error
}

// LockInventoryInTx 在事务中锁定库存并创建流水
func LockInventoryInTx(ctx context.Context, spid string, uid string, quantity int64, orderId string) (string, error) {
	// 生成流水ID
	flowID := fmt.Sprintf("flow_%s_%d", spid, time.Now().UnixNano())

	// 开始事务
	tx := initialize.GetMysql().WithContext(ctx).Begin()
	if tx.Error != nil {
		return "", tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 获取商品并加悲观锁
	var product model.SeckillProduct
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("spid = ? AND is_active = ?", spid, true).
		First(&product).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	// 检查库存
	if product.Stock < quantity {
		tx.Rollback()
		return "", errors.New("库存不足")
	}

	// 更新库存
	if err := tx.Model(&product).
		Update("stock", gorm.Expr("stock - ?", quantity)).
		Update("version", gorm.Expr("version + 1")).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	// 创建库存流水
	flow := &model.InventoryFlow{
		FlowID:     flowID,
		SPID:       spid,
		UID:        uid,
		OrderID:    orderId,
		OpType:     1, // 预扣
		Quantity:   quantity,
		Status:     0, // 处理中
		RetryCount: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := tx.Create(flow).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return "", err
	}

	return flowID, nil
}

// ConfirmInventory 确认库存扣减
func ConfirmInventory(ctx context.Context, flowID string, orderId string) error {
	return initialize.GetMysql().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var flow model.InventoryFlow
		if err := tx.Where("flow_id = ? AND status = ?", flowID, 0).First(&flow).Error; err != nil {
			return err
		}

		flow.Status = 1 // 成功
		flow.OrderID = orderId
		flow.OpType = 2 // 确认
		flow.UpdatedAt = time.Now()

		return tx.Save(&flow).Error
	})
}

// RollbackInventory 回滚库存
func RollbackInventory(ctx context.Context, flowID string) error {
	return initialize.GetMysql().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var flow model.InventoryFlow
		if err := tx.Where("flow_id = ? AND status = ?", flowID, 0).First(&flow).Error; err != nil {
			return err
		}

		// 更新流水状态
		flow.Status = 2 // 失败
		flow.OpType = 3 // 回滚
		flow.UpdatedAt = time.Now()

		if err := tx.Save(&flow).Error; err != nil {
			return err
		}

		// 恢复库存
		if err := tx.Model(&model.SeckillProduct{}).
			Where("spid = ?", flow.SPID).
			Update("stock", gorm.Expr("stock + ?", flow.Quantity)).Error; err != nil {
			return err
		}

		return nil
	})
}

// ListPendingFlows 获取待处理的流水记录
func ListPendingFlows(ctx context.Context, limit int) ([]*model.InventoryFlow, error) {
	var flows []*model.InventoryFlow
	err := initialize.GetMysql().WithContext(ctx).
		Where("status = ? AND retry_count < ?", 0, 3).
		Order("created_at ASC").
		Limit(limit).
		Find(&flows).Error
	return flows, err
}
