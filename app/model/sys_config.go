package model

import (
	"ruoyi-go/framework/datetime"
)

type SysConfig struct {
	ConfigId    int `gorm:"primaryKey;autoIncrement"`
	ConfigName  string
	ConfigKey   string
	ConfigValue string
	ConfigType  string `gorm:"default:N"`
	CreateBy    string
	CreateTime  datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy    string
	UpdateTime  datetime.Datetime `gorm:"autoUpdateTime"`
	Remark      string
}

func (SysConfig) TableName() string {
	return "sys_config"
}
