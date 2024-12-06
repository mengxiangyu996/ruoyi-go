package monitorcontroller

import (
	"regexp"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/framework/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type LogininforController struct{}

// 登录日志列表
func (*LogininforController) List(ctx *gin.Context) {

	var param dto.LogininforListRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	// 排序规则默认为倒序（DESC）
	param.OrderRule = "DESC"
	if strings.HasPrefix(param.IsAsc, "asc") {
		param.OrderRule = ""
	}

	// 排序字段小驼峰转蛇形
	if param.OrderByColumn == "" {
		param.OrderByColumn = "loginTime"
	}
	param.OrderByColumn = strings.ToLower(regexp.MustCompile("([A-Z])").ReplaceAllString(param.OrderByColumn, "_${1}"))

	logininfors, total := (&service.LogininforService{}).GetLogininforList(param)

	response.NewSuccess().SetPageData(logininfors, total).Json(ctx)
}
