package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/dal"
	"strings"
)

type MenuService struct{}

// 新增菜单
func (s *MenuService) CreateMenu(param dto.SaveMenu) error {

	return dal.Gorm.Model(model.SysMenu{}).Create(&model.SysMenu{
		MenuName:  param.MenuName,
		ParentId:  param.ParentId,
		OrderNum:  param.OrderNum,
		Path:      param.Path,
		Component: param.Component,
		Query:     param.Query,
		RouteName: param.RouteName,
		IsFrame:   param.IsFrame,
		IsCache:   param.IsCache,
		MenuType:  param.MenuType,
		Visible:   param.Visible,
		Perms:     param.Perms,
		Icon:      param.Icon,
		Status:    param.Status,
		Remark:    param.Remark,
		CreateBy:  param.CreateBy,
	}).Error
}

// 更新菜单
func (s *MenuService) UpdateMenu(param dto.SaveMenu) error {

	return dal.Gorm.Model(model.SysMenu{}).Where("menu_id = ?", param.MenuId).Updates(&model.SysMenu{
		MenuName:  param.MenuName,
		ParentId:  param.ParentId,
		OrderNum:  param.OrderNum,
		Path:      param.Path,
		Component: param.Component,
		Query:     param.Query,
		RouteName: param.RouteName,
		IsFrame:   param.IsFrame,
		IsCache:   param.IsCache,
		MenuType:  param.MenuType,
		Visible:   param.Visible,
		Perms:     param.Perms,
		Icon:      param.Icon,
		Status:    param.Status,
		UpdateBy:  param.UpdateBy,
		Remark:    param.Remark,
	}).Error
}

// 删除菜单
func (s *MenuService) DeleteMenu(menuId int) error {
	return dal.Gorm.Model(model.SysMenu{}).Where("menu_id = ?", menuId).Delete(&model.SysMenu{}).Error
}

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

// 根据菜单id查询菜单
func (s *MenuService) GetMenuByMenuId(menuId int) dto.MenuDetailResponse {

	var menu dto.MenuDetailResponse

	dal.Gorm.Model(model.SysMenu{}).Where("menu_id = ?", menuId).Last(&menu)

	return menu
}

// 根据菜单名称查询菜单
func (s *MenuService) GetMenuByMenuName(menuName string) dto.MenuDetailResponse {

	var menu dto.MenuDetailResponse

	dal.Gorm.Model(model.SysMenu{}).Where("menu_name = ?", menuName).Last(&menu)

	return menu
}

// 查询是否存在下级菜单
func (s *MenuService) MenuHasChildren(menuId int) bool {

	var count int64

	dal.Gorm.Model(model.SysMenu{}).Where("parent_id = ?", menuId).Count(&count)

	return count > 0
}

// 查询菜单是否已分配到权限
func (s *MenuService) MenuExistRole(menuId int) bool {

	var count int64

	dal.Gorm.Model(model.SysRoleMenu{}).Where("menu_id = ?", menuId).Count(&count)

	return count > 0
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

// 根据角色id查询拥有的菜单id集合
func (s *MenuService) GetMenuIdsByRoleId(roleId int) []int {

	menuIds := make([]int, 0)

	dal.Gorm.Model(model.SysRoleMenu{}).
		Joins("JOIN sys_menu ON sys_menu.menu_id = sys_role_menu.menu_id").
		Where("sys_menu.status = ? AND sys_role_menu.role_id = ?", constant.NORMAL_STATUS, roleId).
		Pluck("sys_menu.menu_id", &menuIds)

	return menuIds
}

// 菜单下拉树列表
func (s *MenuService) MenuSelect() []dto.SeleteTree {

	menus := make([]dto.SeleteTree, 0)

	dal.Gorm.Model(model.SysMenu{}).Order("order_num, menu_id").
		Select("menu_id as id", "menu_name as label", "parent_id").
		Where("status = ?", constant.NORMAL_STATUS).
		Find(&menus)

	return menus
}

// 菜单下拉列表转树形结构
func (s *MenuService) MenuSeleteToTree(menus []dto.SeleteTree, parentId int) []dto.SeleteTree {

	tree := make([]dto.SeleteTree, 0)

	for _, menu := range menus {
		if menu.ParentId == parentId {
			tree = append(tree, dto.SeleteTree{
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
		Distinct("sys_menu.*").
		Order("sys_menu.parent_id, sys_menu.order_num").
		Joins("LEFT JOIN sys_role_menu ON sys_menu.menu_id = sys_role_menu.menu_id").
		Joins("LEFT JOIN sys_role ON sys_role_menu.role_id = sys_role.role_id").
		Joins("LEFT JOIN sys_user_role ON sys_role.role_id = sys_user_role.role_id").
		Where("sys_menu.status = ? AND sys_menu.menu_type IN ?", constant.NORMAL_STATUS, []string{"M", "C"})

	if userId > 1 {
		query = query.Where("sys_user_role.user_id = ? AND sys_role.status = ?", userId, constant.NORMAL_STATUS)
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
				Path:      s.InnerLinkReplacePach(menu.Path),
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
		routePath = s.InnerLinkReplacePach(routePath)
	}

	// 非外链并且是一级目录（类型为目录）
	if menu.ParentId == 0 && menu.MenuType == constant.MENU_TYPE_DIRECTORY && menu.IsFrame == constant.MENU_NO_FRAME {
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
	return menu.ParentId == 0 && constant.MENU_TYPE_MENU == menu.MenuType && menu.IsFrame == constant.MENU_NO_FRAME
}

// 是否为内链组件
func (s *MenuService) IsInnerLink(menu dto.MenuListTreeResponse) bool {
	return menu.IsFrame == constant.MENU_NO_FRAME && strings.HasPrefix(menu.Path, "http")
}

// 是否为parent_view组件
func (s *MenuService) IsParentView(menu dto.MenuListTreeResponse) bool {
	return menu.ParentId != 0 && menu.MenuType == constant.MENU_TYPE_DIRECTORY
}

// 内链域名特殊字符替换
func (s *MenuService) InnerLinkReplacePach(path string) string {
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
