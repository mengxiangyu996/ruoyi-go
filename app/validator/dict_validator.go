package validator

import (
	"errors"
	"ruoyi-go/app/dto"
)

// 创建字典类型验证
func CreateDictTypeValidator(param dto.CreateDictTypeRequest) error {

	if param.DictName == "" {
		return errors.New("请输入字典名称")
	}

	if param.DictType == "" {
		return errors.New("请输入字典类型")
	}

	return nil
}

// 更新字典类型验证
func UpdateDictTypeValidator(param dto.UpdateDictTypeRequest) error {

	if param.DictId <= 0 {
		return errors.New("参数错误")
	}

	if param.DictName == "" {
		return errors.New("请输入字典名称")
	}

	if param.DictType == "" {
		return errors.New("请输入字典类型")
	}

	return nil
}

// 创建字典数据验证
func CreateDictDataValidator(param dto.CreateDictDataRequest) error {

	if param.DictLabel == "" {
		return errors.New("请输入数据标签")
	}

	if param.DictValue == "" {
		return errors.New("请输入数据键值")
	}

	return nil
}

// 更新字典数据验证
func UpdateDictDataValidator(param dto.UpdateDictDataRequest) error {

	if param.DictCode <= 0 {
		return errors.New("参数错误")
	}

	if param.DictLabel == "" {
		return errors.New("请输入数据标签")
	}

	if param.DictValue == "" {
		return errors.New("请输入数据键值")
	}

	return nil
}
