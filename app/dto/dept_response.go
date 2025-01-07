package dto

import "ruoyi-go/framework/datetime"

// 部门列表
type DeptListResponse struct {
	DeptId     int               `json:"deptId"`
	ParentId   int               `json:"parentId"`
	Ancestors  string            `json:"ancestors"`
	DeptName   string            `json:"deptName"`
	OrderNum   int               `json:"orderNum"`
	Status     string            `json:"status"`
	CreateTime datetime.Datetime `json:"createTime"`
}

// 部门列表树形
type DeptTreeListResponse struct {
	DeptListResponse
	Children []DeptTreeListResponse `json:"children"`
}

// 部门详情
type DeptDetailResponse struct {
	DeptId     int               `json:"deptId"`
	ParentId   int               `json:"parentId"`
	Ancestors  string            `json:"ancestors"`
	DeptName   string            `json:"deptName"`
	OrderNum   int               `json:"orderNum"`
	Leader     string            `json:"leader"`
	Phone      string            `json:"phone"`
	Email      string            `json:"email"`
	Status     string            `json:"status"`
	CreateTime datetime.Datetime `json:"createTime"`
}

// 部门树（用户管理树形）
type DeptTreeResponse struct {
	Id       int                `json:"id"`
	Label    string             `json:"label"`
	Children []DeptTreeResponse `json:"children" gorm:"-"`
	ParentId int                `json:"-"`
}
