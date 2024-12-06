package systemcontroller

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/framework/response"

	"github.com/gin-gonic/gin"
)

type DictTypeController struct{}

// 字典类型列表
func (*DictTypeController) List(ctx *gin.Context) {

	var param dto.DictTypeRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	dictTypes, total := (&service.DictService{}).GetDictTypeList(param)

	response.NewSuccess().SetPageData(dictTypes, total).Json(ctx)
}
