package rediskey

import "ruoyi-go/config"

var (
	// 验证码 redis key
	CaptchaCodeKey = config.Data.Ruoyi.Name + ":captcha:code:"

	// 登录账户密码错误次数 redis key
	LoginPasswordErrorKey = config.Data.Ruoyi.Name + ":login:password:error:"

	// 登录用户 redis key
	UserTokenKey = config.Data.Ruoyi.Name + ":user:token:"

	// 防重提交 redis key
	RepeatSubmitKey = config.Data.Ruoyi.Name + ":repeat:submit:"

	// 配置表数据 redis key
	SysConfigKey = config.Data.Ruoyi.Name + ":system:config"

	// 字典表数据 redis key
	SysDictKey = config.Data.Ruoyi.Name + ":system:dict:data"
)
