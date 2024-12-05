package validator

import (
	"errors"
	"ruoyi-go/app/dto"
)

// 登录验证
func LoginValidator(param *dto.LoginRequest) error {

	if param.Username == "" {
		return errors.New("用户名不能为空")
	}

	if param.Password == "" {
		return errors.New("密码不能为空")
	}

	if param.Code == "" {
		return errors.New("验证码不能为空")
	}

	return nil
}
