package main

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSum(t *testing.T) {
	r := calculate([]string{"5", "+", "5"})

	assert.Equal(t, float64(10), r)
}

func TestCalculateSub(t *testing.T) {
	r := calculate([]string{"5", "-", "15"})

	assert.Equal(t, float64(-10), r)
}

func TestCalculateMult(t *testing.T) {
	r := calculate([]string{"10", "x", "10"})

	assert.Equal(t, float64(100), r)
}

func TestCalculateDiv(t *testing.T) {
	r := calculate([]string{"100", "/", "10"})

	assert.Equal(t, float64(10), r)
}

func TestCalculateWithPanic(t *testing.T) {
	defer func() {
		err := recover().(error)
		if err != nil {
			assert.Contains(t, err.Error(), "parsing \"B\": invalid syntax")
		}
	}()

	calculate([]string{"10", "/", "B"})

	t.Errorf("😳 The panic function is not called.")
}

func TestMainWithPanicWithFork(t *testing.T) {
	if os.Getenv("FORK") == "1" {
		calculate([]string{"A", "/", "10"})
	}

	stdout, stderr, err := RunForkTest(t, "TestMainWithPanicWithFork")

	assert.Equal(t, err.Error(), "exit status 2")
	assert.Contains(t, stderr, "parsing \"A\": invalid syntax")
	assert.Contains(t, stdout, "FAIL")
}

func TestMainWithExit(t *testing.T) {
	oldStdout := os.Stdout
	oldArgs := os.Args

	var buf bytes.Buffer
	log.SetOutput(&buf)

	defer func() {
		os.Args = oldArgs
		os.Stdout = oldStdout
	}()

	os.Args = []string{"", "10", "x", "10"}
	main()

	log.SetOutput(os.Stderr)

	assert.Contains(t, buf.String(), "🟰  100.00")
}
