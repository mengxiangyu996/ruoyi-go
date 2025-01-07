package security

import (
	"ruoyi-go/app/service"
	"ruoyi-go/app/token"

	"github.com/gin-gonic/gin"
)

// 获取用户id
func GetAuthUserId(ctx *gin.Context) int {

	authUser, err := token.GetAuhtUser(ctx)
	if err != nil {
		return 0
	}

	return authUser.UserId
}

// 获取部门id
func GetAuthDeptId(ctx *gin.Context) int {

	authUser, err := token.GetAuhtUser(ctx)
	if err != nil {
		return 0
	}

	return authUser.DeptId
}

// 获取用户账户
func GetAuthUserName(ctx *gin.Context) string {

	authUser, err := token.GetAuhtUser(ctx)
	if err != nil {
		return ""
	}

	return authUser.UserName
}

// 获取用户
func GetAuthUser(ctx *gin.Context) *token.UserTokenResponse {

	authUser, err := token.GetAuhtUser(ctx)
	if err != nil {
		return nil
	}

	return authUser
}

// 验证用户是否具备某权限，相当于 @PreAuthorize("@ss.hasPermi('system:user:list')")
func HasPerm(userId int, perm string) bool {
	return (&service.UserService{}).UserHasPerms(userId, []string{perm})
}

// 验证用户是否不具备某权限，与 HasPerm 逻辑相反，相当于 @PreAuthorize("@ss.lacksPermi('system:user:list')")
func LacksPerm(userId int, perm string) bool {
	return !(&service.UserService{}).UserHasPerms(userId, []string{perm})
}

// 验证用户是否具有以下任意一个权限，相当于 @PreAuthorize("@ss.hasAnyPermi('system:user:add,system:user:edit')")
func HasAnyPerms(userId int, perms []string) bool {
	return (&service.UserService{}).UserHasPerms(userId, perms)
}

// 验证用户是否拥有某个角色，相当于 @PreAuthorize("@ss.hasRole('user')")
func HasRole(userId int, roleKey string) bool {
	return (&service.UserService{}).UserHasRoles(userId, []string{roleKey})
}

// 验证用户是否不具备某个角色，与 HasRole 逻辑相反，相当于 @PreAuthorize("@ss.lacksRole('user')")
func LacksRole(userId int, roleKey string) bool {
	return !(&service.UserService{}).UserHasRoles(userId, []string{roleKey})
}

// 验证用户是否具有以下任意一个角色，相当于 @PreAuthorize("@ss.hasAnyRoles('user,admin')")
func HasAnyRoles(userId int, roleKey []string) bool {
	return (&service.UserService{}).UserHasPerms(userId, roleKey)
}
