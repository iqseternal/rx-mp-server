package common

import (
	"testing"
)

func TestRandomIntInRange(t *testing.T) {
	t.Run("测试随机范围整数的生成", func(t *testing.T) {
		minR, maxR := 10, 50

		for i := 0; i < 1000; i++ {
			result := RandomIntInRange(minR, maxR)
			if result < minR || result > maxR {
				t.Errorf("生成的随机数 %d 不在 [%d, %d] 范围内", result, minR, maxR)
			}
		}
	})
}
