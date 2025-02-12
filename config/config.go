package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	// 项目相关配置
	Ruoyi struct {
		// 名称
		Name string `yaml:"name"`
		// 版本
		Version string `yaml:"version"`
		// 版权年份
		Copyright string `yaml:"copyright"`
		// 域名
		Domain string `yaml:"domain"`
		// 启用SSL
		SSL bool `yaml:"ssl"`
		// 文件上传路径
		UploadPath string `yaml:"uploadPath"`
	} `yaml:"ruoyi"`

	// 开发环境配置
	Server struct {
		// 端口
		Port int `yaml:"port"`
		// 模式，可选值：debug、test、release
		Mode string `yaml:"mode"`
	} `yaml:"server"`

	// 数据库配置
	Mysql struct {
		Host string `yaml:"host"`
		// 端口，默认为3306
		Port int `yaml:"port"`
		// 数据库名称
		Database string `yaml:"database"`
		// 用户名
		Username string `yaml:"username"`
		// 密码
		Password string `yaml:"password"`
		// 编码
		Charset string `yaml:"charset"`
		// 连接池最大连接数
		MaxIdleConns int `yaml:"maxIdleConns"`
		// 连接池最大打开连接数
		MaxOpenConns int `yaml:"maxOpenConns"`
	} `yaml:"mysql"`

	// Redis配置
	Redis struct {
		Host string `yaml:"host"`
		// 端口，默认为6379
		Port int `yaml:"port"`
		// 数据库索引
		Database int `yaml:"database"`
		// 密码
		Password string `yaml:"password"`
	} `yaml:"redis"`

	// Token配置
	Token struct {
		// 令牌自定义标识
		Header string `yaml:"header"`
		// 令牌密钥
		Secret string `yaml:"secret"`
		// 令牌有效期（默认30分钟）
		ExpireTime int `yaml:"expireTime"`
	} `yaml:"token"`

	// 用户配置
	User struct {
		Password struct {
			// 密码最大错误次数
			MaxRetryCount int `yaml:"maxRetryCount"`
			// 密码锁定时间（默认10分钟）
			LockTime int `yaml:"lockTime"`
		} `yaml:"password"`
	} `yaml:"user"`
}

var Data *Config

func init() {

	file, err := os.ReadFile("application.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &Data)
	if err != nil {
		panic(err)
	}
}
