package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/dal"
)

type ConfigService struct{}

// 获取参数列表
func (*ConfigService) GetConfigList(param dto.ConfigListRequest) ([]dto.ConfigListResponse, int) {

	var count int64
	configs := make([]dto.ConfigListResponse, 0)

	query := dal.Gorm.Model(model.SysConfig{}).Order("config_id")

	if param.ConfigName != "" {
		query = query.Where("config_name LIKE ?", "%"+param.ConfigName+"%")
	}

	if param.ConfigKey != "" {
		query = query.Where("config_key LIKE ?", "%"+param.ConfigKey+"%")
	}

	if param.ConfigType != "" {
		query = query.Where("config_type = ?", param.ConfigType)
	}

	if param.BeginTime != "" && param.EndTime != "" {
		query = query.Where("create_time BETWEEN ? AND ?", param.BeginTime, param.EndTime)
	}

	query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize).Find(&configs)

	return configs, int(count)
}

// 根据参数ey获取参数值
func (*ConfigService) GetConfigByConfigKey(configKey string) dto.ConfigDetailResponse {

	var config dto.ConfigDetailResponse

	dal.Gorm.Model(model.SysConfig{}).Where("status = ? AND config_key = ?", constant.NORMAL_STATUS, configKey).First(&config)

	return config
}
