package systemcontroller

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/framework/response"

	"github.com/gin-gonic/gin"
)

type DictController struct{}

// 字典类型列表
func (*DictController) List(ctx *gin.Context) {

	var param dto.DictTypeRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).ToJson(ctx)
		return
	}

	dictTypes, total := (&service.DictService{}).GetDictTypeList(param)

	response.NewSuccess().SetPageData(dictTypes, total).ToJson(ctx)
}
