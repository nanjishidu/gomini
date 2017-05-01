// des_test.go
package gocrypto

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestTripleDesDecrypt(t *testing.T) {
	SetDesKey("nanjishidunanjishidunanj") //24
	s, err := TripleDesEncrypt([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("tripledes base64ToString:", base64.StdEncoding.EncodeToString(s))
	d, err := TripleDesDecrypt(s)
	if err != nil {
		t.Fatal(err)
	}
	if string(d) != "hello" {
		t.Log(d)
		t.FailNow()
	}
}

func TestDesDecrypt(t *testing.T) {
	SetDesKey("nanjishi") //8
	s, err := DesEncrypt([]byte("world"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("des base64ToString:", base64.StdEncoding.EncodeToString(s))
	d, err := DesDecrypt(s)
	if err != nil {
		t.Fatal(err)
	}
	if string(d) != "world" {
		t.Log(d)
		t.FailNow()
	}
}
