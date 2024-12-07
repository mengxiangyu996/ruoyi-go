package middleware

import (
	"ruoyi-go/app/service"
	"ruoyi-go/app/token"
	"ruoyi-go/framework/response"

	"github.com/gin-gonic/gin"
)

// 判断api权限是否存在
//
// 为了实现@PreAuthorize("@ss.hasPermi('system:user:list')")注解
//
// 用法：api.GET("/system/user/deptTree", middleware.HasPerm("system:user:list"), (&systemcontroller.UserController{}).DeptTree)
func HasPerm(perm string) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		loginUser, _ := token.GetLoginUser(ctx)
		if loginUser.UserId == 1 {
			ctx.Next()
			return
		}

		if hasPerm := (&service.UserService{}).UserHasPerm(loginUser.UserId, perm); !hasPerm {
			response.NewError().SetCode(601).SetMsg("权限不足").Json(ctx)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
