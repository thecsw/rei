package rei

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// Sha256 hashes given string to sha256.
func Sha256(what []byte) string {
	ans := sha256.Sum256(what)
	return hex.EncodeToString(ans[:])
}

// Btao encodes bytes to base64 string.
func Btao(what []byte) string {
	return base64.StdEncoding.EncodeToString(what)
}

// AtobMust decodes base64 string to bytes, panics on error.
func AtobMust(what string) []byte {
	res, err := base64.StdEncoding.DecodeString(what)
	Try(err)
	return res
}

// Atob decodes base64 string to bytes and an error if any.
func Atob(what string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(what)
}
