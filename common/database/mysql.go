package database

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	UserName string `yaml:"userName"`
	DbName   string `yaml:"dbName"`
	Charset  string `yaml:"charset"`
	Timezone string `yaml:"timeZone"`
}

var db = new(gorm.DB)

func GetDB() *gorm.DB {
	return db
}

func InitMySQL(val any) error {
	var err error

	sql := new(Mysql)

	err = mapstructure.Decode(val, sql) // 为结构体赋值
	if err != nil {
		log.Panicf("error decoding config to Pgsql struct: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		sql.UserName,
		sql.Password,
		sql.Host,
		sql.Port,
		sql.DbName,
		sql.Charset,
	)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}))

	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// TODO : 连接池配置 , 暂时写死
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// TODO : 为Mysql 启动链路追踪
	return nil
}
