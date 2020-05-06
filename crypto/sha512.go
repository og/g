package gcrypto

import (
	"crypto/sha512"
	"encoding/hex"
)

func SHA512(s string) string {
	if s == "" {
		warningLog("SHA521(s string) s is empty string, There may have been some errors")
	}
	hash := sha512.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
func SaltSHA512(password string, salt string) (hash string) {
	return SHA512(password + salt)
}
func CheckSaltSHA512(password string, salt string, hash string) bool {
	return SHA512(password + salt) == hash
}
