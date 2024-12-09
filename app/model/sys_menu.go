package model

import (
	"ruoyi-go/framework/datetime"

	"gorm.io/gorm"
)

type SysMenu struct {
	MenuId     int `gorm:"primaryKey;autoIncrement"`
	MenuName   string
	ParentId   int
	OrderNum   int
	Path       string
	Component  string
	Query      string
	RouteName  string
	IsFrame    *int `gorm:"default:1"`
	IsCache    *int
	MenuType   string
	Visible    string `gorm:"default:0"`
	Perms      string
	Icon       string `gorm:"default:#"`
	Status     string `gorm:"default:0"`
	CreateBy   string
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy   string
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	Remark     string
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
