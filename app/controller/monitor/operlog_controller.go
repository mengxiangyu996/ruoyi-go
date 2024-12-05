package monitorcontroller

import (
	"regexp"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/framework/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type OperlogController struct{}

// 操作记录列表
func (*OperlogController) List(ctx *gin.Context) {

	var param dto.OperLogListRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).ToJson(ctx)
		return
	}

	// 排序规则默认为倒序（DESC）
	param.OrderRule = "DESC"
	if strings.HasPrefix(param.IsAsc, "asc") {
		param.OrderRule = ""
	}

	// 排序字段小驼峰转蛇形
	if param.OrderByColumn == "" {
		param.OrderByColumn = "operTime"
	}
	param.OrderByColumn = strings.ToLower(regexp.MustCompile("([A-Z])").ReplaceAllString(param.OrderByColumn, "_${1}"))

	operLogs, total := (&service.OperLogService{}).GetOperLogList(param)

	response.NewSuccess().SetPageData(operLogs, total).ToJson(ctx)
}
