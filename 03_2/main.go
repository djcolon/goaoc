package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Size of a group of elves sharing a badge.
const groupSize int = 3

// We don't have to track items in the last bag since we do the checks in the
// last bag. Do the calculation here once rather than many times in our loop.
const itemSliceSize = groupSize - 1

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

// Checks whether the given item was present in the other bags.
func wasItemInPastBags(itemSlice *[]bool) (inAllBags bool) {
	for _, pastBag := range *itemSlice {
		if !pastBag {
			return false
		}
	}
	// If we haven't continued, each past bag contained the
	// item!
	return true
}

// Searches a set of bags for an item present in each bag, and returns the item.
func findItemPresentInEachBag(bags *[]string, groupSize int) (item byte, err error) {
	itemMap := make(map[byte]*[]bool)
	// Go over each bag.
	for i, bag := range *bags {
		lastBag := i+1 == groupSize
		// Iterate over each item.
		for _, itemRune := range bag {
			item := byte(itemRune)
			// We have an item in this bag.
			itemSlice, inMap := itemMap[item]
			if inMap {
				if lastBag {
					// If we're in the last bag, check whether the item was in
					// all the other bags.
					if wasItemInPastBags(itemSlice) {
						return item, nil
					}
					// If it wasn't int the past bag, continue to the next item.
					continue
				} else {
					// Update the itemSlice with the found item.
					(*itemSlice)[i] = true
				}
			} else if i == 0 {
				// The item wasn't in the map yet. Add it unless we're in the
				// not in the first bag.
				// Any items that weren't mapped beyond the first bag cannot
				// be present in all bags, so we won't bother adding them to the
				// map!
				newItemSlice := make([]bool, itemSliceSize)
				// Set the index corresponding to our bag to true.
				newItemSlice[i] = true
				itemMap[item] = &newItemSlice
			}
		}
	}
	// Uh oh, we haven't returned.
	return 0, fmt.Errorf("no item found in each bag for group")
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
	var groupBags []string

	for scanner.Scan() {
		if i%groupSize == 0 {
			// (re-)initialise our group's bags slice.
			groupBags = make([]string, 0, groupSize)
		}
		// Grab our bag string.
		bag := scanner.Text()
		// Put it in the group
		groupBags = append(groupBags, bag)

		// If this is the end of the group, process the group's bags.
		if (i+1)%groupSize == 0 {
			// Get the duplicate item.
			itemPresentInEachBag, err := findItemPresentInEachBag(&groupBags, groupSize)
			if err != nil {
				return 0, fmt.Errorf("no item found that was present in each bag for group ending on line %d with error: '%v'", i, err)
			}
			// get the score.
			itemPriority, err := getItemPriority(itemPresentInEachBag)
			if err != nil {
				return 0, fmt.Errorf("couldn't find priority score for item %c for group ending with line %d with error: '%v'", itemPresentInEachBag, i, err)
			}
			// Add it.
			score += itemPriority
		}
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
