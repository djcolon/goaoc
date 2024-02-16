package main

import (
	"reflect"
	"testing"
)

// Tests the function that loads a tree map file into a matrix.
func TestLoadMatrix(t *testing.T) {
	expected := [][]uint8{
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

// Test West get
func TestGetTreeVisibleFromWest(t *testing.T) {
	if !GetTreeVisibleFromWest(0b_0001_0000) {
		t.Fatal("GetTreeVisibleFromWest returned false, should be true.")
	}
	if GetTreeVisibleFromWest(0b_1110_1111) {
		t.Fatal("GetTreeVisibleFromWest returned true, should be false.")
	}
}

// Test North get
func TestGetTreeVisibleFromNorth(t *testing.T) {
	if !GetTreeVisibleFromNorth(0b_0010_0000) {
		t.Fatal("GetTreeVisibleFromNorth returned false, should be true.")
	}
	if GetTreeVisibleFromNorth(0b_1101_1111) {
		t.Fatal("GetTreeVisibleFromNorth returned true, should be false.")
	}
}

// Test East get
func TestGetTreeVisibleFromEast(t *testing.T) {
	if !GetTreeVisibleFromEast(0b_0100_0000) {
		t.Fatal("GetTreeVisibleFromEast returned false, should be true.")
	}
	if GetTreeVisibleFromEast(0b_1011_1111) {
		t.Fatal("GetTreeVisibleFromEast returned true, should be false.")
	}
}

// Test West set
func TestSetTreeVisibleFromWest(t *testing.T) {
	// Set to true
	result := SetTreeVisibleFromWest(0, true)
	if result != 0b0001_0000 {
		t.Fatalf("SetTreeVisibleFromWest(0, true) incorrect, expected 0b0001_0000, got %#b", result)
	}
	result = SetTreeVisibleFromWest(0b0001_0000, true)
	if result != 0b0001_0000 {
		t.Fatalf("SetTreeVisibleFromWest(0b0001_0000, true) incorrect, expected 0b0001_0000, got %#b", result)
	}
	result = SetTreeVisibleFromWest(0b1110_1111, true)
	if result != 0b1111_1111 {
		t.Fatalf("SetTreeVisibleFromWest(0b1110_1111, true) incorrect, expected 0b1111_1111, got %#b", result)
	}

	// Set to false
	result = SetTreeVisibleFromWest(0, false)
	if result != 0b0000_0000 {
		t.Fatalf("SetTreeVisibleFromWest(0, false) incorrect, expected 0b0000_0000, got %#b", result)
	}
	result = SetTreeVisibleFromWest(0b0001_0000, false)
	if result != 0b0000_0000 {
		t.Fatalf("SetTreeVisibleFromWest(0b0001_0000, false) incorrect, expected 0b0000_0000, got %#b", result)
	}
	result = SetTreeVisibleFromWest(0b1110_1111, false)
	if result != 0b1110_1111 {
		t.Fatalf("SetTreeVisibleFromWest(0b1110_1111, false) incorrect, expected 0b1111_1111, got %#b", result)
	}
	result = SetTreeVisibleFromWest(0b1111_1111, false)
	if result != 0b1110_1111 {
		t.Fatalf("SetTreeVisibleFromWest(0b1111_1111, false) incorrect, expected 0b1110_1111, got %#b", result)
	}
}

// Test North set
func TestSetTreeVisibleFromNorth(t *testing.T) {
	// Set to true
	result := SetTreeVisibleFromNorth(0, true)
	if result != 0b0010_0000 {
		t.Fatalf("SetTreeVisibleFromNorth(0, true) incorrect, expected 0b0010_0000, got %#b", result)
	}
	result = SetTreeVisibleFromNorth(0b0010_0000, true)
	if result != 0b0010_0000 {
		t.Fatalf("SetTreeVisibleFromNorth(0b0010_0000, true) incorrect, expected 0b0010_0000, got %#b", result)
	}
	result = SetTreeVisibleFromNorth(0b1101_1111, true)
	if result != 0b1111_1111 {
		t.Fatalf("SetTreeVisibleFromNorth(0b1101_1111, true) incorrect, expected 0b1111_1111, got %#b", result)
	}

	// Set to false
	result = SetTreeVisibleFromNorth(0, false)
	if result != 0b0000_0000 {
		t.Fatalf("SetTreeVisibleFromNorth(0, false) incorrect, expected 0b0000_0000, got %#b", result)
	}
	result = SetTreeVisibleFromNorth(0b0010_0000, false)
	if result != 0b0000_0000 {
		t.Fatalf("SetTreeVisibleFromNorth(0b0010_0000, false) incorrect, expected 0b0000_0000, got %#b", result)
	}
	result = SetTreeVisibleFromNorth(0b1101_1111, false)
	if result != 0b1101_1111 {
		t.Fatalf("SetTreeVisibleFromNorth(0b1101_1111, false) incorrect, expected 0b1101_1111, got %#b", result)
	}
	result = SetTreeVisibleFromNorth(0b1111_1111, false)
	if result != 0b1101_1111 {
		t.Fatalf("SetTreeVisibleFromNorth(0b1111_1111, false) incorrect, expected 0b1101_1111, got %#b", result)
	}
}

func TestSetTreeVisibleFromEast(t *testing.T) {
	// Set to true
	result := SetTreeVisibleFromEast(0, true)
	if result != 0b0100_0000 {
		t.Fatalf("SetTreeVisibleFromEast(0, true) incorrect, expected 0b0100_0000, got %#b", result)
	}
	result = SetTreeVisibleFromEast(0b0100_0000, true)
	if result != 0b0100_0000 {
		t.Fatalf("SetTreeVisibleFromEast(0b0100_0000, true) incorrect, expected 0b0100_0000, got %#b", result)
	}
	result = SetTreeVisibleFromEast(0b1011_1111, true)
	if result != 0b1111_1111 {
		t.Fatalf("SetTreeVisibleFromEast(0b1011_1111, true) incorrect, expected 0b1111_1111, got %#b", result)
	}

	// Set to false
	result = SetTreeVisibleFromEast(0, false)
	if result != 0b0000_0000 {
		t.Fatalf("SetTreeVisibleFromEast(0, false) incorrect, expected 0b0000_0000, got %#b", result)
	}
	result = SetTreeVisibleFromEast(0b0100_0000, false)
	if result != 0b0000_0000 {
		t.Fatalf("SetTreeVisibleFromEast(0b0100_0000, false) incorrect, expected 0b0000_0000, got %#b", result)
	}
	result = SetTreeVisibleFromEast(0b1011_1111, false)
	if result != 0b1011_1111 {
		t.Fatalf("SetTreeVisibleFromEast(0b1011_1111, false) incorrect, expected 0b1011_1111, got %#b", result)
	}
	result = SetTreeVisibleFromEast(0b1111_1111, false)
	if result != 0b1011_1111 {
		t.Fatalf("SetTreeVisibleFromEast(0b1111_1111, false) incorrect, expected 0b1011_1111, got %#b", result)
	}
}

// Test height
func TestGetTreeHeight(t *testing.T) {
	result := GetTreeHeight(0b_0001_0111)
	if result != 7 {
		t.Fatalf("Expected 7, got %d", result)
	}
	result = GetTreeHeight(0b_1111_1001)
	if result != 9 {
		t.Fatalf("Expected 9, got %d", result)
	}
}

func TestGetTreeVisible(t *testing.T) {
	// Visible
	if !getTreeVisible(0b0100_0000) {
		t.Fatalf("getTreeVisible incorrect. 0b0100_0000 is visible.")
	}
	if !getTreeVisible(0b0110_0110) {
		t.Fatalf("getTreeVisible incorrect. 0b0110_0110 is visible.")
	}
	if !getTreeVisible(0b0110_0001) {
		t.Fatalf("getTreeVisible incorrect. 0b0110_0001 is visible.")
	}
	if !getTreeVisible(0b0001_0000) {
		t.Fatalf("getTreeVisible incorrect. 0b0001_0000 is visible.")
	}
	if !getTreeVisible(0b1111_0000) {
		t.Fatalf("getTreeVisible incorrect. 0b1111_0000 is visible.")
	}
	// Not visible
	if getTreeVisible(0b1000_0000) {
		t.Fatalf("getTreeVisible incorrect. 0b1000_0000 is not visible.")
	}
	if getTreeVisible(0b1000_0010) {
		t.Fatalf("getTreeVisible incorrect. 0b1000_0010 is not visible.")
	}
	if getTreeVisible(0b0000_1100) {
		t.Fatalf("getTreeVisible incorrect. 0b0000_1100 is not visible.")
	}
	if getTreeVisible(0b0000_1111) {
		t.Fatalf("getTreeVisible incorrect. 0b0000_1111 is not visible.")
	}
}
