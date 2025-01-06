package dto

import "ruoyi-go/framework/datetime"

// 用户授权
type UserTokenResponse struct {
	UserId   int    `json:"userId"`
	DeptId   int    `json:"deptId"`
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	UserType string `json:"userType"`
	Password string `json:"-"`
	Status   string `json:"status"`
	DeptName string `json:"deptName"`
}

// 用户列表
type UserListResponse struct {
	UserId      int               `json:"userId"`
	DeptId      int               `json:"deptId"`
	UserName    string            `json:"userName"`
	NickName    string            `json:"nickName"`
	Email       string            `json:"email"`
	Phonenumber string            `json:"phonenumber"`
	Sex         string            `json:"sex"`
	LoginIp     string            `json:"loginIp"`
	LoginDate   datetime.Datetime `json:"loginDate"`
	Status      string            `json:"status"`
	CreateTime  datetime.Datetime `json:"createTime"`
	Dept        struct {
		DeptId   int    `json:"deptId"`
		DeptName string `json:"deptName"`
		Leader   string `json:"leader"`
	} `json:"dept" gorm:"-"`
	DeptName string `json:"-"`
	Leader   string `json:"-"`
}

// 用户详情
type UserDetailResponse struct {
	UserId      int               `json:"userId"`
	DeptId      int               `json:"deptId"`
	UserName    string            `json:"userName"`
	NickName    string            `json:"nickName"`
	UserType    string            `json:"userType"`
	Email       string            `json:"email"`
	Phonenumber string            `json:"phonenumber"`
	Sex         string            `json:"sex"`
	Avatar      string            `json:"avatar"`
	Password    string            `json:"-"`
	LoginIP     string            `json:"loginIp"`
	LoginDate   datetime.Datetime `json:"loginDate"`
	Status      string            `json:"status"`
	CreateTime  datetime.Datetime `json:"createTime"`
	Admin       bool              `json:"admin" gorm:"-"`
}

// 授权用户信息
type AuthUserInfoResponse struct {
	UserDetailResponse
	Dept  DeptDetailResponse `json:"dept"`
	Roles []RoleListResponse `json:"roles"`
}

// 用户导出
type UserExportResponse struct {
	UserId      int    `excel:"name:用户序号;"`
	UserName    string `excel:"name:登录名称;"`
	NickName    string `excel:"name:用户名称;"`
	Email       string `excel:"name:用户邮箱;"`
	Phonenumber string `excel:"name:手机号码;"`
	Sex         string `excel:"name:用户性别;replace:0_男,1_女,2_未知;"`
	Status      string `excel:"name:帐号状态;replace:0_正常,1_停用;"`
	LoginIp     string `excel:"name:最后登录IP;"`
	LoginDate   string `excel:"name:最后登录时间;"`
	DeptName    string `excel:"name:部门名称;"`
	DeptLeader  string `excel:"name:部门负责人;"`
}
