package main

import (
	"testing"

	signalprocessor "helder.uk/goaoc/06_1/signal_processor"
)

// Error handling.
func TestGetMarkerPositionForFileError(t *testing.T) {
	_, err := getMarkerPosition("./idontexist", signalprocessor.ArrayProcessorEngine)
	if err == nil {
		t.Fatalf("Expected error for non-existant test file.")
	}
}

func TestGetMarkerPositionForUninplementedEngineError(t *testing.T) {
	_, err := getMarkerPosition("./idontexist", signalprocessor.ProcessorEngine(99))
	if err == nil {
		t.Fatalf("Expected error for non-existant processing engine.")
	}
}

func TestGetMarkerPositionForNonMarkerInput(t *testing.T) {
	_, err := getMarkerPosition("./non_marker_input.txt", signalprocessor.ArrayProcessorEngine)
	if err == nil {
		t.Fatalf("Expected error for input without marker.")
	}
}

// Getmarkerposition.
func testGetMarkerPositionsForExamplesForEngineForFile(t *testing.T, engine signalprocessor.ProcessorEngine, filePath string, expected uint64) {
	markerPosition, err := getMarkerPosition(filePath, engine)
	if err != nil {
		t.Fatalf("Unexpected error for file '%s': '%v'.", filePath, err)
	}
	if markerPosition != expected {
		t.Fatalf("Incorrect result for file '%s' with engine %d, expected %d, got %d.", filePath, engine, expected, markerPosition)
	}
}

func testGetMarkerPositionsForExamplesForEngine(t *testing.T, engine signalprocessor.ProcessorEngine) {
	testGetMarkerPositionsForExamplesForEngineForFile(t, engine, "./example_input_1.txt", 5)
	testGetMarkerPositionsForExamplesForEngineForFile(t, engine, "./example_input_2.txt", 6)
	testGetMarkerPositionsForExamplesForEngineForFile(t, engine, "./example_input_3.txt", 10)
	testGetMarkerPositionsForExamplesForEngineForFile(t, engine, "./example_input_4.txt", 11)
}

// Array engine specific tests.
func TestGetMarkerPositionsForExamplesForArrayEngine(t *testing.T) {
	testGetMarkerPositionsForExamplesForEngine(t, signalprocessor.ArrayProcessorEngine)
}

// Map engine specific tests.
func TestGetMarkerPositionsForExamplesForMapEngine(t *testing.T) {
	testGetMarkerPositionsForExamplesForEngine(t, signalprocessor.MapProcessorEngine)
}
