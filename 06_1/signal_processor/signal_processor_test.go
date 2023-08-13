package signalprocessor

import (
	"math"
	"testing"
)

// Generalised functions for SignalProcessor interfaces.
func testProcessHelper(t *testing.T, processor *SignalProcessor, input byte, expectedOutcome bool) {
	endOfMarker, err := (*processor).Process(input)
	if err != nil {
		t.Fatalf("Unexpected error for input %c on processor %v.", input, *processor)
	}
	if endOfMarker != expectedOutcome {
		t.Fatalf("Unexpected outcome for input %c on processor %v. Expected %v.", input, *processor, expectedOutcome)
	}
}

func testProcessHappyPath(t *testing.T, processor SignalProcessor) {
	processor.Initialise(4, byte('z'-'a')+1, byte('a'))
	testProcessHelper(t, &processor, 'b', false)
	testProcessHelper(t, &processor, 'v', false)
	testProcessHelper(t, &processor, 'w', false)
	testProcessHelper(t, &processor, 'b', false)
	testProcessHelper(t, &processor, 'j', true)

	processor.Initialise(4, byte('z'-'a')+1, byte('a'))
	testProcessHelper(t, &processor, 'a', false)
	testProcessHelper(t, &processor, 'b', false)
	testProcessHelper(t, &processor, 'b', false)
	testProcessHelper(t, &processor, 'c', false)
	testProcessHelper(t, &processor, 'd', false)
	testProcessHelper(t, &processor, 'd', false)
	testProcessHelper(t, &processor, 'e', false)
	testProcessHelper(t, &processor, 'f', false)
	testProcessHelper(t, &processor, 'g', true)
}

func testProcessBounds(t *testing.T, processor SignalProcessor) {
	processor.Initialise(4, byte('z'-'a')+1, byte('a'))
	// test just outside bounds.
	_, err := processor.Process('`')
	if err == nil {
		t.Fatalf("Expected error for passing byte below offset into SignalProcessor.Process.")
	}
	_, err = processor.Process('{')
	if err == nil {
		t.Fatalf("Expected error for passing byte beyond search space into SignalProcessor.Process.")
	}
	// Test just inside bounds.
	_, err = processor.Process('a')
	if err != nil {
		t.Fatalf("Unexpected error for passing byte at lower limit inside search sapce ('a') into SignalProcessor.Process.")
	}
	_, err = processor.Process('z')
	if err != nil {
		t.Fatalf("Unexpected error for passing byte at upper limit inside search space ('z') into SignalProcessor.Process.")
	}
	// And overflowing our counter.
	processor.Initialise(300, byte('z'-'a')+1, byte('a'))
	for i := 0; i < math.MaxUint8; i++ {
		_, _ = processor.Process('a')
	}
	_, err = processor.Process('a')
	if err == nil {
		t.Fatalf("Expected error for filling up searchSpace counter to Uint8Max.")
	}
}

func testGetGetCharactersprocessed(t *testing.T, processor SignalProcessor) {
	processor.Initialise(4, byte('z'-'a')+1, byte('a'))
	if processor.GetCharactersprocessed() != 0 {
		t.Fatalf("Expected GetCharactersprocessed() 0, received %d.", processor.GetCharactersprocessed())
	}
	processor.Process('a')
	if processor.GetCharactersprocessed() != 1 {
		t.Fatalf("Expected GetCharactersprocessed() 1, received %d.", processor.GetCharactersprocessed())
	}
	processor.Process('b')
	if processor.GetCharactersprocessed() != 2 {
		t.Fatalf("Expected GetCharactersprocessed() 2, received %d.", processor.GetCharactersprocessed())
	}
	processor.Process('c')
	if processor.GetCharactersprocessed() != 3 {
		t.Fatalf("Expected GetCharactersprocessed() 3, received %d.", processor.GetCharactersprocessed())
	}
	processor.Process('d')
	if processor.GetCharactersprocessed() != 4 {
		t.Fatalf("Expected GetCharactersprocessed() 4, received %d.", processor.GetCharactersprocessed())
	}
	processor.Process('e')
	if processor.GetCharactersprocessed() != 5 {
		t.Fatalf("Expected GetCharactersprocessed() 5, received %d.", processor.GetCharactersprocessed())
	}
	processor.Process('e')
	if processor.GetCharactersprocessed() != 6 {
		t.Fatalf("Expected GetCharactersprocessed() 6, received %d.", processor.GetCharactersprocessed())
	}
}

// Specific tests for ArrayProcessor.
func TestArrayProcessorHappyPath(t *testing.T) {
	processor := ArrayProcessor{}
	testProcessHappyPath(t, &processor)
}

func TestArrayProcessorBound(t *testing.T) {
	processor := ArrayProcessor{}
	testProcessBounds(t, &processor)
}

func TestArrayProcessorGetCharactersProcessed(t *testing.T) {
	processor := ArrayProcessor{}
	testGetGetCharactersprocessed(t, &processor)
}
