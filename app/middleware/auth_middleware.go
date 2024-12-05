package middleware

import (
	"ruoyi-go/app/service"
	"ruoyi-go/app/token"
	"ruoyi-go/common/types/constant"
	statusCode "ruoyi-go/common/types/status-code"
	"ruoyi-go/common/utils"
	"ruoyi-go/framework/response"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 认证中间件
func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		user, err := token.GetLoginUser(ctx)
		if user == nil || err != nil {
			response.NewError().SetCode(statusCode.Unauthorized).SetMsg(err.Error()).ToJson(ctx)
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
			response.NewError().SetCode(statusCode.Unauthorized).SetMsg("用户被禁用").ToJson(ctx)
			ctx.Abort()
			return
		}

		// 获取用户权限
		perms := (&service.MenuService{}).GetPermsByUserId(user.UserId)
		perm := strings.ReplaceAll(strings.ReplaceAll(ctx.Request.URL.Path, "/api/", ""), "/", ":")
		if utils.Contains(perms, perm) {
			response.NewError().SetCode(statusCode.Unauthorized).ToJson(ctx)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
