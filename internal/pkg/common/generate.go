package common

import (
	"crypto/rand"
	"encoding/hex"
	"math"
)

// GenerateRandomHexStr 生成随机十六进制字符串
func GenerateRandomHexStr(len uint) string {
	byteLen := int(math.Ceil(float64(len) / 2))
	bytes := make([]byte, byteLen)
	_, _ = rand.Read(bytes)
	hexStr := hex.EncodeToString(bytes)
	return hexStr[:len]
}

// GenerateRandomHexColor 生成随机十六进制颜色字符串
func GenerateRandomHexColor() string {
	return "#" + GenerateRandomHexStr(8)
}
