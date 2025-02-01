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

func TestPrintUserData(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call printUserData function
	go func() {
		printUserData("John Doe", 30, "test#emia.com")
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
	expectedOutput := "Name:  John Doe\nAge:  30\nEmail:  test#emia.com\n"
	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}

func TestPrintUserDataEmptyEmail(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call printUserData function with empty email
	go func() {
		printUserData("Jane Doe", 28, "")
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
	expectedOutput := "Name:  Jane Doe\nAge:  28\nEmail:  \n"
	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}

func TestMainFunction(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call main function
	go func() {
		main()
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
	expectedOutput := "Hello World!\nName:  John Doe\nAge:  30\nEmail:  test#emia.com\n"
	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}
