package main

import "testing"

// Tests the computeDirSizes wit the problem's example.
func TestComputeSmallestRemovableDir(t *testing.T) {
	result, err := computeSmallestRemovableDir("./example_input_1.txt")
	if err != nil {
		t.Fatalf("computeSmallestRemovableDir failed for example_input_1.txt with: %v.", err)
	}
	if result != 24933642 {
		t.Fatalf("computeSmallestRemovableDir computed incorrect size, expected 24933642, got: %d.", result)
	}
}

// Test the function calculating how much space we need to free up.
func TestComputeMinimumSpaceToFree(t *testing.T) {
	result := computeMinimumSpaceToFree(48381165)
	if result != 8381165 {
		t.Fatalf("computeMinimumSpaceToFree incorrect for example, expected 48381165, got %d.", result)
	}
}
