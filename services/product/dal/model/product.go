package model

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Strings []string

// Scan 从数据库读取数据，将字符串转换为字符串切片，实现 sql.Scanner 接口
func (i *Strings) Scan(value interface{}) error {
	// 将数据库中的值断言为字符串类型（数据库中的类型为 string）
	switch v := value.(type) {
	case string:
		*i = strings.Split(v, "~")
	case []byte:
		*i = strings.Split(string(v), "~")
	default:
		return errors.New("unsupported data type for Strings")
	}

	return nil
}

// Value 存入数据库，将字符串切片转换为带 "~" 分隔符的字符串，实现 driver.Valuer 接口
func (i *Strings) Value() (driver.Value, error) {
	// 初始化一个字符串，逐个连接切片中的每个元素，并用 "~" 分隔
	// 如果切片为空，直接返回空字符串
	if len(*i) == 0 {
		return "", nil
	}

	// 使用 Join 连接字符串切片
	return strings.Join(*i, "~"), nil
}

type Product struct {
	Pid         string   `gorm:"primaryKey; not null; unique" json:"pid"` // ID 编号
	Bid         string   `json:"bid"`                                     // 品牌id，表示该商品是哪个品牌
	Uid         string   `gorm:"not null; index" json:"uid"`              // 用户id，表示哪个用户发表了这个商品
	Name        string   `json:"name"`                                    // 商品名称
	Categories  *Strings `json:"categories"`                              // 商品类别
	Description string   `json:"description"`                             // 商品描述
	Picture     string   `json:"picture"`                                 // 商品实物图片
	Price       float32  `json:"price"`                                   // 商品价格
	Stock       int64    `gorm:"not null"`                                // 秒杀库存
	// Reviews     *Strings       `json:"reviews,omitempty"`             // 用户评价，存储评价id
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (p *Product) TableName() string {
	return "product"
}
