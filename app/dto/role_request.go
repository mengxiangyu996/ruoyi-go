package dto

// 角色列表
type RoleListRequest struct {
	PageRequest
	RoleName  string `query:"roleName" form:"roleName"`
	RoleKey   string `query:"roleKey" form:"roleKey"`
	Status    string `query:"status" form:"status"`
	BeginTime string `query:"params[beginTime]" form:"params[beginTime]"`
	EndTime   string `query:"params[endTime]" form:"params[endTime]"`
}
