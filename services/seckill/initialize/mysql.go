package initialize

import (
	"fmt"
	"time"

	"gomall/services/seckill/config"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var _db *gorm.DB

// InitMySQL 初始化MySQL连接
func InitMySQL() {
	// 从配置中获取
	conf := config.GetConf().MySQL

	// 构建DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DB,
		conf.Charset)

	// 连接数据库
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		klog.Fatalf("failed to connect database: %v", err)
	}

	// 获取通用数据库对象，用于设置连接池
	sqlDB, err := _db.DB()
	if err != nil {
		klog.Fatalf("failed to get database: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	klog.Infof("MySQL connected: %s:%d/%s", conf.Host, conf.Port, conf.DB)
}

// GetMysql 获取数据库连接
func GetMysql() *gorm.DB {
	return _db
}
