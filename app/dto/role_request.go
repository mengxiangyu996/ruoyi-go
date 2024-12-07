package dto

// 保存角色
type SaveRoleRequest struct {
	RoleId            int    `json:"roleId"`
	RoleName          string `json:"roleName"`
	RoleKey           string `json:"roleKey"`
	RoleSort          int    `json:"roleSort"`
	DataScope         string `json:"dataScope"`
	MenuCheckStrictly int    `json:"menuCheckStrictly"`
	DeptCheckStrictly int    `json:"deptCheckStrictly"`
	Status            string `json:"status"`
	CreateBy          string `json:"createBy"`
	UpdateBy          string `json:"updateBy"`
	Remark            string `json:"remark"`
}

// 角色列表
type RoleListRequest struct {
	PageRequest
	RoleName  string `query:"roleName" form:"roleName"`
	RoleKey   string `query:"roleKey" form:"roleKey"`
	Status    string `query:"status" form:"status"`
	BeginTime string `query:"params[beginTime]" form:"params[beginTime]"`
	EndTime   string `query:"params[endTime]" form:"params[endTime]"`
}

// 新增角色
type CreateRoleRequest struct {
	RoleName          string `json:"roleName"`
	RoleKey           string `json:"roleKey"`
	RoleSort          int    `json:"roleSort"`
	MenuCheckStrictly bool   `json:"menuCheckStrictly"`
	DeptCheckStrictly bool   `json:"deptCheckStrictly"`
	Status            string `json:"status"`
	Remark            string `json:"remark"`
	MenuIds           []int  `json:"menuIds"`
	CreateBy          string `json:"createBy"`
}
