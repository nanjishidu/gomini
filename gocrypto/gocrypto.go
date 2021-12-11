package gocrypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
)

func Md5(s string) string {
	ss, _ := Crypto("md5", s)
	return ss
}
func Sha1(s string) string {
	ss, _ := Crypto("sha1", s)
	return ss
}
func Sha256(s string) string {
	ss, _ := Crypto("sha256", s)
	return ss
}
func Sha512(s string) string {
	ss, _ := Crypto("sha512", s)
	return ss
}
func Crypto(ctype string, s string) (string, error) {
	var h hash.Hash
	switch ctype {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	default:
		return "", errors.New("unsupported types")
	}
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)), nil
}

func HmacSha1(ctype, s, key string) string {
	ss, _ := Hmac(ctype, s, key)
	return ss
}

func HmacSha256(ctype, s, key string) string {
	ss, _ := Hmac(ctype, s, key)
	return ss
}

func Hmac(ctype, s, key string) (string, error) {
	var h hash.Hash
	switch ctype {
	case "sha1":
		h = hmac.New(sha1.New, []byte(key))
	case "sha256":
		h = hmac.New(sha256.New, []byte(key))
	}
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)), nil
}
