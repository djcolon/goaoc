package main

import (
	"reflect"
	"testing"
)

func TestGetElfDoesntError(t *testing.T) {
	path := "./example_input.txt"
	_, err := getElfWithMostCalories(path, 3)
	if err != nil {
		t.Fatalf("getElfWithMostCalories threw an error: %v", err)
	}
}

func TestGetElfIndex(t *testing.T) {
	path := "./example_input.txt"
	topElves, err := getElfWithMostCalories(path, 3)
	if err != nil {
		t.Fatalf("getElfWithMostCalories threw an error: %v", err)
	}

	expected := []int{24000, 11000, 10000}
	if !reflect.DeepEqual(topElves, expected) {
		t.Fatalf("Expected %d, got %d", expected, topElves)
	}
}

func TestSumTopElves(t *testing.T) {
	test := []int{24000, 11000, 10000}
	ans := sumTopElves(test)
	if ans != 45000 {
		t.Fatalf("sumTopElves returned %d, expected 45000", ans)
	}
}
