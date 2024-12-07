package systemcontroller

import (
	"ruoyi-go/app/controller/validator"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/app/token"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/response"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

// 角色列表
func (*RoleController) List(ctx *gin.Context) {

	var param dto.RoleListRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	roles, total := (&service.RoleService{}).GetRoleList(param, true)

	response.NewSuccess().SetPageData(roles, total).Json(ctx)
}

// 新增角色
func (*RoleController) Create(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_INSERT)

	var param dto.CreateRoleRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateRoleValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if role := (&service.RoleService{}).GetRoleByRoleName(param.RoleName); role.RoleId > 0 {
		response.NewError().SetMsg("角色名称已存在").Json(ctx)
		return
	}

	if role := (&service.RoleService{}).GetRoleByRoleKey(param.RoleKey); role.RoleId > 0 {
		response.NewError().SetMsg("权限字符已存在").Json(ctx)
		return
	}

	var menuCheckStrictly, deptCheckStrictly int
	if param.MenuCheckStrictly {
		menuCheckStrictly = 1
	}
	if param.DeptCheckStrictly {
		deptCheckStrictly = 1
	}

	if err := (&service.RoleService{}).CreateRole(dto.SaveRoleRequest{
		RoleName:          param.RoleName,
		RoleKey:           param.RoleKey,
		RoleSort:          param.RoleSort,
		MenuCheckStrictly: menuCheckStrictly,
		DeptCheckStrictly: deptCheckStrictly,
		Status:            param.Status,
		CreateBy:          loginUser.UserName,
		Remark:            param.Remark,
	}, param.MenuIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}
