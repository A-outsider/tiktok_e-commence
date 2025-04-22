package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"

	rpcSeckill "gomall/kitex_gen/seckill"
	"gomall/services/seckill/constants"
	"gomall/services/seckill/dal/cache"
	"gomall/services/seckill/dal/db"
	"gomall/services/seckill/dal/model"
)

// SeckillServiceImpl implements the last service interface defined in the IDL.
type SeckillServiceImpl struct{}

// NewSeckillServiceImpl creates a new SeckillServiceImpl instance
func NewSeckillServiceImpl() *SeckillServiceImpl {
	return &SeckillServiceImpl{}
}

// CreateSeckillProduct implements the SeckillServiceImpl interface.
func (s *SeckillServiceImpl) CreateSeckillProduct(ctx context.Context, req *rpcSeckill.CreateSeckillProductReq) (resp *rpcSeckill.CreateSeckillProductResp, _ error) {
	resp = new(rpcSeckill.CreateSeckillProductResp)
	resp.StatusCode = constants.ServiceErr
	resp.StatusMsg = "服务内部错误"

	// 参数校验
	if req.Pid == "" || req.SeckillPrice <= 0 || req.Stock <= 0 {
		resp.StatusCode = constants.ParamErr
		resp.StatusMsg = "参数错误"
		return resp, nil
	}

	// 转换为数值ID
	pidInt, err := db.StringToInt64(req.Pid)
	if err != nil {
		klog.Errorf("参数错误: %v", err)
		resp.StatusCode = constants.ParamErr
		resp.StatusMsg = "商品ID格式错误"
		return resp, nil
	}

	// 创建秒杀商品
	product := &model.SeckillProduct{
		ProductID:    pidInt,
		SeckillPrice: req.SeckillPrice,
		Stock:        int32(req.Stock),
		StartTime:    time.Unix(req.StartTime, 0),
		EndTime:      time.Unix(req.EndTime, 0),
		LimitPerUser: req.LimitPerUser,
		IsActive:     req.IsActive,
	}

	// 保存到数据库
	spid, err := db.CreateSeckillProduct(ctx, product)
	if err != nil {
		klog.Errorf("创建秒杀商品失败: %v", err)
		return resp, nil
	}

	// 设置 Redis 缓存
	if err := cache.SetProductStock(ctx, spid, int32(req.Stock)); err != nil {
		klog.Warnf("设置库存缓存失败: %v", err)
		// 非致命错误，继续执行
	}

	// 如果商品已激活，添加到活动列表
	if req.IsActive {
		spidStr := fmt.Sprintf("%d", spid)
		if err := cache.AddToActiveList(ctx, spidStr); err != nil {
			klog.Warnf("添加到活动列表失败: %v", err)
			// 非致命错误，继续执行
		}
	}

	resp.StatusCode = constants.Success
	resp.StatusMsg = "创建成功"
	resp.Spid = fmt.Sprintf("%d", spid)

	return resp, nil
}

// GetSeckillProduct implements the SeckillServiceImpl interface.
func (s *SeckillServiceImpl) GetSeckillProduct(ctx context.Context, req *rpcSeckill.GetSeckillProductReq) (resp *rpcSeckill.GetSeckillProductResp, _ error) {
	resp = new(rpcSeckill.GetSeckillProductResp)
	resp.StatusCode = constants.ServiceErr
	resp.StatusMsg = "服务内部错误"

	// 参数校验
	if req.Spid == "" {
		resp.StatusCode = constants.ParamErr
		resp.StatusMsg = "参数错误"
		return resp, nil
	}

	// 转换为数值ID
	spidInt, err := db.StringToInt64(req.Spid)
	if err != nil {
		klog.Errorf("参数错误: %v", err)
		resp.StatusCode = constants.ParamErr
		resp.StatusMsg = "商品ID格式错误"
		return resp, nil
	}

	// 从数据库获取商品
	product, err := db.GetSeckillProduct(ctx, spidInt)
	if err != nil {
		klog.Errorf("获取秒杀商品失败: %v", err)
		return resp, nil
	}

	if product == nil {
		resp.StatusCode = constants.NotFoundErr
		resp.StatusMsg = "商品不存在"
		return resp, nil
	}

	// 获取实时库存
	stock := product.Stock
	redisStock, err := cache.GetProductStock(ctx, spidInt)
	if err == nil {
		// 如果获取到Redis中的库存，使用Redis中的值
		stock = redisStock
	}

	// 构建响应
	resp.Product = &rpcSeckill.SeckillProductInfo{
		Spid:         fmt.Sprintf("%d", product.ID),
		Pid:          fmt.Sprintf("%d", product.ProductID),
		SeckillPrice: product.SeckillPrice,
		Stock:        int64(stock),
		StartTime:    product.StartTime.Unix(),
		EndTime:      product.EndTime.Unix(),
		LimitPerUser: product.LimitPerUser,
		IsActive:     product.IsActive,
	}

	resp.StatusCode = constants.Success
	resp.StatusMsg = "获取成功"

	return resp, nil
}

// ListActiveSeckillProducts implements the SeckillServiceImpl interface.
func (s *SeckillServiceImpl) ListActiveSeckillProducts(ctx context.Context, req *rpcSeckill.ListActiveSeckillProductsReq) (resp *rpcSeckill.ListActiveSeckillProductsResp, _ error) {
	resp = new(rpcSeckill.ListActiveSeckillProductsResp)
	resp.StatusCode = constants.ServiceErr
	resp.StatusMsg = "服务内部错误"

	// 参数校验
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 获取活动商品列表
	products, total, err := db.GetActiveSeckillProducts(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		klog.Errorf("获取活动商品列表失败: %v", err)
		return resp, nil
	}

	// 构建响应
	resp.Products = make([]*rpcSeckill.SeckillProductInfo, 0, len(products))
	for _, product := range products {
		// 获取实时库存
		stock := product.Stock
		redisStock, err := cache.GetProductStock(ctx, product.ID)
		if err == nil {
			// 如果获取到Redis中的库存，使用Redis中的值
			stock = redisStock
		}

		resp.Products = append(resp.Products, &rpcSeckill.SeckillProductInfo{
			Spid:         fmt.Sprintf("%d", product.ID),
			Pid:          fmt.Sprintf("%d", product.ProductID),
			SeckillPrice: product.SeckillPrice,
			Stock:        int64(stock),
			StartTime:    product.StartTime.Unix(),
			EndTime:      product.EndTime.Unix(),
			LimitPerUser: product.LimitPerUser,
			IsActive:     product.IsActive,
		})
	}

	resp.Total = int32(total)
	resp.StatusCode = constants.Success
	resp.StatusMsg = "获取成功"

	return resp, nil
}

// DoSeckill implements the SeckillServiceImpl interface.
func (s *SeckillServiceImpl) DoSeckill(ctx context.Context, req *rpcSeckill.DoSeckillReq) (resp *rpcSeckill.DoSeckillResp, _ error) {
	resp = new(rpcSeckill.DoSeckillResp)
	resp.StatusCode = constants.ServiceErr
	resp.StatusMsg = "服务内部错误"

	// 参数校验
	if req.Spid == "" || req.Uid == "" || req.Quantity <= 0 {
		resp.StatusCode = constants.ParamErr
		resp.StatusMsg = "参数错误"
		return resp, nil
	}

	// 转换为数值ID
	spidInt, err := db.StringToInt64(req.Spid)
	if err != nil {
		klog.Errorf("商品ID格式错误: %v", err)
		resp.StatusCode = constants.ParamErr
		resp.StatusMsg = "商品ID格式错误"
		return resp, nil
	}

	uidInt, err := db.StringToInt64(req.Uid)
	if err != nil {
		klog.Errorf("用户ID格式错误: %v", err)
		resp.StatusCode = constants.ParamErr
		resp.StatusMsg = "用户ID格式错误"
		return resp, nil
	}

	// 幂等性控制 - 生成幂等Token
	idempToken := uuid.New().String()

	// 检查是否已处理
	idempInfo, err := cache.GetIdempotentInfo(ctx, idempToken)
	if err == nil && idempInfo != nil {
		// 如果已存在幂等信息，直接返回之前的结果
		switch idempInfo.Status {
		case constants.IdempStatusPending:
			resp.StatusCode = constants.Success
			resp.StatusMsg = "正在处理中"
			resp.FlowId = idempInfo.FlowID
			return resp, nil
		case constants.IdempStatusSuccess:
			resp.StatusCode = constants.Success
			resp.StatusMsg = "下单成功"
			resp.FlowId = idempInfo.FlowID
			return resp, nil
		case constants.IdempStatusFailed:
			resp.StatusCode = constants.FailedErr
			resp.StatusMsg = "下单失败"
			return resp, nil
		}
	}

	// 获取商品信息并检查
	product, err := db.GetSeckillProduct(ctx, spidInt)
	if err != nil {
		klog.Errorf("获取秒杀商品失败: %v", err)
		return resp, nil
	}

	if product == nil {
		resp.StatusCode = constants.NotFoundErr
		resp.StatusMsg = "商品不存在"
		return resp, nil
	}

	// 检查活动时间
	now := time.Now()
	if now.Before(product.StartTime) {
		resp.StatusCode = constants.SeckillNotStartErr
		resp.StatusMsg = "秒杀活动未开始"
		return resp, nil
	}

	if now.After(product.EndTime) {
		resp.StatusCode = constants.SeckillEndedErr
		resp.StatusMsg = "秒杀活动已结束"
		return resp, nil
	}

	// 检查每人限购
	if product.LimitPerUser > 0 {
		bought, err := cache.HasUserBought(ctx, spidInt, uidInt)
		if err != nil {
			klog.Warnf("检查用户购买记录失败: %v", err)
			// 非致命错误，继续执行
		} else if bought {
			resp.StatusCode = constants.LimitExceededErr
			resp.StatusMsg = "您已参与过此秒杀活动"
			return resp, nil
		}
	}

	// 预减库存
	if success, err := cache.DecrProductStock(ctx, spidInt, int32(req.Quantity)); err != nil {
		klog.Errorf("扣减库存失败: %v", err)
		return resp, nil
	} else if !success {
		resp.StatusCode = constants.StockNotEnoughErr
		resp.StatusMsg = "库存不足"
		return resp, nil
	}

	// 生成流水ID
	flowID := fmt.Sprintf("flow_%s_%d", time.Now().Format("20060102150405"), time.Now().UnixNano()%1000000)

	// 创建幂等信息
	idempInfo = &cache.IdempotentInfo{
		SPID:      spidInt,
		UID:       uidInt,
		Quantity:  int32(req.Quantity),
		FlowID:    flowID,
		Status:    constants.IdempStatusPending,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(time.Hour), // 1小时过期
	}

	// 保存幂等信息
	if err := cache.SaveIdempotentInfo(ctx, idempToken, idempInfo); err != nil {
		klog.Errorf("保存幂等信息失败: %v", err)
		// 回滚库存
		if err := cache.IncrProductStock(ctx, spidInt, int32(req.Quantity)); err != nil {
			klog.Errorf("回滚库存失败: %v", err)
		}
		return resp, nil
	}

	// 创建库存流水
	flow := &model.InventoryFlow{
		FlowID:     flowID,
		SPID:       spidInt,
		UID:        uidInt,
		Quantity:   int32(req.Quantity),
		Status:     constants.FlowStatusPending,
		IdempToken: idempToken,
		CreatedAt:  time.Now(),
	}

	// 保存流水记录
	if err := db.CreateInventoryFlow(ctx, flow); err != nil {
		klog.Errorf("创建库存流水失败: %v", err)
		// 回滚库存
		if err := cache.IncrProductStock(ctx, spidInt, int32(req.Quantity)); err != nil {
			klog.Errorf("回滚库存失败: %v", err)
		}
		return resp, nil
	}

	// 如果限购，标记用户已购买
	if product.LimitPerUser > 0 {
		if err := cache.MarkUserBought(ctx, spidInt, uidInt); err != nil {
			klog.Warnf("标记用户购买记录失败: %v", err)
			// 非致命错误，继续执行
		}
	}

	// 发送MQ消息（这里使用异步处理，实际项目中应该使用RocketMQ等）
	go s.processInventoryFlow(ctx, flow)

	resp.StatusCode = constants.Success
	resp.StatusMsg = "下单成功"
	resp.FlowId = flowID

	klog.Infof("秒杀请求成功: 用户=%s, 商品=%s, 数量=%d, 流水号=%s", req.Uid, req.Spid, req.Quantity, flowID)

	return resp, nil
}

// 自定义的查询秒杀结果请求
type QuerySeckillResultReq struct {
	IdempToken string `json:"idemp_token"`
}

// 自定义的查询秒杀结果响应
type QuerySeckillResultResp struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	FlowId     string `json:"flow_id"`
	OrderId    string `json:"order_id"`
	Status     int32  `json:"status"`
}

// QuerySeckillResult 查询秒杀结果
func (s *SeckillServiceImpl) QuerySeckillResult(ctx context.Context, req *QuerySeckillResultReq) (resp *QuerySeckillResultResp, _ error) {
	resp = new(QuerySeckillResultResp)
	resp.StatusCode = constants.ServiceErr
	resp.StatusMsg = "服务内部错误"

	// 参数校验
	if req.IdempToken == "" {
		resp.StatusCode = constants.ParamErr
		resp.StatusMsg = "参数错误"
		return resp, nil
	}

	// 获取幂等信息
	idempInfo, err := cache.GetIdempotentInfo(ctx, req.IdempToken)
	if err != nil || idempInfo == nil {
		if err != nil {
			klog.Errorf("获取幂等信息失败: %v", err)
		}
		resp.StatusCode = constants.NotFoundErr
		resp.StatusMsg = "订单不存在"
		return resp, nil
	}

	// 设置状态
	resp.Status = int32(idempInfo.Status)
	resp.FlowId = idempInfo.FlowID
	resp.OrderId = idempInfo.OrderID

	// 设置状态信息
	switch idempInfo.Status {
	case constants.IdempStatusPending:
		resp.StatusMsg = "处理中"
	case constants.IdempStatusSuccess:
		resp.StatusMsg = "下单成功"
	case constants.IdempStatusFailed:
		resp.StatusMsg = "下单失败"
	default:
		resp.StatusMsg = "未知状态"
	}

	resp.StatusCode = constants.Success

	return resp, nil
}

// ConfirmSeckillOrder implements the SeckillServiceImpl interface.
func (s *SeckillServiceImpl) ConfirmSeckillOrder(ctx context.Context, req *rpcSeckill.ConfirmSeckillOrderReq) (resp *rpcSeckill.ConfirmSeckillOrderResp, _ error) {
	resp = new(rpcSeckill.ConfirmSeckillOrderResp)
	resp.StatusCode = constants.ServiceErr
	resp.StatusMsg = "服务内部错误"

	// 参数校验
	if req.FlowId == "" || req.OrderId == "" {
		resp.StatusCode = constants.ParamErr
		resp.StatusMsg = "参数错误"
		return resp, nil
	}

	// 获取流水信息
	flow, err := db.GetInventoryFlow(ctx, req.FlowId)
	if err != nil {
		klog.Errorf("获取流水信息失败: %v", err)
		return resp, nil
	}

	if flow == nil {
		resp.StatusCode = constants.NotFoundErr
		resp.StatusMsg = "流水不存在"
		return resp, nil
	}

	// 确认订单，更新流水状态
	if err := db.UpdateInventoryFlowStatus(ctx, req.FlowId, constants.FlowStatusSuccess, req.OrderId); err != nil {
		klog.Errorf("更新流水状态失败: %v", err)
		return resp, nil
	}

	// 更新幂等信息
	if err := cache.UpdateIdempotentStatus(ctx, flow.IdempToken, constants.IdempStatusSuccess, req.OrderId); err != nil {
		klog.Errorf("更新幂等信息失败: %v", err)
		// 非致命错误，继续执行
	}

	resp.StatusCode = constants.Success
	resp.StatusMsg = "确认成功"

	return resp, nil
}

// CancelSeckill implements the SeckillServiceImpl interface.
func (s *SeckillServiceImpl) CancelSeckill(ctx context.Context, req *rpcSeckill.CancelSeckillReq) (resp *rpcSeckill.CancelSeckillResp, _ error) {
	resp = new(rpcSeckill.CancelSeckillResp)
	resp.StatusCode = constants.ServiceErr
	resp.StatusMsg = "服务内部错误"

	// 参数校验
	if req.FlowId == "" {
		resp.StatusCode = constants.ParamErr
		resp.StatusMsg = "参数错误"
		return resp, nil
	}

	// 获取流水信息
	flow, err := db.GetInventoryFlow(ctx, req.FlowId)
	if err != nil {
		klog.Errorf("获取流水信息失败: %v", err)
		return resp, nil
	}

	if flow == nil {
		resp.StatusCode = constants.NotFoundErr
		resp.StatusMsg = "流水不存在"
		return resp, nil
	}

	// 只有处理中的流水才能取消
	if flow.Status != constants.FlowStatusPending {
		resp.StatusCode = constants.FailedErr
		resp.StatusMsg = "流水状态不允许取消"
		return resp, nil
	}

	// 回滚库存
	if err := cache.IncrProductStock(ctx, flow.SPID, flow.Quantity); err != nil {
		klog.Errorf("回滚库存失败: %v", err)
		// 非致命错误，继续执行
	}

	// 更新流水状态
	if err := db.UpdateInventoryFlowStatus(ctx, req.FlowId, constants.FlowStatusFailed, ""); err != nil {
		klog.Errorf("更新流水状态失败: %v", err)
		return resp, nil
	}

	// 更新幂等信息
	if err := cache.UpdateIdempotentStatus(ctx, flow.IdempToken, constants.IdempStatusFailed, ""); err != nil {
		klog.Errorf("更新幂等信息失败: %v", err)
		// 非致命错误，继续执行
	}

	resp.StatusCode = constants.Success
	resp.StatusMsg = "取消成功"

	return resp, nil
}

// 异步处理库存流水
func (s *SeckillServiceImpl) processInventoryFlow(ctx context.Context, flow *model.InventoryFlow) {
	// 模拟实际业务处理
	time.Sleep(500 * time.Millisecond)

	// 在实际项目中，这里应该调用订单系统创建订单
	// 然后通过ConfirmSeckillOrder接口来确认订单

	// 这里为了演示效果，我们假设处理成功，直接标记为成功状态
	orderID := fmt.Sprintf("order_%s", uuid.New().String()[:8])

	// 更新流水状态
	if err := db.UpdateInventoryFlowStatus(ctx, flow.FlowID, constants.FlowStatusSuccess, orderID); err != nil {
		klog.Errorf("更新流水状态失败: %v", err)
		return
	}

	// 更新幂等信息
	if err := cache.UpdateIdempotentStatus(ctx, flow.IdempToken, constants.IdempStatusSuccess, orderID); err != nil {
		klog.Errorf("更新幂等信息失败: %v", err)
	}

	klog.Infof("秒杀流水处理完成: 流水=%s, 订单=%s", flow.FlowID, orderID)
}

// CheckExpiredFlows 定期检查过期订单
func (s *SeckillServiceImpl) CheckExpiredFlows(ctx context.Context) {
	timeout := time.Now().Add(-15 * time.Minute) // 15分钟未处理的流水视为超时

	flows, err := db.GetPendingInventoryFlows(ctx, timeout)
	if err != nil {
		klog.Errorf("获取待处理流水失败: %v", err)
		return
	}

	for _, flow := range flows {
		// 回滚库存
		if err := cache.IncrProductStock(ctx, flow.SPID, flow.Quantity); err != nil {
			klog.Errorf("回滚库存失败: %v", err)
			continue
		}

		// 更新流水状态
		if err := db.UpdateInventoryFlowStatus(ctx, flow.FlowID, constants.FlowStatusTimeout, ""); err != nil {
			klog.Errorf("更新流水状态失败: %v", err)
			continue
		}

		// 更新幂等信息
		if err := cache.UpdateIdempotentStatus(ctx, flow.IdempToken, constants.IdempStatusFailed, ""); err != nil {
			klog.Errorf("更新幂等信息失败: %v", err)
		}

		klog.Infof("过期流水处理完成: 流水=%s", flow.FlowID)
	}
}

// CheckInventoryConsistency 检查库存一致性
func (s *SeckillServiceImpl) CheckInventoryConsistency(ctx context.Context) {
	products, err := db.GetAllSeckillProducts(ctx)
	if err != nil {
		klog.Errorf("获取所有秒杀商品失败: %v", err)
		return
	}

	var fixedCount int
	for _, product := range products {
		// 获取Redis库存
		redisStock, err := cache.GetProductStock(ctx, product.ID)
		if err != nil {
			klog.Errorf("获取Redis库存失败: %v", err)
			continue
		}

		// 比较库存
		if redisStock != product.Stock {
			klog.Infof("发现库存不一致: 商品=%d, MySQL=%d, Redis=%d", product.ID, product.Stock, redisStock)

			// 更新Redis库存
			if err := cache.SetProductStock(ctx, product.ID, product.Stock); err != nil {
				klog.Errorf("更新Redis库存失败: %v", err)
				continue
			}

			fixedCount++
		}
	}

	if fixedCount > 0 {
		klog.Infof("库存一致性检查完成: 修复 %d 条记录", fixedCount)
	}
}
