package password

import (
	"golang.org/x/crypto/bcrypt"
)

// 生成密码
func Generate(password string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hash)
}

// 验证密码
func Verify(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
