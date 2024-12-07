package model

import (
	"ruoyi-go/framework/datetime"
)

type SysDictData struct {
	DictCode   int `gorm:"primaryKey;autoIncrement"`
	DictSort   int
	DictLabel  string
	DictValue  string
	DictType   string
	CssClass   string
	ListClass  string
	IsDefault  string `gorm:"default:N"`
	Status     string `gorm:"default:0"`
	CreateBy   string
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy   string
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	Remark     string
}

func (SysDictData) TableName() string {
	return "sys_dict_data"
}
