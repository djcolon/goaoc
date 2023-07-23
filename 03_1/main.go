package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// gets the score for any item [a-z][A-Z].
func getItemPriority(item byte) (priority int, err error) {
	// Each character can be represented as a byte with an ASCII value.
	// Ranges a-z and A-Z are continuous.
	if item >= 'a' && item <= 'z' {
		// Independent of the value of a, this will score the range of a-z
		// starting at 1
		return int(item - 'a' + 1), nil
	}
	if item >= 'A' && item <= 'Z' {
		// Independent of the value of a, this will score the range of a-z
		// starting at 1
		return int(item - 'A' + 27), nil
	}
	return 0, fmt.Errorf("received out of bounds item in getItemScore: %d", item)
}

// Calculates the first index of the second compartment of a bag.
func calcStartIndexOfSecondCompartment(bag string) (index int, err error) {
	length := len(bag)
	// Do basic checks.
	if length == 0 {
		return 0, fmt.Errorf("bag was empty: %s", bag)
	}
	if length%2 != 0 {
		return 0, fmt.Errorf("bag contained an uneven number of items (%d): '%s'", length, bag)
	}
	// Then divide.
	// As we use a 0 index, dividing by two will yield the an index pointing at
	// the first item in the second half of the string.
	return length / 2, nil
}

// Finds the duplicate item between two compartments of a bag.
// Errors if there are no duplicates.
// Will return after a first duplicate is found.
// This function assumes any character in the bag string can be safely cast to
// a byte. If this is untrue it will lead to undefined behaviour (it will be
//	cast).
func findDuplicatesBetweenCompartments(bag string, secondCompartmentStartIndex int) (duplicate byte, err error) {
	firstCompartmentIndex := map[byte]bool{}
	for i, itemRune := range bag {
		item := byte(itemRune)
		if i < secondCompartmentStartIndex {
			// Build an index of items in the first compartment.
			firstCompartmentIndex[item] = true
		} else {
			// When in the second half, look for duplicates.
			_, inFirstCompartment := firstCompartmentIndex[item]
			if inFirstCompartment {
				return item, nil
			}
		}
	}
	// If we got here, we found no duplicates.
	return 0, fmt.Errorf("failed to find duplicates between compartments for bag '%s', with secondCompartmentStartIndex %d", bag, secondCompartmentStartIndex)
}

// Calculates the sum of priority scores for all the items in all the bags.
func calcPrioritySumForInputFile(filePath string) (prioritySum int, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	log.Printf("Reading input from '%s'.\n", filePath)

	// Initialise
	scanner := bufio.NewScanner(file)
	score := 0
	i := 0

	for scanner.Scan() {
		bag := scanner.Text()
		// Get the index start of our second compartment.
		secondCompartmentStartIndex, err := calcStartIndexOfSecondCompartment(bag)
		if err != nil {
			return 0, fmt.Errorf("failed to determine index of second compartment for bag '%s' on line %d with error: '%v'", bag, i, err)
		}
		// Get the duplicate.
		duplicate, err := findDuplicatesBetweenCompartments(bag, secondCompartmentStartIndex)
		if err != nil {
			return 0, fmt.Errorf("failed to find duplicate for bag '%s' on line %d with error: '%v'", bag, i, err)
		}
		// Get the duplicate's priority.
		priority, err := getItemPriority(duplicate)
		if err != nil {
			return 0, fmt.Errorf("failed to find duplicate priority on line %d with error: '%v'", i, err)
		}
		score += priority
		i++
	}

	return score, nil
}

// Main.
func main() {
	log.Println("Advent of Code 2022 - 3.1")

	// Checks args.
	if len(os.Args) < 2 {
		log.Fatal("Please enter an input file path.")
	}

	score, err := calcPrioritySumForInputFile(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to calculate priority sum with error: '%v'", err)
	}
	log.Printf("=========\n")
	log.Printf("Score for input file '%s' is: %d.", os.Args[1], score)
}
