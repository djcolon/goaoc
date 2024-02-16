package main

import (
	"reflect"
	"testing"
)

// tests the function marking trees as visible from the west.
func TestMarkTreesFromWest(t *testing.T) {
	input := [][]uint8{
		{2, 5, 4, 8, 1, 9},
		{8, 1, 0, 1, 9, 9},
	}
	expected := [][]uint8{
		{2 | 0b0001_0000, 5 | 0b0001_0000, 4, 8 | 0b0001_0000, 1, 9 | 0b0001_0000},
		{8 | 0b0001_0000, 1, 0, 1, 9 | 0b0001_0000, 9},
	}
	markTreesFromWest(input)
	if !reflect.DeepEqual(
		input,
		expected,
	) {
		t.Errorf("markTreesFromWest result incorrect. Expected:\n%#08b\n, received:\n%#08b", expected, input)
	}
}

// tests the function marking trees as visible from the North.
func TestMarkTreesFromNorth(t *testing.T) {
	input := [][]uint8{
		{2, 8},
		{5, 1},
		{4, 0},
		{8, 1},
		{1, 9},
		{9, 9},
	}
	expected := [][]uint8{
		{2 | 0b0010_0000, 8 | 0b0010_0000},
		{5 | 0b0010_0000, 1},
		{4, 0},
		{8 | 0b0010_0000, 1},
		{1, 9 | 0b0010_0000},
		{9 | 0b0010_0000, 9},
	}
	markTreesFromNorth(input)
	if !reflect.DeepEqual(
		input,
		expected,
	) {
		t.Errorf("markTreesFromNorth result incorrect. Expected:\n%#08b\n, received:\n%#08b", expected, input)
	}
}

// tests the function marking trees as visible from the east.
func TestMarkTreesFromEast(t *testing.T) {
	input := [][]uint8{
		{9, 1, 8, 4, 5, 2},
		{9, 9, 1, 0, 1, 8},
	}
	expected := [][]uint8{
		{9 | 0b0100_0000, 1, 8 | 0b0100_0000, 4, 5 | 0b0100_0000, 2 | 0b0100_0000},
		{9, 9 | 0b0100_0000, 1, 0, 1, 8 | 0b0100_0000},
	}
	markTreesFromEast(input)
	if !reflect.DeepEqual(
		input,
		expected,
	) {
		t.Errorf("markTreesFromEast result incorrect. Expected:\n%#08b\n, received:\n%#08b", expected, input)
	}
}

func TestMarkTreesFromSouthAndCount(t *testing.T) {
	input := [][]uint8{
		{9, 9},
		{1, 9},
		{8, 1},
		{4, 0},
		{5, 1},
		{2, 8},
	}
	expected := [][]uint8{
		{9, 9},
		{1, 9},
		{8, 1},
		{4, 0},
		{5, 1},
		{2, 8},
	}
	count := markTreesFromSouthAndCount(input)
	if !reflect.DeepEqual(
		input,
		expected,
	) {
		t.Errorf("markTreesFromSouthAndCount mark result incorrect. Expected:\n%#08b\n, received:\n%#08b", expected, input)
	}
	if count != 6 {
		t.Errorf("markTreesFromSouthAndCount count result incorrect. Expected: 6, received: %d", count)
	}
}

func TestMarkTreesFromSouthAndCountMarked(t *testing.T) {
	input := [][]uint8{
		{9, 9},
		{1 | 0b_0001_0000, 9},
		{8, 1},
		{4, 0 | 0b0100_0000},
		{5, 1},
		{2, 8},
	}
	expected := [][]uint8{
		{9, 9},
		{1 | 0b_0001_0000, 9},
		{8, 1},
		{4, 0 | 0b0100_0000},
		{5, 1},
		{2, 8},
	}
	count := markTreesFromSouthAndCount(input)
	if !reflect.DeepEqual(
		input,
		expected,
	) {
		t.Errorf("TestMarkTreesFromSouthAndCountMarked mark result incorrect. Expected:\n%#08b\n, received:\n%#08b", expected, input)
	}
	if count != 8 {
		t.Errorf("TestMarkTreesFromSouthAndCountMarked count result incorrect. Expected: 8, received: %d", count)
	}
}

// Test full functionality.
func TestComputeVisibleTrees(t *testing.T) {
	result := computeVisibleTrees("./example_input_1.txt", 5)
	if result != 21 {
		t.Fatalf("computeVisibleTrees incorrect for assignment example. Expected 21, got: %d.", result)
	}
}
