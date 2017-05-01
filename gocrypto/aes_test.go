// aes_test.go
package gocrypto

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAesDecrypt(t *testing.T) {
	SetAesKey("nanjishidu170501")
	se, err := AesEncrypt([]byte("aes-20170501-30-1000"))
	if err != nil {
		t.Fatal(err)
	}
	s := base64.StdEncoding.EncodeToString(se)
	fmt.Println("aes base64ToString:", s)
	sed, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	sd, err := AesDecrypt(sed)
	if string(sd) != "aes-20170501-30-1000" {
		t.Log(sd)
		t.FailNow()
	}
}
