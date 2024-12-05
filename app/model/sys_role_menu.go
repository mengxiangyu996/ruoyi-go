package model

type SysRoleMenu struct {
	RoleId int
	MenuId int
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
