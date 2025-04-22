package constants

// Service name
const (
	SeckillServiceName = "seckill"
)

// Key prefixes
const (
	StockPrefix      = "seckill:stock:"
	UserBoughtPrefix = "seckill:user:"
)

// Flow status
const (
	FlowStatusPending = 0
	FlowStatusSuccess = 1
	FlowStatusFailed  = 2
	FlowStatusTimeout = 3
)

// Idempotent status
const (
	IdempStatusPending = 0
	IdempStatusSuccess = 1
	IdempStatusFailed  = 2
)
