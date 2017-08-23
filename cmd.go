package gomini

import (
	"bytes"
	"errors"
	"os/exec"
	"time"
	//"log"
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
func ExecWithTimeout(bin string, timeout time.Duration, args ...string) (string, error) {
	cmd := exec.Command(bin, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err, isTimeout := CmdRunWithTimeout(cmd, timeout)
	if err != nil {
		return "", err
	}
	if isTimeout {
		return "", errors.New("cmd run timeout")
	}
	return out.String(), nil
}
func CmdRunWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		// timeout
		<-done
		err := cmd.Process.Kill()
		if err != nil {
			return err, true
		}
		return nil, true
	case err = <-done:
		return err, false
	}
}
