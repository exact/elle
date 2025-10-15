package secure

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
)

// Hex is similar to `token_hex` in Python and returns a cryptographically random, hex encoded string of length n*2.
func Hex(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// Number returns a cryptographically secure random number between min and max.
func Number(min, max int64) (int64, error) {
	if min > max {
		return 0, fmt.Errorf("invalid range: min(%d) > max (%d)", min, max)
	}

	num, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		return 0, err
	}

	return num.Int64() + min, nil
}
