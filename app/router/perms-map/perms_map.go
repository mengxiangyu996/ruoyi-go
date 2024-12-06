package permsmap

// 路由权限映射
var permsMap = map[string]string{

	"GET:/api/monitor/logininfor/list":             "monitor:logininfor:list",
	"POST:/api/monitor/logininfor/export":          "monitor:logininfor:export",
	"DELETE:/api/monitor/logininfor/:infoIds":      "monitor:logininfor:remove",
	"DELETE:/api/monitor/logininfor/clean":         "monitor:logininfor:remove",
	"GET:/api/monitor/logininfor/unlock/:userName": "monitor:logininfor:unlock",

	"GET:/api/monitor/operlog/list":        "monitor:operlog:list",
	"POST:/api/monitor/operlog/export":     "monitor:operlog:export",
	"DELETE:/api/monitor/operlog/:operIds": "monitor:operlog:remove",
	"DELETE:/api/monitor/operlog/clean":    "monitor:operlog:remove",

	"GET:/api/system/config/list":            "system:config:list",
	"POST:/api/system/config/export":         "system:config:export",
	"GET:/api/system/config/:configId":       "system:config:query",
	"POST:/api/system/config":                "system:config:add",
	"PUT:/api/system/config":                 "system:config:edit",
	"DELETE:/api/system/config/:configIds":   "system:config:remove",
	"DELETE:/api/system/config/refreshCache": "system:config:remove",

	"GET:/api/system/dept/list":                 "system:dept:list",
	"GET:/api/system/dept/list/exclude/:deptId": "system:dept:list",
	"GET:/api/system/dept/:deptId":              "system:dept:query",
	"POST:/api/system/dept":                     "system:dept:add",
	"PUT:/api/system/dept":                      "system:dept:edit",
	"DELETE:/api/system/dept":                   "system:dept:remove",

	"GET:/api/system/dict/data/list":      "system:dict:list",
	"POST:/api/system/dict/data/export":   "system:dict:export",
	"GET:/api/system/dict/data/:dictCode": "system:dict:query",
	"POST:/api/system/dict":               "system:dict:add",
	"PUT:/api/system/dict":                "system:dict:edit",
	"DELETE:/api/system/dict/:dictCodes":  "system:dict:remove",

	"GET:/api/system/dict/type/list":            "system:dict:list",
	"POST:/api/system/dict/type/export":         "system:dict:export",
	"GET:/api/system/dict/type/:dictId":         "system:dict:query",
	"POST:/api/system/dict/type":                "system:dict:add",
	"PUT:/api/system/dict/type":                 "system:dict:edit",
	"DELETE:/api/system/dict/type/:dictIds":     "system:dict:remove",
	"DELETE:/api/system/dict/type/refreshCache": "system:dict:remove",

	"GET:/api/system/menu/list":       "system:menu:list",
	"GET:/api/system/menu/:menuId":    "system:menu:query",
	"POST:/api/system/menu":           "system:menu:add",
	"PUT:/api/system/menu":            "system:menu:edit",
	"DELETE:/api/system/menu/:menuId": "system:menu:remove",

	"GET:/api/system/post/list":        "system:post:list",
	"POST:/api/system/post/export":     "system:post:export",
	"GET:/api/system/post/:postId":     "system:post:query",
	"POST:/api/system/post":            "system:post:add",
	"PUT:/api/system/post":             "system:post:edit",
	"DELETE:/api/system/post/:postIds": "system:post:remove",

	"GET:/api/system/role/list":                     "system:role:list",
	"POST:/api/system/role/export":                  "system:role:export",
	"GET:/api/system/role/:roleId":                  "system:role:query",
	"POST:/api/system/role":                         "system:role:add",
	"PUT:/api/system/role":                          "system:role:edit",
	"PUT:/api/system/role/dataScope":                "system:role:edit",
	"PUT:/api/system/role/changeStatus":             "system:role:edit",
	"DELETE:/api/system/role/:roleIds":              "system:role:remove",
	"GET:/api/system/role/optionselect":             "system:role:query",
	"GET:/api/system/role/authUser/allocatedList":   "system:role:list",
	"GET:/api/system/role/authUser/unallocatedList": "system:role:list",
	"PUT:/api/system/role/authUser/cancel":          "system:role:edit",
	"PUT:/api/system/role/authUser/cancelAll":       "system:role:edit",
	"PUT:/api/system/role/authUser/selectAll":       "system:role:edit",
	"GET:/api/system/role/deptTree/:roleId":         "system:role:query",

	"GET:/api/system/user/list":             "system:user:list",
	"POST:/api/system/user/export":          "system:user:export",
	"POST:/api/system/user/importData":      "system:user:import",
	"GET:/api/system/user/":                 "system:user:query",
	"GET:/api/system/user/:userId":          "system:user:query",
	"POST:/api/system/user":                 "system:user:add",
	"PUT:/api/system/user":                  "system:user:edit",
	"DELETE:/api/system/user/:userIds":      "system:user:remove",
	"PUT:/api/system/user/resetPwd":         "system:user:resetPwd",
	"PUT:/api/system/user/changeStatus":     "system:user:edit",
	"GET:/api/system/user/authRole/:userId": "system:user:query",
	"PUT:/api/system/user/authRole":         "system:user:edit",
	"GET:/api/system/user/deptTree":         "system:user:list",
}

// 判断api权限是否存在
//
// 为了实现@PreAuthorize("@ss.hasPermi('system:user:list')")注解
func HasPermi(api string) string {

	if perm, ok := permsMap[api]; ok {
		return perm
	}

	return ""
}
