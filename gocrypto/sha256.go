package gocrypto

import (
"encoding/hex"
	"crypto/sha256"
)

func Sha256(key string) string {
	h := sha256.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(nil))
}

