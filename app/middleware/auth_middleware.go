package middleware

import (
	"ruoyi-go/app/security"
	"ruoyi-go/app/token"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/response"
	"time"

	"github.com/gin-gonic/gin"
)

// 认证中间件
func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authUser := security.GetAuthUser(ctx)
		if authUser == nil {
			response.NewError().SetCode(401).SetMsg("未登录").Json(ctx)
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

		ctx.Next()
	}
}
