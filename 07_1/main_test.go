package main

import "testing"

// Tests the computeDirSizes wit the problem's example.
func TestComputeDirSizes(t *testing.T) {
	result, err := computeDirSizes("./example_input_1.txt")
	if err != nil {
		t.Fatalf("computeDirSizes failed for example_input_1.txt with: %v.", err)
	}
	if result != 95437 {
		t.Fatalf("computeDirSizes computed incorrect size, expected 95437, got: %d.", result)
	}
}
