package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	signalprocessor "helder.uk/goaoc/06_2/signal_processor"
)

// Gets the first character in the file after a marker of length 4.
func getMarkerPosition(filePath string, engine signalprocessor.ProcessorEngine) (markerPosition uint64, err error) {
	return getMarkerPositionForMarkerLength(filePath, engine, 4)
}

// Gets the first character in the file after a marker of length 14.
func getMessagePosition(filePath string, engine signalprocessor.ProcessorEngine) (markerPosition uint64, err error) {
	return getMarkerPositionForMarkerLength(filePath, engine, 14)
}

// Returns the position after which the configured marker has been completed for
// the given file.
func getMarkerPositionForMarkerLength(filePath string, engine signalprocessor.ProcessorEngine, markerLength uint64) (markerPosition uint64, err error) {
	// Open the input file.
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	log.Printf("Reading input from '%s'.\n", filePath)

	// Instantiate a signal processor.
	var processor signalprocessor.SignalProcessor
	switch engine {
	case signalprocessor.ArrayProcessorEngine:
		processor = &signalprocessor.ArrayProcessor{}
	case signalprocessor.MapProcessorEngine:
		processor = &signalprocessor.MapProcessor{}
	default:
		return 0, errors.New("failed to instatiate SignalProcessor with unimplemented processing engine")
	}
	// And configure it for challenge 1.
	processor.Initialise(markerLength, byte('z'-'a')+1, byte('a'))

	// Now read the file into the processor.
	reader := bufio.NewReader(file)

	for {
		// Get a byte.
		dataItem, err := reader.ReadByte()
		if err != nil && errors.Is(err, io.EOF) {
			return 0, errors.New("encountered end of file before having found marker")
		} else if err != nil {
			return 0, fmt.Errorf("encountered error reading from input file: %v", err)
		}
		// Process it.
		endOfMarker, err := processor.Process(dataItem)
		if err != nil {
			return 0, fmt.Errorf("error processing character at position %d of input file: %v", processor.GetCharactersprocessed(), err)
		}
		// If we've reached the end of marker, return the position.
		if endOfMarker {
			return processor.GetCharactersprocessed(), nil
		}
	}
}

func main() {
	log.Println("Advent of Code 2022 - 6.2")

	// Checks args.
	if len(os.Args) < 2 {
		log.Fatal("Please enter an input file path.")
	}
	// And get the result.
	markerPosition, err := getMessagePosition(os.Args[1], signalprocessor.ArrayProcessorEngine)
	if err != nil {
		log.Fatalf("Failed to get message position for file %s with error: '%v'.", os.Args[1], err)
	}
	log.Printf("First message in file '%s' after character %d.", os.Args[1], markerPosition)
}
