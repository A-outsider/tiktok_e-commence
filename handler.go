package main

import (
	"context"
	seckill "gomall/kitex_gen/seckill"
)

// SeckillServiceImpl implements the last service interface defined in the IDL.
type SeckillServiceImpl struct{}

// CreateSeckillProduct implements the SeckillServiceImpl interface.
func (s *SeckillServiceImpl) CreateSeckillProduct(ctx context.Context, req *seckill.CreateSeckillProductReq) (resp *seckill.CreateSeckillProductResp, err error) {
	// TODO: Your code here...
	return
}

// GetSeckillProduct implements the SeckillServiceImpl interface.
func (s *SeckillServiceImpl) GetSeckillProduct(ctx context.Context, req *seckill.GetSeckillProductReq) (resp *seckill.GetSeckillProductResp, err error) {
	// TODO: Your code here...
	return
}

// ListActiveSeckillProducts implements the SeckillServiceImpl interface.
func (s *SeckillServiceImpl) ListActiveSeckillProducts(ctx context.Context, req *seckill.ListActiveSeckillProductsReq) (resp *seckill.ListActiveSeckillProductsResp, err error) {
	// TODO: Your code here...
	return
}

// DoSeckill implements the SeckillServiceImpl interface.
func (s *SeckillServiceImpl) DoSeckill(ctx context.Context, req *seckill.DoSeckillReq) (resp *seckill.DoSeckillResp, err error) {
	// TODO: Your code here...
	return
}

// ConfirmSeckillOrder implements the SeckillServiceImpl interface.
func (s *SeckillServiceImpl) ConfirmSeckillOrder(ctx context.Context, req *seckill.ConfirmSeckillOrderReq) (resp *seckill.ConfirmSeckillOrderResp, err error) {
	// TODO: Your code here...
	return
}

// CancelSeckill implements the SeckillServiceImpl interface.
func (s *SeckillServiceImpl) CancelSeckill(ctx context.Context, req *seckill.CancelSeckillReq) (resp *seckill.CancelSeckillResp, err error) {
	// TODO: Your code here...
	return
}
