package main

import (
	"testing"

	"go-reloaded/punctuations"
)

func TestFormatPunctuation(t *testing.T) {
	got := punctuations.FormatPunctuation("I was thinking ... You were right")
	expect := "I was thinking... You were right"

	if got != expect {
		t.Errorf("got: %s, expect: %s", got, expect)
	}
}
