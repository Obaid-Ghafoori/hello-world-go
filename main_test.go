package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

// Test main output
func TestMainOutput(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Capture log output
	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)

	// Store original log output and restore it after test
	originalLogOutput := log.Writer()
	defer log.SetOutput(originalLogOutput)

	// Run main
	main()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read captured output: %v", err)
	}

	// Test stdout output
	output := buf.String()
	expectedOutput := "Hello World!\n"
	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}

	// Test log output
	logOutput := logBuffer.String()

	// Check if log contains timestamp
	if !strings.Contains(logOutput, "Application started at:") {
		t.Error("Log output doesn't contain 'Application started at:'")
	}

	// Check if the timestamp in the log is recent (within last second)
	timeStr := strings.TrimPrefix(logOutput, "Application started at: ")
	timeStr = strings.TrimSpace(timeStr)
	logTime, err := time.Parse("2006/01/02 15:04:05", timeStr)
	if err == nil { // Only check if we can parse the time
		if time.Since(logTime) > time.Second {
			t.Error("Log timestamp is too old")
		}
	}
}
