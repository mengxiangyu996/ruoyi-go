package validator

import (
	"errors"
	"ruoyi-go/app/dto"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/common/utils"
	"strings"
)

// 添加菜单验证
func CreateMenuValidator(param dto.CreateMenuRequest) error {

	if param.MenuName == "" {
		return errors.New("请输入菜单名称")
	}

	if utils.Contains([]string{constant.MENU_TYPE_DIRECTORY, constant.MENU_TYPE_MENU}, param.Path) && param.Path == "" {
		return errors.New("请输入路由地址")
	}

	if param.IsFrame == constant.MENU_YES_FRAME && !strings.HasPrefix(param.Path, "http") {
		return errors.New("新增菜单" + param.MenuName + "失败，地址必须以http(s)://开头")
	}

	return nil
}

// 更新菜单验证
func UpdateMenuValidator(param dto.UpdateMenuRequest) error {

	if param.MenuId <= 0 {
		return errors.New("参数错误")
	}

	if param.MenuName == "" {
		return errors.New("请输入菜单名称")
	}

	if utils.Contains([]string{constant.MENU_TYPE_DIRECTORY, constant.MENU_TYPE_MENU}, param.Path) && param.Path == "" {
		return errors.New("请输入路由地址")
	}

	if param.IsFrame == constant.MENU_YES_FRAME && !strings.HasPrefix(param.Path, "http") {
		return errors.New("修改菜单" + param.MenuName + "失败，地址必须以http(s)://开头")
	}

	if param.MenuId == param.ParentId {
		return errors.New("修改菜单" + param.MenuName + "失败，上级菜单不能选择自己")
	}

	return nil
}
