package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int {
	return int(min + rand.Int63n(max - min + 1))
}

func RandomString(size int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < size; i++ {
		sb.WriteByte(alphabet[rand.Intn(k)])
	}
	return sb.String()
}

func RandomBool() bool {
	num := rand.Int63n(2);
	if num ==0 {
		return false
	} else {
		return true;
	}
}