package cron

import (
	"dream_program/config"
	"os"
	"time"

	"github.com/jasonlvhit/gocron"
	"go.uber.org/zap"
)

func init() {
	c := gocron.NewScheduler()

	// 每天凌晨4点扫描上传文件临时文件夹并清理垃圾文件
	c.Every(1).Days().At("4:00").Do(cleanRedundant)

	// 每天凌晨4点同步播放量数据
	// 从缓存同步到数据库
	// ...
}

// 移除冗余文件
func cleanRedundant() {
	fileConf := config.Get().File
	// 扫描临时图片和视频文件夹
	//removeExpired(fileConf.ImageTempPath, 48)
	removeExpired(fileConf.VideoTempPath, 48)
}

// 移除过期的文件
// params: expired 过期时间(小时)
func removeExpired(dirPath string, expired float64) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		zap.L().Error("定时任务：移除冗余文件,获取全部文件信息错误！")
		return
	}

	for _, file := range files {
		info, _ := file.Info()
		LastAccessedTimes := info.ModTime()
		Now := time.Now()
		// 计算时间差
		TimeDifference := Now.Sub(LastAccessedTimes)

		if TimeDifference.Hours() > expired {
			err = os.Remove(dirPath + "/" + file.Name())
			if err != nil {
				zap.L().Error("移除文件失败！文件：" + dirPath + "/" + file.Name() + err.Error())
			}
		}
	}
	zap.L().Info("清理文件夹冗余完成" + dirPath)
}
