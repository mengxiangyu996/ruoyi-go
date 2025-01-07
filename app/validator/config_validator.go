package validator

import (
	"errors"
	"ruoyi-go/app/dto"
)

// 添加参数验证
func CreateConfigValidator(param dto.CreateConfigRequest) error {

	if param.ConfigName == "" {
		return errors.New("请输入参数名称")
	}

	if param.ConfigKey == "" {
		return errors.New("请输入参数键名")
	}

	if param.ConfigValue == "" {
		return errors.New("请输入参数键值")
	}

	return nil
}

// 更新参数验证
func UpdateConfigValidator(param dto.UpdateConfigRequest) error {

	if param.ConfigId <= 0 {
		return errors.New("参数错误")
	}

	if param.ConfigName == "" {
		return errors.New("请输入参数名称")
	}

	if param.ConfigKey == "" {
		return errors.New("请输入参数键名")
	}

	if param.ConfigValue == "" {
		return errors.New("请输入参数键值")
	}

	return nil
}