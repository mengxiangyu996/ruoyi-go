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

// 角色导出
type RoleExportResponse struct {
	RoleId    int    `excel:"name:角色序号;"`
	RoleName  string `excel:"name:角色名称;"`
	RoleKey   string `excel:"name:角色权限;"`
	RoleSort  int    `excel:"name:角色排序;"`
	DataScope string `excel:"name:数据范围;replace:1_全部数据权限,2_自定数据权限,3_本部门数据权限,4_本部门及以下数据权限,5_仅本人数据权限;"`
	Status    string `excel:"name:角色状态;replace:0_正常,1_停用;"`
}
