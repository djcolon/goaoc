package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Loads the matrix of trees into a matrix of uint16s.
func LoadMatrix(filePath string, matrixSize int) [][]uint16 {
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
	result := make([][]uint16, matrixSize)
	var n, i int
	var s string
	for scanner.Scan() {
		s = scanner.Text()
		if s == "\n" {
			n++
		} else if s != "\r" {
			i, _ = strconv.Atoi(s)
			result[n] = append(result[n], uint16(i))
		}
	}
	return result
}

// We're going to pack the tallest tree that is in one of the cardinal
// directions from a given tree into the 16 bit integer. These functions
// will retrieve them.
func GetHighestTreeWest(tree uint16) uint16 {
	return tree & 0b_1111_0000_0000_0000 >> 12
}

// We're going to pack the tallest tree that is in one of the cardinal
// directions from a given tree into the 16 bit integer. These functions
// will retrieve them.
func GetHighestTreeNorth(tree uint16) uint16 {
	return tree & 0b_0000_1111_0000_0000 >> 8
}

// We're going to pack the tallest tree that is in one of the cardinal
// directions from a given tree into the 16 bit integer. These functions
// will retrieve them.
func GetHighestTreeEast(tree uint16) uint16 {
	return tree & 0b_0000_0000_1111_0000 >> 4
}

func GetTreeHeight(tree uint16) uint16 {
	return tree & 0b_0000_0000_0000_1111
}
