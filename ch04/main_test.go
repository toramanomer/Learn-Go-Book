package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestExercise1(t *testing.T) {
	ints := exercise1()
	if len(ints) != 100 {
		t.Fatalf("wanted length of 100, got: %d", len(ints))
	}
	for i, v := range ints {
		switch {
		case v < 0:
			t.Errorf("wanted value to be bigger than zero, got %d at index %d", v, i)
		case v > 100:
			t.Errorf("wanted value to be smaller than 100, got %d at index %d", v, i)
		}
	}
}

func TestExercise2(t *testing.T) {
	var buf bytes.Buffer
	var (
		two   = "Two!"
		three = "Three!"
		six   = "Six!"
		nvm   = "Never mind"
	)
	var ints = []int{1, 2, 3, 4, 5, 6, 7}
	var want = strings.Join([]string{nvm, two, three, two, nvm, six, nvm}, "")

	exercise2(&buf, ints)
	if got := buf.String(); got != want {
		t.Errorf("want: %q, got %q", want, got)
	}
}
