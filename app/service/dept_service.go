package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/dal"
)

type DeptService struct{}

// 创建部门
func (s *DeptService) CreateDept(param dto.SaveDept) error {

	return dal.Gorm.Model(model.SysDept{}).Create(&model.SysDept{
		ParentId:  param.ParentId,
		Ancestors: param.Ancestors,
		DeptName:  param.DeptName,
		OrderNum:  param.OrderNum,
		Leader:    param.Leader,
		Phone:     param.Phone,
		Email:     param.Email,
		Status:    param.Status,
		CreateBy:  param.CreateBy,
	}).Error
}

// 更新部门
func (s *DeptService) UpdateDept(param dto.SaveDept) error {

	return dal.Gorm.Model(model.SysDept{}).Where("dept_id = ?", param.DeptId).Updates(&model.SysDept{
		ParentId:  param.ParentId,
		Ancestors: param.Ancestors,
		DeptName:  param.DeptName,
		OrderNum:  param.OrderNum,
		Leader:    param.Leader,
		Phone:     param.Phone,
		Email:     param.Email,
		Status:    param.Status,
		UpdateBy:  param.UpdateBy,
	}).Error
}

// 删除部门
func (s *DeptService) DeleteDept(deptId int) error {
	return dal.Gorm.Model(model.SysDept{}).Where("dept_id = ?", deptId).Delete(&model.SysDept{}).Error
}

// 获取部门列表
func (s *DeptService) GetDeptList(param dto.DeptListRequest, userId int) []dto.DeptListResponse {

	depts := make([]dto.DeptListResponse, 0)

	query := dal.Gorm.Model(model.SysDept{}).Order("order_num, dept_id").Scopes(GetDataScope("sys_dept", userId, ""))

	if param.DeptName != "" {
		query.Where("dept_name LIKE ?", "%"+param.DeptName+"%")
	}

	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}

	query.Find(&depts)

	return depts
}

// 根据部门id查询部门信息
func (s *DeptService) GetDeptByDeptId(deptId int) dto.DeptDetailResponse {

	var dept dto.DeptDetailResponse

	dal.Gorm.Model(model.SysDept{}).Where("dept_id = ?", deptId).Last(&dept)

	return dept
}

// 根据部门名称查询部门信息
func (s *DeptService) GetDeptByDeptName(deptName string) dto.DeptDetailResponse {

	var dept dto.DeptDetailResponse

	dal.Gorm.Model(model.SysDept{}).Where("dept_name = ?", deptName).Last(&dept)

	return dept
}

// 获取部门树
func (s *DeptService) GetUserDeptTree(userId int) []dto.DeptTreeResponse {

	depts := make([]dto.DeptTreeResponse, 0)

	dal.Gorm.Model(model.SysDept{}).
		Select(
			"dept_id as id",
			"dept_name as label",
			"parent_id",
		).
		Order("order_num, dept_id").
		Where("status = ?", constant.NORMAL_STATUS).
		Scopes(GetDataScope("sys_dept", userId, "")).
		Find(&depts)

	return depts
}

// 根据角色id获取部门id集合
func (s *DeptService) GetDeptIdsByRoleId(roleId int) []int {

	deptIds := make([]int, 0)

	dal.Gorm.Model(model.SysRoleDept{}).
		Joins("JOIN sys_dept ON sys_dept.dept_id = sys_role_dept.dept_id").
		Where("sys_dept.status = ? AND sys_role_dept.role_id = ?", constant.NORMAL_STATUS, roleId).
		Pluck("sys_dept.dept_id", &deptIds)

	return deptIds
}

// 部门下拉树列表
func (s *DeptService) DeptSelect() []dto.SeleteTree {

	depts := make([]dto.SeleteTree, 0)

	dal.Gorm.Model(model.SysDept{}).Order("order_num, dept_id").
		Select("dept_id as id", "dept_name as label", "parent_id").
		Where("status = ?", constant.NORMAL_STATUS).
		Find(&depts)

	return depts
}

// 部门下拉列表转树形结构
func (s *DeptService) DeptSeleteToTree(depts []dto.SeleteTree, parentId int) []dto.SeleteTree {

	tree := make([]dto.SeleteTree, 0)

	for _, dept := range depts {
		if dept.ParentId == parentId {
			tree = append(tree, dto.SeleteTree{
				Id:       dept.Id,
				Label:    dept.Label,
				ParentId: dept.ParentId,
				Children: s.DeptSeleteToTree(depts, dept.Id),
			})
		}
	}

	return tree
}

// 查询部门是否存在下级
func (s *DeptService) DeptHasChildren(deptId int) bool {

	var count int64

	dal.Gorm.Model(model.SysDept{}).Where("parent_id = ?", deptId).Count(&count)

	return count > 0
}
