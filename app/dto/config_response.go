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
}
