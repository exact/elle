package secure

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	digestSalt    int    = 32     // 256-bit salt (x8)
	digestSize    int    = 32     // Length of the generated hash (256 bits) 10k - 1mil
	digestPrefix  string = "elle" // pbkdf2, hmac, sha256 hashing for security
	digestVersion string = "v1"
)

// Hash takes the FNV-1a hash of a string
func Hash(data string) string {
	var hash uint64 = 14695981039346656037
	for _, b := range []byte(data) {
		hash ^= uint64(b)
		hash *= 1099511628211
	}
	return strconv.FormatUint(hash, 10)
}

// Digest hashes the password using pbkdf2 and returns a secure hash
func Digest(password string, iterations int) string {
	//hashMutex.Lock()
	//defer hashMutex.Unlock()

	salt := make([]byte, digestSalt)
	rand.Read(salt) // result in a combined format of "$prefix$iterations$hash$salt$version"
	return fmt.Sprintf(
		"$%s$%d$%s$%s$%v",
		digestPrefix,
		iterations,
		base64.StdEncoding.EncodeToString(pbkdf2(password, salt, iterations)),
		base64.StdEncoding.EncodeToString(salt),
		digestVersion,
	)
}

// VerifyDigest verifies that the password matches the stored hash using constant-time comparison
func VerifyDigest(storedHash, password string) (bool, error) {
	//hashMutex.Lock()
	//defer hashMutex.Unlock()

	parts := strings.Split(storedHash, "$")

	if len(parts) != 6 || parts[1] != digestPrefix {
		return false, fmt.Errorf("invalid hash format")
	}

	iterations := 0
	if _, err := fmt.Sscanf(parts[2], "%d", &iterations); err != nil {
		return false, fmt.Errorf("invalid iterations")
	}

	salt, err := base64.StdEncoding.DecodeString(parts[4]) // Decode the salt and hash from base64
	if err != nil {
		return false, fmt.Errorf("invalid salt encoding")
	}

	storedHashBytes, err := base64.StdEncoding.DecodeString(parts[3])
	if err != nil {
		return false, fmt.Errorf("invalid hash encoding")
	}

	return subtle.ConstantTimeCompare(pbkdf2(password, salt, iterations), storedHashBytes) == 1, nil
}

// main key derivation function (HMAC-SHA-256)
func pbkdf2(password string, salt []byte, iterations int) []byte {
	var result []byte
	for i := 1; i <= int(math.Ceil(float64(digestSize)/float64(64))); i++ { // SHA-256 output length is 32 bytes
		block := append(salt, byte(i>>24), byte(i>>16), byte(i>>8), byte(i)) // Create the block input: salt + block index

		hmacResult := hmac.New(sha256.New, []byte(password)) // Apply HMAC for the current block
		hmacResult.Write(block)
		hash := hmacResult.Sum(nil)

		previousResult := hash // Apply the iterations
		for j := 1; j < iterations; j++ {
			hmacResult.Reset()
			hmacResult.Write(previousResult)
			previousResult = hmacResult.Sum(nil)
			for k := range hash {
				hash[k] ^= previousResult[k]
			}
		}

		result = append(result, hash...) // Append the result to the final hash output
	}

	return result // Return the first bytes of the final result[:digestSize]
}
