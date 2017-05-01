// aes.go
package gocrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"sync"
)

var (
	syncAesMutex sync.Mutex
	commonAeskey []byte
)

func SetAesKey(key string) {
	syncAesMutex.Lock()
	defer syncAesMutex.Unlock()
	commonAeskey = []byte(key)
}
func AesEncrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(commonAeskey)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(ciphertext[aes.BlockSize:],
		plaintext)
	return ciphertext, nil

}
func AesDecrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(commonAeskey)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}
