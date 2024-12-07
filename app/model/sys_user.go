package model

import (
	"ruoyi-go/framework/datetime"

	"gorm.io/gorm"
)

type SysUser struct {
	UserId      int `gorm:"primaryKey;autoIncrement"`
	DeptId      int
	UserName    string
	NickName    string
	UserType    string `gorm:"default:00"`
	Email       string
	Phonenumber string
	Sex         string `gorm:"default:0"`
	Avatar      string
	Password    string
	LoginIP     string
	LoginDate   datetime.Datetime
	Status      string `gorm:"default:0"`
	CreateBy    string
	CreateTime  datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy    string
	UpdateTime  datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime  gorm.DeletedAt
	Remark      string
}

func (SysUser) TableName() string {
	return "sys_user"
}
