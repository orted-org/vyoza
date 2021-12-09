package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

const alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const symbols = `~!@#$%^&*()_-+={[}]|\:;"'<,>.?/`
const alphaNumeric = numbers + alphabets
const alphaNumericSymbol = numbers + alphabets + symbols

func RandomInt(min, max int64) int {
	return int(min + rand.Int63n(max-min+1))
}

func RandomString(size int) string {
	var sb strings.Builder
	k := len(alphabets)
	for i := 0; i < size; i++ {
		sb.WriteByte(alphabets[rand.Intn(k)])
	}
	return sb.String()
}

func RandomBool() bool {
	num := rand.Int63n(2)
	if num == 0 {
		return false
	} else {
		return true
	}
}

func RandomAlphaNumeric(size int) string {
	var sb strings.Builder
	k := len(alphaNumeric)
	for i := 0; i < size; i++ {
		sb.WriteByte(alphaNumeric[rand.Intn(k)])
	}
	return sb.String()
}

func RandomAlphaNumericSymbolString(size int) string {
	var sb strings.Builder
	k := len(alphaNumericSymbol)
	for i := 0; i < size; i++ {
		sb.WriteByte(alphaNumericSymbol[rand.Intn(k)])
	}
	return sb.String()
}
