package constants

// API 响应码
const (
	// Success 成功
	Success int64 = 0
	// ParamErr 参数错误
	ParamErr int64 = 10001
	// NotFoundErr 资源不存在
	NotFoundErr int64 = 10002
	// ServiceErr 服务内部错误
	ServiceErr int64 = 10003
	// FailedErr 操作失败
	FailedErr int64 = 10004
	// LimitExceededErr 超出限制
	LimitExceededErr int64 = 10005
	// StockNotEnoughErr 库存不足
	StockNotEnoughErr int64 = 10006
	// SeckillNotStartErr 秒杀未开始
	SeckillNotStartErr int64 = 10007
	// SeckillEndedErr 秒杀已结束
	SeckillEndedErr int64 = 10008
)
