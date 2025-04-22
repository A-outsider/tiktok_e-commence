package database

import (
	"fmt"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

// DBConfig 数据库配置
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// InitDB 初始化数据库连接
func InitDB(config interface{}) {
	dbOnce.Do(func() {
		var dsn string

		switch cfg := config.(type) {
		case DBConfig:
			dsn = fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
			)
		case map[string]interface{}:
			// 尝试从map配置中提取DSN
			if dsnVal, ok := cfg["dsn"]; ok {
				if dsnStr, ok := dsnVal.(string); ok {
					dsn = dsnStr
				}
			} else {
				// 尝试从各个字段构建DSN
				host, _ := cfg["host"].(string)
				port, _ := cfg["port"].(float64) // JSON解析数字为float64
				user, _ := cfg["user"].(string)
				password, _ := cfg["password"].(string)
				dbName, _ := cfg["db_name"].(string)

				dsn = fmt.Sprintf(
					"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
					user, password, host, int(port), dbName,
				)
			}
		case string:
			// 直接使用DSN字符串
			dsn = cfg
		default:
			klog.Fatal("不支持的数据库配置类型")
		}

		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名
			},
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			klog.Fatalf("数据库连接失败: %v", err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			klog.Fatalf("获取数据库连接失败: %v", err)
		}

		// 设置连接池
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)

		klog.Info("数据库连接成功")
	})
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	if db == nil {
		klog.Fatal("数据库未初始化")
	}
	return db
}
