package validator

import (
	"errors"
	"ruoyi-go/app/dto"
)

// 新增角色验证
func CreateRoleValidator(param dto.CreateRoleRequest) error {

	if param.RoleName == "" {
		return errors.New("请输入角色名称")
	}

	if param.RoleKey == "" {
		return errors.New("请输入权限字符")
	}

	return nil
}