package main

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestCalculate(t *testing.T) {
	testCases := []struct {
		name       string
		expression []string
		wantInt    int
		wantErr    error
	}{
		{
			name:       "addition",
			expression: []string{"3", "+", "2"},
			wantInt:    5,
			wantErr:    nil,
		},
		{
			name:       "sub",
			expression: []string{"3", "-", "2"},
			wantInt:    1,
			wantErr:    nil,
		},
		{
			name:       "multiplication",
			expression: []string{"3", "*", "2"},
			wantInt:    6,
			wantErr:    nil,
		},
		{
			name:       "division",
			expression: []string{"10", "/", "2"},
			wantInt:    5,
			wantErr:    nil,
		},
		{
			name:       "unsupported op",
			expression: []string{"10", "...", "2"},
			wantInt:    0,
			wantErr:    ErrUnsupportedOperation,
		},
		{
			name:       "multiplication by zero",
			expression: []string{"10", "*", "0"},
			wantInt:    0,
			wantErr:    nil,
		},
		{
			name:       "dividing by zero",
			expression: []string{"10", "/", "0"},
			wantInt:    0,
			wantErr:    ErrDivisionByZero,
		},
		{
			name:       "op1 not int",
			expression: []string{"abc", "+", "0"},
			wantInt:    0,
			wantErr:    ErrOperandNotInt,
		},
		{
			name:       "op2 not int",
			expression: []string{"0", "+", "abc"},
			wantInt:    0,
			wantErr:    ErrOperandNotInt,
		},
		{
			name:       "empty slice",
			expression: []string{},
			wantInt:    0,
			wantErr:    ErrNotEnoughArgs,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			i, err := calculate(testCase.expression)

			if testCase.wantInt != i {
				t.Errorf("want int: %d, got: %d", testCase.wantInt, i)
			}
			if testCase.wantErr != err {
				t.Errorf("want err: %v, got: %v", testCase.wantErr, err)
			}
		})
	}
}

func TestFileLen(t *testing.T) {
	t.Run("correct length", func(t *testing.T) {

		tempFile, err := os.CreateTemp("", "testfile")
		if err != nil {
			t.Fatalf("error creating temp file for test: %v", err)
		}
		defer os.Remove(tempFile.Name())
		defer tempFile.Close()

		lengthWritten, err := io.WriteString(tempFile, "hello world")
		if err != nil {
			t.Fatalf("error writing to testfile: %v", err)
		}

		filelength, err := fileLen(tempFile.Name())

		if lengthWritten != filelength {
			t.Errorf("expected length: %d, got: %d", lengthWritten, filelength)
		}
		if err != nil {
			t.Errorf("did not expect err, got one: %v", err)
		}
	})

	t.Run("returns error on nonexistent file", func(t *testing.T) {
		_, err := fileLen("nonexistent.txt")
		if err == nil {
			t.Errorf("expected an error for nonexistent file")
		}
	})

	t.Run("using testdata", func(t *testing.T) {
		testdataDir := "testdata"
		if err := os.MkdirAll(testdataDir, 0755); err != nil {
			t.Fatalf("failed to create testdata dir: %v", err)
		}
		defer os.RemoveAll(testdataDir)

		filePath := filepath.Join(testdataDir, "temp.txt")
		content := []byte("hello world")
		if err := os.WriteFile(filePath, content, 0655); err != nil {
			t.Fatalf("failed to write test file: %v", err)
		}

		length, err := fileLen(filePath)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if length != len(content) {
			t.Errorf("expected length %d, got %d", len(content), length)
		}
	})
}

func TestPrefixer(t *testing.T) {
	helloPrefix := prefixer("Hello")
	if want, got := "Hello Bob", helloPrefix("Bob"); got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}
