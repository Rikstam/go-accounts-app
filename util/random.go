package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt function takes 2 int64 numbers: min and max as input.
// And it returns a random int64 number between min and max.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString  declares a new string builder object sb, get the total number of characters in the alphabet and assign it to k.
//
//Then we will use a simple for loop to generate n random characters.
//We use rand.Intn(k) to get a random position from 0 to k-1, and take the corresponding character at that position in the alphabet, assign it to variable c.
//
//We call sb.WriteByte() to write that character c to the string builder.
//Finally we just return sb.ToString() to the caller.

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
