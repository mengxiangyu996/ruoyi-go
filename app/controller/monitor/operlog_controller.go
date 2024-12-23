package monitorcontroller

import (
	"regexp"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/common/utils"
	"ruoyi-go/framework/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type OperlogController struct{}

// 操作日志列表
func (*OperlogController) List(ctx *gin.Context) {

	var param dto.OperLogListRequest

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
		param.OrderByColumn = "operTime"
	}
	param.OrderByColumn = strings.ToLower(regexp.MustCompile("([A-Z])").ReplaceAllString(param.OrderByColumn, "_${1}"))

	operLogs, total := (&service.OperLogService{}).GetOperLogList(param)

	response.NewSuccess().SetPageData(operLogs, total).Json(ctx)
}

// 删除操作日志
func (*OperlogController) Remove(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_DELETE)

	operIds, err := utils.StringToIntSlice(ctx.Param("operIds"), ",")
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err = (&service.OperLogService{}).DeleteOperLog(operIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 清空操作日志
func (*OperlogController) Clean(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_DELETE)

	if err := (&service.OperLogService{}).DeleteOperLog(nil); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}
