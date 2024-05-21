package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func RandomUserName() string {
	return RandomString(10)
}

func RandomAmount() int64 {
	return RandomInt(100, 1000)
}

func RandomCurrency() string {
	currencies := []string{"BDT", "USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
