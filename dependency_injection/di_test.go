package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Jake")

	got := buffer.String()
	want := "Hello, Jake"

	if got != want {
		t.Errorf("want: %q, got: %q.", want, got)
	}
}
