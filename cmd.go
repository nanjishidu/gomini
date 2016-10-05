package gomini

import (
	"bytes"
	"os/exec"
)

func Exec(bin string, args ...string) (string, error) {
	cmd := exec.Command(bin, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
