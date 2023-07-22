package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

// Returns the score for the player's move. Error for non-player move input.
func getScoreForPlayerMove(input int) (score int, err error) {
	switch input {
	case X:
		return XScore, nil
	case Y:
		return YScore, nil
	case Z:
		return ZScore, nil
	}
	return 0, errors.New("received invalid input to getScoreForPlayerMove")
}

// Returns the score for the match outcome.
func getScoreForMatchOutcome(opponentMove, playerMove int) (score int, err error) {
	if opponentMove < 0 || opponentMove > 2 {
		return 0, fmt.Errorf("received invalid value for opponent move in getScoreForMoves: %d", opponentMove)
	}
	if playerMove-3 < 0 || playerMove-3 > 2 {
		return 0, fmt.Errorf("received invalid value for player move in getScoreForMoves: %d", playerMove)
	}
	return winTable[opponentMove][playerMove-3], nil
}

// Parses a string into an opponent and player move.
// String must have format `^[ABC] [XYZ]$` or an error is thrown.
func parseStrategyLine(line string) (opponentMove, playerMove int, err error) {
	if len(line) != 3 {
		return 0, 0, fmt.Errorf("strategy line should be 3 characters, received: '%s'", line)
	}
	switch line[0] {
	case 'A':
		opponentMove = A
	case 'B':
		opponentMove = B
	case 'C':
		opponentMove = C
	default:
		return 0, 0, fmt.Errorf("opponent move should be one of [ABC], received: '%q'", line[0])
	}
	switch line[2] {
	case 'X':
		playerMove = X
	case 'Y':
		playerMove = Y
	case 'Z':
		playerMove = Z
	default:
		return 0, 0, fmt.Errorf("player move should be one of [XYZ], received: '%q'", line[2])
	}
	return opponentMove, playerMove, nil
}

// Calculates the score for a strategy file:
// - Tries to open the file
// - Reads the file line-by-line
// - For each line:
//   - Parses the contents of the line into a player and opponent move
//   - gets the score for the move and outcome
//   - Adds them to the total score
// Returns the score.
// Errors if the file can't be read or contains invalid content.
func calcScoreForStrategyFile(filePath string) (score int, err error) {
	// Try to open the file
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	fmt.Printf("Reading from input file: '%s'.\n", filePath)

	// Init our file scanner.
	scanner := bufio.NewScanner(file)

	// Init our score, and track the line in case of errors.
	score = 0
	i := 0

	// Go through lines.
	for scanner.Scan() {
		// Read a line.
		line := scanner.Text()

		// Parse it.
		opponentMove, playerMove, err := parseStrategyLine(line)
		if err != nil {
			return 0, fmt.Errorf("failed parsing '%s' on line %d with error: '%v'", filePath, i, err)
		}
		// Get the score for the playermove.
		scoreForPlayerMove, err := getScoreForPlayerMove(playerMove)
		if err != nil {
			return 0, fmt.Errorf("failed determining score for player move for entry '%s' on line %d with error: '%v'", line, i, err)
		}
		// Get the score for the match outcome.
		scoreForMatchOutcome, err := getScoreForMatchOutcome(opponentMove, playerMove)
		if err != nil {
			return 0, fmt.Errorf("failed determining score for match outcome for entry '%s' on line %d with error: '%v'", line, i, err)
		}
		// Add up the result and rpint them.
		roundScore := scoreForPlayerMove + scoreForMatchOutcome
		score += roundScore
		fmt.Printf("%d %s - %d %d %d\n", i, line, scoreForPlayerMove, scoreForMatchOutcome, roundScore)
		// Increase line number.
		i++
	}

	return score, nil
}

// Main.
func main() {
	fmt.Println("Advent of Code 2022 - 2.1")

	// Check args.
	if len(os.Args) < 2 {
		log.Fatal("Please enter an input filename.")
	}

	// Then calculate our result.
	score, err := calcScoreForStrategyFile(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=======")
	fmt.Printf("Score for strategy file '%s' is %d.\n", os.Args[1], score)
}
