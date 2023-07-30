package main

import (
	"reflect"
	"testing"
)

// moveCrates
func testMoveCrateErrorHelper(t *testing.T, testStacks *stacks, move move) {
	err := (*testStacks).moveCrates(move)
	if err == nil {
		t.Fatalf("Expected error for stacks '%v' for move: %v.", *testStacks, move)
	}
}

func TestMoveCrateError(t *testing.T) {
	testStacks := stacks{stacks: [][]byte{{'Z', 'N'}, {'M', 'C', 'D'}, {'P'}}, numberOfStacks: 3}
	// Invalid moves.
	testMoveCrateErrorHelper(t, &testStacks, move{1, -1, 1})
	testMoveCrateErrorHelper(t, &testStacks, move{1, 1, -1})
	testMoveCrateErrorHelper(t, &testStacks, move{-1, 1, 1})
	testMoveCrateErrorHelper(t, &testStacks, move{-1, 1, 1})
	testMoveCrateErrorHelper(t, &testStacks, move{0, 1, 1})
	// Out of range
	testMoveCrateErrorHelper(t, &testStacks, move{1, 4, 1})
	testMoveCrateErrorHelper(t, &testStacks, move{1, 1, 4})
	// Not enough crates
	testMoveCrateErrorHelper(t, &testStacks, move{3, 1, 2})
}

func testMoveCrateHelper(t *testing.T, testStacks *stacks, move move, expected stacks) {
	err := (*testStacks).moveCrates(move)
	if err != nil {
		t.Fatalf("Unexpected error '%v' after move.", err)
	}
	if !reflect.DeepEqual(*testStacks, expected) {
		t.Fatalf("testStacks incorrect after move. Expected %v, got %v.", testStacks, expected)
	}
}

func TestMoveCrates(t *testing.T) {
	testStacks := stacks{stacks: [][]byte{{'Z', 'N'}, {'M', 'C', 'D'}, {'P'}}, numberOfStacks: 3}
	// Move one from example
	expectedAfterOne := stacks{stacks: [][]byte{{'Z', 'N', 'D'}, {'M', 'C'}, {'P'}}, numberOfStacks: 3}
	testMoveCrateHelper(t, &testStacks, move{1, 2, 1}, expectedAfterOne)
	// Two
	expectedAfterTwo := stacks{stacks: [][]byte{{}, {'M', 'C'}, {'P', 'D', 'N', 'Z'}}, numberOfStacks: 3}
	testMoveCrateHelper(t, &testStacks, move{3, 1, 3}, expectedAfterTwo)
	// Three
	expectedAfterThree := stacks{stacks: [][]byte{{'C', 'M'}, {}, {'P', 'D', 'N', 'Z'}}, numberOfStacks: 3}
	testMoveCrateHelper(t, &testStacks, move{2, 2, 1}, expectedAfterThree)
}

// getTopCrates
func TestGetTopCrates(t *testing.T) {
	testStacks := stacks{stacks: [][]byte{{'Z', 'N'}, {'M', 'C', 'D'}, {'P'}}, numberOfStacks: 3}
	result := testStacks.getTopCrates()
	if result != "NDP" {
		t.Fatalf("Incorrect output for getTopcrates. Expected \"NDP\", got %s for %v.", result, testStacks)
	}

	testStacks = stacks{stacks: [][]byte{{}, {'M', 'C', 'D'}, {'P'}}, numberOfStacks: 3}
	result = testStacks.getTopCrates()
	if result != "DP" {
		t.Fatalf("Incorrect output for getTopcrates. Expected \"DP\", got %s for %v.", result, testStacks)
	}
}

// parseIndexLine
func TestParseIndexLineError(t *testing.T) {
	_, err := parseIndexLine(" 1 2 3 4 5 6 7 8 9 10 11")
	if err == nil {
		t.Fatalf("Expected error for parseIndexLine with index > 9.")
	}

	_, err = parseIndexLine(" 1 2 3 4 5 a 7 8 9 ")
	if err == nil {
		t.Fatalf("Expected error for parseIndexLine with non-integer index.")
	}
}

func testParseIndexLineHelper(t *testing.T, input string, expected []int) {
	indices, err := parseIndexLine(input)
	if err != nil {
		t.Fatalf("Unexpected error for parseIndexLine(\"%s\").", input)
	}
	if !reflect.DeepEqual(indices, expected) {
		t.Fatalf("Expected: %v, got %v.", expected, indices)
	}
}

func TestParseIndexLine(t *testing.T) {
	testParseIndexLineHelper(t, " 1 2 3 4 5 6 7 8 9 ", []int{1, 3, 5, 7, 9, 11, 13, 15, 17})
	testParseIndexLineHelper(t, " 1 2 3 4 5 6 7  8    9", []int{1, 3, 5, 7, 9, 11, 13, 16, 21})
	testParseIndexLineHelper(t, "1 2 3 4 5 6 7  8    9", []int{0, 2, 4, 6, 8, 10, 12, 15, 20})
}

// makeStacks
func testMakeStacksHelper(t *testing.T, testStrings []string, expected stacks) {
	output, err := makeStacks(testStrings)
	if err != nil {
		t.Fatalf("Unexpected error '%v' in makeStacks for input '%v'.", err, testStrings)
	}
	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("Expected: %v, got %v for input %v.", expected, output, testStrings)
	}
}
func TestMakeStacks(t *testing.T) {
	testStrings := []string{"[Z] [M] [P]", " 1   2   3 "}
	expected := stacks{stacks: [][]byte{{'Z'}, {'M'}, {'P'}}, numberOfStacks: 3}
	testMakeStacksHelper(t, testStrings, expected)
	// test with appending spaces.
	testStrings = []string{"[N] [C]    ", "[Z] [M] [P]", " 1   2   3 "}
	expected = stacks{stacks: [][]byte{{'Z', 'N'}, {'M', 'C'}, {'P'}}, numberOfStacks: 3}
	testMakeStacksHelper(t, testStrings, expected)
	// test without appending spaces.
	testStrings = []string{"[N] [C]", "[Z] [M] [P]", " 1   2   3 "}
	expected = stacks{stacks: [][]byte{{'Z', 'N'}, {'M', 'C'}, {'P'}}, numberOfStacks: 3}
	testMakeStacksHelper(t, testStrings, expected)
	// test the full example
	testStrings = []string{"    [D]    ", "[N] [C]", "[Z] [M] [P]", " 1   2   3 "}
	expected = stacks{stacks: [][]byte{{'Z', 'N'}, {'M', 'C', 'D'}, {'P'}}, numberOfStacks: 3}
	testMakeStacksHelper(t, testStrings, expected)
}
