package dto

// 保存角色
type SaveRole struct {
	RoleId            int    `json:"roleId"`
	RoleName          string `json:"roleName"`
	RoleKey           string `json:"roleKey"`
	RoleSort          int    `json:"roleSort"`
	DataScope         string `json:"dataScope"`
	MenuCheckStrictly *int   `json:"menuCheckStrictly"`
	DeptCheckStrictly *int   `json:"deptCheckStrictly"`
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
}

// 更新角色
type UpdateRoleRequest struct {
	RoleId            int    `json:"roleId"`
	RoleName          string `json:"roleName"`
	RoleKey           string `json:"roleKey"`
	RoleSort          int    `json:"roleSort"`
	DataScope         string `json:"dataScope"`
	MenuCheckStrictly bool   `json:"menuCheckStrictly"`
	DeptCheckStrictly bool   `json:"deptCheckStrictly"`
	Status            string `json:"status"`
	Remark            string `json:"remark"`
	MenuIds           []int  `json:"menuIds"`
	DeptIds           []int  `json:"deptIds"`
}

// 查询已分配用户角色列表
type RoleAuthUserAllocatedListRequest struct {
	PageRequest
	RoleId      int    `query:"roleId" form:"roleId"`
	UserName    string `query:"userName" form:"userName"`
	Phonenumber string `query:"phonenumber" form:"phonenumber"`
}

// 批量选择用户授权
type RoleAuthUserSelectAllRequest struct {
	RoleId  int    `query:"roleId" form:"roleId"`
	UserIds string `query:"userIds" form:"userIds"`
}

// 取消用户授权
type RoleAuthUserCancelRequest struct {
	RoleId int `json:"roleId,string"`
	UserId int `json:"userId"`
}

// 批量取消用户授权
type RoleAuthUserCancelAllRequest struct {
	RoleId  int    `query:"roleId" form:"roleId"`
	UserIds string `query:"userIds" form:"userIds"`
}
