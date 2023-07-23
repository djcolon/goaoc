package main

import (
	"testing"
)

// getScoreForPlayerMove
func TestGetScoreForRock(t *testing.T) {
	score, _ := getScoreForPlayerMove(rock)
	if score != 1 {
		t.Fatalf("Expected score of 1 for rock, received: %d", score)
	}
}

func TestGetScoreForPaper(t *testing.T) {
	score, _ := getScoreForPlayerMove(paper)
	if score != 2 {
		t.Fatalf("Expected score of 2 for paper, received: %d", score)
	}
}

func TestGetScoreForScissors(t *testing.T) {
	score, _ := getScoreForPlayerMove(scissors)
	if score != 3 {
		t.Fatalf("Expected score of 3 for scissors, received: %d", score)
	}
}

func TestGetScoreErrors(t *testing.T) {
	_, err := getScoreForPlayerMove(3)
	if err == nil {
		t.Fatalf("Expected error for out of bounds input to getScoreForInput.")
	}
}

// getScoreForOutcome
func TestGetScoreForLoss(t *testing.T) {
	score, _ := getScoreForMatchOutcome(loss)
	if score != 0 {
		t.Fatalf("Expected score of 0 for loss, received: %d", score)
	}
}

func TestGetScoreForDraw(t *testing.T) {
	score, _ := getScoreForMatchOutcome(draw)
	if score != 3 {
		t.Fatalf("Expected score of 3 for draw, received: %d", score)
	}
}

func TestGetScoreForWin(t *testing.T) {
	score, _ := getScoreForMatchOutcome(win)
	if score != 6 {
		t.Fatalf("Expected score of 6 for win, received: %d", score)
	}
}

func TestGetScoreForOutcomeErrors(t *testing.T) {
	_, err := getScoreForMatchOutcome(3)
	if err == nil {
		t.Fatalf("Expected error for out of bounds input to getScoreForMatchOutcome.")
	}
}

// getPlayerMoveForOpponentMoveAndOutcome
func TestGetPlayerMoveForOpponentMoveAndOutcomeWins(t *testing.T) {
	playerMove, _ := getPlayerMoveForOpponentMoveAndOutcome(rock, win)
	if playerMove != paper {
		t.Fatalf("Expected paper in TestGetScoreForMovesErrors for input rock, win.")
	}
	playerMove, _ = getPlayerMoveForOpponentMoveAndOutcome(paper, win)
	if playerMove != scissors {
		t.Fatalf("Expected scissors in TestGetScoreForMovesErrors for input paper, win")
	}
	playerMove, _ = getPlayerMoveForOpponentMoveAndOutcome(scissors, win)
	if playerMove != rock {
		t.Fatalf("Expected rock in TestGetScoreForMovesErrors for input scissors, win.")
	}
}

func TestGetPlayerMoveForOpponentMoveAndOutcomeDraws(t *testing.T) {
	playerMove, _ := getPlayerMoveForOpponentMoveAndOutcome(rock, draw)
	if playerMove != rock {
		t.Fatalf("Expected rock in TestGetScoreForMovesErrors for input rock, draw.")
	}
	playerMove, _ = getPlayerMoveForOpponentMoveAndOutcome(paper, draw)
	if playerMove != paper {
		t.Fatalf("Expected paper in TestGetScoreForMovesErrors for input paper, draw")
	}
	playerMove, _ = getPlayerMoveForOpponentMoveAndOutcome(scissors, draw)
	if playerMove != scissors {
		t.Fatalf("Expected scissors in TestGetScoreForMovesErrors for input scissors, draw.")
	}
}

func TestGetPlayerMoveForOpponentMoveAndOutcomeLosses(t *testing.T) {
	playerMove, _ := getPlayerMoveForOpponentMoveAndOutcome(rock, loss)
	if playerMove != scissors {
		t.Fatalf("Expected scissors in TestGetScoreForMovesErrors for input rock, loss.")
	}
	playerMove, _ = getPlayerMoveForOpponentMoveAndOutcome(paper, loss)
	if playerMove != rock {
		t.Fatalf("Expected rock in TestGetScoreForMovesErrors for input paper, loss")
	}
	playerMove, _ = getPlayerMoveForOpponentMoveAndOutcome(scissors, loss)
	if playerMove != paper {
		t.Fatalf("Expected paper in TestGetScoreForMovesErrors for input scissors, loss.")
	}
}

// parseStrategyLine
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
		t.Fatalf("Expected error from parseStrategyLine for line with invalid desiredOutcome value.")
	}
	_, _, err = parseStrategyLine("AXX")
	if err != nil {
		t.Fatalf("Didn't expect error from parseStrategyLine for irrelevant middle character not being a space.")
	}
}

func TestParseStrategyLine(t *testing.T) {
	opponentMove, desiredOutcome, err := parseStrategyLine("A X")
	if err != nil {
		t.Fatalf("Unexpected error for parseStrategyLine(\"A X\"): '%v'", err)
	}
	if opponentMove != rock {
		t.Fatalf("Expected rock as opponentMove value for parseStrategyLine(\"A X\"), received %d", opponentMove)
	}
	if desiredOutcome != loss {
		t.Fatalf("Expected loss as desiredOutcome value for parseStrategyLine(\"A X\"), received %d", desiredOutcome)
	}

	opponentMove, desiredOutcome, err = parseStrategyLine("C Y")
	if err != nil {
		t.Fatalf("Unexpected error for parseStrategyLine(\"C Y\"): '%v'", err)
	}
	if opponentMove != scissors {
		t.Fatalf("Expected scissors as opponentMove value for parseStrategyLine(\"C Y\"), received %d", opponentMove)
	}
	if desiredOutcome != draw {
		t.Fatalf("Expected draw as desiredOutcome value for parseStrategyLine(\"C Y\"), received %d", desiredOutcome)
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
	if score != 12 {
		t.Fatalf("Expected score of 12 for example file, calculated: %d.", score)
	}
}
