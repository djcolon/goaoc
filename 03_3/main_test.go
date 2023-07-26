package main

import (
	"testing"
)

// getItemPriority #############################################################
func TestGetItemPriorityErrors(t *testing.T) {
	_, err := getItemPriority('-')
	if err == nil {
		t.Fatalf("Expected error for getItemPriority('-').")
	}
	_, err = getItemPriority('@')
	if err == nil {
		t.Fatalf("Expected error for getItemPriority('@').")
	}
}

// helper function to test getItemPriority.
func testGetItemPriority(t *testing.T, item byte, expected int) {
	priority, err := getItemPriority(item)
	if priority != expected {
		t.Fatalf("Expected priority %d for getItemPriority('%c'). Got %d.", expected, item, priority)
	}
	if err != nil {
		t.Fatalf("Unexpected error for getItemPriority('%c'): '%v'.", item, err)
	}
}

func TestGetItemPriority(t *testing.T) {
	// Edges.
	testGetItemPriority(t, 'a', 1)
	testGetItemPriority(t, 'z', 26)
	testGetItemPriority(t, 'A', 27)
	testGetItemPriority(t, 'Z', 52)
	// From example.
	testGetItemPriority(t, 'p', 16)
	testGetItemPriority(t, 'L', 38)
	testGetItemPriority(t, 'P', 42)
	testGetItemPriority(t, 'v', 22)
	testGetItemPriority(t, 't', 20)
	testGetItemPriority(t, 's', 19)
}

// wasItemInPastBags ###########################################
func TestWasItemInPastBags(t *testing.T) {
	testSlice := []bool{true, true, true}
	if !wasItemInPastBags(&testSlice) {
		t.Fatalf("Expected wasItemInPastBags to return true for all true slice.")
	}
	testSlice = []bool{true, false, true}
	if wasItemInPastBags(&testSlice) {
		t.Fatalf("Expected wasItemInPastBags to return false for slice containing a false as middle item.")
	}
	testSlice = []bool{false, true, true, true}
	if wasItemInPastBags(&testSlice) {
		t.Fatalf("Expected wasItemInPastBags to return false for slice containing a false as first item.")
	}
	testSlice = []bool{true, true, true, false}
	if wasItemInPastBags(&testSlice) {
		t.Fatalf("Expected wasItemInPastBags to return false for slice containing a false as last item.")
	}
}

// findItemPresentInEachBag #################################################
func TestFindItemPresentInEachBagError(t *testing.T) {
	groupBags := []string{
		"aBCDEFG123",
		"HIJKLM456",
		"NNOPQaRSTUVW",
	}
	item, err := findItemPresentInEachBag(&groupBags, len(groupBags))
	if err == nil {
		t.Fatalf("Expected error for findItemPresentInEachBag for group bags with no duplicates. Received: '%c'.", item)
	}
}

// Helper.
func testFindItemPresentInEachBag(t *testing.T, expectedAnswer byte, groupBags []string) {
	itemPresentInEachBag, err := findItemPresentInEachBag(&groupBags, len(groupBags))
	if err != nil {
		t.Fatalf("Unexpected error for findItemPresentInEachBag(%v): '%v'.", groupBags, err)
	}
	if itemPresentInEachBag != expectedAnswer {
		t.Fatalf("Expected return of '%c' for findItemPresentInEachBag(%v), received: '%c'.", expectedAnswer, groupBags, itemPresentInEachBag)
	}
}

func TestFindItemPresentInEachBag(t *testing.T) {
	testFindItemPresentInEachBag(t, 'a', []string{
		"aBCDEFG123",
		"HIJaKLM456",
		"NNOPQaRSTUVW",
	})
	// From the puzzle example:
	testFindItemPresentInEachBag(t, 'r', []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	})
	testFindItemPresentInEachBag(t, 'Z', []string{
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	})
}

// calcPrioritySumForInputFile
func TestCalcPrioritySumForInputFileError(t *testing.T) {
	_, err := calcPrioritySumForInputFile("./idontexist.txt")
	if err == nil {
		t.Fatalf("Expected error for calcPrioritySumForInputFile for non-existant file.")
	}
}

func TestCalcPrioritySumForInputFile(t *testing.T) {
	prioritySum, err := calcPrioritySumForInputFile("./example_input.txt")
	if err != nil {
		t.Fatalf("Unexpected error for calcPrioritySumForInputFile with test file: %v.", err)
	}
	if prioritySum != 70 {
		t.Fatalf("Expected prioritySum 70 for calcPrioritySumForInputFile with test file, got: %d.", prioritySum)
	}
}
