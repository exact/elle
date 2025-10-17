package random

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

// Hex is similar to `token_hex` in Python and returns a cryptographically random, hex encoded string of length n*2.
func Hex(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// Number returns a cryptographically secure random number between min and max.
func Number(min int64, max int) int {
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(max)-min+1))
	return int(num.Int64() + min)
}
