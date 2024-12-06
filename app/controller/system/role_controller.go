package systemcontroller

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
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
