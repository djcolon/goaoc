package main

import (
	"fmt"
	"log"
	"os"
)

// const dirSizeLimit int = 100000

// Finds any directories under dirSizeLimit and returns the sum of their sizes.
func computeDirSize(filePath string) (int, error) {
	result := 0

	// Open the input file.
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	log.Printf("Reading input from '%s'.\n", filePath)

	// Create root dir.
	// rootDir := newRootDir()

	return result, nil
}

// Parses cli args and kicks off the program.
func main() {
	log.Println("Advent of Code 2022 - 7.1")

	// Checks args.
	if len(os.Args) < 2 {
		log.Fatal("Please enter an input file path.")
	}

	// Pass it to our main function.
	result, err := computeDirSize(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to compute dir sizes for file '%s' with error: '%v'.\n", os.Args[1], err)
	}
	fmt.Printf("Result: %d", result)
}
