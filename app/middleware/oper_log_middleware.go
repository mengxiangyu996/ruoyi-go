package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/security"
	"ruoyi-go/app/service"
	ipaddress "ruoyi-go/common/ip-address"
	responsewriter "ruoyi-go/common/response-writer"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/datetime"
	"ruoyi-go/framework/response"
	"time"

	"github.com/gin-gonic/gin"
)

// 操作日志中间件
// title 操作模块标题
// businessType 操作类型 constant.REQUEST_BUSINESS_TYPE_*
func OperLogMiddleware(title string, businessType int) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var operName, deptName string

		if authUser := security.GetAuthUser(ctx); authUser != nil {
			operName = authUser.NickName
			deptName = authUser.DeptName
		}

		// 记录请求时间，用于获取请求耗时
		requestStartTime := time.Now()

		// 因读取请求体后，请求体的数据流会被消耗完毕，未避免EOF错误，需要缓存请求体，并且每次使用后需要重新赋值给ctx.Request.Body
		bodyBytes, _ := ctx.GetRawData()
		// 将缓存的请求体重新赋值给ctx.Request.Body，供下方ctx.ShouldBind使用
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		rw := &responsewriter.ResponseWriter{
			ResponseWriter: ctx.Writer,
			Body:           bytes.NewBufferString(""),
		}

		param := make(map[string]interface{}, 0)

		ctx.ShouldBind(&param)

		// 因ctx.ShouldBind后，请求体的数据流会被消耗完毕，需要将缓存的请求体重新赋值给ctx.Request.Body
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// 将query参数转为map并添加到请求参数中，用query-key的形式以便区分
		for key, value := range ctx.Request.URL.Query() {
			param[key] = value
		}

		operParam, _ := json.Marshal(&param)

		ipaddress := ipaddress.GetAddress(ctx.ClientIP(), ctx.Request.UserAgent())

		sysOperLog := dto.SaveOperLogRequest{
			Title:         title,
			BusinessType:  businessType,
			Method:        ctx.HandlerName(),
			RequestMethod: ctx.Request.Method,
			OperName:      operName,
			DeptName:      deptName,
			OperUrl:       ctx.Request.URL.Path,
			OperIp:        ipaddress.Ip,
			OperLocation:  ipaddress.Addr,
			OperParam:     string(operParam),
			JsonResult:    "",
			Status:        constant.NORMAL_STATUS,
			ErrorMsg:      "",
			OperTime:      datetime.Datetime{Time: time.Now()},
			CostTime:      0,
		}

		ctx.Writer = rw

		ctx.Next()

		sysOperLog.JsonResult = rw.Body.String()

		// 解析响应
		var body response.Response
		json.Unmarshal(rw.Body.Bytes(), &body)

		if body.Code != 200 {
			sysOperLog.Status = constant.EXCEPTION_STATUS
			sysOperLog.ErrorMsg = body.Msg
		}

		duration := time.Since(requestStartTime)
		sysOperLog.CostTime = int(duration.Milliseconds())

		(&service.OperLogService{}).CreateSysOperLog(sysOperLog)
	}
}
