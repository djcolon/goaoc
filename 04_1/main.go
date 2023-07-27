package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Settings.
const groupSize int = 2

// Parses an assignment form a string. Expects format ^[0-9]+-[0-9]+$
func assignmentFromString(assignmentString string) (output assignment, err error) {
	// First split the string.
	assignmentSlice := strings.Split(assignmentString, "-")
	if len(assignmentSlice) != 2 {
		return assignment{}, fmt.Errorf("invalid assignment string: '%s' - expected two boundaries", assignmentString)
	}
	first, err := strconv.Atoi(assignmentSlice[0])
	if err != nil {
		return assignment{}, fmt.Errorf("invalid assignment string: '%s' - first boundary not an integer", assignmentString)
	}
	second, err := strconv.Atoi(assignmentSlice[1])
	if err != nil {
		return assignment{}, fmt.Errorf("invalid assignment string: '%s' - second boundary not an integer", assignmentString)
	}
	// gracefully cope with reverse order.
	if first < second {
		return assignment{
			LowerBound: first,
			UpperBound: second,
		}, nil
	}
	return assignment{
		LowerBound: second,
		UpperBound: first,
	}, nil
}

// Parses a line into a number of assignmentstrings.
// Unless expectedCount is set to -1, will check how many are in the string and
// error if it doesn't match.
func assignmentStringsFromLine(line string, expectedCount int) (assignmentStrings []string, err error) {
	if line == "" {
		return []string{}, fmt.Errorf("didn't expect empty string")
	}
	// Split.
	assignmentStrings = strings.Split(line, ",")
	count := len(assignmentStrings)
	if expectedCount != -1 && count != expectedCount {
		return []string{}, fmt.Errorf("expected %d assignmentStrings in '%s', found: %d", expectedCount, line, count)
	}
	return assignmentStrings, nil
}

// Tests whether one of the assignments fully contains the other.
func doesOneContainOther(one, other assignment) (oneContainsOther bool) {
	return (other.LowerBound <= one.LowerBound && other.UpperBound >= one.UpperBound) || (one.LowerBound <= other.LowerBound && one.UpperBound >= other.UpperBound)
}

// Get indices of every combination not with the same index for an array/slice
//  of a certain length.
func getExclusiveCombinations(length int) (result [][]int) {
	// We could use the combin package for this, but as this is a learning
	// exercise...
	// This shows a combination of every assignment with every other assignment
	// but itself for an example.
	//   0 1 2 3 4
	// 0 x o o o o
	// 1 x x o o o
	// 2 x x x o o
	// 3 x x x x o
	// 4 x x x x x
	// We can pre-allocate our slice as the sum of the series 0 .. length - 1
	resultSize := int(((length - 1) * (length)) / 2)
	result = make([][]int, resultSize)
	counter := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			result[counter] = []int{i, j}
			counter++
		}
	}
	return result
}

// Works out whether there is at least one combination of assignments in the
// group that contains another.
func doesGroupContainRedundantElves(groupAssignments *[]assignment) (containsRedundantElves bool) {
	// Look at every combination.
	exclusiveCombinations := getExclusiveCombinations(len(*groupAssignments))
	for _, combination := range exclusiveCombinations {
		if doesOneContainOther((*groupAssignments)[combination[0]], (*groupAssignments)[combination[1]]) {
			return true
		}
	}
	return false
}

// Work out the total umber of pairs where one assignment fully contains the
// other by reading an input file.
func calcNumberOfContainedPairs(filePath string) (redundantElves int, err error) {
	redundantElves = 0
	// Open the file.
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	log.Printf("Reading input from '%s'.\n", filePath)

	// Open a scanner
	scanner := bufio.NewScanner(file)
	i := 0

	// Read it line by line.
	for scanner.Scan() {
		line := scanner.Text()
		// Extract assignment strings from the line.
		assignmentStrings, err := assignmentStringsFromLine(line, groupSize)
		if err != nil {
			return 0, fmt.Errorf("failed to parse assignmentStrings from line %d: '%s' with error: '%v'", i, line, err)
		}
		// Get each elf's assignment.
		// We're writing this generalised for a larger group-size.
		groupAssignments := make([]assignment, groupSize)
		for j, assignmentString := range assignmentStrings {
			assignment, err := assignmentFromString(assignmentString)
			if err != nil {
				return 0, fmt.Errorf("failed to get assignment from string for assignment %d on line %d with error: '%v'", j, i, err)
			}
			// We know this will never be out-of-bounds as it was initialised to
			// groupSize, and we checked against groupSize in
			// assignmentFromString
			groupAssignments[j] = assignment
		}
		// We now have assignments for each elf in the group.
		// If any combination (not permutation) barring a combination with the
		// same group member completewly contains another we have at least one
		// redundant elf. We could build a graph of assignments containing each
		// other to remove multiple redundancy, but we will do so if it is
		// required for part two.
		containsRedundantElves := doesGroupContainRedundantElves(&groupAssignments)
		if err != nil {
			return 0, fmt.Errorf("failed to find redundant elves from string on line %d with error: '%v'", i, err)
		}
		if containsRedundantElves {
			redundantElves++
		}
		// Then test
		i++
	}

	return redundantElves, nil
}

// Main.
func main() {
	log.Println("Advent of Code 2022 - 4.1")

	// Checks args.
	if len(os.Args) < 2 {
		log.Fatal("Please enter an input file path.")
	}

	score, err := calcNumberOfContainedPairs(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to calculate number of contained pairs: '%v'", err)
	}
	log.Printf("=========\n")
	log.Printf("Score for input file '%s' is: %d.", os.Args[1], score)
}
