package dto

import (
	"ruoyi-go/framework/datetime"
)

// 角色列表
type RoleListResponse struct {
	RoleId            int               `json:"roleId"`
	RoleName          string            `json:"roleName"`
	RoleKey           string            `json:"roleKey"`
	RoleSort          int               `json:"roleSort"`
	DataScope         string            `json:"dataScope"`
	MenuCheckStrictly bool              `json:"menuCheckStrictly"`
	DeptCheckStrictly bool              `json:"deptCheckStrictly"`
	Status            string            `json:"status"`
	CreateTime        datetime.Datetime `json:"createTime"`
	Flag              bool              `json:"flag" gorm:"-"`
}

// 角色详情
type RoleDetailResponse struct {
	RoleId            int    `json:"roleId"`
	RoleName          string `json:"roleName"`
	RoleKey           string `json:"roleKey"`
	RoleSort          int    `json:"roleSort"`
	DataScope         string `json:"dataScope"`
	MenuCheckStrictly bool   `json:"menuCheckStrictly"`
	DeptCheckStrictly bool   `json:"deptCheckStrictly"`
	Status            string `json:"status"`
	Remark            string `json:"remark"`
}
