package model

type SysUserRole struct {
	UserId int
	RoleId int
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
