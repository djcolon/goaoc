package main

import (
	"fmt"
	"strconv"
)

// Struct that represents the boatyard populated with stacks of crates.
type stacks struct {
	stacks         [][]byte
	numberOfStacks int
}

// Alters the stacks according to the move instruction.
func (stacks stacks) moveCrates(instruction move) (err error) {
	// Validate the move.
	if instruction.count < 1 || instruction.from < 1 || instruction.to < 1 {
		return fmt.Errorf("move.count %d, move.from %d or move.to %d out of range - must be > 0", instruction.count, instruction.from, instruction.to)
	}
	if stacks.numberOfStacks < instruction.from {
		return fmt.Errorf("move.from %d out of range of stack's numberOfStacks %d", instruction.from, stacks.numberOfStacks)
	}
	if stacks.numberOfStacks < instruction.to {
		return fmt.Errorf("move.to %d out of range of stack's numberOfStacks %d", instruction.to, stacks.numberOfStacks)
	}
	if len(stacks.stacks[instruction.from-1]) < instruction.count {
		return fmt.Errorf("not enough crates on stack %d to satisfy move.count", instruction.from)
	}
	for moveNo := 0; moveNo < instruction.count; moveNo++ {
		// Get a crate.
		// Note move from and to are not 0-indexed.
		fromLength := len(stacks.stacks[instruction.from-1])
		crate := stacks.stacks[instruction.from-1][fromLength-1]
		// Shorten the from array.
		stacks.stacks[instruction.from-1] = stacks.stacks[instruction.from-1][:fromLength-1]
		// And append to the to.
		stacks.stacks[instruction.to-1] = append(stacks.stacks[instruction.to-1], crate)
	}
	return nil
}

// returns a string conatining the identifiers of the crates at the top of each
// stack.
func (stacks stacks) getTopCrates() (topCrates string) {
	topCrates = ""
	for _, stack := range stacks.stacks {
		if len(stack) > 0 {
			topCrates = fmt.Sprintf("%s%c", topCrates, stack[len(stack)-1])
		}
	}
	return topCrates
}

// Parses the index line of the stacks definition into a slice of indices
// for each crate character's position in the definition strings.
// The maximum stack number supported is 9.
func parseIndexLine(indexLine string) (indices []int, err error) {
	// Note the letters align with the numbers on the stack. Use that to index
	// the position in the string of each crate's character.
	// We'll pre-allocate an estimate of the number of stacks as our index
	// length.
	indices = make([]int, 0, len(indexLine)/3)
	previousCharacter := " "
	for i, character := range indexLine {
		charString := string(character)
		if character != ' ' {
			// Is it an integer?
			_, err := strconv.Atoi(charString)
			if err != nil {
				return []int{}, fmt.Errorf("invalid stack index, encountered non-integer at position %d, in line '%s'", i, indexLine)
			}
			// This functions doesn't support indices over 9
			if previousCharacter != " " {
				return []int{}, fmt.Errorf("invalid stack index, encountered index > 9 at position %d, in line '%s'", i, indexLine)
			}
			indices = append(indices, i)
		}
		previousCharacter = charString
	}
	return indices, nil
}

// Initialises stacks from the input file strings of format:
// [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3     <- index line
func makeStacks(inputFileStrings []string) (result stacks, err error) {
	// First grab the number mapping from the last line.
	result = stacks{}
	indices, err := parseIndexLine(inputFileStrings[len(inputFileStrings)-1])
	if err != nil {
		return result, err
	}
	// Make a rough guess of the highest stack we're likely to encounter.
	result.numberOfStacks = len(indices)
	highestStack := len(inputFileStrings) - 1
	highestExpectedStack := (result.numberOfStacks * highestStack) / 2
	// The instantiate our stacks.
	result.stacks = make([][]byte, result.numberOfStacks)
	for i := 0; i < result.numberOfStacks; i++ {
		result.stacks[i] = make([]byte, 0, highestExpectedStack)
	}
	// Then parse each line in the definition.
	// We're going from bottom to top so the top of each stack is at the highest
	// index in each array (which makes sense when we start moving them around)
	// Skip the last line as it is indices.
	for i := len(inputFileStrings) - 2; i >= 0; i-- {
		line := inputFileStrings[i]
		maxLineIndex := len(line) - 1
		for j, index := range indices {
			// Don't accidentally go out of bounds.
			if index > maxLineIndex {
				break
			}
			// otherwise grab the character.
			crateIdentifier := byte(line[index])
			// Space is not valid!
			if crateIdentifier != ' ' {
				result.stacks[j] = append(result.stacks[j], crateIdentifier)
			}
		}
	}
	return result, nil
}
