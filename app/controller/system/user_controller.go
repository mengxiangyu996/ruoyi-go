package systemcontroller

import (
	"io/ioutil"
	"os"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/app/token"
	"ruoyi-go/app/validator"
	"ruoyi-go/common/password"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/common/upload"
	"ruoyi-go/common/utils"
	"ruoyi-go/config"
	"ruoyi-go/framework/response"
	"strconv"
	"strings"
	"time"

	"gitee.com/hanshuangjianke/go-excel/excel"
	"github.com/gin-gonic/gin"
	excelize "github.com/xuri/excelize/v2"
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

	users, total := (&service.UserService{}).GetUserList(param, loginUser.UserId, true)

	for key, user := range users {
		users[key].Dept.DeptName = user.DeptName
		users[key].Dept.Leader = user.Leader
	}

	response.NewSuccess().SetPageData(users, total).Json(ctx)
}

// 用户详情
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
func (*UserController) Create(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_INSERT)

	var param dto.CreateUserRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateUserValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if user := (&service.UserService{}).GetUserByUsername(param.UserName); user.UserId > 0 {
		response.NewError().SetMsg("新增用户" + param.UserName + "失败，用户名已存在").Json(ctx)
		return
	}

	if param.Email != "" {
		if user := (&service.UserService{}).GetUserByEmail(param.Email); user.UserId > 0 {
			response.NewError().SetMsg("新增用户" + param.UserName + "失败，邮箱已存在").Json(ctx)
			return
		}
	}

	if param.Phonenumber != "" {
		if user := (&service.UserService{}).GetUserByPhonenumber(param.Phonenumber); user.UserId > 0 {
			response.NewError().SetMsg("新增用户" + param.UserName + "失败，手机号已存在").Json(ctx)
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

	response.NewSuccess().Json(ctx)
}

// 更新用户
func (*UserController) Update(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

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
			response.NewError().SetMsg("修改用户" + param.UserName + "失败，邮箱已存在").Json(ctx)
			return
		}
	}

	if param.Phonenumber != "" {
		if user := (&service.UserService{}).GetUserByPhonenumber(param.Phonenumber); user.UserId > 0 && user.UserId != param.UserId {
			response.NewError().SetMsg("修改用户" + param.UserName + "失败，手机号已存在").Json(ctx)
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

	response.NewSuccess().Json(ctx)
}

// 删除用户
func (*UserController) Remove(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_DELETE)

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

	if err = (&service.UserService{}).DeleteUser(userIds); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更改用户状态
func (*UserController) ChangeStatus(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

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

	response.NewSuccess().Json(ctx)
}

// 重置用户密码
func (*UserController) ResetPwd(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

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

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

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

	response.NewSuccess().Json(ctx)
}

// 导入用户模板
func (*UserController) ImportTemplate(ctx *gin.Context) {

	list := make([]dto.UserImportRequest, 0)

	list = append(list, dto.UserImportRequest{
		DeptId:      1,
		UserName:    "example",
		NickName:    "模板",
		Email:       "example@example.com",
		Phonenumber: "12345678901",
		Sex:         "1",
		Status:      "0",
	})

	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	excel.DownLoadExcel("user_template_"+time.Now().Format("20060102150405"), ctx.Writer, file)
}

// 导入用户数据
func (*UserController) ImportData(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_IMPORT)

	file, err := ctx.FormFile("file")
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	fileName := config.Data.Ruoyi.UploadPath + file.Filename

	// 临时保存文件
	ctx.SaveUploadedFile(file, fileName)
	defer os.Remove(fileName)

	// 是否更新已经存在的用户数据
	updateSupport, _ := strconv.ParseBool(ctx.Query("updateSupport"))

	excelFile, err := excelize.OpenFile(fileName)
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	list := make([]dto.UserImportRequest, 0)

	if err = excel.ImportExcel(excelFile, &list, 0, 1); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if len(list) <= 0 {
		response.NewError().SetMsg("导入用户数据不能为空").Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	var successNum, failNum int
	var failMsg []string

	for _, item := range list {
		user := (&service.UserService{}).GetUserByUsername(item.UserName)

		// 插入新用户
		if user.UserId <= 0 {
			if err = validator.ImportUserValidator(dto.CreateUserRequest{
				DeptId:      item.DeptId,
				UserName:    item.UserName,
				NickName:    item.NickName,
				Email:       item.Email,
				Phonenumber: item.Phonenumber,
				Sex:         item.Sex,
				Status:      item.Status,
			}); err != nil {
				failNum = failNum + 1
				failMsg = append(failMsg, strconv.Itoa(failNum)+"、账号 "+item.UserName+" 新增失败："+err.Error())
				continue
			}
			if err = (&service.UserService{}).CreateUser(dto.SaveUser{
				DeptId:      item.DeptId,
				UserName:    item.UserName,
				NickName:    item.NickName,
				Email:       item.Email,
				Phonenumber: item.Phonenumber,
				Sex:         item.Sex,
				Password:    password.Generate((&service.ConfigService{}).GetConfigByConfigKey("sys.user.initPassword").ConfigValue),
				Status:      item.Status,
				CreateBy:    loginUser.UserName,
			}, nil, nil); err != nil {
				failNum = failNum + 1
				failMsg = append(failMsg, strconv.Itoa(failNum)+"、账号 "+item.UserName+" 新增失败："+err.Error())
				continue
			}
			successNum = successNum + 1
			continue
		} else if updateSupport {
			if err = validator.UpdateUserValidator(dto.UpdateUserRequest{
				UserId:      user.UserId,
				DeptId:      item.DeptId,
				NickName:    item.NickName,
				Email:       item.Email,
				Phonenumber: item.Phonenumber,
				Sex:         item.Sex,
				Status:      item.Status,
			}); err != nil {
				failNum = failNum + 1
				failMsg = append(failMsg, strconv.Itoa(failNum)+"、账号 "+item.UserName+" 更新失败："+err.Error())
				continue
			}
			// 更新已经存在的用户
			if err = (&service.UserService{}).UpdateUser(dto.SaveUser{
				UserId:      user.UserId,
				DeptId:      item.DeptId,
				NickName:    item.NickName,
				Email:       item.Email,
				Phonenumber: item.Phonenumber,
				Sex:         item.Sex,
				Status:      item.Status,
				UpdateBy:    loginUser.UserName,
			}, nil, nil); err != nil {
				failNum = failNum + 1
				failMsg = append(failMsg, strconv.Itoa(failNum)+"、账号 "+item.UserName+" 更新失败："+err.Error())
				continue
			}
			successNum = successNum + 1
			// successMsg = append(successMsg, strconv.Itoa(successNum)+"、账号 "+item.UserName+" 更新成功")
			continue
		} else {
			failNum = failNum + 1
			failMsg = append(failMsg, strconv.Itoa(failNum)+"、账号 "+item.UserName+" 已存在")
		}
	}

	if failNum > 0 {
		response.NewError().SetMsg("导入失败，共 " + strconv.Itoa(failNum) + " 条数据错误，错误如下：" + strings.Join(failMsg, "<br/>")).Json(ctx)
		return
	}

	response.NewSuccess().SetMsg("导入成功，共 " + strconv.Itoa(successNum) + " 条数据").Json(ctx)
}

// 导出用户数据
func (*UserController) Export(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_EXPORT)

	var param dto.UserListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	list := make([]dto.UserExportResponse, 0)

	users, _ := (&service.UserService{}).GetUserList(param, loginUser.UserId, false)
	for _, user := range users {

		loginDate := user.LoginDate.Format("2006-01-02 15:04:05")
		if user.LoginDate.IsZero() {
			loginDate = ""
		}

		list = append(list, dto.UserExportResponse{
			UserId:      user.UserId,
			UserName:    user.UserName,
			NickName:    user.NickName,
			Email:       user.Email,
			Phonenumber: user.Phonenumber,
			Sex:         user.Sex,
			Status:      user.Status,
			LoginIp:     user.LoginIp,
			LoginDate:   loginDate,
			DeptName:    user.DeptName,
			DeptLeader:  user.Leader,
		})
	}

	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	excel.DownLoadExcel("user_"+time.Now().Format("20060102150405"), ctx.Writer, file)
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

// 上传头像
func (*UserController) UserProfileUpdateAvatar(ctx *gin.Context) {

	fileHeader, err := ctx.FormFile("avatarfile")
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	fileResult, err := upload.New(
		upload.SetLimitType([]string{
			"image/jpeg",
			"image/png",
			"image/svg+xml",
		}),
	).SetFile(&upload.File{
		FileName:    fileHeader.Filename,
		FileType:    fileHeader.Header.Get("Content-Type"),
		FileHeader:  fileHeader.Header,
		FileContent: fileContent,
	}).Save()
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	imgUrl := "/" + fileResult.UrlPath + fileResult.FileName

	if err = (&service.UserService{}).UpdateUser(dto.SaveUser{
		UserId: loginUser.UserId,
		Avatar: imgUrl,
	}, nil, nil); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().SetData("imgUrl", imgUrl).Json(ctx)
}
