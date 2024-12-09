package model

import (
	"ruoyi-go/framework/datetime"

	"gorm.io/gorm"
)

type SysRole struct {
	RoleId            int `gorm:"primaryKey;autoIncrement"`
	RoleName          string
	RoleKey           string
	RoleSort          int
	DataScope         string `gorm:"default:1"`
	MenuCheckStrictly *int   `gorm:"default:1"`
	DeptCheckStrictly *int   `gorm:"default:1"`
	Status            string `gorm:"default:0"`
	CreateBy          string
	CreateTime        datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy          string
	UpdateTime        datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime        gorm.DeletedAt
	Remark            string
}

func (SysRole) TableName() string {
	return "sys_role"
}
