package main

import "testing"

func TestGetTopCratesAfterMovesForFileError(t *testing.T) {
	_, err := getTopCratesAfterMovesForFile("./idontexist")
	if err == nil {
		t.Fatalf("Expected error for non-existent test file.")
	}
}

func TestGetTopCratesAfterMovesForFile(t *testing.T) {
	topCrates, err := getTopCratesAfterMovesForFile("example_input.txt")
	if err != nil {
		t.Fatalf("Unexpected error for test file: '%v'.", err)
	}
	if topCrates != "CMZ" {
		t.Fatalf("Incorrect result for example input, expected \"CMZ\", got \"%s\".", topCrates)
	}
}
