package signalprocessor

import (
	"math"
	"testing"
)

// Generalised functions for SignalProcessor interfaces.
func testProcessHelper(t *testing.T, processor *SignalProcessor, input byte, expectedOutcome bool) {
	endOfBuffer, err := (*processor).Process(input)
	if err != nil {
		t.Fatalf("Unexpected error for input %c on processor %v.", input, *processor)
	}
	if endOfBuffer != expectedOutcome {
		t.Fatalf("Unexpected outcome for input %c on processor %v. Expected %v.", input, *processor, expectedOutcome)
	}
}

func testProcessHappyPath(t *testing.T, processor SignalProcessor) {
	processor.Initialise(4, byte('z'-'a'), byte('a'))
	testProcessHelper(t, &processor, 'b', false)
	testProcessHelper(t, &processor, 'v', false)
	testProcessHelper(t, &processor, 'w', false)
	testProcessHelper(t, &processor, 'b', false)
	testProcessHelper(t, &processor, 'j', true)

	processor.Initialise(4, byte('z'-'a'), byte('a'))
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
	processor.Initialise(4, byte('z'-'a'), byte('a'))
	_, err := processor.Process('`')
	if err == nil {
		t.Fatalf("Expected error for passing byte below offset into SignalProcessor.Process.")
	}
	_, err = processor.Process('{')
	if err == nil {
		t.Fatalf("Expected error for passing byte beyond search space into SignalProcessor.Process.")
	}
	processor.Initialise(300, byte('z'-'a'), byte('a'))
	for i := 0; i < math.MaxUint8; i++ {
		_, _ = processor.Process('a')
	}
	_, err = processor.Process('a')
	if err == nil {
		t.Fatalf("Expected error for filling up searchSpace counter to Uint8Max.")
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
