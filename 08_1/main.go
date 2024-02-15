package main

import (
	"fmt"
	"log"
	"os"
)

// Computes the number of visible trees from the map in the given file path.
func computeVisibleTrees(filePath string, matrixSize int) (int, error) {
	//matrix := LoadMatrix(filePath, matrixSize)
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
