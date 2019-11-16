package main

import "testing"

func TestCreateIdentifier(t *testing.T) {
	for i := 1; i < 100; i++ {
		identifier := createIdentifier(i)

		if len(identifier) != i {
			t.Errorf("Failed to create an identifier of length %d.", i)
		}
	}
}
