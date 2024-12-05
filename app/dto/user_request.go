package dto

// 用户列表
type UserListRequest struct {
	PageRequest
	UserName    string `query:"userName" form:"userName"`
	Phonenumber string `query:"phonenumber" form:"phonenumber"`
	Status      string `query:"status" form:"status"`
	DeptId      int    `query:"deptId" form:"deptId"`
	BeginTime   string `query:"params[beginTime]" form:"params[beginTime]"`
	EndTime     string `query:"params[endTime]" form:"params[endTime]"`
}
