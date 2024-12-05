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
