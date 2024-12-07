package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/dal"
)

type UserService struct{}

// 新增用户
func (s *UserService) CreateUser(param dto.SaveUser, roleIds, postIds []int) error {

	tx := dal.Gorm.Begin()

	user := model.SysUser{
		DeptId:      param.DeptId,
		UserName:    param.UserName,
		NickName:    param.NickName,
		UserType:    param.UserType,
		Email:       param.Email,
		Phonenumber: param.Phonenumber,
		Sex:         param.Sex,
		Avatar:      param.Avatar,
		Password:    param.Password,
		LoginIP:     param.LoginIP,
		LoginDate:   param.LoginDate,
		Status:      param.Status,
		CreateBy:    param.CreateBy,
		Remark:      param.Remark,
	}

	if err := tx.Model(model.SysUser{}).Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(roleIds) > 0 {
		for _, roleId := range roleIds {
			if err := tx.Model(model.SysUserRole{}).Create(&model.SysUserRole{
				UserId: user.UserId,
				RoleId: roleId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if len(postIds) > 0 {
		for _, postId := range postIds {
			if err := tx.Model(model.SysUserPost{}).Create(&model.SysUserPost{
				UserId: user.UserId,
				PostId: postId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// 更新用户
func (s *UserService) UpdateUser(param dto.SaveUser, roleIds, postIds []int) error {

	tx := dal.Gorm.Begin()

	if err := tx.Model(model.SysUser{}).Where("user_id = ?", param.UserId).Updates(&model.SysUser{
		DeptId:      param.DeptId,
		NickName:    param.NickName,
		UserType:    param.UserType,
		Email:       param.Email,
		Phonenumber: param.Phonenumber,
		Sex:         param.Sex,
		Avatar:      param.Avatar,
		Password:    param.Password,
		LoginIP:     param.LoginIP,
		LoginDate:   param.LoginDate,
		Status:      param.Status,
		UpdateBy:    param.UpdateBy,
		Remark:      param.Remark,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 清理角色和岗位旧数据
	if roleIds != nil {
		if err := tx.Model(model.SysUserRole{}).Where("user_id = ?", param.UserId).Delete(&model.SysUserRole{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if postIds != nil {
		if err := tx.Model(model.SysUserPost{}).Where("user_id = ?", param.UserId).Delete(&model.SysUserPost{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(roleIds) > 0 {
		for _, roleId := range roleIds {
			if err := tx.Model(model.SysUserRole{}).Create(&model.SysUserRole{
				UserId: param.UserId,
				RoleId: roleId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if len(postIds) > 0 {
		for _, postId := range postIds {
			if err := tx.Model(model.SysUserPost{}).Create(&model.SysUserPost{
				UserId: param.UserId,
				PostId: postId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// 删除用户
func (s *UserService) DeleteUser(userIds []int) error {

	tx := dal.Gorm.Begin()

	if err := tx.Model(model.SysUser{}).Where("user_id IN ?", userIds).Delete(&model.SysUser{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(model.SysUserRole{}).Where("user_id IN ?", userIds).Delete(&model.SysUserRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(model.SysUserPost{}).Where("user_id IN ?", userIds).Delete(&model.SysUserPost{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 用户授权角色
func (s *UserService) AddAuthRole(userId int, roleIds []int) error {

	tx := dal.Gorm.Begin()

	// 清理用户角色
	if err := tx.Model(model.SysUserRole{}).Where("user_id = ?", userId).Delete(&model.SysUserRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 重新插入分配的角色
	if len(roleIds) > 0 {
		for _, roleId := range roleIds {
			if err := tx.Model(model.SysUserRole{}).Create(&model.SysUserRole{
				UserId: userId,
				RoleId: roleId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// 获取用户列表
func (s *UserService) GetUserList(param dto.UserListRequest, userId int, isPaging bool) ([]dto.UserListResponse, int) {

	var count int64
	users := make([]dto.UserListResponse, 0)

	query := dal.Gorm.Model(model.SysUser{}).
		Select("sys_user.*", "sys_dept.dept_name", "sys_dept.leader").
		Joins("LEFT JOIN sys_dept ON sys_user.dept_id = sys_dept.dept_id").
		Scopes(GetDataScope("sys_dept", userId, "sys_user"))

	if param.UserName != "" {
		query = query.Where("sys_user.user_name LIKE ?", "%"+param.UserName+"%")
	}

	if param.Phonenumber != "" {
		query = query.Where("sys_user.phonenumber LIKE ?", "%"+param.Phonenumber+"%")
	}

	if param.Status != "" {
		query = query.Where("sys_user.status = ?", param.Status)
	}

	if param.DeptId != 0 {
		query = query.Where("sys_user.dept_id = ?", param.DeptId)
	}

	if param.BeginTime != "" && param.EndTime != "" {
		query = query.Where("sys_user.create_time BETWEEN ? AND ?", param.BeginTime, param.EndTime)
	}

	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	query.Find(&users)

	return users, int(count)
}

// 根据用户id查询用户信息
func (s *UserService) GetUserByUserId(userId int) dto.UserDetailResponse {

	var user dto.UserDetailResponse

	dal.Gorm.Model(model.SysUser{}).Where("user_id = ?", userId).Last(&user)

	return user
}

// 根据用户名查询用户信息
func (s *UserService) GetUserByUsername(userName string) dto.UserTokenResponse {

	var user dto.UserTokenResponse

	dal.Gorm.Model(model.SysUser{}).
		Select(
			"sys_user.user_id",
			"sys_user.dept_id",
			"sys_user.user_name",
			"sys_user.nick_name",
			"sys_user.user_type",
			"sys_user.password",
			"sys_user.status",
			"sys_dept.dept_name",
		).
		Joins("LEFT JOIN sys_dept ON sys_user.dept_id = sys_dept.dept_id").
		Where("sys_user.user_name = ?", userName).
		Last(&user)

	return user
}

// 根据邮箱查询用户信息
func (s *UserService) GetUserByEmail(email string) dto.UserTokenResponse {

	var user dto.UserTokenResponse

	dal.Gorm.Model(model.SysUser{}).
		Select(
			"sys_user.user_id",
			"sys_user.dept_id",
			"sys_user.user_name",
			"sys_user.nick_name",
			"sys_user.user_type",
			"sys_user.password",
			"sys_user.status",
			"sys_dept.dept_name",
		).
		Joins("LEFT JOIN sys_dept ON sys_user.dept_id = sys_dept.dept_id").
		Where("sys_user.email = ?", email).
		Last(&user)

	return user
}

// 根据手机号码查询用户信息
func (s *UserService) GetUserByPhonenumber(phonenumber string) dto.UserTokenResponse {

	var user dto.UserTokenResponse

	dal.Gorm.Model(model.SysUser{}).
		Select(
			"sys_user.user_id",
			"sys_user.dept_id",
			"sys_user.user_name",
			"sys_user.nick_name",
			"sys_user.user_type",
			"sys_user.password",
			"sys_user.status",
			"sys_dept.dept_name",
		).
		Joins("LEFT JOIN sys_dept ON sys_user.dept_id = sys_dept.dept_id").
		Where("sys_user.phonenumber = ?", phonenumber).
		Last(&user)

	return user
}

// 部门列表转树形
func (s *UserService) DeptListToTree(depts []dto.DeptTreeResponse) []*dto.DeptTreeResponse {

	treeMap := make(map[int]*dto.DeptTreeResponse)

	// 构建查找表
	for _, dept := range depts {
		treeMap[dept.Id] = &dto.DeptTreeResponse{
			Id:       dept.Id,
			Label:    dept.Label,
			ParentId: dept.ParentId,
			Children: make([]*dto.DeptTreeResponse, 0), // 初始化子节点切片
		}
	}

	tree := make([]*dto.DeptTreeResponse, 0)

	for _, dept := range treeMap {
		if _, exists := treeMap[dept.ParentId]; !exists {
			tree = append(tree, dept)
		} else {
			if parent, exists := treeMap[dept.ParentId]; exists {
				parent.Children = append(parent.Children, dept)
			}
		}
	}

	return tree
}

// 查询用户是否拥有权限，拥有返回true
func (s *UserService) UserHasPerm(userId int, perm string) bool {

	var count int64

	dal.Gorm.Model(model.SysUserRole{}).
		Joins("JOIN sys_role ON sys_user_role.role_id = sys_role.role_id").
		Joins("JOIN sys_role_menu ON sys_role_menu.role_id = sys_role.role_id").
		Joins("JOIN sys_menu ON sys_menu.menu_id = sys_role_menu.menu_id").
		Where("sys_role.delete_time IS NULL AND sys_menu.delete_time IS NULL").
		Where("sys_user_role.user_id = ? AND sys_menu.status = ? AND sys_menu.perms = ?", userId, constant.NORMAL_STATUS, perm).
		Count(&count)

	return count > 0
}
