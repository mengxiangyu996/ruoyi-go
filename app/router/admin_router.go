package router

import (
	"ruoyi-go/app/controller"
	monitorcontroller "ruoyi-go/app/controller/monitor"
	systemcontroller "ruoyi-go/app/controller/system"
	"ruoyi-go/app/middleware"
	"ruoyi-go/common/types/constant"

	"github.com/gin-gonic/gin"
)

// 后台路由组
func RegisterAdminGroupApi(api *gin.RouterGroup) {

	api.Use(middleware.Cors())                                                                  // 跨域中间件
	api.GET("/captchaImage", (&controller.AuthController{}).CaptchaImage)                       // 获取验证码
	api.POST("/register", (&controller.AuthController{}).Register)                              // 注册
	api.POST("/login", middleware.LogininforMiddleware(), (&controller.AuthController{}).Login) // 登录
	api.POST("/logout", (&controller.AuthController{}).Logout)                                  // 退出登录

	// 启用鉴权中间件，下面的路由都需要鉴权中间件验证通过后才可访问
	api.Use(middleware.AuthMiddleware())

	api.GET("/getInfo", (&controller.AuthController{}).GetInfo)       // 获取用户信息
	api.GET("/getRouters", (&controller.AuthController{}).GetRouters) // 获取路由信息

	api.GET("/system/user/profile", (&systemcontroller.UserController{}).GetProfile)                      // 个人信息
	api.PUT("/system/user/profile", (&systemcontroller.UserController{}).UpdateProfile)                   // 修改用户
	api.PUT("/system/user/profile/updatePwd", (&systemcontroller.UserController{}).UserProfileUpdatePwd)  // 重置密码
	api.POST("/system/user/profile/avatar", (&systemcontroller.UserController{}).UserProfileUpdateAvatar) // 更新头像

	api.GET("/system/user/deptTree", middleware.HasPerm("system:user:list"), (&systemcontroller.UserController{}).DeptTree)          // 获取部门树列表
	api.GET("/system/user/list", middleware.HasPerm("system:user:list"), (&systemcontroller.UserController{}).List)                  // 获取用户列表
	api.GET("/system/user/", middleware.HasPerm("system:user:query"), (&systemcontroller.UserController{}).Detail)                   // 根据用户编号获取详细信息
	api.GET("/system/user/:userId", middleware.HasPerm("system:user:query"), (&systemcontroller.UserController{}).Detail)            // 根据用户编号获取详细信息
	api.GET("/system/user/authRole/:userId", middleware.HasPerm("system:user:query"), (&systemcontroller.UserController{}).AuthRole) // 根据用户编号获取详细信息

	api.POST("/system/user", middleware.HasPerm("system:user:add"), middleware.OperLogMiddleware("新增用户", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.UserController{}).Create)
	api.PUT("/system/user", middleware.HasPerm("system:user:edit"), middleware.OperLogMiddleware("更新用户", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.UserController{}).Update)
	api.DELETE("/system/user/:userIds", middleware.HasPerm("system:user:remove"), middleware.OperLogMiddleware("删除用户", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.UserController{}).Remove)
	api.PUT("/system/user/changeStatus", middleware.HasPerm("system:user:edit"), middleware.OperLogMiddleware("修改用户状态", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.UserController{}).ChangeStatus)
	api.PUT("/system/user/resetPwd", middleware.HasPerm("system:user:edit"), middleware.OperLogMiddleware("修改用户密码", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.UserController{}).ResetPwd)
	api.PUT("/system/user/authRole", middleware.HasPerm("system:user:edit"), middleware.OperLogMiddleware("用户授权角色", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.UserController{}).AddAuthRole)
	api.POST("/system/user/export", middleware.HasPerm("system:user:export"), middleware.OperLogMiddleware("导出用户", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&systemcontroller.UserController{}).Export)
	api.POST("/system/user/importData", middleware.HasPerm("system:user:import"), middleware.OperLogMiddleware("导入用户", constant.REQUEST_BUSINESS_TYPE_IMPORT), (&systemcontroller.UserController{}).ImportData)
	api.POST("/system/user/importTemplate", middleware.OperLogMiddleware("导入用户模板", constant.REQUEST_BUSINESS_TYPE_IMPORT), (&systemcontroller.UserController{}).ImportTemplate)

	api.GET("/system/role/list", middleware.HasPerm("system:role:list"), (&systemcontroller.RoleController{}).List)                                            // 获取角色列表
	api.GET("/system/role/:roleId", middleware.HasPerm("system:role:query"), (&systemcontroller.RoleController{}).Detail)                                      // 获取角色详情
	api.GET("/system/role/deptTree/:roleId", middleware.HasPerm("system:role:query"), (&systemcontroller.RoleController{}).DeptTree)                           // 获取部门树
	api.GET("/system/role/authUser/allocatedList", middleware.HasPerm("system:role:list"), (&systemcontroller.RoleController{}).RoleAuthUserAllocatedList)     // 查询已分配用户角色列表
	api.GET("/system/role/authUser/unallocatedList", middleware.HasPerm("system:role:list"), (&systemcontroller.RoleController{}).RoleAuthUserUnallocatedList) // 查询未分配用户角色列表

	api.POST("/system/role", middleware.HasPerm("system:role:add"), middleware.OperLogMiddleware("新增角色", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.RoleController{}).Create)
	api.PUT("/system/role", middleware.HasPerm("system:role:edit"), middleware.OperLogMiddleware("更新角色", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.RoleController{}).Update)
	api.DELETE("/system/role/:roleIds", middleware.HasPerm("system:role:remove"), middleware.OperLogMiddleware("删除角色", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.RoleController{}).Remove)
	api.PUT("/system/role/changeStatus", middleware.HasPerm("system:role:edit"), middleware.OperLogMiddleware("修改角色状态", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.RoleController{}).ChangeStatus)
	api.PUT("/system/role/dataScope", middleware.HasPerm("system:role:edit"), middleware.OperLogMiddleware("分配数据权限", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.RoleController{}).DataScope)
	api.PUT("/system/role/authUser/selectAll", middleware.HasPerm("system:role:edit"), middleware.OperLogMiddleware("批量选择用户授权", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.RoleController{}).RoleAuthUserSelectAll)
	api.PUT("/system/role/authUser/cancel", middleware.HasPerm("system:role:edit"), middleware.OperLogMiddleware("取消授权用户", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.RoleController{}).RoleAuthUserCancel)
	api.PUT("/system/role/authUser/cancelAll", middleware.HasPerm("system:role:edit"), middleware.OperLogMiddleware("批量取消授权用户", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.RoleController{}).RoleAuthUserCancelAll)
	api.POST("/system/role/export", middleware.HasPerm("system:role:export"), middleware.OperLogMiddleware("导出角色", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&systemcontroller.RoleController{}).Export)

	api.GET("/system/menu/list", middleware.HasPerm("system:menu:list"), (&systemcontroller.MenuController{}).List)       // 获取菜单列表
	api.GET("/system/menu/treeselect", (&systemcontroller.MenuController{}).Treeselect)                                   // 获取菜单下拉树列表
	api.GET("/system/menu/roleMenuTreeselect/:roleId", (&systemcontroller.MenuController{}).RoleMenuTreeselect)           // 加载对应角色菜单列表树
	api.GET("/system/menu/:menuId", middleware.HasPerm("system:menu:query"), (&systemcontroller.MenuController{}).Detail) // 获取菜单详情

	api.POST("/system/menu", middleware.HasPerm("system:menu:add"), middleware.OperLogMiddleware("新增菜单", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.MenuController{}).Create)
	api.PUT("/system/menu", middleware.HasPerm("system:menu:edit"), middleware.OperLogMiddleware("更新菜单", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.MenuController{}).Update)
	api.DELETE("/system/menu/:menuId", middleware.HasPerm("system:menu:remove"), middleware.OperLogMiddleware("删除菜单", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.MenuController{}).Remove)

	api.GET("/system/dept/list", middleware.HasPerm("system:dept:list"), (&systemcontroller.DeptController{}).List)                        // 获取部门列表
	api.GET("/system/dept/list/exclude/:deptId", middleware.HasPerm("system:dept:list"), (&systemcontroller.DeptController{}).ListExclude) // 查询部门列表（排除节点）
	api.GET("/system/dept/:deptId", middleware.HasPerm("system:dept:query"), (&systemcontroller.DeptController{}).Detail)                  // 获取部门详情

	api.POST("/system/dept", middleware.HasPerm("system:dept:add"), middleware.OperLogMiddleware("新增部门", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.DeptController{}).Create)
	api.PUT("/system/dept", middleware.HasPerm("system:dept:edit"), middleware.OperLogMiddleware("更新部门", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.DeptController{}).Update)
	api.DELETE("/system/dept/:deptId", middleware.HasPerm("system:dept:remove"), middleware.OperLogMiddleware("删除部门", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.DeptController{}).Remove)

	api.GET("/system/post/list", middleware.HasPerm("system:post:list"), (&systemcontroller.PostController{}).List)       // 获取岗位列表
	api.GET("/system/post/:postId", middleware.HasPerm("system:post:query"), (&systemcontroller.PostController{}).Detail) // 获取岗位详情

	api.POST("/system/post", middleware.HasPerm("system:post:add"), middleware.OperLogMiddleware("新增岗位", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.PostController{}).Create)
	api.PUT("/system/post", middleware.HasPerm("system:post:edit"), middleware.OperLogMiddleware("更新岗位", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.PostController{}).Update)
	api.DELETE("/system/post/:postIds", middleware.HasPerm("system:post:remove"), middleware.OperLogMiddleware("删除岗位", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.PostController{}).Remove)
	api.POST("/system/post/export", middleware.HasPerm("system:post:export"), middleware.OperLogMiddleware("导出岗位", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&systemcontroller.PostController{}).Export)

	api.GET("/system/dict/list", middleware.HasPerm("system:dict:list"), (&systemcontroller.DictTypeController{}).List)            // 获取字典类型列表
	api.GET("/system/dict/type/:dictId", middleware.HasPerm("system:dict:query"), (&systemcontroller.DictTypeController{}).Detail) // 获取字典类型详情
	api.GET("/system/dict/type/optionselect", (&systemcontroller.DictTypeController{}).Optionselect)                               // 获取字典选择框列表

	api.POST("/system/dict/type", middleware.HasPerm("system:dict:add"), middleware.OperLogMiddleware("新增字典类型", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.DictTypeController{}).Create)
	api.PUT("/system/dict/type", middleware.HasPerm("system:dict:edit"), middleware.OperLogMiddleware("更新字典类型", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.DictTypeController{}).Update)
	api.DELETE("/system/dict/type/:dictIds", middleware.HasPerm("system:dict:remove"), middleware.OperLogMiddleware("删除字典类型", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.DictTypeController{}).Remove)
	api.POST("/system/dict/type/export", middleware.HasPerm("system:dict:export"), middleware.OperLogMiddleware("导出字典类型", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&systemcontroller.DictTypeController{}).Export)
	api.DELETE("/system/dict/type/refreshCache", middleware.HasPerm("system:dict:remove"), middleware.OperLogMiddleware("刷新字典类型缓存", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.DictTypeController{}).RefreshCache)

	api.GET("/system/dict/data/list", middleware.HasPerm("system:dict:list"), (&systemcontroller.DictDataController{}).List)         // 获取字典数据列表
	api.GET("/system/dict/data/:dictCode", middleware.HasPerm("system:dict:query"), (&systemcontroller.DictDataController{}).Detail) // 获取字典数据详情
	api.GET("/system/dict/data/type/:dictType", (&systemcontroller.DictDataController{}).Type)                                       // 根据字典类型查询字典数据

	api.POST("/system/dict/data", middleware.HasPerm("system:dict:add"), middleware.OperLogMiddleware("新增字典数据", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.DictDataController{}).Create)
	api.PUT("/system/dict/data", middleware.HasPerm("system:dict:edit"), middleware.OperLogMiddleware("更新字典数据", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.DictDataController{}).Update)
	api.DELETE("/system/dict/data/:dictCodes", middleware.HasPerm("system:dict:remove"), middleware.OperLogMiddleware("删除字典数据", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.DictDataController{}).Remove)
	api.POST("/system/dict/data/export", middleware.HasPerm("system:dict:export"), middleware.OperLogMiddleware("导出字典数据", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&systemcontroller.DictDataController{}).Export)

	api.GET("/system/config/list", middleware.HasPerm("system:config:list"), (&systemcontroller.ConfigController{}).List)         // 获取参数配置列表
	api.GET("/system/config/:configId", middleware.HasPerm("system:config:query"), (&systemcontroller.ConfigController{}).Detail) // 获取参数配置详情
	api.GET("/system/config/configKey/:configKey", (&systemcontroller.ConfigController{}).ConfigKey)                              // 根据参数键名查询参数值

	api.POST("/system/config", middleware.HasPerm("system:config:add"), middleware.OperLogMiddleware("新增参数配置", constant.REQUEST_BUSINESS_TYPE_INSERT), (&systemcontroller.ConfigController{}).Create)
	api.PUT("/system/config", middleware.HasPerm("system:config:edit"), middleware.OperLogMiddleware("更新参数配置", constant.REQUEST_BUSINESS_TYPE_UPDATE), (&systemcontroller.ConfigController{}).Update)
	api.DELETE("/system/config/:configIds", middleware.HasPerm("system:config:remove"), middleware.OperLogMiddleware("删除参数配置", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.ConfigController{}).Remove)
	api.POST("/system/config/export", middleware.HasPerm("system:config:export"), middleware.OperLogMiddleware("导出参数配置", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&systemcontroller.ConfigController{}).Export)
	api.DELETE("/system/config/refreshCache", middleware.HasPerm("system:config:remove"), middleware.OperLogMiddleware("刷新参数配置缓存", constant.REQUEST_BUSINESS_TYPE_DELETE), (&systemcontroller.ConfigController{}).RefreshCache)

	api.GET("/monitor/logininfor/list", middleware.HasPerm("monitor:operlog:list"), (&monitorcontroller.LogininforController{}).List) // 获取登录日志列表

	api.DELETE("/monitor/logininfor/:infoIds", middleware.HasPerm("monitor:logininfor:remove"), middleware.OperLogMiddleware("删除登录日志", constant.REQUEST_BUSINESS_TYPE_DELETE), (&monitorcontroller.LogininforController{}).Remove)
	api.DELETE("/monitor/logininfor/clean", middleware.HasPerm("monitor:logininfor:remove"), middleware.OperLogMiddleware("清空登录日志", constant.REQUEST_BUSINESS_TYPE_DELETE), (&monitorcontroller.LogininforController{}).Clean)
	api.GET("/monitor/logininfor/unlock/:userName", middleware.HasPerm("monitor:logininfor:unlock"), middleware.OperLogMiddleware("账户解锁", constant.REQUEST_BUSINESS_TYPE_DELETE), (&monitorcontroller.LogininforController{}).Unlock)
	api.POST("/monitor/logininfor/export", middleware.HasPerm("monitor:logininfor:export"), middleware.OperLogMiddleware("导出登录日志", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&monitorcontroller.LogininforController{}).Export)

	api.GET("/monitor/operlog/list", middleware.HasPerm("monitor:logininfor:list"), (&monitorcontroller.OperlogController{}).List) // 获取操作日志列表

	api.DELETE("/monitor/operlog/:operIds", middleware.HasPerm("monitor:operlog:remove"), middleware.OperLogMiddleware("删除操作日志", constant.REQUEST_BUSINESS_TYPE_DELETE), (&monitorcontroller.OperlogController{}).Remove)
	api.DELETE("/monitor/operlog/clean", middleware.HasPerm("monitor:operlog:remove"), middleware.OperLogMiddleware("清空操作日志", constant.REQUEST_BUSINESS_TYPE_DELETE), (&monitorcontroller.OperlogController{}).Clean)
	api.POST("/monitor/operlog/export", middleware.HasPerm("monitor:operlog:export"), middleware.OperLogMiddleware("导出操作日志", constant.REQUEST_BUSINESS_TYPE_EXPORT), (&monitorcontroller.OperlogController{}).Export)
}
