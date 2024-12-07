package systemcontroller

import (
	"ruoyi-go/app/controller/validator"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/app/token"
	"ruoyi-go/common/password"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/common/utils"
	"ruoyi-go/framework/response"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// 获取部门树
func (*UserController) DeptTree(ctx *gin.Context) {

	loginUser, _ := token.GetLoginUser(ctx)

	depts := (&service.DeptService{}).GetUserDeptTree(loginUser.UserId)

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

	loginUser, _ := token.GetLoginUser(ctx)

	users, total := (&service.UserService{}).GetUserList(param, loginUser.UserId)

	for key, user := range users {
		users[key].Dept.DeptName = user.DeptName
		users[key].Dept.Leader = user.Leader
	}

	response.NewSuccess().SetPageData(users, total).Json(ctx)
}

// 获取用户详情
func (*UserController) Detail(ctx *gin.Context) {

	userId, _ := strconv.Atoi(ctx.Param("userId"))

	response := response.NewSuccess()

	if userId > 0 {
		user := (&service.UserService{}).GetUserByUserId(userId)

		user.Admin = user.UserId == 1

		dept := (&service.DeptService{}).GetDeptByDeptId(user.DeptId)

		roles := (&service.RoleService{}).GetRoleListByUserId(user.UserId)

		response.SetData("data", dto.AuthUserInfoResponse{
			UserDetailResponse: user,
			Dept:               dept,
			Roles:              roles,
		})

		roleIds := make([]int, 0)
		for _, role := range roles {
			roleIds = append(roleIds, role.RoleId)
		}
		response.SetData("roleIds", roleIds)

		postIds := (&service.PostService{}).GetPostIdsByUserId(user.UserId)
		response.SetData("postIds", postIds)
	}

	roles, _ := (&service.RoleService{}).GetRoleList(dto.RoleListRequest{}, false)
	if userId != 1 {
		roles = utils.Filter(roles, func(role dto.RoleListResponse) bool {
			return role.RoleId != 1
		})
	}
	response.SetData("roles", roles)

	posts, _ := (&service.PostService{}).GetPostList(dto.PostListRequest{}, false)
	response.SetData("posts", posts)

	response.Json(ctx)
}

// 新增用户
func (*UserController) Add(ctx *gin.Context) {

	var param dto.AddUserRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.AddUserValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if user := (&service.UserService{}).GetUserByUsername(param.UserName); user.UserId > 0 {
		response.NewError().SetMsg("用户名称已存在").Json(ctx)
		return
	}

	if param.Email != "" {
		if user := (&service.UserService{}).GetUserByEmail(param.Email); user.UserId > 0 {
			response.NewError().SetMsg("邮箱账号已存在").Json(ctx)
			return
		}
	}

	if param.Phonenumber != "" {
		if user := (&service.UserService{}).GetUserByPhonenumber(param.Phonenumber); user.UserId > 0 {
			response.NewError().SetMsg("手机号码已存在").Json(ctx)
			return
		}
	}

	if err := (&service.UserService{}).CreateUser(dto.SaveUser{
		DeptId:      param.DeptId,
		UserName:    param.UserName,
		NickName:    param.NickName,
		Email:       param.Email,
		Phonenumber: param.Phonenumber,
		Sex:         param.Sex,
		Password:    password.Generate(param.Password),
		Status:      param.Status,
		Remark:      param.Remark,
		CreateBy:    loginUser.UserName,
	}, param.RoleIds, param.PostIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_INSERT)

	response.NewSuccess().Json(ctx)
}

// 更新用户
func (*UserController) Update(ctx *gin.Context) {

	var param dto.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateUserValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if param.Email != "" {
		if user := (&service.UserService{}).GetUserByEmail(param.Email); user.UserId > 0 && user.UserId != param.UserId {
			response.NewError().SetMsg("邮箱账号已存在").Json(ctx)
			return
		}
	}

	if param.Phonenumber != "" {
		if user := (&service.UserService{}).GetUserByPhonenumber(param.Phonenumber); user.UserId > 0 && user.UserId != param.UserId {
			response.NewError().SetMsg("手机号码已存在").Json(ctx)
			return
		}
	}

	if err := (&service.UserService{}).UpdateUser(dto.SaveUser{
		UserId:      param.UserId,
		DeptId:      param.DeptId,
		NickName:    param.NickName,
		Email:       param.Email,
		Phonenumber: param.Phonenumber,
		Sex:         param.Sex,
		Status:      param.Status,
		Remark:      param.Remark,
		UpdateBy:    loginUser.UserName,
	}, param.RoleIds, param.PostIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

	response.NewSuccess().Json(ctx)
}

// 删除用户
func (*UserController) Remove(ctx *gin.Context) {

	userIds, err := utils.StringToIntSlice(ctx.Param("userIds"), ",")
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if err = validator.RemoveUserValidator(userIds, loginUser.UserId); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := (&service.UserService{}).DeleteUser(userIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_DELETE)

	response.NewSuccess().Json(ctx)
}

// 更改用户状态
func (*UserController) ChangeStatus(ctx *gin.Context) {

	var param dto.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.ChangeUserStatusValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if err := (&service.UserService{}).UpdateUser(dto.SaveUser{
		UserId:   param.UserId,
		Status:   param.Status,
		UpdateBy: loginUser.UserName,
	}, nil, nil); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

	response.NewSuccess().Json(ctx)
}

// 重置用户密码
func (*UserController) ResetPwd(ctx *gin.Context) {

	var param dto.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.ResetUserPwdValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if err := (&service.UserService{}).UpdateUser(dto.SaveUser{
		UserId:   param.UserId,
		Password: password.Generate(param.Password),
		UpdateBy: loginUser.UserName,
	}, nil, nil); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

	response.NewSuccess().Json(ctx)
}

// 根据用户编号获取授权角色
func (*UserController) AuthRole(ctx *gin.Context) {

	userId, _ := strconv.Atoi(ctx.Param("userId"))

	response := response.NewSuccess()

	var userHasRoleIds []int

	if userId > 0 {
		user := (&service.UserService{}).GetUserByUserId(userId)

		user.Admin = user.UserId == 1

		dept := (&service.DeptService{}).GetDeptByDeptId(user.DeptId)

		roles := (&service.RoleService{}).GetRoleListByUserId(user.UserId)
		for _, role := range roles {
			userHasRoleIds = append(userHasRoleIds, role.RoleId)
		}

		response.SetData("user", dto.AuthUserInfoResponse{
			UserDetailResponse: user,
			Dept:               dept,
			Roles:              roles,
		})
	}

	roles, _ := (&service.RoleService{}).GetRoleList(dto.RoleListRequest{}, false)
	if userId != 1 {
		roles = utils.Filter(roles, func(role dto.RoleListResponse) bool {
			return role.RoleId != 1
		})
		// 设置角色选中标识，如果角色在用户所拥有的角色列表中设置标识为true
		for key, role := range roles {
			if utils.Contains(userHasRoleIds, role.RoleId) {
				roles[key].Flag = true
			}
		}
	}
	response.SetData("roles", roles)

	response.Json(ctx)
}

// 用户授权角色
func (*UserController) AddAuthRole(ctx *gin.Context) {

	var param dto.AddUserAuthRoleRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	roleIds, err := utils.StringToIntSlice(param.RoleIds, ",")
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := (&service.UserService{}).AddAuthRole(param.UserId, roleIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

	response.NewSuccess().Json(ctx)
}

// 个人信息
func (*UserController) GetProfile(ctx *gin.Context) {

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
func (*UserController) UpdateProfile(ctx *gin.Context) {

	var param dto.UpdateProfileRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateProfileValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if err := (&service.UserService{}).UpdateUser(dto.SaveUser{
		UserId:      loginUser.UserId,
		NickName:    param.NickName,
		Email:       param.Email,
		Phonenumber: param.Phonenumber,
		Sex:         param.Sex,
	}, nil, nil); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 修改个人密码
func (*UserController) UserProfileUpdatePwd(ctx *gin.Context) {

	var param dto.UserProfileUpdatePwdRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UserProfileUpdatePwdValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	user := (&service.UserService{}).GetUserByUserId(loginUser.UserId)
	if !password.Verify(user.Password, param.OldPassword) {
		response.NewError().SetMsg("旧密码输入错误").Json(ctx)
		return
	}

	if err := (&service.UserService{}).UpdateUser(dto.SaveUser{
		UserId:   user.UserId,
		Password: password.Generate(param.NewPassword),
	}, nil, nil); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}
