package middleware

import (
	"ruoyi-go/app/security"
	"ruoyi-go/app/token"
	"ruoyi-go/common/types/constant"
	statusCode "ruoyi-go/common/types/status-code"
	"ruoyi-go/framework/response"
	"time"

	"github.com/gin-gonic/gin"
)

// 认证中间件
func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authUser := security.GetAuthUser(ctx)
		if authUser == nil {
			response.NewError().SetCode(statusCode.Unauthorized).SetMsg("未登录").Json(ctx)
			ctx.Abort()
			return
		}

		// 判断token临期，小于20分钟刷新
		if authUser.ExpireTime.Time.Before(time.Now().Add(time.Minute * 20)) {
			token.RefreshToken(ctx, authUser.UserTokenResponse)
		}

		if authUser.Status != constant.NORMAL_STATUS {
			response.NewError().SetCode(601).SetMsg("用户被禁用").Json(ctx)
			ctx.Abort()
			return
		}

		// 去路由权限映射表中查询权限（转为在路由中加入权限中间件方法）
		// perm := permsmap.HasPerm(ctx.Request.Method + ":" + ctx.FullPath())
		// if perm != "" {
		// 	// 获取用户权限
		// 	perms := (&service.MenuService{}).GetPermsByUserId(authUser.UserId)
		// 	// 查询用户是否拥有权限
		// 	if !utils.Contains(perms, perm) {
		// 		response.NewError().SetCode(601).SetMsg("权限不足").Json(ctx)
		// 		ctx.Abort()
		// 		return
		// 	}
		// }

		ctx.Next()
	}
}
