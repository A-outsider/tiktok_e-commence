package middleware

import (
	"github.com/juju/ratelimit"
)

var bucket *ratelimit.Bucket

// 令牌桶限流策略
//func RateLimitInit() {
//	Interval := parse.Duration(config.Get().Auth.RateLimitInterval)
//	caps := config.Get().Auth.RateLimitCap
//	bucket = ratelimit.NewBucket(Interval, int64(caps))
//}
//
//func RateLimitMiddleware() func(c *gin.Context) {
//	return func(c *gin.Context) {
//		res := new(apiCode.Response)
//		if bucket.TakeAvailable(1) < 1 {
//			c.JSON(http.StatusOK, res.NoDataResponse(apiCode.CodeVisitLimitExceeded))
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}
