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

// 登陆日志导出
type LogininforExportResponse struct {
	InfoId        int    `excel:"name:序号;"`
	UserName      string `excel:"name:用户账号;"`
	Status        string `excel:"name:登陆状态;replace:0_成功,1_失败;"`
	Ipaddr        string `excel:"name:登录地址;"`
	LoginLocation string `excel:"name:登录地点;"`
	Browser       string `excel:"name:浏览器;"`
	Os            string `excel:"name:操作系统;"`
	Msg           string `excel:"name:提示消息;"`
	LoginTime     string `excel:"name:访问时间;"`
}
