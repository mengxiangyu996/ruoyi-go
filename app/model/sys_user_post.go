package model

type SysUserPost struct {
	UserId int
	PostId int
}

func (SysUserPost) TableName() string {
	return "sys_user_post"
}
