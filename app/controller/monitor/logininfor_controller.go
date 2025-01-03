package monitorcontroller

import (
	"context"
	"regexp"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/common/types/constant"
	rediskey "ruoyi-go/common/types/redis-key"
	"ruoyi-go/common/utils"
	"ruoyi-go/framework/dal"
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

// 删除登录日志
func (*LogininforController) Remove(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_DELETE)

	infoIds, err := utils.StringToIntSlice(ctx.Param("infoIds"), ",")
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err = (&service.LogininforService{}).DeleteLogininfor(infoIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 清空登录日志
func (*LogininforController) Clean(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_DELETE)

	if err := (&service.LogininforService{}).DeleteLogininfor(nil); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 账户解锁（删除登录错误次数限制10分钟缓存）
func (*LogininforController) Unlock(ctx *gin.Context) {

	userName := ctx.Param("userName")

	if _, err := dal.Redis.Del(context.Background(), rediskey.LoginPasswordErrorKey+userName).Result(); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}
