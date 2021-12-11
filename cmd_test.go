package gomini

import (
	"testing"
	"time"
)

func TestExec(t *testing.T) {
	_, err := Exec("ping", "8.8.8.8")
	if err != nil {
		t.Fatal(err)
	}
}
func TestExecWithTimeout(t *testing.T) {
	_, err := ExecWithTimeout("ping", 5*time.Second, "8.8.8.8")
	if err != nil {
		t.Fatal(err)
	}
}
