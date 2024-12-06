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
