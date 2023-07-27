package main

import (
	"reflect"
	"testing"
)

// assignmentFromString
func testAssignmentFromStringErrorsHelper(t *testing.T, input string) {
	_, err := assignmentFromString(input)
	if err == nil {
		t.Fatalf("Expected error for assignmentFromString(\"%s\").", input)
	}
}
func TestAssignmentFromStringErrors(t *testing.T) {
	testAssignmentFromStringErrorsHelper(t, "123")
	testAssignmentFromStringErrorsHelper(t, "123_456")
	testAssignmentFromStringErrorsHelper(t, "123-456-789")
	testAssignmentFromStringErrorsHelper(t, "a-10")
	testAssignmentFromStringErrorsHelper(t, "1-z")
	testAssignmentFromStringErrorsHelper(t, "a-b")
	testAssignmentFromStringErrorsHelper(t, "a-b-b")
	testAssignmentFromStringErrorsHelper(t, "%-10")
}

func testAssignmentFromStringHelper(t *testing.T, input string, expectedLower, expectedUpper int) {
	assignment, err := assignmentFromString(input)
	if err != nil {
		t.Fatalf("Unexpected error for assignmentFromString(\"%s\"): '%v'.", input, err)
	}
	if assignment.LowerBound != expectedLower {
		t.Fatalf("Unexpected LowerBound for assignmentFromString(\"%s\"): got %d, expected %d.", input, assignment.LowerBound, expectedLower)
	}
	if assignment.UpperBound != expectedUpper {
		t.Fatalf("Unexpected LowerBound for assignmentFromString(\"%s\"): got %d, expected %d.", input, assignment.UpperBound, expectedUpper)
	}
}

func TestAssignmentFromString(t *testing.T) {
	testAssignmentFromStringHelper(t, "12-20", 12, 20)
	testAssignmentFromStringHelper(t, "20-12", 12, 20)
	testAssignmentFromStringHelper(t, "0-9999", 0, 9999)
	testAssignmentFromStringHelper(t, "9999-0", 0, 9999)
}

// assignmentStringsFromLine
func testAssignmentStringsFromLineErrorsHelper(t *testing.T, input string, expectedCount int) {
	_, err := assignmentStringsFromLine(input, expectedCount)
	if err == nil {
		t.Fatalf("Expected error for assignmentStringsFromLine(\"%s\").", input)
	}
}

func TestAssignmentStringsFromLineErrors(t *testing.T) {
	testAssignmentStringsFromLineErrorsHelper(t, "1-4,2-3,6-7", 2)
	testAssignmentStringsFromLineErrorsHelper(t, "1-4,2-3,6-7", 4)
	testAssignmentStringsFromLineErrorsHelper(t, "1-4", 0)
	testAssignmentStringsFromLineErrorsHelper(t, "", 1)
}

// doesOneContainOther
func testDoesOneContainOtherHelper(t *testing.T, oneLower, oneUpper, otherLower, otherUpper int, expected bool) {
	one := assignment{LowerBound: oneLower, UpperBound: oneUpper}
	other := assignment{LowerBound: otherLower, UpperBound: otherUpper}
	if doesOneContainOther(one, other) != expected {
		t.Fatalf("Expected %t, but received %t for doesOneContainOther(%v, %v).", expected, !expected, one, other)
	}
}

func TestDoesOneContainOther(t *testing.T) {
	// From exercise example.
	testDoesOneContainOtherHelper(t, 2, 4, 6, 8, false)
	testDoesOneContainOtherHelper(t, 2, 3, 4, 5, false)
	testDoesOneContainOtherHelper(t, 5, 7, 7, 9, false)
	testDoesOneContainOtherHelper(t, 2, 8, 3, 7, true)
	testDoesOneContainOtherHelper(t, 3, 7, 2, 8, true)
	testDoesOneContainOtherHelper(t, 6, 6, 4, 6, true)
	testDoesOneContainOtherHelper(t, 4, 6, 6, 6, true)
	testDoesOneContainOtherHelper(t, 2, 6, 4, 8, false)
	testDoesOneContainOtherHelper(t, 4, 8, 2, 6, false)
}

// getExclusiveCombinations
func TestGetExclusiveCombinations(t *testing.T) {
	result := getExclusiveCombinations(2)
	if !reflect.DeepEqual(result, [][]int{{0, 1}}) {
		t.Fatalf("getExclusiveCombinations gave incorrect result for length 2: %v.", result)
	}
	result = getExclusiveCombinations(3)
	if !reflect.DeepEqual(result, [][]int{{0, 1}, {0, 2}, {1, 2}}) {
		t.Fatalf("getExclusiveCombinations gave incorrect result for length 3: %v.", result)
	}
	result = getExclusiveCombinations(4)
	if !reflect.DeepEqual(result, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}) {
		t.Fatalf("getExclusiveCombinations gave incorrect result for length 4: %v.", result)
	}
}

// doesGroupContainRedundantElves
func TestDoesGroupContainRedundantElves(t *testing.T) {
	testAssignments := []assignment{assignment{2, 4}, assignment{6, 8}}
	containsRedundantElves := doesGroupContainRedundantElves(&testAssignments)
	if containsRedundantElves {
		t.Fatalf("Expected to find no redundant elves in %v", testAssignments)
	}
	testAssignments = []assignment{assignment{2, 8}, assignment{3, 7}}
	containsRedundantElves = doesGroupContainRedundantElves(&testAssignments)
	if !containsRedundantElves {
		t.Fatalf("Expected to find redundant elves in %v", testAssignments)
	}
}

// calcNumberOfContainedPairs
func TestCalcNumberOfContainedPairsErrors(t *testing.T) {
	_, err := calcNumberOfContainedPairs("./idontexist.txt")
	if err == nil {
		t.Fatalf("Expected error from calcNumberOfContainedPairs for non-existant input file.")
	}
}

func TestCalcNumberOfContainedPairs(t *testing.T) {
	total, err := calcNumberOfContainedPairs("example_input.txt")
	if err != nil {
		t.Fatalf("Failed calcNumberOfContainedPairs for example input with error '%v'.", err)
	}
	if total != 2 {
		t.Fatalf("Expected total of 2 for example_input.txt, received: %d.", total)
	}
}
