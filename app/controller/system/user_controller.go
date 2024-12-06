package systemcontroller

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/app/token"
	"ruoyi-go/framework/response"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// 获取部门树
func (*UserController) DeptTree(ctx *gin.Context) {

	user, _ := token.GetLoginUser(ctx)

	depts := (&service.DeptService{}).GetUserDeptTree(user.UserId)

	tree := (&service.UserService{}).DeptListToTree(depts)

	response.NewSuccess().SetData("data", tree).Json(ctx)
}

// 获取用户列表
func (*UserController) List(ctx *gin.Context) {

	var param dto.UserListRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	user, _ := token.GetLoginUser(ctx)

	users, total := (&service.UserService{}).GetUserList(param, user.UserId)

	for key, user := range users {
		users[key].Dept.DeptName = user.DeptName
		users[key].Dept.Leader = user.Leader
	}

	response.NewSuccess().SetPageData(users, total).Json(ctx)
}
