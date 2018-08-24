// ras_test.go
package gocrypto

import (
	"encoding/hex"
	"testing"
)

func TestRSADecrypt(t *testing.T) {
	err := GenRsaKey(1024, "/tmp/private.pem", "/tmp/public.pem")
	if err != nil {
		t.Fatal(err)
	}
	var pass = "hello world"
	e, err := RSAEncrypt([]byte(pass), "/tmp/public.pem")
	if err != nil {
		t.Fatal(err)
	}
	s := hex.EncodeToString(e)
	es, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	d, err := RSADecrypt(es, "/tmp/private.pem")
	if err != nil {
		t.Fatal(err)
	}
	if string(d) != "hello world" {
		t.Log(string(d))
		t.FailNow()
	}
}
