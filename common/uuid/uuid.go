package uuid

import (
	"crypto/rand"
	"fmt"
)

func New() (string, error) {

	uuid := make([]byte, 16)

	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	// 设置版本号为 4 (0000 0100)
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // 0100 0000
	// 设置变体为 10 (10xxxxxx)
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // 1000 0000

	// 返回格式化后的UUID
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
