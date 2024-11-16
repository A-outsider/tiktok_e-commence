package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         string `gorm:"primarykey;NOT NULL;comment:用户ID" json:"id"`
	Email      string `gorm:"type:varchar(25)" json:"email"`
	Phone      string `gorm:"unique;NOT NULL" json:"phone"`
	Name       string `json:"name"`
	Password   string `json:"-"`           // 用户密码
	AvatarPath string `json:"avatar_path"` // 用户头像的Url
	DefaultId  string `json:"default_id"`  // 默认地址id
	Gender     int64  `json:"gender"`
	Signature  string `json:"signature"`
	Birthday   string `json:"birthday"`

	Role int64 `json:"role"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Address struct {
	Aid     string `gorm:"primarykey;NOT NULL;comment:地址ID" json:"aid"`
	Uid     string `gorm:"NOT NULL" json:"uid"`
	Address string `json:"address"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	//IsDefault bool   `json:"is_default"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// 用户角色
const (
	ConstRoleOfUser = iota
	ConstRoleOfSeller
	ConstRoleOfAdmin
)
