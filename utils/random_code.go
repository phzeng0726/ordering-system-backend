package utils

import "math/rand"

// 生成指定长度的随机数字串
func GenerateRandomCode(length int) string {
	const digits = "0123456789"

	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = digits[rand.Intn(len(digits))]
	}

	return string(code)
}
