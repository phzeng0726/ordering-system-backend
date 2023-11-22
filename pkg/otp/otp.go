package otp

import "math/rand"

// 生成隨機字串
func GenerateRandomCode(length int) string {
	const digits = "0123456789"

	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = digits[rand.Intn(len(digits))]
	}

	return string(code)
}
