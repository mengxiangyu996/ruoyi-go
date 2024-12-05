package service

import (
	"ruoyi-go/common/types/constant"
	"strings"

	"gorm.io/gorm"
)

// 获取数据范围
//
// 在需要数据权限语句中调用此方法，
// 如果需要实现数据权限，需要有dept_id和user_id这两个字段
//
// 必填项：deptAlias、userId，其中deptAlias为dept表别名，userId为当前授权登录的用户id
//
// 例如：dal.Grom.Model(model.User{}).Scopes(GetDataScope(deptAlias, userId, userAlias)...).Find(&[]model.User{})
//
// 数据范围：1-全部数据权限；2-自定数据权限；3-本部门数据权限；4-本部门及以下数据权限；5-仅本人数据权限
func GetDataScope(deptAlias string, userId int, userAlias string) func(*gorm.DB) *gorm.DB {

	// 超级管理员不进行数据权限过滤
	if userId == 1 {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}

	if deptAlias == "" {
		deptAlias = "sys_dept"
	}

	// 获取用户信息
	user := (&UserService{}).GetUserByUserId(userId)

	var roleIds []int

	// 获取当前用户的角色
	roles := (&RoleService{}).GetRoleListByUserId(user.UserId)
	for _, role := range roles {
		if role.DataScope == "2" && role.Status == constant.NORMAL_STATUS {
			roleIds = append(roleIds, role.RoleId)
		}
	}

	return func(db *gorm.DB) *gorm.DB {

		var sqlCondition []string
		var sqlArg []interface{}

		for _, role := range roles {

			// 全部数据权限
			if role.DataScope == "1" {
				return db
			}

			// 自定数据权限
			if role.DataScope == "2" {
				if len(roleIds) > 0 {
					sqlCondition = append(sqlCondition, deptAlias+".dept_id IN (SELECT dept_id FROM sys_role_dept WHERE role_id IN (?))")
					sqlArg = append(sqlArg, roleIds)
				} else {
					sqlCondition = append(sqlCondition, deptAlias+".dept_id IN (SELECT dept_id FROM sys_role_dept WHERE role_id = ?)")
					sqlArg = append(sqlArg, role.RoleId)
				}
			}

			// 本部门数据权限
			if role.DataScope == "3" {
				sqlCondition = append(sqlCondition, deptAlias+".dept_id = ?")
				sqlArg = append(sqlArg, user.DeptId)
			}

			// 本部门及以下数据权限
			if role.DataScope == "4" {
				sqlCondition = append(sqlCondition, deptAlias+".dept_id IN ( SELECT dept_id FROM sys_dept WHERE dept_id = ? OR find_in_set(?, ancestors) )")
				sqlArg = append(sqlArg, user.DeptId, user.DeptId)
			}

			// 仅本人数据权限
			if role.DataScope == "5" {
				if userAlias != "" {
					sqlCondition = append(sqlCondition, userAlias+".user_id = ?")
					sqlArg = append(sqlArg, user.UserId)
				} else {
					// 数据权限为仅本人且没有userAlias别名不查询任何数据
					sqlCondition = append(sqlCondition, deptAlias+".dept_id = ?")
					sqlArg = append(sqlArg, 0)
				}
			}
		}

		if len(sqlCondition) > 0 {
			return db.Where(strings.Join(sqlCondition, " OR "), sqlArg...)
		}

		return db
	}
}
