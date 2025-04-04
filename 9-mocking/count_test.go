package main

import (
	"bytes"
	"testing"
)

// Using dependency inject to facilitate testing this in the DI section
func TestCount(t *testing.T) {

	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
