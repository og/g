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
type SaltEncode struct {
	Password string
	Salt string
}
func SaltSHA512(encode SaltEncode) (hash string) {
	return SHA512(encode.Password + encode.Salt)
}
type SaltDecode struct {
	Password string
	Salt string
	Hash string
}
func CheckSaltSHA512(decode SaltDecode) bool {
	return SHA512(decode.Password + decode.Salt) == decode.Hash
}
