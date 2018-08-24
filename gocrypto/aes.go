// aes.go
package gocrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"sync"
)

var (
	syncAesMutex sync.Mutex
	commonAeskey []byte
)

func SetAesKey(key string) (err error) {
	syncAesMutex.Lock()
	defer syncAesMutex.Unlock()
	b := []byte(key)
	if len(b) == 16 || len(b) == 24 || len(b) == 32 {
		commonAeskey = b
		return nil
	}
	return fmt.Errorf("key size is not 16 or 24 or 32, but %d", len(b))

}
func AesCFBEncrypt(plaintext []byte, paddingType ...string) (ciphertext []byte,
	err error) {
	block, err := aes.NewCipher(commonAeskey)
	if err != nil {
		return nil, err
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroPadding":
			plaintext = ZeroPadding(plaintext, aes.BlockSize)
		case "PKCS5Padding":
			plaintext = PKCS5Padding(plaintext, aes.BlockSize)
		}
	} else {
		plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	}

	ciphertext = make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(ciphertext[aes.BlockSize:],
		plaintext)
	return ciphertext, nil

}
func AesCFBDecrypt(ciphertext []byte, paddingType ...string) (plaintext []byte,
	err error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	block, err := aes.NewCipher(commonAeskey)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(ciphertext, ciphertext)
	if int(ciphertext[len(ciphertext)-1]) > len(ciphertext) {
		return nil, errors.New("aes decrypt failed")
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroUnPadding":
			plaintext = ZeroUnPadding(ciphertext)
		case "PKCS5UnPadding":
			plaintext = ZeroUnPadding(ciphertext)
		}
	} else {
		plaintext = PKCS5UnPadding(ciphertext)
	}
	return plaintext, nil
}

func AesCBCEncrypt(plaintext []byte, paddingType ...string) (ciphertext []byte,
	err error) {
	block, err := aes.NewCipher(commonAeskey)
	if err != nil {
		return nil, err
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroPadding":
			plaintext = ZeroPadding(plaintext, aes.BlockSize)
		case "PKCS5Padding":
			plaintext = PKCS5Padding(plaintext, aes.BlockSize)
		}
	} else {
		plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	}
	ciphertext = make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

func AesCBCDecrypt(ciphertext []byte, paddingType ...string) (plaintext []byte,
	err error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	block, err := aes.NewCipher(commonAeskey)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
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

func AesECBEncrypt(plaintext []byte, paddingType ...string) (ciphertext []byte, err error) {
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroPadding":
			plaintext = ZeroPadding(plaintext, aes.BlockSize)
		case "PKCS5Padding":
			plaintext = PKCS5Padding(plaintext, aes.BlockSize)
		}
	} else {
		plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	}
	if len(plaintext)%aes.BlockSize != 0 {
		return nil, errors.New("plaintext is not a multiple of the block size")
	}
	block, err := aes.NewCipher(commonAeskey)
	if err != nil {
		return nil, err
	}
	ciphertext = make([]byte, len(plaintext))
	NewECBEncrypter(block).CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

func AesECBDecrypt(ciphertext []byte, paddingType ...string) (plaintext []byte, err error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	// ECB mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	block, err := aes.NewCipher(commonAeskey)
	if err != nil {
		return nil, err
	}
	NewECBDecrypter(block).CryptBlocks(ciphertext, ciphertext)
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
