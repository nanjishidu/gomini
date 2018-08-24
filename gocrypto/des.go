// des.go
package gocrypto

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"sync"
)

var (
	syncDesMutex sync.Mutex
	commonDeskey []byte
)

func SetDesKey(key string) error {
	syncDesMutex.Lock()
	defer syncDesMutex.Unlock()
	b := []byte(key)
	if len(b) == 8 || len(b) == 24 {
		commonDeskey = b
		return nil
	}
	return fmt.Errorf("key size is not 8 or 24, but %d", len(b))
}

func DesCBCEncrypt(plaintext []byte, paddingType ...string) (ciphertext []byte,
	err error) {
	block, err := des.NewCipher(commonDeskey)
	if err != nil {
		return nil, err
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroPadding":
			plaintext = ZeroPadding(plaintext, des.BlockSize)
		case "PKCS5Padding":
			plaintext = PKCS5Padding(plaintext, des.BlockSize)
		}
	} else {
		plaintext = PKCS5Padding(plaintext, des.BlockSize)
	}
	ciphertext = make([]byte, des.BlockSize+len(plaintext))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext[des.BlockSize:],
		plaintext)
	return ciphertext, nil
}

func DesCBCDecrypt(ciphertext []byte, paddingType ...string) (plaintext []byte,
	err error) {
	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	block, err := des.NewCipher(commonDeskey)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:des.BlockSize]
	ciphertext = ciphertext[des.BlockSize:]
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(ciphertext, ciphertext)
	if int(ciphertext[len(ciphertext)-1]) > len(ciphertext) {
		return nil, errors.New("des decrypt failed")
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroUnPadding":
			plaintext = ZeroUnPadding(ciphertext)
		case "PKCS5UnPadding":
			plaintext = PKCS5UnPadding(ciphertext)
		}
	} else {
		plaintext = PKCS5UnPadding(ciphertext)
	}
	return plaintext, nil
}

// 3DES加密
func TripleDesCBCEncrypt(plaintext []byte, paddingType ...string) (ciphertext []byte,
	err error) {
	block, err := des.NewTripleDESCipher(commonDeskey)
	if err != nil {
		return nil, err
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroPadding":
			plaintext = ZeroPadding(plaintext, des.BlockSize)
		case "PKCS5Padding":
			plaintext = PKCS5Padding(plaintext, des.BlockSize)
		}
	} else {
		plaintext = PKCS5Padding(plaintext, des.BlockSize)
	}
	ciphertext = make([]byte, des.BlockSize+len(plaintext))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext[des.BlockSize:],
		plaintext)
	return ciphertext, nil
}

// 3DES解密
func TripleDesCBCDecrypt(ciphertext []byte, paddingType ...string) (plaintext []byte,
	err error) {
	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	block, err := des.NewTripleDESCipher(commonDeskey)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:des.BlockSize]
	ciphertext = ciphertext[des.BlockSize:]
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(ciphertext, ciphertext)
	if int(ciphertext[len(ciphertext)-1]) > len(ciphertext) {
		return nil, errors.New("aes decrypt failed")
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroUnPadding":
			plaintext = ZeroUnPadding(ciphertext)
		case "PKCS5UnPadding":
			plaintext = PKCS5UnPadding(ciphertext)
		}
	} else {
		plaintext = PKCS5UnPadding(ciphertext)
	}
	return plaintext, nil
}

func DesCFBEncrypt(plaintext []byte, paddingType ...string) (ciphertext []byte,
	err error) {
	block, err := des.NewCipher(commonDeskey)
	if err != nil {
		return nil, err
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroPadding":
			plaintext = ZeroPadding(plaintext, des.BlockSize)
		case "PKCS5Padding":
			plaintext = PKCS5Padding(plaintext, des.BlockSize)
		}
	} else {
		plaintext = PKCS5Padding(plaintext, des.BlockSize)
	}
	ciphertext = make([]byte, des.BlockSize+len(plaintext))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(ciphertext[des.BlockSize:],
		plaintext)
	return ciphertext, nil
}

func DesCFBDecrypt(ciphertext []byte, paddingType ...string) (plaintext []byte,
	err error) {
	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	block, err := des.NewCipher(commonDeskey)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:des.BlockSize]
	ciphertext = ciphertext[des.BlockSize:]
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(ciphertext, ciphertext)
	if int(ciphertext[len(ciphertext)-1]) > len(ciphertext) {
		return nil, errors.New("aes decrypt failed")
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroUnPadding":
			plaintext = ZeroUnPadding(ciphertext)
		case "PKCS5UnPadding":
			plaintext = PKCS5UnPadding(ciphertext)
		}
	} else {
		plaintext = PKCS5UnPadding(ciphertext)
	}
	return plaintext, nil
}

// 3DES加密
func TripleDesCFBEncrypt(plaintext []byte, paddingType ...string) (ciphertext []byte,
	err error) {
	block, err := des.NewTripleDESCipher(commonDeskey)
	if err != nil {
		return nil, err
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroPadding":
			plaintext = ZeroPadding(plaintext, des.BlockSize)
		case "PKCS5Padding":
			plaintext = PKCS5Padding(plaintext, des.BlockSize)
		}
	} else {
		plaintext = PKCS5Padding(plaintext, des.BlockSize)
	}
	ciphertext = make([]byte, des.BlockSize+len(plaintext))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(ciphertext[des.BlockSize:],
		plaintext)
	return ciphertext, nil
}

// 3DES解密
func TripleDesCFBDecrypt(ciphertext []byte, paddingType ...string) (plaintext []byte,
	err error) {
	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	block, err := des.NewTripleDESCipher(commonDeskey)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:des.BlockSize]
	ciphertext = ciphertext[des.BlockSize:]
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(ciphertext, ciphertext)
	if int(ciphertext[len(ciphertext)-1]) > len(ciphertext) {
		return nil, errors.New("aes decrypt failed")
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroUnPadding":
			plaintext = ZeroUnPadding(ciphertext)
		case "PKCS5UnPadding":
			plaintext = PKCS5UnPadding(ciphertext)
		}
	} else {
		plaintext = PKCS5UnPadding(ciphertext)
	}
	return plaintext, nil
}
