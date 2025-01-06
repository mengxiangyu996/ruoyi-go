package constant

// 平台内系统用户的唯一标志
const SYS_USER = "SYS_USER"

// 正常状态
const NORMAL_STATUS = "0"

// 异常状态
const EXCEPTION_STATUS = "1"

// 是否系统默认（是）
const IS_DEFAULT_YES = "Y"

// 是否系统默认（否）
const IS_DEFAULT_NO = "N"

// 是否菜单外链（是）
const MENU_YES_FRAME = 0

// 是否菜单外链（否）
const MENU_NO_FRAME = 1

// 菜单类型（目录）
const MENU_TYPE_DIRECTORY = "M"

// 菜单类型（菜单）
const MENU_TYPE_MENU = "C"

// 菜单类型（按钮）
const MENU_TYPE_BUTTON = "F"

// Layout组件
const LAYOUT_COMPONENT = "Layout"

// ParentView组件标识
const PARENT_VIEW_COMPONENT = "ParentView"

// InnerLink组件标识
const INNER_LINK_COMPONENT = "InnerLink"

// 请求的操作标题key（操作日志需要）
const REQUEST_TITLE = "businessTitle"

// 请求的操作类型key（操作日志需要）
const REQUEST_BUSINESS_TYPE = "businessType"

// 请求的操作类型（具体类型）（操作日志需要）
const (
	REQUEST_BUSINESS_TYPE_INSERT = 1 // 新增
	REQUEST_BUSINESS_TYPE_UPDATE = 2 // 修改
	REQUEST_BUSINESS_TYPE_DELETE = 3 // 删除
	REQUEST_BUSINESS_TYPE_GRANT  = 4 // 授权
	REQUEST_BUSINESS_TYPE_EXPORT = 5 // 导出
	REQUEST_BUSINESS_TYPE_IMPORT = 6 // 导入
	REQUEST_BUSINESS_TYPE_FORCE  = 7 // 强退
	REQUEST_BUSINESS_TYPE_GENCOD = 8 // 生成代码
	REQUEST_BUSINESS_TYPE_CLEAN  = 9 // 清空数据
)
