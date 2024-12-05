package dto

import "ruoyi-go/framework/datetime"

// 登录日志列表
type LogininforListRequest struct {
	PageRequest
	Ipaddr        string `query:"ipaddr" form:"ipaddr"`
	UserName      string `query:"userName" form:"userName"`
	Status        string `query:"status" form:"status"`
	BeginTime     string `query:"params[beginTime]" form:"params[beginTime]"`
	EndTime       string `query:"params[endTime]" form:"params[endTime]"`
	OrderByColumn string `query:"orderByColumn" form:"orderByColumn"`
	IsAsc         string `query:"isAsc" form:"isAsc"`
	OrderRule     string
}

// 保存登录日志信息
type SaveLogininforRequest struct {
	UserName      string            `json:"userName"`
	Ipaddr        string            `json:"ipaddr"`
	LoginLocation string            `json:"loginLocation"`
	Browser       string            `json:"browser"`
	Os            string            `json:"os"`
	Status        string            `json:"status"`
	Msg           string            `json:"msg"`
	LoginTime     datetime.Datetime `json:"loginTime"`
}
