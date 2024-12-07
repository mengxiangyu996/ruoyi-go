package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/dal"
)

type RoleService struct{}

// 新增角色
func (s *RoleService) CreateRole(param dto.SaveRoleRequest, menuIds []int) error {

	tx := dal.Gorm.Begin()

	role := model.SysRole{
		RoleName:          param.RoleName,
		RoleKey:           param.RoleKey,
		RoleSort:          param.RoleSort,
		MenuCheckStrictly: param.MenuCheckStrictly,
		DeptCheckStrictly: param.DeptCheckStrictly,
		Status:            param.Status,
		CreateBy:          param.CreateBy,
		Remark:            param.Remark,
	}

	if err := tx.Model(model.SysRole{}).Create(&role).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(menuIds) > 0 {
		for _, menuId := range menuIds {
			if err := tx.Model(model.SysRoleMenu{}).Create(&model.SysRoleMenu{
				RoleId: role.RoleId,
				MenuId: menuId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// 获取角色列表
func (s *RoleService) GetRoleList(param dto.RoleListRequest, isPaging bool) ([]dto.RoleListResponse, int) {

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

	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	query.Find(&roles)

	return roles, int(count)
}

// 根据用户id查询角色列表
func (s *RoleService) GetRoleListByUserId(userId int) []dto.RoleListResponse {

	roles := make([]dto.RoleListResponse, 0)

	dal.Gorm.Model(model.SysRole{}).Select("sys_role.*").
		Joins("JOIN sys_user_role ON sys_role.role_id = sys_user_role.role_id").
		Where("sys_user_role.user_id = ? AND sys_role.status = ?", userId, constant.NORMAL_STATUS).
		Find(&roles)

	return roles
}

// 根据用户id查询角色key
func (s *RoleService) GetRoleKeysByUserId(userId int) []string {

	roleKeys := make([]string, 0)

	dal.Gorm.Model(model.SysRole{}).
		Joins("JOIN sys_user_role ON sys_user_role.role_id = sys_role.role_id").
		Where("sys_user_role.user_id = ? AND sys_role.status = ?", userId, constant.NORMAL_STATUS).
		Pluck("sys_role.role_key", &roleKeys)

	return roleKeys
}

// 根据用户id查询角色名
func (s *RoleService) GetRoleNamesByUserId(userId int) []string {

	var roleNames []string

	dal.Gorm.Model(model.SysRole{}).
		Joins("JOIN sys_user_role ON sys_user_role.role_id = sys_role.role_id").
		Where("sys_user_role.user_id = ? AND sys_role.status = ?", userId, constant.NORMAL_STATUS).
		Pluck("sys_role.role_name", &roleNames)

	return roleNames
}

// 根据角色key查询角色
func (s *RoleService) GetRoleByRoleName(roleName string) dto.RoleDetailResponse {

	var role dto.RoleDetailResponse

	dal.Gorm.Model(model.SysRole{}).Where("role_name = ?", roleName).Last(&role)

	return role
}

// 根据角色名称查询角色
func (s *RoleService) GetRoleByRoleKey(roleKey string) dto.RoleDetailResponse {

	var role dto.RoleDetailResponse

	dal.Gorm.Model(model.SysRole{}).Where("role_key = ?", roleKey).Last(&role)

	return role
}
