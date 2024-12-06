package controller

import (
	"ruoyi-go/app/service"
	"ruoyi-go/framework/response"

	"github.com/gin-gonic/gin"
)

type ConfigController struct{}

// 根据配置key获取配置值
func (*ConfigController) GetConfigValueByKey(ctx *gin.Context) {

	configKey := ctx.Param("configKey")

	config := (&service.ConfigService{}).GetConfigByConfigKey(configKey)

	response.NewSuccess().SetMsg(config.ConfigValue).Json(ctx)
}
