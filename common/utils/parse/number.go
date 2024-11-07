package parse

import (
	"github.com/dustin/go-humanize"
	"go.uber.org/zap"
	"log"
	"strconv"
	"strings"
	"time"
)

// 处理时间标准 : 1s 1m 1h 1d
// TODO 响应带error
func Duration(d string) time.Duration {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr
	}

	// 解析day
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)

		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr
		}
		return dr + ndr
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	zap.L().Error("Error parsing duration:", zap.Error(err))

	return time.Duration(dv)
}

// TODO : 这两个需要重构掉
func TimeStrToMilli(timeStr string) int64 {
	if len(strings.TrimSpace(timeStr)) == 0 {
		return 3000
	}

	// 处理没有单位的情况
	tem := timeStr
	if tem, err := strconv.ParseInt(tem, 10, 64); err == nil {
		return tem
	}
	duration, err := time.ParseDuration(timeStr)
	if err != nil {
		log.Println("parseTimeStr Error:", err)
		return 3000
	}
	return duration.Milliseconds()
}

func MemoryToBytes(memoryStr string) int64 {
	if len(strings.TrimSpace(memoryStr)) == 0 {
		return 536870912
	} // Default to 512 MB

	// 处理没有单位的情况
	tem := memoryStr
	if tem, err := strconv.ParseInt(tem, 10, 64); err == nil {
		return tem
	}

	bytes, err := humanize.ParseBytes(memoryStr)
	if err != nil {
		zap.L().Error("Error parsing memory string:", zap.Error(err))
		return 536870912
	}
	return int64(bytes)
}
