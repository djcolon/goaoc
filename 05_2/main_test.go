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
	if topCrates != "MCD" {
		t.Fatalf("Incorrect result for example input, expected \"MCD\", got \"%s\".", topCrates)
	}
}

// test the error groups.
func TestInvalidDefinitionErrors(t *testing.T) {
	_, err := getTopCratesAfterMovesForFile("example_input_invalid_definition.txt")
	if err == nil {
		t.Fatalf("Expected error for test file: '%v'.", err)
	}
}

func TestInvalidMoveErrors(t *testing.T) {
	_, err := getTopCratesAfterMovesForFile("example_input_invalid_move.txt")
	if err == nil {
		t.Fatalf("Expected error for test file: '%v'.", err)
	}
}
