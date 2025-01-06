package dto

import "ruoyi-go/framework/datetime"

// 操作日志列表
type OperLogListResponse struct {
	OperId        int               `json:"operId"`
	Title         string            `json:"title"`
	BusinessType  int               `json:"businessType"`
	Method        string            `json:"method"`
	RequestMethod string            `json:"requestMethod"`
	OperName      string            `json:"operName"`
	DeptName      string            `json:"deptName"`
	OperUrl       string            `json:"operUrl"`
	OperIp        string            `json:"operIp"`
	OperLocation  string            `json:"operLocation"`
	OperParam     string            `json:"operParam"`
	JsonResult    string            `json:"jsonResult"`
	Status        int               `json:"status"`
	ErrorMsg      string            `json:"errorMsg"`
	OperTime      datetime.Datetime `json:"operTime"`
	CostTime      int               `json:"costTime"`
}

// 操作日志导出
type OperLogExportResponse struct {
	OperId        int    `excel:"name:操作序号;"`
	Title         string `excel:"name:操作模块;"`
	BusinessType  int    `excel:"name:业务类型;replace:0_其它,1_新增,2_修改,3_删除,4_授权,5_导出,6_导入,7_强退,8_生成代码,9_清空数据;"`
	Method        string `excel:"name:请求方法;"`
	RequestMethod string `excel:"name:请求方式;"`
	OperName      string `excel:"name:操作人员;"`
	DeptName      string `excel:"name:部门名称;"`
	OperUrl       string `excel:"name:请求地址;"`
	OperIp        string `excel:"name:操作地址;"`
	OperLocation  string `excel:"name:操作地点;"`
	OperParam     string `excel:"name:请求参数;"`
	JsonResult    string `excel:"name:返回参数;"`
	Status        int    `excel:"name:操作状态;replace:0_正常,1_异常;"`
	ErrorMsg      string `excel:"name:错误消息;"`
	OperTime      string `excel:"name:操作时间;"`
	CostTime      string `excel:"name:消耗时间;"`
}
