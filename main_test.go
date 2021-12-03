package main

import "testing"

const NAME = "Ganne"

func TestGetName(t *testing.T) {
	name := GetName()

	if name != NAME {
		t.Fatalf("Got name: %s, expected: %s\n", name, NAME)
	}
}
