package systemcontroller

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/security"
	"ruoyi-go/app/service"
	"ruoyi-go/app/validator"
	"ruoyi-go/framework/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuController struct{}

// 菜单列表
func (*MenuController) List(ctx *gin.Context) {

	var param dto.MenuListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	menus := (&service.MenuService{}).GetMenuList(param)

	response.NewSuccess().SetData("data", menus).Json(ctx)
}

// 菜单详情
func (*MenuController) Detail(ctx *gin.Context) {

	menuId, _ := strconv.Atoi(ctx.Param("menuId"))

	menu := (&service.MenuService{}).GetMenuByMenuId(menuId)

	response.NewSuccess().SetData("data", menu).Json(ctx)
}

// 获取菜单下拉树列表
func (*MenuController) Treeselect(ctx *gin.Context) {

	menus := (&service.MenuService{}).MenuSelect()

	tree := (&service.MenuService{}).MenuSeleteToTree(menus, 0)

	response.NewSuccess().SetData("data", tree).Json(ctx)
}

// 加载对应角色菜单列表树
func (*MenuController) RoleMenuTreeselect(ctx *gin.Context) {

	roleId, _ := strconv.Atoi(ctx.Param("roleId"))
	roleHasMenuIds := (&service.MenuService{}).GetMenuIdsByRoleId(roleId)

	menus := (&service.MenuService{}).MenuSelect()
	tree := (&service.MenuService{}).MenuSeleteToTree(menus, 0)

	response.NewSuccess().SetData("menus", tree).SetData("checkedKeys", roleHasMenuIds).Json(ctx)
}

// 新增菜单
func (*MenuController) Create(ctx *gin.Context) {

	var param dto.CreateMenuRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateMenuValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if menu := (&service.MenuService{}).GetMenuByMenuName(param.MenuName); menu.MenuId > 0 {
		response.NewError().SetMsg("新增菜单" + param.MenuName + "失败，菜单名称已存在").Json(ctx)
		return
	}

	if err := (&service.MenuService{}).CreateMenu(dto.SaveMenu{
		MenuName:  param.MenuName,
		ParentId:  param.ParentId,
		OrderNum:  param.OrderNum,
		Path:      param.Path,
		Component: param.Component,
		Query:     param.Query,
		RouteName: param.RouteName,
		IsFrame:   &param.IsFrame,
		IsCache:   &param.IsCache,
		MenuType:  param.MenuType,
		Visible:   param.Visible,
		Perms:     param.Perms,
		Icon:      param.Icon,
		Status:    param.Status,
		CreateBy:  security.GetAuthUserName(ctx),
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更新菜单
func (*MenuController) Update(ctx *gin.Context) {

	var param dto.UpdateMenuRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateMenuValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if menu := (&service.MenuService{}).GetMenuByMenuName(param.MenuName); menu.MenuId > 0 && menu.MenuId != param.MenuId {
		response.NewError().SetMsg("修改菜单" + param.MenuName + "失败，菜单名称已存在").Json(ctx)
		return
	}

	if err := (&service.MenuService{}).UpdateMenu(dto.SaveMenu{
		MenuId:    param.MenuId,
		MenuName:  param.MenuName,
		ParentId:  param.ParentId,
		OrderNum:  param.OrderNum,
		Path:      param.Path,
		Component: param.Component,
		Query:     param.Query,
		RouteName: param.RouteName,
		IsFrame:   &param.IsFrame,
		IsCache:   &param.IsCache,
		MenuType:  param.MenuType,
		Visible:   param.Visible,
		Perms:     param.Perms,
		Icon:      param.Icon,
		Status:    param.Status,
		UpdateBy:  security.GetAuthUserName(ctx),
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 删除菜单
func (*MenuController) Remove(ctx *gin.Context) {

	menuId, _ := strconv.Atoi(ctx.Param("menuId"))

	if (&service.MenuService{}).MenuHasChildren(menuId) {
		response.NewError().SetMsg("存在子菜单，不允许删除").Json(ctx)
		return
	}

	if (&service.MenuService{}).MenuExistRole(menuId) {
		response.NewError().SetMsg("菜单已分配，不允许删除").Json(ctx)
		return
	}

	if err := (&service.MenuService{}).DeleteMenu(menuId); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}
