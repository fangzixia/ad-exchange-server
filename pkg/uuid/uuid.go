package uuid

import (
	"fmt"
	"github.com/google/uuid"
)

// GenerateUUID 生成带连字符的UUID V4（标准格式，推荐默认使用）
func GenerateUUID() string {
	// 生成UUID V4原始对象
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return fmt.Sprintf("生成UUID失败：%w", err.Error())
	}
	return uuidObj.String()
}
