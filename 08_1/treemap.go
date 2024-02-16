package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Loads the matrix of trees into a matrix of uint8s.
func LoadMatrix(filePath string, matrixSize int) [][]uint8 {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file at: '%s'.", filePath)
	}
	defer file.Close()
	log.Printf("Reading input from '%s'.\n", filePath)

	// Initialise scanner
	scanner := bufio.NewScanner(file)
	// We'll read bytes as we'll only ever get [0-9\n]
	scanner.Split(bufio.ScanBytes)

	// Set up our result.
	result := make([][]uint8, matrixSize)
	var n, i int
	var s string
	for scanner.Scan() {
		s = scanner.Text()
		if s == "\n" {
			n++
		} else if s != "\r" {
			i, _ = strconv.Atoi(s)
			result[n] = append(result[n], uint8(i))
		}
	}
	return result
}

// Gets whether the tree was marked visible from the west.
func GetTreeVisibleFromWest(tree uint8) bool {
	return tree&0b_0001_0000 > 0
}

// Gets whether the tree was marked visible from the north.
func GetTreeVisibleFromNorth(tree uint8) bool {
	return tree&0b_0010_0000 > 0
}

// Gets whether the tree was marked visible from the east.
func GetTreeVisibleFromEast(tree uint8) bool {
	return tree&0b_0100_0000 > 0
}

// Sets whether the tree was marked visible from the west.
func SetTreeVisibleFromWest(tree uint8, visible bool) uint8 {
	var mask uint8 = 0b_0001_0000
	if visible {
		// If visible we set the bit to 1.
		return tree | mask
	} else {
		// Otherwise we strip it out.
		// This is tree & 0b_1110_1111
		return tree & ^mask
	}
}

// Sets whether the tree was marked visible from the north.
func SetTreeVisibleFromNorth(tree uint8, visible bool) uint8 {
	var mask uint8 = 0b_0010_0000
	if visible {
		return tree | mask
	} else {
		return tree & ^mask
	}
}

// Sets whether the tree was marked visible from the east.
func SetTreeVisibleFromEast(tree uint8, visible bool) uint8 {
	var mask uint8 = 0b_0100_0000
	if visible {
		return tree | mask
	} else {
		return tree & ^mask
	}
}

// Gets the height of the tree.
func GetTreeHeight(tree uint8) uint8 {
	return tree & 0b_0000_1111
}
