package validator

import (
	"errors"
	"ruoyi-go/app/dto"
)

// 添加部门验证
func CreateDeptValidator(param dto.CreateDeptRequest) error {

	if param.ParentId <= 0 {
		return errors.New("请选择上级部门")
	}

	if param.DeptName == "" {
		return errors.New("请输入部门名称")
	}

	return nil
}

// 更新部门验证
func UpdateDeptValidator(param dto.UpdateDeptRequest) error {

	if param.DeptId <= 0 {
		return errors.New("参数错误")
	}

	if param.DeptId != 100 && param.ParentId <= 0 {
		return errors.New("请选择上级部门")
	}

	if param.DeptName == "" {
		return errors.New("请输入部门名称")
	}

	if param.DeptId == param.ParentId {
		return errors.New("修改菜单" + param.DeptName + "失败，上级部门不能选择自己")
	}

	return nil
}
