package common

import (
	"testing"
)

func TestGenerateRandomHexStr(t *testing.T) {
	t.Run("检测随机字符串生成", func(t *testing.T) {
		for i := uint(1); i < 100; i++ {
			str := GenerateRandomHexStr(i)

			if len(str) != int(i) {
				t.Errorf(`GenerateRandomHexStr(%d) = %s; 期望长度为 %d, 实际长度为 %d`, i, str, i, len(str))
			}
		}
	})
}

func TestGenerateRandomHexColor(t *testing.T) {
	t.Run("检测随机十六进制颜色生成", func(t *testing.T) {
		for i := uint(1); i < 50; i++ {
			str := GenerateRandomHexColor()

			if len(str) != 9 && len(str) != 7 {
				t.Errorf(`GenerateRandomHexStr(%d) = %s; 期望长度为 (9|7), 实际长度为 %d`, i, str, len(str))
			}
		}
	})
}
