package middleware

import (
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

		loginUser, err := token.GetLoginUser(ctx)
		if loginUser == nil || err != nil {
			response.NewError().SetCode(statusCode.Unauthorized).SetMsg(err.Error()).Json(ctx)
			ctx.Abort()
			return
		}

		// 判断token临期，小于20分钟刷新
		if loginUser.ExpireTime.Time.Before(time.Now().Add(time.Minute * 20)) {
			token.RefreshToken(ctx, loginUser.UserTokenResponse)
		}

		ctx.Set("userId", loginUser.UserId)
		ctx.Set("nickName", loginUser.NickName)

		// 超级管理员跳过后续验证
		if loginUser.UserId == 1 {
			ctx.Next()
			return
		}

		if loginUser.Status != constant.NORMAL_STATUS {
			response.NewError().SetCode(601).SetMsg("用户被禁用").Json(ctx)
			ctx.Abort()
			return
		}

		// 去路由权限映射表中查询权限（转为在路由中加入权限中间件方法）
		// perm := permsmap.HasPerm(ctx.Request.Method + ":" + ctx.FullPath())
		// if perm != "" {
		// 	// 获取用户权限
		// 	perms := (&service.MenuService{}).GetPermsByUserId(loginUser.UserId)
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
