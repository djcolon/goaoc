package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Struct to hold a move instruction.
type move struct {
	count int
	from  int
	to    int
}

// Parses an input file line and returns a move.
func makeMove(inputLine string) (result move, err error) {
	// Line sare structured as follows:
	// move 1 from 2 to 1
	// All lines have the same format, so just grabbing the numbers from their
	// position will be easiest.
	splitLine := strings.Split(inputLine, " ")
	if len(splitLine) != 6 {
		return move{}, fmt.Errorf("inputline invalid, expected 6 words")
	}
	count, err := strconv.Atoi(splitLine[1])
	if err != nil {
		return move{}, fmt.Errorf("invalid 'count' in move string: '%s' expected integer", splitLine[1])
	}
	from, err := strconv.Atoi(splitLine[3])
	if err != nil {
		return move{}, fmt.Errorf("invalid 'from' in move string: '%s' expected integer", splitLine[3])
	}
	to, err := strconv.Atoi(splitLine[5])
	if err != nil {
		return move{}, fmt.Errorf("invalid 'to' in move string: '%s' expected integer", splitLine[5])
	}
	return move{count: count, from: from, to: to}, nil
}
