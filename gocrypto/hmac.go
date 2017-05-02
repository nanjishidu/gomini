package gocrypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

func Hmac(s string,key string)string{
	mac:=hmac.New(sha1.New,[]byte(key))
	mac.Write([]byte(s))
	return  hex.EncodeToString(mac.Sum(nil))
}