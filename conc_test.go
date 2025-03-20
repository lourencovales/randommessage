package main

import "testing"

func TestMessageOutput(t *testing.T) {
	test := messageOutput(5)
	if len(test) != 5 {
		t.Errorf("Wanted 5, got something else")
	}
}
