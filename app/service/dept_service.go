package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/dal"
)

type DeptService struct{}

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

	dal.Gorm.Model(model.SysDept{}).Where("status = ? AND dept_id = ?", constant.NORMAL_STATUS, deptId).Last(&dept)

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
