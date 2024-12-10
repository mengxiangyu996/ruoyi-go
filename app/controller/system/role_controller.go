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

// 角色详情
func (*RoleController) Detail(ctx *gin.Context) {

	roleId, _ := strconv.Atoi(ctx.Param("roleId"))

	role := (&service.RoleService{}).GetRoleByRoleId(roleId)

	response.NewSuccess().SetData("data", role).Json(ctx)
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
		response.NewError().SetMsg("新增角色" + param.RoleName + "失败，角色名已存在").Json(ctx)
		return
	}

	if role := (&service.RoleService{}).GetRoleByRoleKey(param.RoleKey); role.RoleId > 0 {
		response.NewError().SetMsg("新增角色" + param.RoleName + "失败，权限字符已存在").Json(ctx)
		return
	}

	menuCheckStrictly, deptCheckStrictly := 0, 0
	if param.MenuCheckStrictly {
		menuCheckStrictly = 1
	}
	if param.DeptCheckStrictly {
		deptCheckStrictly = 1
	}

	if err := (&service.RoleService{}).CreateRole(dto.SaveRole{
		RoleName:          param.RoleName,
		RoleKey:           param.RoleKey,
		RoleSort:          param.RoleSort,
		MenuCheckStrictly: &menuCheckStrictly,
		DeptCheckStrictly: &deptCheckStrictly,
		Status:            param.Status,
		CreateBy:          loginUser.UserName,
		Remark:            param.Remark,
	}, param.MenuIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更新角色
func (*RoleController) Update(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

	var param dto.UpdateRoleRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateRoleValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if role := (&service.RoleService{}).GetRoleByRoleName(param.RoleName); role.RoleId > 0 && role.RoleId != param.RoleId {
		response.NewError().SetMsg("修改角色" + param.RoleName + "失败，角色名已存在").Json(ctx)
		return
	}

	if role := (&service.RoleService{}).GetRoleByRoleKey(param.RoleKey); role.RoleId > 0 && role.RoleId != param.RoleId {
		response.NewError().SetMsg("修改角色" + param.RoleName + "失败，权限字符已存在").Json(ctx)
		return
	}

	menuCheckStrictly, deptCheckStrictly := 0, 0
	if param.MenuCheckStrictly {
		menuCheckStrictly = 1
	}
	if param.DeptCheckStrictly {
		deptCheckStrictly = 1
	}

	if err := (&service.RoleService{}).UpdateRole(dto.SaveRole{
		RoleId:            param.RoleId,
		RoleName:          param.RoleName,
		RoleKey:           param.RoleKey,
		RoleSort:          param.RoleSort,
		MenuCheckStrictly: &menuCheckStrictly,
		DeptCheckStrictly: &deptCheckStrictly,
		Status:            param.Status,
		UpdateBy:          loginUser.UserName,
		Remark:            param.Remark,
	}, param.MenuIds, nil); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 删除角色
func (*RoleController) Remove(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_DELETE)

	roleIds, err := utils.StringToIntSlice(ctx.Param("roleIds"), ",")
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	roles := (&service.RoleService{}).GetRoleListByUserId(loginUser.UserId)

	for _, role := range roles {
		if err = validator.RemoveRoleValidator(roleIds, role.RoleId, role.RoleName); err != nil {
			response.NewError().SetMsg(err.Error()).Json(ctx)
			return
		}
	}

	if err := (&service.RoleService{}).DeleteRole(roleIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 修改角色状态
func (*RoleController) ChangeStatus(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

	var param dto.UpdateRoleRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.ChangeRoleStatusValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if err := (&service.RoleService{}).UpdateRole(dto.SaveRole{
		RoleId:   param.RoleId,
		Status:   param.Status,
		UpdateBy: loginUser.UserName,
	}, nil, nil); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 部门树
func (*RoleController) DeptTree(ctx *gin.Context) {

	roleId, _ := strconv.Atoi(ctx.Param("roleId"))
	roleHasDeptIds := (&service.DeptService{}).GetDeptIdsByRoleId(roleId)

	depts := (&service.DeptService{}).DeptSelect()
	tree := (&service.DeptService{}).DeptSeleteToTree(depts, 0)

	response.NewSuccess().SetData("depts", tree).SetData("checkedKeys", roleHasDeptIds).Json(ctx)
}

// 分配数据权限
func (*RoleController) DataScope(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

	var param dto.UpdateRoleRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	deptCheckStrictly := 0
	if param.DeptCheckStrictly {
		deptCheckStrictly = 1
	}

	if err := (&service.RoleService{}).UpdateRole(dto.SaveRole{
		RoleId:            param.RoleId,
		DataScope:         param.DataScope,
		DeptCheckStrictly: &deptCheckStrictly,
		UpdateBy:          loginUser.UserName,
	}, nil, param.DeptIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 查询已分配用户角色列表
func (*RoleController) RoleAuthUserAllocatedList(ctx *gin.Context) {

	var param dto.RoleAuthUserAllocatedListRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	users, total := (&service.UserService{}).GetUserListByRoleId(param, loginUser.UserId, true)

	response.NewSuccess().SetPageData(users, total).Json(ctx)
}

// 查询未分配用户角色列表
func (*RoleController) RoleAuthUserUnallocatedList(ctx *gin.Context) {

	var param dto.RoleAuthUserAllocatedListRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	users, total := (&service.UserService{}).GetUserListByRoleId(param, loginUser.UserId, false)

	response.NewSuccess().SetPageData(users, total).Json(ctx)
}

// 批量选择用户授权
func (*RoleController) RoleAuthUserSelectAll(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

	var param dto.RoleAuthUserSelectAllRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	userIds, err := utils.StringToIntSlice(param.UserIds, ",")
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err = (&service.RoleService{}).AuthUserSelectAll(param.RoleId, userIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 取消授权用户
func (*RoleController) RoleAuthUserCancel(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

	var param dto.RoleAuthUserCancelRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := (&service.RoleService{}).AuthUserDelete(param.RoleId, []int{param.UserId}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 批量取消授权用户
func (*RoleController) RoleAuthUserCancelAll(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

	var param dto.RoleAuthUserCancelAllRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	userIds, err := utils.StringToIntSlice(param.UserIds, ",")
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := (&service.RoleService{}).AuthUserDelete(param.RoleId, userIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 导出角色数据
func (*RoleController) Export(ctx *gin.Context) {

	// TODO

}
