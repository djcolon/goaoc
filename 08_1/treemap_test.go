package main

import (
	"reflect"
	"testing"
)

// Tests the function that loads a tree map file into a matrix.
func TestLoadMatrix(t *testing.T) {
	expected := [][]uint16{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}
	matrix := LoadMatrix("./example_input_1.txt", 5)
	if !reflect.DeepEqual(
		matrix,
		expected,
	) {
		t.Errorf("loadMatrix result incorrect. Expected %#v, received: %#v", expected, matrix)
	}
}

// Test West
func TestGetHighestTreeWest(t *testing.T) {
	result := GetHighestTreeWest(0b_0111_0001_0001_0001)
	if result != 7 {
		t.Fatalf("Expected 7, got %d", result)
	}
	result = GetHighestTreeWest(0b_1001_0000_0000_0000)
	if result != 9 {
		t.Fatalf("Expected 9, got %d", result)
	}
}

// Test North
func TestGetHighestTreeNorth(t *testing.T) {
	result := GetHighestTreeNorth(0b_0001_0111_0001_0001)
	if result != 7 {
		t.Fatalf("Expected 7, got %d", result)
	}
	result = GetHighestTreeNorth(0b_0000_1001_0000_0000)
	if result != 9 {
		t.Fatalf("Expected 9, got %d", result)
	}
}

// Test East
func TestGetHighestTreeEast(t *testing.T) {
	result := GetHighestTreeEast(0b_0001_0001_0111_0001)
	if result != 7 {
		t.Fatalf("Expected 7, got %d", result)
	}
	result = GetHighestTreeEast(0b_0000_0000_1001_0000)
	if result != 9 {
		t.Fatalf("Expected 9, got %d", result)
	}
}

// Test height
func TestGetTreeHeight(t *testing.T) {
	result := GetTreeHeight(0b_0001_0001_0001_0111)
	if result != 7 {
		t.Fatalf("Expected 7, got %d", result)
	}
	result = GetTreeHeight(0b_0000_0000_0000_1001)
	if result != 9 {
		t.Fatalf("Expected 9, got %d", result)
	}
}
