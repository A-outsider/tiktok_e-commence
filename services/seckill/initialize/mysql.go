package initialize

import (
	"log"
	"os"
	"time"

	"gomall/services/seckill/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

// InitMySQL 初始化MySQL连接
func InitMySQL() {
	// 从配置中获取DSN
	dsn := config.GetConf().MySQL.DSN

	// 自定义日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// 获取通用数据库对象，用于设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database: " + err.Error())
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(config.GetConf().MySQL.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.GetConf().MySQL.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	_db = db
}

// GetMysql 获取数据库连接
func GetMysql() *gorm.DB {
	return _db
}
