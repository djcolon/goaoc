package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const dirSizeLimit int = 100000

func scanFileIntoProcessor(filePath string, processorIn chan<- string) {
	// Open the input file.
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file at: '%s'.", filePath)
	}
	defer file.Close()
	log.Printf("Reading input from '%s'.\n", filePath)

	// Initialise scanner
	scanner := bufio.NewScanner(file)

	// read the file.
	for scanner.Scan() {
		processorIn <- scanner.Text()
	}
	close(processorIn)
}

// Finds any directories under dirSizeLimit and returns the sum of their sizes.
func computeDirSizes(filePath string) (int, error) {
	// Instantiate the processor.
	processor := NewProcessor()
	// And start reading files.
	go scanFileIntoProcessor(filePath, processor.in)
	// And then wait until the processor has finished processing.
	processor.process()
	// Finally compute our result.
	return processor.SumDirSizesUpTo(dirSizeLimit), nil
}

// Parses cli args and kicks off the program.
func main() {
	log.Println("Advent of Code 2022 - 7.1")

	// Checks args.
	path := ""
	if len(os.Args) < 2 {
		//log.Fatal("Please enter an input file path.")
		path = "./input.txt"
	} else {
		path = os.Args[1]
	}

	// Pass it to our main function.
	result, err := computeDirSizes(path)
	if err != nil {
		log.Fatalf("Failed to compute dir sizes for file '%s' with error: '%v'.\n", os.Args[1], err)
	}
	fmt.Printf("Result: %d", result)
}
