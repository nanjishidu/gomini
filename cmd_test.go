package gomini

import (
	"testing"
	"time"
)

func TestExec(t *testing.T) {
	out, err := Exec("ls", "-al")
	t.Log(string(out))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
func TestExecWithTimeout(t *testing.T) {
	out, err := ExecWithTimeout("ls", 3*time.Second, "-al")
	t.Log(string(out))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
