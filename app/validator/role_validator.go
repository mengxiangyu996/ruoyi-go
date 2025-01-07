package validator

import (
	"errors"
	"ruoyi-go/app/dto"
	"ruoyi-go/common/utils"
)

// 添加角色验证
func CreateRoleValidator(param dto.CreateRoleRequest) error {

	if param.RoleName == "" {
		return errors.New("请输入角色名称")
	}

	if param.RoleKey == "" {
		return errors.New("请输入权限字符")
	}

	return nil
}

// 更新角色验证
func UpdateRoleValidator(param dto.UpdateRoleRequest) error {

	if param.RoleId <= 0 {
		return errors.New("参数错误")
	}

	if param.RoleName == "" {
		return errors.New("请输入角色名称")
	}

	if param.RoleKey == "" {
		return errors.New("请输入权限字符")
	}

	return nil
}

// 删除角色验证
func RemoveRoleValidator(roleIds []int, roleId int, roleName string) error {

	if utils.Contains(roleIds, 1) {
		return errors.New("超级管理员无法删除")
	}

	if utils.Contains(roleIds, roleId) {
		return errors.New(roleName + "角色无法删除")
	}

	return nil
}

// 修改用户状态验证
func ChangeRoleStatusValidator(param dto.UpdateRoleRequest) error {

	if param.RoleId <= 0 {
		return errors.New("参数错误")
	}

	if param.Status == "" {
		return errors.New("请选择状态")
	}

	return nil
}
