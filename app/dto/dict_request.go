package dto

// 字典类型列表
type DictTypeRequest struct {
	PageRequest
	DictName  string `query:"dictName" form:"dictName"`
	DictType  string `query:"dictType" form:"dictType"`
	Status    string `query:"status" form:"status"`
	BeginTime string `query:"params[beginTime]" form:"params[beginTime]"`
	EndTime   string `query:"params[endTime]" form:"params[endTime]"`
}
