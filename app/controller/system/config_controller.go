package systemcontroller

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/framework/response"

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

// 根据配置key获取配置值
func (*ConfigController) ConfigKey(ctx *gin.Context) {

	configKey := ctx.Param("configKey")

	config := (&service.ConfigService{}).GetConfigByConfigKey(configKey)

	response.NewSuccess().SetMsg(config.ConfigValue).Json(ctx)
}
