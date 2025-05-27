package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

type MockWriter struct {
	vals []any
}

func (m *MockWriter) Write(p []byte) (n int, err error) {

	return 0, nil
}

func TestExercise1(t *testing.T) {
	originalStdout := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("error creating a pipe: %v", err)
	}

	os.Stdout = w

	exercise1()

	if err := w.Close(); err != nil {
		t.Fatalf("error closing writer: %v", err)
	}
	os.Stdout = originalStdout

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("error copying bytes: %v", err)
	}

	expected := "20\n20.00\n"
	actual := buf.String()
	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
