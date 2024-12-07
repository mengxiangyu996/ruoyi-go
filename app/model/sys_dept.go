package model

import (
	"ruoyi-go/framework/datetime"

	"gorm.io/gorm"
)

type SysDept struct {
	DeptId     int `gorm:"primaryKey;autoIncrement"`
	ParentId   int
	Ancestors  string
	DeptName   string
	OrderNum   int
	Leader     string
	Phone      string
	Email      string
	Status     string `gorm:"default:0"`
	CreateBy   string
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy   string
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
}

func (SysDept) TableName() string {
	return "sys_dept"
}
