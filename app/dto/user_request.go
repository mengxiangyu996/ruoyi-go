package dto

import "ruoyi-go/framework/datetime"

// 保存用户
type SaveUser struct {
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
	CreateBy    string            `json:"createBy"`
	UpdateBy    string            `json:"updateBy"`
	Remark      string            `json:"remark"`
}

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

// 新增用户
type CreateUserRequest struct {
	DeptId      int    `json:"deptId"`
	UserName    string `json:"userName"`
	NickName    string `json:"nickName"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Sex         string `json:"sex"`
	Password    string `json:"password"`
	Status      string `json:"status"`
	Remark      string `json:"remark"`
	PostIds     []int  `json:"postIds"`
	RoleIds     []int  `json:"roleIds"`
}

// 更新用户
type UpdateUserRequest struct {
	UserId      int    `json:"userId"`
	DeptId      int    `json:"deptId"`
	UserName    string `json:"userName"`
	NickName    string `json:"nickName"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Sex         string `json:"sex"`
	Password    string `json:"password"`
	Status      string `json:"status"`
	Remark      string `json:"remark"`
	PostIds     []int  `json:"postIds"`
	RoleIds     []int  `json:"roleIds"`
}

// 用户授权角色
type AddUserAuthRoleRequest struct {
	UserId  int    `query:"userId" form:"userId"`
	RoleIds string `query:"roleIds" form:"roleIds"`
}

// 修改个人信息
type UpdateProfileRequest struct {
	NickName    string `json:"nickName"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Sex         string `json:"sex"`
}

// 更新个人密码
type UserProfileUpdatePwdRequest struct {
	OldPassword string `query:"oldPassword" form:"oldPassword"`
	NewPassword string `query:"newPassword" form:"newPassword"`
}

// 用户导入
type UserImportRequest struct {
	DeptId      int    `excel:"name:部门编号;"`
	UserName    string `excel:"name:登录名称;"`
	NickName    string `excel:"name:用户名称;"`
	Email       string `excel:"name:用户邮箱;"`
	Phonenumber string `excel:"name:手机号码;"`
	Sex         string `excel:"name:用户性别;replace:0_男,1_女,2_未知;"`
	Status      string `excel:"name:帐号状态;replace:0_正常,1_停用;"`
}
