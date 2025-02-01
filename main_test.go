package main

import (
	"bytes"
	"os"
	"testing"
)

func TestPrintHelloWorld(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call printHelloWorld function
	go func() {
		printHelloWorld()
		w.Close()
	}()
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	os.Stdout = oldStdout

	if err != nil {
		t.Fatalf("Failed to read captured output: %v", err)
	}

	// Test stdout output
	output := buf.String()
	expectedOutput := "Hello World!\n"
	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}
