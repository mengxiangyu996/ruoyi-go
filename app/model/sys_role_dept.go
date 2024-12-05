package model

type SysRoleDept struct {
	RoleId int
	DeptId int
}

func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}
