package dto

// 部门列表
type DeptListRequest struct {
	DeptName string `query:"deptName" form:"deptName"`
	Status   string `query:"status" form:"status"`
}
