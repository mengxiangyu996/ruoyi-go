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
	"strings"

	"github.com/gin-gonic/gin"
)

type DeptController struct{}

// 部门列表
func (*DeptController) List(ctx *gin.Context) {

	var param dto.DeptListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	depts := (&service.DeptService{}).GetDeptList(param, security.GetAuthUserId(ctx))

	response.NewSuccess().SetData("data", depts).Json(ctx)
}

// 查询部门列表（排除节点）
func (*DeptController) ListExclude(ctx *gin.Context) {

	deptId, _ := strconv.Atoi(ctx.Param("deptId"))

	data := make([]dto.DeptListResponse, 0)

	depts := (&service.DeptService{}).GetDeptList(dto.DeptListRequest{}, security.GetAuthUserId(ctx))
	for _, dept := range depts {
		if dept.DeptId == deptId || utils.Contains(strings.Split(dept.Ancestors, ","), strconv.Itoa(deptId)) {
			continue
		}
		data = append(data, dept)
	}

	response.NewSuccess().SetData("data", data).Json(ctx)
}

// 获取部门详情
func (*DeptController) Detail(ctx *gin.Context) {

	deptId, _ := strconv.Atoi(ctx.Param("deptId"))

	dept := (&service.DeptService{}).GetDeptByDeptId(deptId)

	response.NewSuccess().SetData("data", dept).Json(ctx)
}

// 新增部门
func (*DeptController) Create(ctx *gin.Context) {

	var param dto.CreateDeptRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateDeptValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if dept := (&service.DeptService{}).GetDeptByDeptName(param.DeptName); dept.DeptId > 0 {
		response.NewError().SetMsg("新增部门" + param.DeptName + "失败，部门名称已存在").Json(ctx)
		return
	}

	// 拼接ancestors，获取上级的祖级列表
	parentDept := (&service.DeptService{}).GetDeptByDeptId(param.ParentId)
	if parentDept.Status == constant.EXCEPTION_STATUS {
		response.NewError().SetMsg("部门停用，不允许新增").Json(ctx)
		return
	}
	ancestors := parentDept.Ancestors + "," + strconv.Itoa(parentDept.DeptId)

	if err := (&service.DeptService{}).CreateDept(dto.SaveDept{
		ParentId:  param.ParentId,
		Ancestors: ancestors,
		DeptName:  param.DeptName,
		OrderNum:  param.OrderNum,
		Leader:    param.Leader,
		Phone:     param.Phone,
		Email:     param.Email,
		Status:    param.Status,
		CreateBy:  security.GetAuthUserName(ctx),
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更新部门
func (*DeptController) Update(ctx *gin.Context) {

	var param dto.UpdateDeptRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateDeptValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if dept := (&service.DeptService{}).GetDeptByDeptName(param.DeptName); dept.DeptId > 0 && dept.DeptId != param.DeptId {
		response.NewError().SetMsg("修改部门" + param.DeptName + "失败，部门名称已存在").Json(ctx)
		return
	}

	if dept := (&service.DeptService{}).GetDeptByDeptId(param.DeptId); dept.ParentId != param.ParentId && (&service.DeptService{}).DeptHasChildren(param.DeptId) {
		response.NewError().SetMsg("存在子级部门，无法直接修改所属部门").Json(ctx)
		return
	}

	// 拼接ancestors，获取上级的祖级列表
	parentDept := (&service.DeptService{}).GetDeptByDeptId(param.ParentId)
	if parentDept.Status == constant.EXCEPTION_STATUS {
		response.NewError().SetMsg("部门停用，不允许新增").Json(ctx)
		return
	}
	ancestors := parentDept.Ancestors + "," + strconv.Itoa(parentDept.DeptId)

	if err := (&service.DeptService{}).UpdateDept(dto.SaveDept{
		DeptId:    param.DeptId,
		ParentId:  param.ParentId,
		Ancestors: ancestors,
		DeptName:  param.DeptName,
		OrderNum:  param.OrderNum,
		Leader:    param.Leader,
		Phone:     param.Phone,
		Email:     param.Email,
		Status:    param.Status,
		UpdateBy:  security.GetAuthUserName(ctx),
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 删除部门
func (*DeptController) Remove(ctx *gin.Context) {

	deptId, _ := strconv.Atoi(ctx.Param("deptId"))

	if (&service.DeptService{}).DeptHasChildren(deptId) {
		response.NewError().SetMsg("存在下级部门，不允许删除").Json(ctx)
		return
	}

	if (&service.UserService{}).UserHasDeptByDeptId(deptId) {
		response.NewError().SetMsg("部门存在用户，不允许删除").Json(ctx)
		return
	}

	if err := (&service.DeptService{}).DeleteDept(deptId); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}
