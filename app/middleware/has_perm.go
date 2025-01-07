package middleware

import (
	"ruoyi-go/app/security"
	"ruoyi-go/framework/response"

	"github.com/gin-gonic/gin"
)

// 验证用户是否具备某权限
//
// 为了实现@PreAuthorize("@ss.hasPermi('system:user:list')")注解
//
// 用法：api.GET("/system/user/deptTree", middleware.HasPerm("system:user:list"), (&systemcontroller.UserController{}).DeptTree)
func HasPerm(perm string) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authUserId := security.GetAuthUserId(ctx)
		if authUserId == 1 {
			ctx.Next()
			return
		}

		if hasPerm := security.HasPerm(authUserId, perm); !hasPerm {
			response.NewError().SetCode(601).SetMsg("权限不足").Json(ctx)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
