package handler

import (
	"context"
	"time"

	"gomall/gateway/types/resp/common"
	seckill "gomall/kitex_gen/seckill"
	"gomall/services/seckill/dal/model"
)

// KitexSeckillServiceImpl 实现Kitex生成的接口
type KitexSeckillServiceImpl struct {
	seckillService *SeckillServiceImpl
}

// NewKitexSeckillServiceImpl 创建Kitex服务实现实例
func NewKitexSeckillServiceImpl() *KitexSeckillServiceImpl {
	return &KitexSeckillServiceImpl{
		seckillService: NewSeckillServiceImpl(),
	}
}

// CreateSeckillProduct 创建秒杀商品
func (s *KitexSeckillServiceImpl) CreateSeckillProduct(ctx context.Context, req *seckill.CreateSeckillProductReq) (resp *seckill.CreateSeckillProductResp, err error) {
	resp = &seckill.CreateSeckillProductResp{
		StatusCode: common.CodeServerBusy,
	}

	// 转换请求参数
	product := &model.SeckillProduct{
		PID:          req.Pid,
		SeckillPrice: req.SeckillPrice,
		Stock:        req.Stock,
		StartTime:    time.Unix(req.StartTime, 0),
		EndTime:      time.Unix(req.EndTime, 0),
		LimitPerUser: int(req.LimitPerUser),
		IsActive:     req.IsActive,
	}

	// 调用内部服务创建秒杀商品
	err = s.seckillService.CreateSeckillProduct(ctx, product)
	if err != nil {
		resp.StatusMsg = err.Error()
		return
	}

	resp.StatusCode = common.CodeSuccess
	resp.Spid = product.SPID
	return
}

// GetSeckillProduct 获取秒杀商品详情
func (s *KitexSeckillServiceImpl) GetSeckillProduct(ctx context.Context, req *seckill.GetSeckillProductReq) (resp *seckill.GetSeckillProductResp, err error) {
	resp = &seckill.GetSeckillProductResp{
		StatusCode: common.CodeServerBusy,
	}

	// 调用内部服务获取秒杀商品
	product, err := s.seckillService.GetSeckillProduct(ctx, req.Spid)
	if err != nil {
		resp.StatusMsg = err.Error()
		return
	}

	// 转换为响应格式
	resp.Product = &seckill.SeckillProductInfo{
		Spid:         product.SPID,
		Pid:          product.PID,
		SeckillPrice: product.SeckillPrice,
		Stock:        product.Stock,
		StartTime:    product.StartTime.Unix(),
		EndTime:      product.EndTime.Unix(),
		LimitPerUser: int32(product.LimitPerUser),
		IsActive:     product.IsActive,
		// 这里需要从产品服务获取商品名称和图片
		ProductName:  "",
		ProductImage: "",
	}

	resp.StatusCode = common.CodeSuccess
	return
}

// ListActiveSeckillProducts 获取活动中的秒杀商品列表
func (s *KitexSeckillServiceImpl) ListActiveSeckillProducts(ctx context.Context, req *seckill.ListActiveSeckillProductsReq) (resp *seckill.ListActiveSeckillProductsResp, err error) {
	resp = &seckill.ListActiveSeckillProductsResp{
		StatusCode: common.CodeServerBusy,
	}

	// 调用内部服务获取活动中的秒杀商品列表
	products, err := s.seckillService.ListActiveSeckillProducts(ctx)
	if err != nil {
		resp.StatusMsg = err.Error()
		return
	}

	// 转换为响应格式
	resp.Products = make([]*seckill.SeckillProductInfo, 0, len(products))
	for _, product := range products {
		resp.Products = append(resp.Products, &seckill.SeckillProductInfo{
			Spid:         product.SPID,
			Pid:          product.PID,
			SeckillPrice: product.SeckillPrice,
			Stock:        product.Stock,
			StartTime:    product.StartTime.Unix(),
			EndTime:      product.EndTime.Unix(),
			LimitPerUser: int32(product.LimitPerUser),
			IsActive:     product.IsActive,
			// 这里需要从产品服务获取商品名称和图片
			ProductName:  "",
			ProductImage: "",
		})
	}

	resp.Total = int32(len(products))
	resp.StatusCode = common.CodeSuccess
	return
}

// DoSeckill 执行秒杀
func (s *KitexSeckillServiceImpl) DoSeckill(ctx context.Context, req *seckill.DoSeckillReq) (resp *seckill.DoSeckillResp, err error) {
	resp = &seckill.DoSeckillResp{
		StatusCode: common.CodeServerBusy,
	}

	// 调用内部服务执行秒杀
	flowID, err := s.seckillService.DoSeckill(ctx, req.Spid, req.Uid, req.Quantity)
	if err != nil {
		resp.StatusMsg = err.Error()
		return
	}

	resp.FlowId = flowID
	resp.StatusCode = common.CodeSuccess
	return
}

// ConfirmSeckillOrder 确认秒杀订单
func (s *KitexSeckillServiceImpl) ConfirmSeckillOrder(ctx context.Context, req *seckill.ConfirmSeckillOrderReq) (resp *seckill.ConfirmSeckillOrderResp, err error) {
	resp = &seckill.ConfirmSeckillOrderResp{
		StatusCode: common.CodeServerBusy,
	}

	// 调用内部服务确认秒杀订单
	err = s.seckillService.ConfirmSeckillOrder(ctx, req.FlowId, req.OrderId)
	if err != nil {
		resp.StatusMsg = err.Error()
		return
	}

	resp.StatusCode = common.CodeSuccess
	return
}

// CancelSeckill 取消秒杀
func (s *KitexSeckillServiceImpl) CancelSeckill(ctx context.Context, req *seckill.CancelSeckillReq) (resp *seckill.CancelSeckillResp, err error) {
	resp = &seckill.CancelSeckillResp{
		StatusCode: common.CodeServerBusy,
	}

	// 调用内部服务取消秒杀
	err = s.seckillService.CancelSeckill(ctx, req.FlowId)
	if err != nil {
		resp.StatusMsg = err.Error()
		return
	}

	resp.StatusCode = common.CodeSuccess
	return
}
