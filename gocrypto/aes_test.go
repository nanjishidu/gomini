// aes_test.go
package gocrypto

import (
	"encoding/base64"
	"testing"
)

// key长度 16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
func TestAesCFBDecrypt(t *testing.T) {
	err := SetAesKey("nanjishidu170501")
	if err != nil {
		t.Fatal(err)
	}
	se, err := AesCFBEncrypt([]byte("aes-20170501-30-1000"))
	if err != nil {
		t.Fatal(err)
	}
	s := base64.StdEncoding.EncodeToString(se)
	sed, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	sd, err := AesCFBDecrypt(sed)
	if string(sd) != "aes-20170501-30-1000" {
		t.Log(sd)
		t.FailNow()
	}
}

func TestAeCBCDecrypt(t *testing.T) {
	err := SetAesKey("nanjishidu170502")
	if err != nil {
		t.Fatal(err)
	}
	se, err := AesCBCEncrypt([]byte("aes-20170501-30-1002"))
	if err != nil {
		t.Fatal(err)
	}
	s := base64.StdEncoding.EncodeToString(se)
	sed, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	sd, err := AesCBCDecrypt(sed)
	if string(sd) != "aes-20170501-30-1002" {
		t.Log(sd)
		t.FailNow()
	}
}
