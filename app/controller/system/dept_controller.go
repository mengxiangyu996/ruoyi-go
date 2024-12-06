package systemcontroller

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/app/token"
	"ruoyi-go/framework/response"

	"github.com/gin-gonic/gin"
)

type DeptController struct{}

// 部门列表
func (*DeptController) List(ctx *gin.Context) {

	var param dto.DeptListRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	user, _ := token.GetLoginUser(ctx)

	depts := (&service.DeptService{}).GetDeptList(param, user.UserId)

	response.NewSuccess().SetData("data", depts).Json(ctx)
}
