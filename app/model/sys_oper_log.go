package model

import "ruoyi-go/framework/datetime"

type SysOperLog struct {
	OperId        int `gorm:"primaryKey;autoIncrement"`
	Title         string
	BusinessType  int
	Method        string
	RequestMethod string
	OperName      string
	DeptName      string
	OperUrl       string
	OperIp        string
	OperLocation  string
	OperParam     string
	JsonResult    string
	Status        string `gorm:"default:0"`
	ErrorMsg      string
	OperTime      datetime.Datetime
	CostTime      int
}

func (SysOperLog) TableName() string {
	return "sys_oper_log"
}
