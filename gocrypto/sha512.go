package gocrypto

import (
	"encoding/hex"
	"crypto/sha512"
)

func Sha512(key string) string {
	h := sha512.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(nil))
}

