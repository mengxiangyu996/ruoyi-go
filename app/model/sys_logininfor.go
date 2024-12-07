package model

import "ruoyi-go/framework/datetime"

type SysLogininfor struct {
	InfoId        int `gorm:"primaryKey;autoIncrement"`
	UserName      string
	Ipaddr        string
	LoginLocation string
	Browser       string
	Os            string
	Status        string `gorm:"default:0"`
	Msg           string
	LoginTime     datetime.Datetime
}

func (SysLogininfor) TableName() string {
	return "sys_logininfor"
}
