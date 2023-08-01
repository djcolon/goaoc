package main

import (
	"bufio"
	"log"
	"os"
)

// Loads the input file, parses it into an initial stacks definition and moves,
// processes the moves on the stacks and then returns the top crates.
func getTopCratesAfterMovesForFile(filePath string) (topCrates string, err error) {
	// Open the input file.
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	log.Printf("Reading input from '%s'.\n", filePath)

	// Initialise channels.
	definitionLinesChannel := make(chan string)
	moveStringChannel := make(chan string, 10)
	moveChannel := make(chan move, 10)
	topCratesChannel := make(chan string)

	// And our workers.
	go stacksWorker(definitionLinesChannel, moveChannel, topCratesChannel)
	go movesWorker(moveStringChannel, moveChannel)

	// Initialise scanner
	scanner := bufio.NewScanner(file)
	i := 0
	inStackDefinition := true

	// read the file.
	for scanner.Scan() {
		line := scanner.Text()
		if inStackDefinition {
			// First get the stack definition.
			if line == "" {
				inStackDefinition = false
				close(definitionLinesChannel)
			} else {
				definitionLinesChannel <- line
			}
		} else {
			// After that, process moves.
			moveStringChannel <- line
		}
		i++
	}
	close(moveStringChannel)
	// Done parsing the file. Wait for a result.
	topCrates = <-topCratesChannel
	return topCrates, nil
}

func main() {
	log.Println("Advent of Code 2022 - 5.1")

	// Checks args.
	if len(os.Args) < 2 {
		log.Fatal("Please enter an input file path.")
	}

	topCrates, err := getTopCratesAfterMovesForFile(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to get top crates with error: '%v'", err)
	}
	log.Printf("=========\n")
	log.Printf("Top crates after moves for file '%s' is: %s.", os.Args[1], topCrates)
}
