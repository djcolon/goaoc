package main

import (
	"testing"
)

func TestGetScoreForRock(t *testing.T) {
	score, _ := getScoreForPlayerMove(X)
	if score != 1 {
		t.Fatalf("Expected score of 1 for rock, received: %d", score)
	}
}

func TestGetScoreForPaper(t *testing.T) {
	score, _ := getScoreForPlayerMove(Y)
	if score != 2 {
		t.Fatalf("Expected score of 2 for paper, received: %d", score)
	}
}

func TestGetScoreForScissors(t *testing.T) {
	score, _ := getScoreForPlayerMove(Z)
	if score != 3 {
		t.Fatalf("Expected score of 3 for scissors, received: %d", score)
	}
}

func TestGetScoreErrors(t *testing.T) {
	_, err := getScoreForPlayerMove(A)
	if err == nil {
		t.Fatalf("Expected error for opponent move input to getScoreForInput.")
	}
}

func TestGetScoreForMovesErrors(t *testing.T) {
	_, err := getScoreForMatchOutcome(X, X)
	if err == nil {
		t.Fatalf("Expected error for player move input to opponent move in TestGetScoreForMovesErrors.")
	}
	_, err = getScoreForMatchOutcome(A, C)
	if err == nil {
		t.Fatalf("Expected error for opponent move input to player move in TestGetScoreForMovesErrors.")
	}
}

func TestGetScoreForMovesWins(t *testing.T) {
	// Rock loses to paper
	score, _ := getScoreForMatchOutcome(A, Y)
	if score != 6 {
		t.Fatalf("Expected win in TestGetScoreForMovesErrors for input A, Y.")
	}
	// Paper loses to scissors
	score, _ = getScoreForMatchOutcome(B, Z)
	if score != 6 {
		t.Fatalf("Expected win in TestGetScoreForMovesErrors for input B, Z.")
	}
	// Scissors loses to rock
	score, _ = getScoreForMatchOutcome(C, X)
	if score != 6 {
		t.Fatalf("Expected win in TestGetScoreForMovesErrors for input C, X.")
	}
}

func TestGetScoreForMovesDraws(t *testing.T) {
	score, _ := getScoreForMatchOutcome(A, X)
	if score != 3 {
		t.Fatalf("Expected draw in TestGetScoreForMovesErrors for input A, X.")
	}
	score, _ = getScoreForMatchOutcome(B, Y)
	if score != 3 {
		t.Fatalf("Expected draw in TestGetScoreForMovesErrors for input B, Y.")
	}
	score, _ = getScoreForMatchOutcome(C, Z)
	if score != 3 {
		t.Fatalf("Expected draw in TestGetScoreForMovesErrors for input C, Z.")
	}
}

func TestGetScoreForMovesLosses(t *testing.T) {
	// Rock beats scissors
	score, _ := getScoreForMatchOutcome(A, Z)
	if score != 0 {
		t.Fatalf("Expected loss in TestGetScoreForMovesErrors for input A, Z.")
	}
	// Paper beats rock
	score, _ = getScoreForMatchOutcome(B, X)
	if score != 0 {
		t.Fatalf("Expected loss in TestGetScoreForMovesErrors for input B, X.")
	}
	// Scissors beats paper
	score, _ = getScoreForMatchOutcome(C, Y)
	if score != 0 {
		t.Fatalf("Expected loss in TestGetScoreForMovesErrors for input C, Y.")
	}
}

func TestParseStrategyLineErrors(t *testing.T) {
	_, _, err := parseStrategyLine("A X ")
	if err == nil {
		t.Fatalf("Expected error from parseStrategyLine for line of len 4.")
	}
	_, _, err = parseStrategyLine("D X")
	if err == nil {
		t.Fatalf("Expected error from parseStrategyLine for line with invalid opponentMove value.")
	}
	_, _, err = parseStrategyLine("A B")
	if err == nil {
		t.Fatalf("Expected error from parseStrategyLine for line with invalid playerMove value.")
	}
	_, _, err = parseStrategyLine("AXX")
	if err != nil {
		t.Fatalf("Didn't expect error from parseStrategyLine for irrelevant middle character not being a space.")
	}
}

func TestParseStrategyLine(t *testing.T) {
	opponentMove, playerMove, err := parseStrategyLine("A X")
	if err != nil {
		t.Fatalf("Unexpected error for parseStrategyLine(\"A X\"): '%v'", err)
	}
	if opponentMove != A {
		t.Fatalf("Expected 0 as opponentMove value for parseStrategyLine(\"A X\"), received %d", opponentMove)
	}
	if playerMove != X {
		t.Fatalf("Expected 3 as playerMove value for parseStrategyLine(\"A X\"), received %d", playerMove)
	}

	opponentMove, playerMove, err = parseStrategyLine("C Y")
	if err != nil {
		t.Fatalf("Unexpected error for parseStrategyLine(\"C Y\"): '%v'", err)
	}
	if opponentMove != C {
		t.Fatalf("Expected 2 as opponentMove value for parseStrategyLine(\"C Y\"), received %d", opponentMove)
	}
	if playerMove != Y {
		t.Fatalf("Expected 4 as playerMove value for parseStrategyLine(\"C Y\"), received %d", playerMove)
	}
}

func TestFileNotFound(t *testing.T) {
	_, err := calcScoreForStrategyFile("./idontexist")
	if err == nil {
		t.Fatalf("Expected error for non-exitent file.")
	}
}

func TestExample(t *testing.T) {
	score, err := calcScoreForStrategyFile("./input_example.txt")
	if err != nil {
		t.Fatalf("Unexpected error for calcScoreForStrategyFile(\"./input_example.txt\"): '%v'.", err)
	}
	if score != 15 {
		t.Fatalf("Expected score of 15 for example file, calculated: %d.", score)
	}
}
