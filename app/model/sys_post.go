package model

import (
	"ruoyi-go/framework/datetime"

	"gorm.io/gorm"
)

type SysPost struct {
	PostId     int `gorm:"primaryKey;autoIncrement"`
	PostCode   string
	PostName   string
	PostSort   int
	Status     string `gorm:"default:0"`
	CreateBy   string
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy   string
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
	Remark     string
}

func (SysPost) TableName() string {
	return "sys_post"
}
