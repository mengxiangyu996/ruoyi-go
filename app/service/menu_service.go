package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/dal"
	"strings"
)

type MenuService struct{}

// 菜单列表
func (s *MenuService) GetMenuList(param dto.MenuListRequest) []dto.MenuListResponse {

	menus := make([]dto.MenuListResponse, 0)

	query := dal.Gorm.Model(model.SysMenu{}).Order("sys_menu.parent_id, sys_menu.order_num, sys_menu.menu_id")

	if param.MenuName != "" {
		query.Where("menu_name LIKE ?", "%"+param.MenuName+"%")
	}

	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}

	query.Find(&menus)

	return menus
}

// 根据用户id查询菜单权限perms
func (s *MenuService) GetPermsByUserId(userId int) []string {

	perms := make([]string, 0)

	// 超级管理员拥有所有权限
	if userId == 1 {
		perms = append(perms, "*:*:*")
	} else {
		dal.Gorm.Model(model.SysMenu{}).
			Joins("JOIN sys_role_menu ON sys_menu.menu_id = sys_role_menu.menu_id").
			Joins("JOIN sys_role ON sys_role_menu.role_id = sys_role.role_id").
			Joins("JOIN sys_user_role ON sys_role.role_id = sys_user_role.role_id").
			Where("sys_user_role.user_id = ? AND sys_menu.status = ?", userId, constant.NORMAL_STATUS).
			Pluck("sys_menu.perms", &perms)
	}

	return perms
}

// 菜单下拉树列表
func (s *MenuService) Menuselect() []dto.MenuSeleteTree {

	menus := make([]dto.MenuSeleteTree, 0)

	dal.Gorm.Model(model.SysMenu{}).Order("order_num, menu_id").
		Select("menu_id as id", "menu_name as label", "parent_id").
		Where("status = ?", constant.NORMAL_STATUS).
		Find(&menus)

	return menus
}

// 菜单下拉列表转树形结构
func (s *MenuService) MenuSeleteToTree(menus []dto.MenuSeleteTree, parentId int) []dto.MenuSeleteTree {

	tree := make([]dto.MenuSeleteTree, 0)

	for _, menu := range menus {
		if menu.ParentId == parentId {
			tree = append(tree, dto.MenuSeleteTree{
				Id:       menu.Id,
				Label:    menu.Label,
				ParentId: menu.ParentId,
				Children: s.MenuSeleteToTree(menus, menu.Id),
			})
		}
	}

	return tree
}

// 根据用户id查询拥有的菜单权限（M-目录；C-菜单；F-按钮）
func (s *MenuService) GetMenuMCListByUserId(userId int) []dto.MenuListResponse {

	menus := make([]dto.MenuListResponse, 0)

	query := dal.Gorm.Model(model.SysMenu{}).
		Select("sys_menu.*").
		Order("sys_menu.parent_id, sys_menu.order_num").
		Joins("JOIN sys_role_menu ON sys_menu.menu_id = sys_role_menu.menu_id").
		Joins("JOIN sys_role ON sys_role_menu.role_id = sys_role.role_id").
		Joins("JOIN sys_user_role ON sys_role.role_id = sys_user_role.role_id").
		Where("sys_menu.status = ? AND sys_role.status = ? AND sys_menu.menu_type IN ?", constant.NORMAL_STATUS, constant.NORMAL_STATUS, []string{"M", "C"})

	if userId > 1 {
		query = query.Where("sys_user_role.user_id = ?", userId)
	}

	query.Find(&menus)

	return menus
}

// 菜单权限列表转树形结构
func (s *MenuService) MenusToTree(menus []dto.MenuListResponse, parentId int) []dto.MenuListTreeResponse {

	tree := make([]dto.MenuListTreeResponse, 0)

	for _, menu := range menus {
		if menu.ParentId == parentId {
			tree = append(tree, dto.MenuListTreeResponse{
				MenuListResponse: menu,
				Children:         s.MenusToTree(menus, menu.MenuId),
			})
		}
	}

	return tree
}

// 构建前端路由所需要的菜单
func (s *MenuService) BuildRouterMenus(menus []dto.MenuListTreeResponse) []dto.MenuMetaTreeResponse {

	routers := make([]dto.MenuMetaTreeResponse, 0)

	for _, menu := range menus {
		router := dto.MenuMetaTreeResponse{
			Name:      s.GetRouteName(menu),
			Path:      s.GetRoutePath(menu),
			Component: s.GetComponent(menu),
			Hidden:    menu.Visible == "1",
			Meta: dto.MenuMetaResponse{
				Title:   menu.MenuName,
				Icon:    menu.Icon,
				NoCache: menu.IsCache == 1,
			},
		}

		if len(menu.Children) > 0 && menu.MenuType == constant.MENU_TYPE_DIRECTORY {
			router.AlwaysShow = true
			router.Redirect = "noRedirect"
			router.Children = s.BuildRouterMenus(menu.Children)
		} else if s.IsMenuFrame(menu) {
			children := dto.MenuMetaTreeResponse{
				Path:      menu.Path,
				Component: menu.Component,
				Name:      s.GetRouteNameOrDefault(menu.RouteName, menu.Path),
				Meta: dto.MenuMetaResponse{
					Title:   menu.MenuName,
					Icon:    menu.Icon,
					NoCache: menu.IsCache == 1,
				},
				Query: menu.Query,
			}
			router.Children = append(router.Children, children)
		} else if menu.ParentId == 0 && s.IsInnerLink(menu) {
			router.Meta = dto.MenuMetaResponse{
				Title: menu.MenuName,
				Icon:  menu.Icon,
			}
			router.Path = "/"
			children := dto.MenuMetaTreeResponse{
				Path:      s.InnerLinkReplaceEach(menu.Path),
				Component: constant.INNER_LINK_COMPONENT,
				Name:      s.GetRouteNameOrDefault(menu.RouteName, menu.Path),
				Meta: dto.MenuMetaResponse{
					Title: menu.MenuName,
					Icon:  menu.Icon,
					Link:  menu.Path,
				},
			}
			router.Children = append(router.Children, children)
		}

		routers = append(routers, router)
	}

	return routers
}

// 获取路由名称
func (s *MenuService) GetRouteName(menu dto.MenuListTreeResponse) string {

	if s.IsMenuFrame(menu) {
		return ""
	}

	return s.GetRouteNameOrDefault(menu.RouteName, menu.Path)
}

// 获取路由名称，如没有配置路由名称则取路由地址
func (s *MenuService) GetRouteNameOrDefault(name, path string) string {

	if name == "" {
		name = path
	}

	return strings.ToUpper(string(name[0])) + name[1:]
}

// 获取路由地址
func (s *MenuService) GetRoutePath(menu dto.MenuListTreeResponse) string {

	routePath := menu.Path

	// 内链打开外网方式
	if menu.ParentId != 0 && !s.IsInnerLink(menu) {
		routePath = s.InnerLinkReplaceEach(routePath)
	}

	// 非外链并且是一级目录（类型为目录）
	if menu.ParentId == 0 && menu.MenuType == constant.MENU_TYPE_DIRECTORY && menu.IsFrame == constant.IS_MENU_INNER_LINK {
		routePath = "/" + routePath
	} else if s.IsMenuFrame(menu) {
		// 非外链并且是一级目录（类型为菜单）
		routePath = "/"
	}

	return routePath
}

// 获取组件信息
func (s *MenuService) GetComponent(menu dto.MenuListTreeResponse) string {

	component := constant.LAYOUT_COMPONENT

	if menu.Component != "" && !s.IsMenuFrame(menu) {
		component = menu.Component
	} else if menu.Component == "" && menu.ParentId != 0 && s.IsInnerLink(menu) {
		component = constant.INNER_LINK_COMPONENT
	} else if menu.Component == "" && s.IsParentView(menu) {
		component = constant.PARENT_VIEW_COMPONENT
	}

	return component
}

// 是否为菜单内部跳转
func (s *MenuService) IsMenuFrame(menu dto.MenuListTreeResponse) bool {
	return menu.ParentId == 0 && constant.MENU_TYPE_MENU == menu.MenuType && menu.IsFrame == constant.IS_MENU_INNER_LINK
}

// 是否为内链组件
func (s *MenuService) IsInnerLink(menu dto.MenuListTreeResponse) bool {
	return menu.IsFrame == constant.IS_MENU_INNER_LINK && strings.HasPrefix(menu.Path, "http")
}

// 是否为parent_view组件
func (s *MenuService) IsParentView(menu dto.MenuListTreeResponse) bool {
	return menu.ParentId != 0 && menu.MenuType == constant.MENU_TYPE_DIRECTORY
}

// 内链域名特殊字符替换
func (s *MenuService) InnerLinkReplaceEach(path string) string {
	// 去掉 http:// 和 https://
	path = strings.ReplaceAll(path, "http://", "")
	path = strings.ReplaceAll(path, "https://", "")
	path = strings.ReplaceAll(path, "www.", "")

	// 将 . 替换为 /
	path = strings.ReplaceAll(path, ".", "/")

	// 将 : 替换为 /
	path = strings.ReplaceAll(path, ":", "/")

	return path
}
