package main

import (
	"testing"
)

// getItemPriority #############################################################
func TestGetItemPriorityErrors(t *testing.T) {
	_, err := getItemPriority('-')
	if err == nil {
		t.Fatalf("Expected error for getItemPriority('-').")
	}
	_, err = getItemPriority('@')
	if err == nil {
		t.Fatalf("Expected error for getItemPriority('@').")
	}
}

// helper function to test getItemPriority.
func testGetItemPriority(t *testing.T, item byte, expected int) {
	priority, err := getItemPriority(item)
	if priority != expected {
		t.Fatalf("Expected priority %d for getItemPriority('%c'). Got %d.", expected, item, priority)
	}
	if err != nil {
		t.Fatalf("Unexpected error for getItemPriority('%c'): '%v'.", item, err)
	}
}

func TestGetItemPriority(t *testing.T) {
	// Edges.
	testGetItemPriority(t, 'a', 1)
	testGetItemPriority(t, 'z', 26)
	testGetItemPriority(t, 'A', 27)
	testGetItemPriority(t, 'Z', 52)
	// From example.
	testGetItemPriority(t, 'p', 16)
	testGetItemPriority(t, 'L', 38)
	testGetItemPriority(t, 'P', 42)
	testGetItemPriority(t, 'v', 22)
	testGetItemPriority(t, 't', 20)
	testGetItemPriority(t, 's', 19)
}

// calcStartIndexOfSecondCompartment ###########################################
func TestCalcStartIndexOfSecondCompartmentErrors(t *testing.T) {
	_, err := calcStartIndexOfSecondCompartment("")
	if err == nil {
		t.Fatalf("Expected error for calcStartIndexOfSecondCompartment(\"\").")
	}
	_, err = calcStartIndexOfSecondCompartment("123")
	if err == nil {
		t.Fatalf("Expected error for calcStartIndexOfSecondCompartment(\"123\").")
	}
}

// Helper function to test calcStartIndexOfSecondCompartment.
func testCalcStartIndexOfSecondCompartment(t *testing.T, bag string, expected int, expectedByte byte) {
	index, err := calcStartIndexOfSecondCompartment(bag)
	selectedByte := bag[index]
	if index != expected {
		t.Fatalf("Expected index %d for calcStartIndexOfSecondCompartment(\"%s\"). Got %d.", expected, bag, index)
	}
	if selectedByte != expectedByte {
		t.Fatalf("Expected byte %c for calcStartIndexOfSecondCompartment(\"%s\"). Got %c.", expectedByte, bag, selectedByte)
	}
	if err != nil {
		t.Fatalf("Unexpected error for calcStartIndexOfSecondCompartment(\"%s\"): '%v'.", bag, err)
	}
}

func TestCalcStartIndexOfSecondCompartment(t *testing.T) {
	testCalcStartIndexOfSecondCompartment(t, "1234", 2, '3')
	testCalcStartIndexOfSecondCompartment(t, "12345678", 4, '5')
}

// findDuplicatesBetweenCompartments ###########################################
func TestFindDuplicatesBetweenCompartmentsErrors(t *testing.T) {
	_, err := findDuplicatesBetweenCompartments("aaabbb", 3)
	if err == nil {
		t.Fatalf("Expected not found error for findDuplicatesBetweenCompartments(\"aaabbb\", 3).")
	}
	_, err = findDuplicatesBetweenCompartments("", 0)
	if err == nil {
		t.Fatalf("Expected not found error for findDuplicatesBetweenCompartments(\"aaabbb\", 3).")
	}
}

// Helper function to test findDuplicatesBetweenCompartments.
func testFindDuplicatesBetweenCompartments(t *testing.T, bag string, expected byte) {
	secondCompartmentStartIndex, _ := calcStartIndexOfSecondCompartment(bag)
	duplicate, err := findDuplicatesBetweenCompartments(bag, secondCompartmentStartIndex)
	if duplicate != expected {
		t.Fatalf("Expected duplicate %c for findDuplicatesBetweenCompartments(\"%s\", %d). Got %c.", expected, bag, secondCompartmentStartIndex, duplicate)
	}
	if err != nil {
		t.Fatalf("Unexpected error for findDuplicatesBetweenCompartments(\"%s\"): '%v'.", bag, err)
	}
}

func TestFindDuplicatesBetweenCompartments(t *testing.T) {
	// Use examples from puzzle.
	testFindDuplicatesBetweenCompartments(t, "vJrwpWtwJgWrhcsFMMfFFhFp", 'p')
	testFindDuplicatesBetweenCompartments(t, "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 'L')
	testFindDuplicatesBetweenCompartments(t, "PmmdzqPrVvPwwTWBwg", 'P')
	testFindDuplicatesBetweenCompartments(t, "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 'v')
	testFindDuplicatesBetweenCompartments(t, "ttgJtRGJQctTZtZT", 't')
	testFindDuplicatesBetweenCompartments(t, "CrZsJsPPZsGzwwsLwLmpwMDw", 's')
}

// calcPrioritySumForInputFile #################################################
func TestCalcPrioritySumForInputFileFileNotFound(t *testing.T) {
	_, err := calcPrioritySumForInputFile("./idontexist.txt")
	if err == nil {
		t.Fatalf("Expected error for non-existent input file.")
	}
}

func TestCalcPrioritySumForInputFile(t *testing.T) {
	score, err := calcPrioritySumForInputFile("example_input.txt")
	if err != nil {
		t.Fatalf("Unexpected error for example input file: '%v'.", err)
	}
	if score != 157 {
		t.Fatalf("Expected score of 157 for example input file, but got: %d.", score)
	}
}
