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
	score, inMap := playerMoveScoreLookup[input]
	if !inMap {
		return 0, errors.New("received invalid input to getScoreForPlayerMove")
	}
	return score, nil
}

// Returns the score for the match outcome.
func getScoreForMatchOutcome(outcome int) (score int, err error) {
	score, inMap := outcomeScoreLookup[outcome]
	if !inMap {
		return 0, errors.New("received invalid input to getScoreForMatchOutcome")
	}
	return score, nil
}

// Gets the player mvoe from the oppoentn's move and the desired outcome.
func getPlayerMoveForOpponentMoveAndOutcome(opponentMove, desiredOutcome int) (playerMove int, err error) {
	if playerMove < 0 || playerMove > 2 {
		return 0, fmt.Errorf("playerMove out of bounds in getPlayerMoveForOpponentMoveAndOutcome")
	}
	if desiredOutcome < 0 || desiredOutcome > 2 {
		return 0, fmt.Errorf("desiredOutcome out of bounds in getPlayerMoveForOpponentMoveAndOutcome")
	}
	return moveTable[opponentMove][desiredOutcome], nil
}

// Parses a string into an opponent move and desired outcome.
// String must have format `^[ABC] [XYZ]$` or an error is thrown.
func parseStrategyLine(line string) (opponentMove, desiredOutcome int, err error) {
	if len(line) != 3 {
		return 0, 0, fmt.Errorf("strategy line should be 3 characters, received: '%s'", line)
	}
	opponentMove, inMap := moveLookup[line[0]]
	if !inMap {
		return 0, 0, fmt.Errorf("opponent move should be one of [ABC], received: '%q'", line[0])
	}
	desiredOutcome, inMap = outcomeLookup[line[2]]
	if !inMap {
		return 0, 0, fmt.Errorf("desired outcome move should be one of [XYZ], received: '%q'", line[0])
	}
	return opponentMove, desiredOutcome, nil
}

// Calculates the score for a strategy file:
// - Tries to open the file
// - Reads the file line-by-line
// - For each line:
//   - Parses the contents of the line into an opponent move and desired outcome
//   - Gets the player move to achieve the outcome
//   - Gets the score for the move and outcome
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
		opponentMove, desiredOutcome, err := parseStrategyLine(line)
		if err != nil {
			return 0, fmt.Errorf("failed parsing '%s' on line %d with error: '%v'", filePath, i, err)
		}
		// Determine th eplayer move.
		playerMove, err := getPlayerMoveForOpponentMoveAndOutcome(opponentMove, desiredOutcome)
		if err != nil {
			return 0, fmt.Errorf("failed determining playerMove for entry '%s' on line %d with error: '%v'", line, i, err)
		}
		// Get the score for the playermove.
		scoreForPlayerMove, err := getScoreForPlayerMove(playerMove)
		if err != nil {
			return 0, fmt.Errorf("failed determining score for player move for entry '%s' on line %d with error: '%v'", line, i, err)
		}
		// Get the score for the match outcome.
		scoreForMatchOutcome, err := getScoreForMatchOutcome(desiredOutcome)
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
