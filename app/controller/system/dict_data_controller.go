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

type DictDataController struct{}

// 获取字典数据列表
func (*DictDataController) List(ctx *gin.Context) {

	var param dto.DictDataListRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
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

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_INSERT)

	var param dto.CreateDictDataRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateDictDataValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if err := (&service.DictDataService{}).CreateDictData(dto.SaveDictData{
		DictSort:  param.DictSort,
		DictLabel: param.DictLabel,
		DictValue: param.DictValue,
		DictType:  param.DictType,
		CssClass:  param.CssClass,
		ListClass: param.ListClass,
		IsDefault: param.IsDefault,
		Status:    param.Status,
		CreateBy:  loginUser.UserName,
		Remark:    param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更新字典数据
func (*DictDataController) Update(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

	var param dto.UpdateDictDataRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateDictDataValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

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
		UpdateBy:  loginUser.UserName,
		Remark:    param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 删除字典数据
func (*DictDataController) Remove(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_DELETE)

	dictCodes, _ := utils.StringToIntSlice(ctx.Param("dictCodes"), ",")

	if err := (&service.DictDataService{}).DeleteDictData(dictCodes); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 根据字典类型查询字典数据
func (*DictDataController) Type(ctx *gin.Context) {

	dictType := ctx.Param("dictType")

	dictDatas := (&service.DictDataService{}).GetDictDataByDictType(dictType)

	for key, dictData := range dictDatas {
		dictDatas[key].Default = dictData.IsDefault == constant.IS_DEFAULT_YES
	}

	response.NewSuccess().SetData("data", dictDatas).Json(ctx)
}
