package common

import "math/rand"

// RandomIntInRange 生成范围内的随机整数
func RandomIntInRange(min int, max int) int {
	if max < min {
		t := max
		max = min
		min = t
	}
	minR := max - min + 1

	return rand.Intn(minR) + min
}
