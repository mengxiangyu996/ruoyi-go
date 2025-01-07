package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	ipaddress "ruoyi-go/common/ip-address"
	responsewriter "ruoyi-go/common/response-writer"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/datetime"
	"ruoyi-go/framework/response"
	"time"

	"github.com/gin-gonic/gin"
)

// 登录信息记录
func LogininforMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// 因读取请求体后，请求体的数据流会被消耗完毕，未避免EOF错误，需要缓存请求体，并且每次使用后需要重新赋值给ctx.Request.Body
		bodyBytes, _ := ctx.GetRawData()
		// 将缓存的请求体重新赋值给ctx.Request.Body，供下方ctx.ShouldBind使用
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		rw := &responsewriter.ResponseWriter{
			ResponseWriter: ctx.Writer,
			Body:           bytes.NewBufferString(""),
		}

		var param dto.LoginRequest

		if err := ctx.ShouldBind(&param); err != nil {
			response.NewError().SetCode(400).SetMsg(err.Error()).Json(ctx)
			ctx.Abort()
			return
		}

		// 因ctx.ShouldBind后，请求体的数据流会被消耗完毕，需要将缓存的请求体重新赋值给ctx.Request.Body
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		ipaddress := ipaddress.GetAddress(ctx.ClientIP(), ctx.Request.UserAgent())

		logininfor := dto.SaveLogininforRequest{
			UserName:      param.Username,
			Ipaddr:        ipaddress.Ip,
			LoginLocation: ipaddress.Addr,
			Browser:       ipaddress.Browser,
			Os:            ipaddress.Os,
			Status:        constant.NORMAL_STATUS,
			LoginTime:     datetime.Datetime{Time: time.Now()},
		}

		ctx.Writer = rw

		ctx.Next()

		// 解析响应
		var body response.Response
		json.Unmarshal(rw.Body.Bytes(), &body)

		if body.Code != 200 {
			logininfor.Status = constant.EXCEPTION_STATUS
		}
		logininfor.Msg = body.Msg

		(&service.LogininforService{}).CreateSysLogininfor(logininfor)
	}
}
