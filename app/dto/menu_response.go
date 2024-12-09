package dto

import "ruoyi-go/framework/datetime"

// 菜单权限列表
type MenuListResponse struct {
	MenuId     int               `json:"menuId"`
	MenuName   string            `json:"menuName"`
	ParentId   int               `json:"parentId"`
	OrderNum   int               `json:"orderNum"`
	Path       string            `json:"path"`
	Component  string            `json:"component"`
	Query      string            `json:"query"`
	RouteName  string            `json:"routeName"`
	IsFrame    int               `json:"isFrame"`
	IsCache    int               `json:"isCache"`
	MenuType   string            `json:"menuType"`
	Visible    string            `json:"visible"`
	Perms      string            `json:"perms"`
	Icon       string            `json:"icon"`
	Status     string            `json:"status"`
	CreateTime datetime.Datetime `json:"createTime"`
}

// 菜单权限树形结构
type MenuListTreeResponse struct {
	MenuListResponse
	Children []MenuListTreeResponse `json:"children" gorm:"-"`
}

// 菜单列表树形结构
type MenuMetaTreeResponse struct {
	Name       string                 `json:"name"`
	Path       string                 `json:"path"`
	Redirect   string                 `json:"redirect"`
	Component  string                 `json:"component"`
	Hidden     bool                   `json:"hidden"`
	AlwaysShow bool                   `json:"alwaysShow"`
	Meta       MenuMetaResponse       `json:"meta"`
	Children   []MenuMetaTreeResponse `json:"children"`
	Query      string                 `json:"query"`
	MenuName   string                 `json:"-"`
	ParentId   int                    `json:"-"`
	MenuType   string                 `json:"-"`
	IsFrame    int                    `json:"-"`
	RouteName  string                 `json:"-"`
}

type MenuMetaResponse struct {
	Title   string `json:"title"`
	Icon    string `json:"icon"`
	Link    string `json:"link"`
	NoCache bool   `json:"noCache"`
}

// 菜单详情
type MenuDetailResponse struct {
	MenuId    int    `json:"menuId"`
	MenuName  string `json:"menuName"`
	ParentId  int    `json:"parentId"`
	OrderNum  int    `json:"orderNum"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Query     string `json:"query"`
	RouteName string `json:"routeName"`
	IsFrame   int    `json:"isFrame"`
	IsCache   int    `json:"isCache"`
	MenuType  string `json:"menuType"`
	Visible   string `json:"visible"`
	Perms     string `json:"perms"`
	Icon      string `json:"icon"`
	Status    string `json:"status"`
}
