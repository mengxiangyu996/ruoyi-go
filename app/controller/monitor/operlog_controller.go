package monitorcontroller

import (
	"regexp"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/common/utils"
	"ruoyi-go/framework/response"
	"strconv"
	"strings"
	"time"

	"gitee.com/hanshuangjianke/go-excel/excel"
	"github.com/gin-gonic/gin"
)

type OperlogController struct{}

// 操作日志列表
func (*OperlogController) List(ctx *gin.Context) {

	var param dto.OperLogListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	// 排序规则默认为倒序（DESC）
	param.OrderRule = "DESC"
	if strings.HasPrefix(param.IsAsc, "asc") {
		param.OrderRule = "ASC"
	}

	// 排序字段小驼峰转蛇形
	if param.OrderByColumn == "" {
		param.OrderByColumn = "operTime"
	}
	param.OrderByColumn = strings.ToLower(regexp.MustCompile("([A-Z])").ReplaceAllString(param.OrderByColumn, "_${1}"))

	operLogs, total := (&service.OperLogService{}).GetOperLogList(param, true)

	response.NewSuccess().SetPageData(operLogs, total).Json(ctx)
}

// 删除操作日志
func (*OperlogController) Remove(ctx *gin.Context) {

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

	if err := (&service.OperLogService{}).DeleteOperLog(nil); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 数据导出
func (*OperlogController) Export(ctx *gin.Context) {

	var param dto.OperLogListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	// 排序规则默认为倒序（DESC）
	param.OrderRule = "DESC"
	if strings.HasPrefix(param.IsAsc, "asc") {
		param.OrderRule = "ASC"
	}

	// 排序字段小驼峰转蛇形
	if param.OrderByColumn == "" {
		param.OrderByColumn = "operTime"
	}
	param.OrderByColumn = strings.ToLower(regexp.MustCompile("([A-Z])").ReplaceAllString(param.OrderByColumn, "_${1}"))

	list := make([]dto.OperLogExportResponse, 0)

	operLogs, _ := (&service.OperLogService{}).GetOperLogList(param, false)
	for _, operLog := range operLogs {
		list = append(list, dto.OperLogExportResponse{
			OperId:        operLog.OperId,
			Title:         operLog.Title,
			BusinessType:  operLog.BusinessType,
			Method:        operLog.Method,
			RequestMethod: operLog.RequestMethod,
			OperName:      operLog.OperName,
			DeptName:      operLog.DeptName,
			OperUrl:       operLog.OperUrl,
			OperIp:        operLog.OperIp,
			OperLocation:  operLog.OperLocation,
			OperParam:     operLog.OperParam,
			JsonResult:    operLog.JsonResult,
			Status:        operLog.Status,
			ErrorMsg:      operLog.ErrorMsg,
			OperTime:      operLog.OperTime.Format("2006-01-02 15:04:05"),
			CostTime:      strconv.Itoa(operLog.CostTime) + "毫秒",
		})
	}

	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)

	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	excel.DownLoadExcel("operlog_"+time.Now().Format("20060102150405"), ctx.Writer, file)
}
