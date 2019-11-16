package main

import "testing"

func TestCreateIdentifier(t *testing.T) {
	identifier := createIdentifier()

	if len(identifier) != 5 {
		t.Error("Failed to create an identifier of length 5.")
	}
}
