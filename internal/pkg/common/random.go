package common

import "math/rand"

// RandomInt 生成随机 [0,n) 整数
func RandomInt(n int) int {
	return rand.Intn(n)
}

// RandomIntInRange 生成范围内的随机整数
func RandomIntInRange(min int, max int) int {
	if max < min {
		t := max
		max = min
		min = t
	}
	minR := max - min + 1

	return RandomInt(minR) + min
}
