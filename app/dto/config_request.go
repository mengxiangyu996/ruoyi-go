package dto

// 参数列表
type ConfigListRequest struct {
	PageRequest
	ConfigName string `query:"configName" form:"configName"`
	ConfigKey  string `query:"configKey" form:"configKey"`
	ConfigType string `query:"configType" form:"configType"`
	BeginTime  string `query:"params[beginTime]" form:"params[beginTime]"`
	EndTime    string `query:"params[endTime]" form:"params[endTime]"`
}
