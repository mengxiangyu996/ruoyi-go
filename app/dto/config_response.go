package dto

import "ruoyi-go/framework/datetime"

// 参数列表
type ConfigListResponse struct {
	ConfigId    int               `json:"configId"`
	ConfigName  string            `json:"configName"`
	ConfigKey   string            `json:"configKey"`
	ConfigValue string            `json:"configValue"`
	ConfigType  string            `json:"configType"`
	CreateTime  datetime.Datetime `json:"createTime"`
	Remark      string            `json:"remark"`
}

// 参数详情
type ConfigDetailResponse struct {
	ConfigId    int    `json:"configId"`
	ConfigName  string `json:"configName"`
	ConfigKey   string `json:"configKey"`
	ConfigValue string `json:"configValue"`
	ConfigType  string `json:"configType"`
	Remark      string `json:"remark"`
}

// 参数导出
type ConfigExportResponse struct {
	ConfigId    int    `excel:"name:参数主键;"`
	ConfigName  string `excel:"name:参数名称;"`
	ConfigKey   string `excel:"name:参数键名;"`
	ConfigValue string `excel:"name:参数键值;"`
	ConfigType  string `excel:"name:系统内置;replace:Y_是,N_否;"`
}
