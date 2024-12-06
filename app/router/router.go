package router

import (
	"ruoyi-go/app/controller"
	monitorcontroller "ruoyi-go/app/controller/monitor"
	systemcontroller "ruoyi-go/app/controller/system"
	"ruoyi-go/app/middleware"

	"github.com/gin-gonic/gin"
)

// 注册路由
func Register(server *gin.Engine) {

	// 未授权并且不需要记录操作日志
	api := server.Group("/api")
	{
		api.GET("/auth/captchaImage", (&controller.AuthController{}).CaptchaImage)                       // 获取验证码
		api.POST("/auth/login", middleware.LogininforMiddleware(), (&controller.AuthController{}).Login) // 登录
		api.POST("/auth/logout", (&controller.AuthController{}).Logout)                                  // 退出登录

	}

	// 已授权并且不需要记录操作日志
	api = server.Group("/api", middleware.AuthMiddleware())
	{
		api.GET("/auth/getAuthInfo", (&controller.AuthController{}).GetAuthInfo) // 获取授权信息
		api.GET("/auth/getRoutes", (&controller.AuthController{}).GetRoutes)     // 获取授权路由

		api.GET("/dictData/getDictDataByDictType/:dictType", (&controller.DictDataController{}).GetDictDataByDictType) // 根据字典类型查询字典数据

		api.GET("/config/getConfigValueByKey/:configKey", (&controller.ConfigController{}).GetConfigValueByKey) // 根据配置key获取配置值

		// 系统管理
		api.GET("/system/user/deptTree", (&systemcontroller.UserController{}).DeptTree)                           // 获取用户管理中的部门树
		api.GET("/system/user/list", (&systemcontroller.UserController{}).List)                                   // 获取用户列表
		api.GET("/system/user/profile", (&systemcontroller.UserController{}).GetUserProfile)                      // 获取用户信息
		api.PUT("/system/user/profile", (&systemcontroller.UserController{}).UpdateUserProfile)                   // 更新个人信息
		api.PUT("/system/user/profile/updatePwd", (&systemcontroller.UserController{}).UpdateUserProfilePassword) // 更新个人密码

		api.GET("/system/role/list", (&systemcontroller.RoleController{}).List)     // 获取角色列表
		api.GET("/system/menu/list", (&systemcontroller.MenuController{}).List)     // 获取菜单列表
		api.GET("/system/dept/list", (&systemcontroller.DeptController{}).List)     // 获取部门列表
		api.GET("/system/post/list", (&systemcontroller.PostController{}).List)     // 获取岗位列表
		api.GET("/system/dict/list", (&systemcontroller.DictController{}).List)     // 获取字典类型列表
		api.GET("/system/config/list", (&systemcontroller.ConfigController{}).List) // 获取参数列表

		// 日志管理
		api.GET("/monitor/operlog/list", (&monitorcontroller.OperlogController{}).List)       // 获取操作日志列表
		api.GET("/monitor/logininfor/list", (&monitorcontroller.LogininforController{}).List) // 获取登录日志列表
	}

	// 已授权并且需要记录操作日志
	api = server.Group("/api", middleware.AuthMiddleware(), middleware.OperLogMiddleware())
	{
	}
}
