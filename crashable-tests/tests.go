package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// Run a fork test that may crash using os.exit.
func RunForkTest(t *testing.T, testName string) (string, string, error) {
	cmd := exec.Command(os.Args[0], fmt.Sprintf("-test.run=%v", testName))
	cmd.Env = append(os.Environ(), "FORK=1")

	var stdoutB, stderrB bytes.Buffer
	cmd.Stdout = &stdoutB
	cmd.Stderr = &stderrB

	err := cmd.Run()

	return stdoutB.String(), stderrB.String(), err
}
