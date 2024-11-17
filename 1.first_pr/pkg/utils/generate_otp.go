package utils

import "math/rand"

func GenerateOTP() int {
	return rand.Int() % 100000
}
