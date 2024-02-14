package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Loads the matrix of trees into a matrix of uint8s.
func loadMatrix(filePath string, matrixSize int) [][]uint8 {
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

// Computes the number of visible trees from the map in the given file path.
func computeVisibleTrees(filePath string, matrixSize int) (int, error) {
	//matrix := loadMatrix(filePath, matrixSize)
	return 0, nil
}

// Parses cli args and kicks off the program.
func main() {
	log.Println("Advent of Code 2022 - 8.1")

	// Checks args.
	path := ""
	if len(os.Args) < 2 {
		path = "./input.txt"
	} else {
		path = os.Args[1]
	}

	// Pass it to our main function.
	result, err := computeVisibleTrees(path, 99)
	if err != nil {
		log.Fatalf("Failed to compute number of trees for file '%s' with error: '%v'.\n", path, err)
	}
	fmt.Printf("Result: %d", result)
}
