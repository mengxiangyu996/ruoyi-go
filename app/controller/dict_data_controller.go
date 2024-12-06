package controller

import (
	"ruoyi-go/app/service"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/response"

	"github.com/gin-gonic/gin"
)

type DictDataController struct{}

// 根据字典类型查询字典数据
func (*DictDataController) GetDictDataByDictType(ctx *gin.Context) {

	dictType := ctx.Param("dictType")

	dictDatas := (&service.DictService{}).GetDictDataByType(dictType)

	for key, dictData := range dictDatas {
		dictDatas[key].Default = dictData.IsDefault == constant.IS_DEFAULT_YES
	}

	response.NewSuccess().SetData("data", dictDatas).Json(ctx)
}
