package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       string `gorm:"primarykey;NOT NULL;comment:用户ID" json:"id"`
	Email    string `gorm:"type:varchar(25)" json:"email"`
	Phone    string `gorm:"unique;NOT NULL" json:"phone"`
	Username string `json:"username"`
	Password string `json:"-"`           // 用户密码
	Avatar   string `json:"avatar_path"` // 用户头像的Url
	Aid      int64  `json:"aid"`         // 地址id

	Role int64 `json:"role"`

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
