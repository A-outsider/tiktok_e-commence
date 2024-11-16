package check

import (
	"regexp"
)

// 校验图片格式
func PhotoType(suffix string) bool {
	pattern := `.(png|jpeg|jpg)`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(suffix)
}

// TODO 目前限制图片大小 5MB 考虑写入配置文件
const MaxPhotoSize = 5

func PhotoSize(contentLength int64) bool {

	//限制整体大小为 目标大小 + 1 MB
	if contentLength > (MaxPhotoSize+1)*1024*1024 {
		return false
	}
	return true
}
