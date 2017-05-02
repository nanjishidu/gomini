// rsa.go
package gocrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

//openssl genrsa -out private.pem 1024
//openssl rsa -in private.pem -pubout -out public.pem

func GenRsaKey(bits int, pemFile ...string) error {
	var privatePem, publicPem = "private.pem","public.pem"
	if len(pemFile)>=2{
		privatePem = pemFile[0]
		publicPem = pemFile[1]
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create(privatePem)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create(publicPem)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

func RSAEncrypt(data []byte, publicPem ...string) ([]byte, error) {
	var publicPemPath = "public.pem"
	if len(publicPem)>= 1 {
		publicPemPath = publicPem[0]
	}
	publicKey, err := readFileBytes(publicPemPath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pubInterface.(*rsa.PublicKey), data)
}
func RSADecrypt(data []byte, privatePem ...string) ([]byte, error) {
	var privatePemPath  = "private.pem"
	if len(privatePem) >= 1 {
		privatePemPath = privatePem[0]
	}
	privateKey, err := readFileBytes(privatePemPath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, data)
}

