package systemcontroller

import (
	"ruoyi-go/app/controller/validator"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/app/token"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/common/utils"
	"ruoyi-go/framework/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ConfigController struct{}

// 参数列表
func (*ConfigController) List(ctx *gin.Context) {

	var param dto.ConfigListRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	configs, total := (&service.ConfigService{}).GetConfigList(param)

	response.NewSuccess().SetPageData(configs, total).Json(ctx)
}

// 参数详情
func (*ConfigController) Detail(ctx *gin.Context) {

	configId, _ := strconv.Atoi(ctx.Param("configId"))

	config := (&service.ConfigService{}).GetConfigByConfigId(configId)

	response.NewSuccess().SetData("data", config).Json(ctx)
}

// 新增参数
func (*ConfigController) Create(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_INSERT)

	var param dto.CreateConfigRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateConfigValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if config := (&service.ConfigService{}).GetConfigByConfigKey(param.ConfigKey); config.ConfigId > 0 {
		response.NewError().SetMsg("新增参数" + param.ConfigName + "失败，参数键名已存在").Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if err := (&service.ConfigService{}).CreateConfig(dto.SaveConfig{
		ConfigName:  param.ConfigName,
		ConfigKey:   param.ConfigKey,
		ConfigValue: param.ConfigValue,
		ConfigType:  param.ConfigType,
		Remark:      param.Remark,
		CreateBy:    loginUser.UserName,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更新参数
func (*ConfigController) Update(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

	var param dto.UpdateConfigRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateConfigValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if config := (&service.ConfigService{}).GetConfigByConfigKey(param.ConfigKey); config.ConfigId > 0 && config.ConfigId != param.ConfigId {
		response.NewError().SetMsg("修改参数" + param.ConfigName + "失败，参数键名已存在").Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if err := (&service.ConfigService{}).UpdateConfig(dto.SaveConfig{
		ConfigId:    param.ConfigId,
		ConfigName:  param.ConfigName,
		ConfigKey:   param.ConfigKey,
		ConfigValue: param.ConfigValue,
		ConfigType:  param.ConfigType,
		Remark:      param.Remark,
		UpdateBy:    loginUser.UserName,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 删除参数
func (*ConfigController) Remove(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_DELETE)

	configIds, err := utils.StringToIntSlice(ctx.Param("configIds"), ",")
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err = (&service.ConfigService{}).DeleteConfig(configIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 根据配置key获取配置值
func (*ConfigController) ConfigKey(ctx *gin.Context) {

	configKey := ctx.Param("configKey")

	config := (&service.ConfigService{}).GetConfigByConfigKey(configKey)

	response.NewSuccess().SetMsg(config.ConfigValue).Json(ctx)
}
