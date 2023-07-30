package main

import "testing"

// makeMove
func testMakeMoveErrorHelper(t *testing.T, inputString string) {
	_, err := makeMove(inputString)
	if err == nil {
		t.Fatalf("Expected error for makeMove(\"%s\")", inputString)
	}
}

func TestMakeMoveError(t *testing.T) {
	testMakeMoveErrorHelper(t, "move 1 from 2 to")
	testMakeMoveErrorHelper(t, "move 1 from 2 to 1 now")
	testMakeMoveErrorHelper(t, "move a from 2 to 1")
	testMakeMoveErrorHelper(t, "move 1 from 2a to 1")
	testMakeMoveErrorHelper(t, "move 1 from b to 1")
	testMakeMoveErrorHelper(t, "move 1 from 2 to c")
}

func testMakeMoveHelper(t *testing.T, inputString string, count, from, to int) {
	move, err := makeMove(inputString)
	if err != nil {
		t.Fatalf("Unexpected error for makeMove(\"%s\"): '%v'.", inputString, err)
	}
	if move.count != count {
		t.Fatalf("Incorrect count for  makeMove(\"%s\"). Expected %d, got %d.", inputString, count, move.count)
	}
	if move.from != from {
		t.Fatalf("Incorrect from for  makeMove(\"%s\"). Expected %d, got %d.", inputString, from, move.from)
	}
	if move.to != to {
		t.Fatalf("Incorrect to for  makeMove(\"%s\"). Expected %d, got %d.", inputString, to, move.to)
	}
}

func TestMakeMove(t *testing.T) {
	testMakeMoveHelper(t, "move 22 from 8 to 4", 22, 8, 4)
	// This is expected behaviour.
	testMakeMoveHelper(t, "bob 22 francis 8 rupert 4", 22, 8, 4)
	testMakeMoveHelper(t, "move 9995 from 1 to 9", 9995, 1, 9)
}
