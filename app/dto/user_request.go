package dto

import "ruoyi-go/framework/datetime"

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

// 修改用户
type UpdateUser struct {
	UserId      int               `json:"userId"`
	DeptId      int               `json:"deptId"`
	UserName    string            `json:"userName"`
	NickName    string            `json:"nickName"`
	UserType    string            `json:"userType"`
	Email       string            `json:"email"`
	Phonenumber string            `json:"phonenumber"`
	Sex         string            `json:"sex"`
	Avatar      string            `json:"avatar"`
	Password    string            `json:"password"`
	LoginIP     string            `json:"loginIp"`
	LoginDate   datetime.Datetime `json:"loginDate"`
	Status      string            `json:"status"`
}

// 修改个人信息
type UpdateUserProfile struct {
	NickName    string `json:"nickName"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Sex         string `json:"sex"`
}

// 更新个人密码
type UpdateUserProfilePassword struct {
	OldPassword string `query:"oldPassword" form:"oldPassword"`
	NewPassword string `query:"newPassword" form:"newPassword"`
}
