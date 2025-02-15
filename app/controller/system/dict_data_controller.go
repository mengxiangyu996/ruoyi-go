package systemcontroller

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/security"
	"ruoyi-go/app/service"
	"ruoyi-go/app/validator"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/common/utils"
	"ruoyi-go/framework/response"
	"strconv"
	"time"

	"gitee.com/hanshuangjianke/go-excel/excel"
	"github.com/gin-gonic/gin"
)

type DictDataController struct{}

// 获取字典数据列表
func (*DictDataController) List(ctx *gin.Context) {

	var param dto.DictDataListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	dictDatas, total := (&service.DictDataService{}).GetDictDataList(param, true)

	response.NewSuccess().SetPageData(dictDatas, total).Json(ctx)
}

// 获取字典数据详情
func (*DictDataController) Detail(ctx *gin.Context) {

	dictCode, _ := strconv.Atoi(ctx.Param("dictCode"))

	dictData := (&service.DictDataService{}).GetDictDataByDictCode(dictCode)

	response.NewSuccess().SetData("data", dictData).Json(ctx)
}

// 新增字典数据
func (*DictDataController) Create(ctx *gin.Context) {

	var param dto.CreateDictDataRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateDictDataValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := (&service.DictDataService{}).CreateDictData(dto.SaveDictData{
		DictSort:  param.DictSort,
		DictLabel: param.DictLabel,
		DictValue: param.DictValue,
		DictType:  param.DictType,
		CssClass:  param.CssClass,
		ListClass: param.ListClass,
		IsDefault: param.IsDefault,
		Status:    param.Status,
		CreateBy:  security.GetAuthUserName(ctx),
		Remark:    param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更新字典数据
func (*DictDataController) Update(ctx *gin.Context) {

	var param dto.UpdateDictDataRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateDictDataValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := (&service.DictDataService{}).UpdateDictData(dto.SaveDictData{
		DictCode:  param.DictCode,
		DictSort:  param.DictSort,
		DictLabel: param.DictLabel,
		DictValue: param.DictValue,
		DictType:  param.DictType,
		CssClass:  param.CssClass,
		ListClass: param.ListClass,
		IsDefault: param.IsDefault,
		Status:    param.Status,
		UpdateBy:  security.GetAuthUserName(ctx),
		Remark:    param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 删除字典数据
func (*DictDataController) Remove(ctx *gin.Context) {

	dictCodes, err := utils.StringToIntSlice(ctx.Param("dictCodes"), ",")
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err = (&service.DictDataService{}).DeleteDictData(dictCodes); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 根据字典类型查询字典数据
func (*DictDataController) Type(ctx *gin.Context) {

	dictType := ctx.Param("dictType")

	dictDatas := (&service.DictDataService{}).GetDictDataCacheByDictType(dictType)

	for key, dictData := range dictDatas {
		dictDatas[key].Default = dictData.IsDefault == constant.IS_DEFAULT_YES
	}

	response.NewSuccess().SetData("data", dictDatas).Json(ctx)
}

// 数据导出
func (*DictDataController) Export(ctx *gin.Context) {

	var param dto.DictDataListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	list := make([]dto.DictDataExportResponse, 0)

	dictDatas, _ := (&service.DictDataService{}).GetDictDataList(param, false)
	for _, dictData := range dictDatas {
		list = append(list, dto.DictDataExportResponse{
			DictCode:  dictData.DictCode,
			DictSort:  dictData.DictSort,
			DictLabel: dictData.DictLabel,
			DictValue: dictData.DictValue,
			DictType:  dictData.DictType,
			IsDefault: dictData.IsDefault,
			Status:    dictData.Status,
		})
	}

	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	excel.DownLoadExcel("data_"+time.Now().Format("20060102150405"), ctx.Writer, file)
}
