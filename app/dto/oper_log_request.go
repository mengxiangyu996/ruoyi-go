package dto

import "ruoyi-go/framework/datetime"

type OperLogListRequest struct {
	PageRequest
	OperIp        string `query:"operIp" form:"operIp"`
	Title         string `query:"title" form:"title"`
	OperName      string `query:"operName" form:"operName"`
	BusinessType  string `query:"businessType" form:"businessType"`
	Status        string `query:"status" form:"status"`
	BeginTime     string `query:"params[beginTime]" form:"params[beginTime]"`
	EndTime       string `query:"params[endTime]" form:"params[endTime]"`
	OrderByColumn string `query:"orderByColumn" form:"orderByColumn"`
	IsAsc         string `query:"isAsc" form:"isAsc"`
	OrderRule     string
}

// 保存操作日志请求
type SaveOperLogRequest struct {
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
	Status        string            `json:"status"`
	ErrorMsg      string            `json:"errorMsg"`
	OperTime      datetime.Datetime `json:"operTime"`
	CostTime      int               `json:"costTime"`
}
