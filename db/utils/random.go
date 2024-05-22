package utils

import (
	"math/rand"
	"strconv"
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
	return RandomInt(1000, 10000)
}

func RandomCurrency() string {
	currencies := []string{"BDT", "USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomAccountNumber() string {
	// only contains digits
	var accountNo string
	for i := 0; i < 16; i++ {
		if i == 4 {
			accountNo += "-"
			continue
		}
		digit := rand.Intn(10)
		accountNo += strconv.Itoa(digit)
	}
	return accountNo
}
