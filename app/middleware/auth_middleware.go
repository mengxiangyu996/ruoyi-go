package middleware

import (
	permsmap "ruoyi-go/app/router/perms-map"
	"ruoyi-go/app/service"
	"ruoyi-go/app/token"
	"ruoyi-go/common/types/constant"
	statusCode "ruoyi-go/common/types/status-code"
	"ruoyi-go/common/utils"
	"ruoyi-go/framework/response"
	"time"

	"github.com/gin-gonic/gin"
)

// 认证中间件
func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		user, err := token.GetLoginUser(ctx)
		if user == nil || err != nil {
			response.NewError().SetCode(statusCode.Unauthorized).SetMsg(err.Error()).Json(ctx)
			ctx.Abort()
			return
		}

		// 判断token临期，小于20分钟刷新
		if user.ExpireTime.Time.Before(time.Now().Add(time.Minute * 20)) {
			token.RefreshToken(ctx, user.UserTokenResponse)
		}

		ctx.Set("userId", user.UserId)
		ctx.Set("nickName", user.NickName)

		// 超级管理员跳过后续验证
		if user.UserId == 1 {
			ctx.Next()
			return
		}

		if user.Status != constant.NORMAL_STATUS {
			response.NewError().SetCode(601).SetMsg("用户被禁用").Json(ctx)
			ctx.Abort()
			return
		}

		// 去路由权限映射表中查询权限
		perm := permsmap.HasPermi(ctx.Request.Method + ":" + ctx.FullPath())
		if perm != "" {
			// 获取用户权限
			perms := (&service.MenuService{}).GetPermsByUserId(user.UserId)
			// 查询用户是否拥有权限
			if !utils.Contains(perms, perm) {
				response.NewError().SetCode(601).SetMsg("权限不足").Json(ctx)
			}
		}

		ctx.Next()
	}
}
