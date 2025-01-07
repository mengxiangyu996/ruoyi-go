package validator

import (
	"errors"
	"ruoyi-go/app/dto"
	"ruoyi-go/common/types/regexp"
	"ruoyi-go/common/utils"
)

// 更新个人资料验证
func UpdateProfileValidator(param dto.UpdateProfileRequest) error {

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
func UserProfileUpdatePwdValidator(param dto.UserProfileUpdatePwdRequest) error {

	if param.OldPassword == "" {
		return errors.New("请输入旧密码")
	}

	if param.NewPassword == "" {
		return errors.New("请输入新密码")
	}

	return nil
}

// 添加用户验证
func CreateUserValidator(param dto.CreateUserRequest) error {

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

// 更新用户验证
func UpdateUserValidator(param dto.UpdateUserRequest) error {

	if param.UserId <= 0 {
		return errors.New("参数错误")
	}

	if param.NickName == "" {
		return errors.New("请输入用户昵称")
	}

	if param.Phonenumber != "" && !utils.CheckRegex(regexp.PHONE, param.Phonenumber) {
		return errors.New("手机号码格式错误")
	}

	if param.Email != "" && !utils.CheckRegex(regexp.EMAIL, param.Email) {
		return errors.New("邮箱账号格式错误")
	}

	return nil
}

// 删除用户验证
func RemoveUserValidator(userIds []int, authUserId int) error {

	if utils.Contains(userIds, 1) {
		return errors.New("超级管理员无法删除")
	}

	if utils.Contains(userIds, authUserId) {
		return errors.New("当前用户无法删除")
	}

	return nil
}

// 修改用户状态验证
func ChangeUserStatusValidator(param dto.UpdateUserRequest) error {

	if param.UserId <= 0 {
		return errors.New("参数错误")
	}

	if param.Status == "" {
		return errors.New("请选择状态")
	}

	return nil
}

// 重置用户密码验证
func ResetUserPwdValidator(param dto.UpdateUserRequest) error {

	if param.UserId <= 0 {
		return errors.New("参数错误")
	}

	if param.Password == "" {
		return errors.New("请输入用户密码")
	}

	return nil
}

// 导入用户验证
func ImportUserValidator(param dto.CreateUserRequest) error {

	if param.NickName == "" {
		return errors.New("请输入用户昵称")
	}

	if param.UserName == "" {
		return errors.New("请输入用户名称")
	}

	if param.Phonenumber != "" && !utils.CheckRegex(regexp.PHONE, param.Phonenumber) {
		return errors.New("手机号码格式错误")
	}

	if param.Email != "" && !utils.CheckRegex(regexp.EMAIL, param.Email) {
		return errors.New("邮箱账号格式错误")
	}

	return nil
}
