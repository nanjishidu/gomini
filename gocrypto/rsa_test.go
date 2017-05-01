// ras_test.go
package gocrypto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestRSADecrypt(t *testing.T) {
	GenRsaKey(1024)
	e, err := RSAEncrypt([]byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	s := hex.EncodeToString(e)
	fmt.Println(s)
	es, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	d, err := RSADecrypt(es)
	if err != nil {
		t.Fatal(err)
	}
	if string(d) != "hello world" {
		t.Log(string(d))
		t.FailNow()
	}
	fmt.Println(string(d))

}
