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

type DictTypeController struct{}

// 字典类型列表
func (*DictTypeController) List(ctx *gin.Context) {

	var param dto.DictTypeListRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	dictTypes, total := (&service.DictTypeService{}).GetDictTypeList(param, true)

	response.NewSuccess().SetPageData(dictTypes, total).Json(ctx)
}

// 字典类型详情
func (*DictTypeController) Detail(ctx *gin.Context) {

	dictId, _ := strconv.Atoi(ctx.Param("dictId"))

	dictType := (&service.DictTypeService{}).GetDictTypeByDictId(dictId)

	response.NewSuccess().SetData("data", dictType).Json(ctx)
}

// 新增字典类型
func (*DictTypeController) Create(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_INSERT)

	var param dto.CreateDictTypeRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateDictTypeValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if dictType := (&service.DictTypeService{}).GetDcitTypeByDictType(param.DictType); dictType.DictId > 0 {
		response.NewError().SetMsg("新增字典" + param.DictName + "失败，字典类型已存在").Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if err := (&service.DictTypeService{}).CreateDictType(dto.SaveDictType{
		DictName: param.DictName,
		DictType: param.DictType,
		Status:   param.Status,
		CreateBy: loginUser.UserName,
		Remark:   param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更新字典类型
func (*DictTypeController) Update(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

	var param dto.UpdateDictTypeRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateDictTypeValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if dictType := (&service.DictTypeService{}).GetDcitTypeByDictType(param.DictType); dictType.DictId > 0 && dictType.DictId != param.DictId {
		response.NewError().SetMsg("修改字典" + param.DictName + "失败，字典类型已存在").Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if err := (&service.DictTypeService{}).UpdateDictType(dto.SaveDictType{
		DictId:   param.DictId,
		DictName: param.DictName,
		DictType: param.DictType,
		Status:   param.Status,
		UpdateBy: loginUser.UserName,
		Remark:   param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 删除字典类型
func (*DictTypeController) Remove(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_DELETE)

	dictIds, _ := utils.StringToIntSlice(ctx.Param("dictIds"), ",")

	if err := (&service.DictTypeService{}).DeleteDictType(dictIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 获取字典选择框列表
func (*DictTypeController) Optionselect(ctx *gin.Context) {

	dictTypes, _ := (&service.DictTypeService{}).GetDictTypeList(dto.DictTypeListRequest{
		Status: "0",
	}, false)

	response.NewSuccess().SetData("data", dictTypes).Json(ctx)
}
