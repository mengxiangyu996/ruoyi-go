package dto

import "ruoyi-go/framework/datetime"

// 登录日志列表
type LogininforListResponse struct {
	InfoId        int               `json:"infoId"`
	UserName      string            `json:"userName"`
	Ipaddr        string            `json:"ipaddr"`
	LoginLocation string            `json:"loginLocation"`
	Browser       string            `json:"browser"`
	Os            string            `json:"os"`
	Status        string            `json:"status"`
	Msg           string            `json:"msg"`
	LoginTime     datetime.Datetime `json:"loginTime"`
}
