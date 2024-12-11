package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
)

type ConfigService struct{}

// 创建参数
func (s *ConfigService) CreateConfig(param dto.SaveConfig) error {

	return dal.Gorm.Model(model.SysConfig{}).Create(&model.SysConfig{
		ConfigName:  param.ConfigName,
		ConfigKey:   param.ConfigKey,
		ConfigValue: param.ConfigValue,
		ConfigType:  param.ConfigType,
		CreateBy:    param.CreateBy,
		Remark:      param.Remark,
	}).Error
}

// 更新参数
func (s *ConfigService) UpdateConfig(param dto.SaveConfig) error {

	return dal.Gorm.Model(model.SysConfig{}).Where("config_id = ?", param.ConfigId).Updates(&model.SysConfig{
		ConfigName:  param.ConfigName,
		ConfigKey:   param.ConfigKey,
		ConfigValue: param.ConfigValue,
		ConfigType:  param.ConfigType,
		UpdateBy:    param.UpdateBy,
		Remark:      param.Remark,
	}).Error
}

// 删除参数
func (s *ConfigService) DeleteConfig(configIds []int) error {
	return dal.Gorm.Model(model.SysConfig{}).Where("config_id IN ?", configIds).Delete(&model.SysConfig{}).Error
}

// 获取参数列表
func (s *ConfigService) GetConfigList(param dto.ConfigListRequest) ([]dto.ConfigListResponse, int) {

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

// 获取参数详情
func (s *ConfigService) GetConfigByConfigId(configId int) dto.ConfigDetailResponse {

	var config dto.ConfigDetailResponse

	dal.Gorm.Model(model.SysConfig{}).Where("config_id = ?", configId).Last(&config)

	return config
}

// 根据参数key获取参数值
func (s *ConfigService) GetConfigByConfigKey(configKey string) dto.ConfigDetailResponse {

	var config dto.ConfigDetailResponse

	dal.Gorm.Model(model.SysConfig{}).Where("config_key = ?", configKey).Last(&config)

	return config
}
