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
		api.GET("/captchaImage", (&controller.AuthController{}).CaptchaImage)                       // 获取验证码
		api.POST("/login", middleware.LogininforMiddleware(), (&controller.AuthController{}).Login) // 登录
		api.POST("/logout", (&controller.AuthController{}).Logout)                                  // 退出登录

	}

	// 已授权并且不需要记录操作日志
	api = server.Group("/api", middleware.AuthMiddleware())
	{
		api.GET("/getInfo", (&controller.AuthController{}).GetInfo)       // 获取用户信息
		api.GET("/getRouters", (&controller.AuthController{}).GetRouters) // 获取路由信息

		// 系统管理
		api.GET("/system/user/profile", (&systemcontroller.UserController{}).GetProfile)                     // 个人信息
		api.PUT("/system/user/profile", (&systemcontroller.UserController{}).UpdateProfile)                  // 修改用户
		api.PUT("/system/user/profile/updatePwd", (&systemcontroller.UserController{}).UserProfileUpdatePwd) // 重置密码

		api.GET("/system/user/deptTree", middleware.HasPerm("system:user:list"), (&systemcontroller.UserController{}).DeptTree)          // 获取部门树列表
		api.GET("/system/user/list", middleware.HasPerm("system:user:list"), (&systemcontroller.UserController{}).List)                  // 获取用户列表
		api.GET("/system/user/", middleware.HasPerm("system:user:query"), (&systemcontroller.UserController{}).Detail)                   // 根据用户编号获取详细信息
		api.GET("/system/user/:userId", middleware.HasPerm("system:user:query"), (&systemcontroller.UserController{}).Detail)            // 根据用户编号获取详细信息
		api.GET("/system/user/authRole/:userId", middleware.HasPerm("system:user:query"), (&systemcontroller.UserController{}).AuthRole) // 根据用户编号获取详细信息

		api.GET("/system/role/list", middleware.HasPerm("system:role:list"), (&systemcontroller.RoleController{}).List)                  // 获取角色列表
		api.GET("/system/role/:roleId", middleware.HasPerm("system:role:query"), (&systemcontroller.RoleController{}).Detail)            // 获取角色详情
		api.GET("/system/role/deptTree/:roleId", middleware.HasPerm("system:role:query"), (&systemcontroller.RoleController{}).DeptTree) // 获取部门树

		api.GET("/system/menu/list", middleware.HasPerm("system:menu:list"), (&systemcontroller.MenuController{}).List) // 获取菜单列表
		api.GET("/system/menu/treeselect", (&systemcontroller.MenuController{}).Treeselect)                             // 获取菜单下拉树列表
		api.GET("/system/menu/roleMenuTreeselect/:roleId", (&systemcontroller.MenuController{}).RoleMenuTreeselect)     // 加载对应角色菜单列表树

		api.GET("/system/dept/list", middleware.HasPerm("system:dept:list"), (&systemcontroller.DeptController{}).List) // 获取部门列表
		api.GET("/system/post/list", middleware.HasPerm("system:post:list"), (&systemcontroller.PostController{}).List) // 获取岗位列表

		api.GET("/system/dict/type/list", middleware.HasPerm("system:dict:list"), (&systemcontroller.DictTypeController{}).List) // 获取字典类型列表
		api.GET("/system/dict/data/type/:dictType", (&systemcontroller.DictDataController{}).Type)                               // 根据字典类型查询字典数据

		api.GET("/system/config/list", middleware.HasPerm("system:config:list"), (&systemcontroller.ConfigController{}).List) // 获取参数配置列表
		api.GET("/system/config/configKey/:configKey", (&systemcontroller.ConfigController{}).ConfigKey)                      // 根据参数键名查询参数值

		// 日志管理
		api.GET("/monitor/operlog/list", middleware.HasPerm("monitor:logininfor:list"), (&monitorcontroller.OperlogController{}).List)    // 获取操作日志列表
		api.GET("/monitor/logininfor/list", middleware.HasPerm("monitor:operlog:list"), (&monitorcontroller.LogininforController{}).List) // 获取登录日志列表
	}

	// 已授权并且需要记录操作日志
	api = server.Group("/api", middleware.AuthMiddleware(), middleware.OperLogMiddleware())
	{
		api.POST("/system/user", middleware.HasPerm("system:user:add"), (&systemcontroller.UserController{}).Create)                    // 新增用户
		api.PUT("/system/user", middleware.HasPerm("system:user:edit"), (&systemcontroller.UserController{}).Update)                    // 更新用户
		api.DELETE("/system/user/:userIds", middleware.HasPerm("system:user:remove"), (&systemcontroller.UserController{}).Remove)      // 删除用户
		api.PUT("/system/user/changeStatus", middleware.HasPerm("system:user:edit"), (&systemcontroller.UserController{}).ChangeStatus) // 修改用户状态
		api.PUT("/system/user/resetPwd", middleware.HasPerm("system:user:edit"), (&systemcontroller.UserController{}).ResetPwd)         // 重置用密码
		api.PUT("/system/user/authRole", middleware.HasPerm("system:user:edit"), (&systemcontroller.UserController{}).AddAuthRole)      // 用户授权角色

		api.POST("/system/role", middleware.HasPerm("system:role:add"), (&systemcontroller.RoleController{}).Create)                    // 新增角色
		api.PUT("/system/role", middleware.HasPerm("system:role:edit"), (&systemcontroller.RoleController{}).Update)                    // 更新角色
		api.DELETE("/system/role/:roleIds", middleware.HasPerm("system:role:remove"), (&systemcontroller.RoleController{}).Remove)      // 删除角色
		api.PUT("/system/role/changeStatus", middleware.HasPerm("system:role:edit"), (&systemcontroller.RoleController{}).ChangeStatus) // 修改角色状态
		api.PUT("/system/role/dataScope", middleware.HasPerm("system:role:edit"), (&systemcontroller.RoleController{}).DataScope)       // 分配数据权限
	}
}
