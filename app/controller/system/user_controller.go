package systemcontroller

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/app/token"
	"ruoyi-go/app/validator"
	"ruoyi-go/common/password"
	"ruoyi-go/framework/response"
	"strings"

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

// 个人信息
func (*UserController) GetUserProfile(ctx *gin.Context) {

	loginUser, _ := token.GetLoginUser(ctx)

	user := (&service.UserService{}).GetUserByUserId(loginUser.UserId)

	user.Admin = user.UserId == 1

	dept := (&service.DeptService{}).GetDeptByDeptId(user.DeptId)

	roles := (&service.RoleService{}).GetRoleListByUserId(user.UserId)

	data := dto.AuthUserInfoResponse{
		UserDetailResponse: user,
		Dept:               dept,
		Roles:              roles,
	}

	// 获取角色组
	roleGroup := (&service.RoleService{}).GetRoleNamesByUserId(user.UserId)

	// 获取岗位组
	postGroup := (&service.PostService{}).GetPostNamesByUserId(user.UserId)

	response.NewSuccess().
		SetData("data", data).
		SetData("roleGroup", strings.Join(roleGroup, ",")).
		SetData("postGroup", strings.Join(postGroup, ",")).
		Json(ctx)
}

// 修改个人信息
func (*UserController) UpdateUserProfile(ctx *gin.Context) {

	var param dto.UpdateUserProfile

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateUserProfileValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	user, _ := token.GetLoginUser(ctx)

	if err := (&service.UserService{}).UpdateUser(dto.UpdateUser{
		UserId:      user.UserId,
		NickName:    param.NickName,
		Email:       param.Email,
		Phonenumber: param.Phonenumber,
		Sex:         param.Sex,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 修改个人密码
func (*UserController) UpdateUserProfilePassword(ctx *gin.Context) {

	var param dto.UpdateUserProfilePassword

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateUserProfilePasswordValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	user := (&service.UserService{}).GetUserByUserId(loginUser.UserId)
	if !password.Verify(user.Password, param.OldPassword) {
		response.NewError().SetMsg("旧密码输入错误").Json(ctx)
		return
	}

	if err := (&service.UserService{}).UpdateUser(dto.UpdateUser{
		UserId:   user.UserId,
		Password: password.Generate(param.NewPassword),
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}
