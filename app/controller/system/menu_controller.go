package systemcontroller

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/framework/response"

	"github.com/gin-gonic/gin"
)

type MenuController struct{}

// 菜单列表
func (*MenuController) List(ctx *gin.Context) {

	var param dto.MenuListRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	menus := (&service.MenuService{}).GetMenuList(param)

	response.NewSuccess().SetData("data", menus).Json(ctx)
}

// 获取菜单下拉树列表
func (*MenuController) Treeselect(ctx *gin.Context) {

	menus := (&service.MenuService{}).Menuselect()

	tree := (&service.MenuService{}).MenuSeleteToTree(menus, 0)

	response.NewSuccess().SetData("data", tree).Json(ctx)
}
