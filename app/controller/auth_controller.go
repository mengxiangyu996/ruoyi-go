package controller

import (
	"context"
	"ruoyi-go/app/controller/validator"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/app/token"
	"ruoyi-go/common/captcha"
	"ruoyi-go/common/password"
	"ruoyi-go/common/types/constant"
	rediskey "ruoyi-go/common/types/redis-key"
	statusCode "ruoyi-go/common/types/status-code"
	"ruoyi-go/config"
	"ruoyi-go/framework/dal"
	"ruoyi-go/framework/datetime"
	"ruoyi-go/framework/response"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

// 获取验证码
func (*AuthController) CaptchaImage(ctx *gin.Context) {

	captcha := captcha.NewCaptcha()

	id, b64s := captcha.Generate()

	b64s = strings.Replace(b64s, "data:image/png;base64,", "", 1)

	response.NewSuccess().SetData("uuid", id).SetData("img", b64s).SetData("captchaEnabled", true).Json(ctx)
}

// 登录
func (*AuthController) Login(ctx *gin.Context) {

	var param dto.LoginRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetCode(statusCode.BadRequest).SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.LoginValidator(&param); err != nil {
		response.NewError().SetCode(statusCode.BadRequest).SetMsg(err.Error()).Json(ctx)
		return
	}

	if !captcha.NewCaptcha().Verify(param.Uuid, param.Code) {
		response.NewError().SetMsg("验证码错误").Json(ctx)
		return
	}

	user := (&service.UserService{}).GetUserByUsername(param.Username)
	if user.UserId <= 0 || user.Status != constant.NORMAL_STATUS {
		response.NewError().SetMsg("用户不存在或被禁用").Json(ctx)
		return
	}

	// 登陆密码错误次数超过限制，锁定账号10分钟
	redisKey := rediskey.LoginPasswordErrorKey + param.Username
	count, _ := dal.Redis.Get(context.Background(), redisKey).Int()
	if count >= config.Data.User.Password.MaxRetryCount {
		response.NewError().SetMsg("密码错误次数超过限制，请" + strconv.Itoa(config.Data.User.Password.LockTime) + "分钟后重试").Json(ctx)
		return
	}

	if !password.Verify(user.Password, param.Password) {
		// 密码错误次数加1，并设置缓存过期时间为锁定时间
		dal.Redis.Set(context.Background(), redisKey, count+1, time.Minute*time.Duration(config.Data.User.Password.LockTime))
		response.NewError().SetMsg("密码错误").Json(ctx)
		return
	}

	// 登录成功，删除错误次数
	dal.Redis.Del(context.Background(), redisKey)

	token, err := token.GetClaims().GenerateToken(user)
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	// 更新登录的ip和时间
	(&service.UserService{}).UpdateUser(dto.SaveUser{
		UserId:    user.UserId,
		LoginIP:   ctx.ClientIP(),
		LoginDate: datetime.Datetime{Time: time.Now()},
	}, nil, nil)

	response.NewSuccess().SetData("token", token).Json(ctx)
}

// 获取授权信息
func (*AuthController) GetInfo(ctx *gin.Context) {

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

	roleKeys := (&service.RoleService{}).GetRoleKeysByUserId(loginUser.UserId)

	perms := (&service.MenuService{}).GetPermsByUserId(loginUser.UserId)

	response.NewSuccess().SetData("user", data).SetData("roles", roleKeys).SetData("permissions", perms).Json(ctx)
}

// 获取授权路由
func (*AuthController) GetRouters(ctx *gin.Context) {

	loginUser, _ := token.GetLoginUser(ctx)

	menus := (&service.MenuService{}).GetMenuMCListByUserId(loginUser.UserId)

	tree := (&service.MenuService{}).MenusToTree(menus, 0)

	routers := (&service.MenuService{}).BuildRouterMenus(tree)

	response.NewSuccess().SetData("data", routers).Json(ctx)
}

// 退出登录
func (*AuthController) Logout(ctx *gin.Context) {

	token.DeleteToken(ctx)

	response.NewSuccess().Json(ctx)
}
