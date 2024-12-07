package validator

import (
	"errors"
	"ruoyi-go/app/dto"
	"ruoyi-go/common/types/regexp"
	"ruoyi-go/common/utils"
)

// 更新个人资料验证
func UpdateProfileValidator(param dto.UpdateProfile) error {

	if param.NickName == "" {
		return errors.New("请输入用户昵称")
	}

	if !utils.CheckRegex(regexp.EMAIL, param.Email) {
		return errors.New("邮箱格式错误")
	}

	if !utils.CheckRegex(regexp.PHONE, param.Phonenumber) {
		return errors.New("手机号格式错误")
	}

	return nil
}

// 更新个人密码验证
func UserProfileUpdatePwdValidator(param dto.UserProfileUpdatePwd) error {

	if param.OldPassword == "" {
		return errors.New("请输入旧密码")
	}

	if param.NewPassword == "" {
		return errors.New("请输入新密码")
	}

	return nil
}

// 添加用户验证
func AddUserValidator(param dto.AddUserRequest) error {

	if param.NickName == "" {
		return errors.New("请输入用户昵称")
	}

	if param.UserName == "" {
		return errors.New("请输入用户名称")
	}

	if param.Password == "" {
		return errors.New("请输入用户密码")
	}

	if param.Phonenumber != "" && !utils.CheckRegex(regexp.PHONE, param.Phonenumber) {
		return errors.New("手机号码格式错误")
	}

	if param.Email != "" && !utils.CheckRegex(regexp.EMAIL, param.Email) {
		return errors.New("邮箱账号格式错误")
	}

	return nil
}