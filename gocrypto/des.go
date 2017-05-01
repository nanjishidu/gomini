// des.go
package gocrypto

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"sync"
)

var (
	syncDesMutex                 sync.Mutex
	commonDeskey, commonDesIvkey []byte
)

func SetDesKey(key string, ivkey ...string) {
	syncDesMutex.Lock()
	defer syncDesMutex.Unlock()
	commonDeskey = []byte(key)
	if len(ivkey) > 0 {
		commonDesIvkey = []byte(ivkey[0])
	} else {
		commonDesIvkey = []byte(key)
	}
}

func DesEncrypt(origData []byte, zeroPadding ...bool) ([]byte, error) {
	block, err := des.NewCipher(commonDeskey)
	if err != nil {
		return nil, err
	}
	if len(zeroPadding) > 0 && zeroPadding == true {
		origData = ZeroPadding(origData, block.BlockSize())
	} else {
		origData = PKCS5Padding(origData, block.BlockSize())
	}

	blockMode := cipher.NewCBCEncrypter(block, commonDesIvkey)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func DesDecrypt(crypted []byte, zeroPadding ...bool) ([]byte, error) {
	block, err := des.NewCipher(commonDeskey)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, commonDesIvkey)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	if len(zeroPadding) > 0 && zeroPadding == true {
		origData = ZeroUnPadding(origData)
	} else {
		origData = PKCS5UnPadding(origData)
	}
	return origData, nil
}

// 3DES加密
func TripleDesEncrypt(origData []byte, zeroPadding ...bool) ([]byte, error) {
	block, err := des.NewTripleDESCipher(commonDeskey)
	if err != nil {
		return nil, err
	}
	if len(zeroPadding) > 0 && zeroPadding == true {
		origData = ZeroPadding(origData, block.BlockSize())
	} else {
		origData = PKCS5Padding(origData, block.BlockSize())
	}
	blockMode := cipher.NewCBCEncrypter(block, commonDesIvkey[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 3DES解密
func TripleDesDecrypt(crypted []byte, zeroPadding ...bool) ([]byte, error) {
	block, err := des.NewTripleDESCipher(commonDeskey)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, commonDesIvkey[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	if len(zeroPadding) > 0 && zeroPadding == true {
		origData = ZeroUnPadding(origData)
	} else {
		origData = PKCS5UnPadding(origData)
	}
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
