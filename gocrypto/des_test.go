// des_test.go
package gocrypto

import (
	"testing"
)

func TestTripleDesCBCDecrypt(t *testing.T) {
	SetDesKey("nanjishidunanjishidunanj") //24
	s, err := TripleDesCBCEncrypt([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	d, err := TripleDesCBCDecrypt(s)
	if err != nil {
		t.Fatal(err)
	}
	if string(d) != "hello" {
		t.Log(d)
		t.FailNow()
	}
}

func TestDesCBCDecrypt(t *testing.T) {
	SetDesKey("nanjishi") //8
	s, err := DesCBCEncrypt([]byte("world"))
	if err != nil {
		t.Fatal(err)
	}
	d, err := DesCBCDecrypt(s)
	if err != nil {
		t.Fatal(err)
	}
	if string(d) != "world" {
		t.Log(d)
		t.FailNow()
	}
}

func TestTripleDesCFBDecrypt(t *testing.T) {
	SetDesKey("nanjishidunanjishidunanj") //24
	s, err := TripleDesCFBEncrypt([]byte("hello CFB"))
	if err != nil {
		t.Fatal(err)
	}
	d, err := TripleDesCFBDecrypt(s)
	if err != nil {
		t.Fatal(err)
	}
	if string(d) != "hello CFB" {
		t.Log(d)
		t.FailNow()
	}
}

func TestDesCFBDecrypt(t *testing.T) {
	SetDesKey("nanjishi") //8
	s, err := DesCFBEncrypt([]byte("world CFB"))
	if err != nil {
		t.Fatal(err)
	}
	d, err := DesCFBDecrypt(s)
	if err != nil {
		t.Fatal(err)
	}
	if string(d) != "world CFB" {
		t.Log(d)
		t.FailNow()
	}
}
