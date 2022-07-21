package main

import (
	"testing"
)

func TestWrap(t *testing.T) {
	line := "Hello, World. It's a beautiful day"
	wrapWidth := 15

	got := Wrap(line, wrapWidth)
	want := "Hello, World. \nIt's a \nbeautiful day"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
