package invalid

import "crypto/sha256"

// Bypass attempt: using SHA-256 instead of MD5, but still not bcrypt/argon2
// ruleid: go.security.no-weak-password-hashing
func hashPassword(pwd string) []byte {
	h := sha256.Sum256([]byte(pwd))
	return h[:]
}
