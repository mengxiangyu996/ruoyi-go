package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/dal"
)

type RoleService struct{}

// 获取角色列表
func (s *RoleService) GetRoleList(param dto.RoleListRequest) ([]dto.RoleListResponse, int) {

	var count int64
	roles := make([]dto.RoleListResponse, 0)

	query := dal.Gorm.Model(model.SysRole{}).Order("role_sort, role_id")

	if param.RoleName != "" {
		query.Where("role_name LIKE ?", "%"+param.RoleName+"%")
	}

	if param.RoleKey != "" {
		query.Where("role_key LIKE ?", "%"+param.RoleKey+"%")
	}

	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}

	if param.BeginTime != "" && param.EndTime != "" {
		query = query.Where("sys_user.create_time BETWEEN ? AND ?", param.BeginTime, param.EndTime)
	}

	query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize).Find(&roles)

	return roles, int(count)
}

// 根据用户id查询角色key
func (s *RoleService) GetRoleKeysByUserId(userId int) []string {

	roleKeys := make([]string, 0)

	dal.Gorm.Model(model.SysRole{}).
		Joins("JOIN sys_user_role ON sys_role.role_id = sys_user_role.role_id").
		Where("sys_role.delete_time IS NULL AND sys_user_role.user_id = ? AND sys_role.status = ?", userId, constant.NORMAL_STATUS).
		Pluck("sys_role.role_key", &roleKeys)

	return roleKeys
}

// 根据用户id查询角色列表
func (s *RoleService) GetRoleListByUserId(userId int) []dto.RoleListResponse {

	roles := make([]dto.RoleListResponse, 0)

	dal.Gorm.Model(model.SysRole{}).Select("sys_role.*").
		Joins("JOIN sys_user_role ON sys_role.role_id = sys_user_role.role_id").
		Where("sys_role.delete_time IS NULL AND sys_user_role.user_id = ? AND sys_role.status = ?", userId, constant.NORMAL_STATUS).
		Find(&roles)

	return roles
}
