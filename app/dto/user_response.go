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
	Phonenumber string            `json:"phonenumber"`
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
