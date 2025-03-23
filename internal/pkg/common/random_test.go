package common

import (
	"testing"
)

func TestRandomInt(t *testing.T) {
	t.Run("测试生成随机数整数", func(t *testing.T) {
		for i := 1; i < 1000; i++ {
			value := RandomInt(i)

			if value < 0 {
				t.Errorf("生成的整数小于0")
			}

			if value >= i {
				t.Errorf("生成的整数大于等于n")
			}
		}
	})
}

func TestRandomIntInRange(t *testing.T) {
	t.Run("测试随机范围整数的生成", func(t *testing.T) {
		minR, maxR := 10, 50

		for i := 1; i <= 1000; i++ {
			result := RandomIntInRange(minR, maxR)
			if result < minR || result > maxR {
				t.Errorf("生成的随机数 %d 不在 [%d, %d] 范围内", result, minR, maxR)
			}
		}
	})
}
