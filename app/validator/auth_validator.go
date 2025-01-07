package validator

import (
	"errors"
	"ruoyi-go/app/dto"
)

// 注册验证
func RegisterValidator(param dto.RegisterRequest) error {

	if param.Username == "" {
		return errors.New("用户名不能为空")
	}

	if param.Password == "" {
		return errors.New("密码不能为空")
	}

	if param.ConfirmPassword != param.Password {
		return errors.New("两次密码不一致")
	}

	if len(param.Username) < 2 || len(param.Username) > 20 {
		return errors.New("用户名长度必须在2-20之间")
	}

	if len(param.Password) < 5 || len(param.Password) > 20 {
		return errors.New("密码长度必须在5-20之间")
	}

	return nil
}

// 登录验证
func LoginValidator(param dto.LoginRequest) error {

	if param.Username == "" {
		return errors.New("用户名不能为空")
	}

	if param.Password == "" {
		return errors.New("密码不能为空")
	}

	return nil
}
