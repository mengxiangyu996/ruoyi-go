package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
)

type UserService struct{}

// 更新个人资料
func (s *UserService) UpdateUser(param dto.UpdateUser) error {
	return dal.Gorm.Model(model.SysUser{}).Where("user_id = ?", param.UserId).Updates(&model.SysUser{
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
	}).Error
}

// 获取用户列表
func (s *UserService) GetUserList(param dto.UserListRequest, userId int) ([]dto.UserListResponse, int) {

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

	query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize).Find(&users)

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
