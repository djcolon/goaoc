package main

import (
	"reflect"
	"testing"
)

// tests the function that loads a file into a matrix.
func TestLoadMatrix(t *testing.T) {
	expected := [][]uint8{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}
	matrix := loadMatrix("./example_input_1.txt", 5)
	if !reflect.DeepEqual(
		matrix,
		expected,
	) {
		t.Errorf("loadMatrix result incorrect. Expected %#v, received: %#v", expected, matrix)
	}
}
