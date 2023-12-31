package signalprocessor

// "Enum" of types.
type ProcessorEngine int

const (
	ArrayProcessorEngine ProcessorEngine = iota
	MapProcessorEngine
)

// Defines the interface for a signal processor.
// A signal processor is instantiated to a marker length, and search space size.
// A signal processor takes in one byte at a time, and returns true when the
// last <marker length> bytes were different.
// It assumes that bytes fed in are continuous from the same stream.
// Offset is subtracted from byte before it is mapped, such that the value
// of the lowest byte in the seacrh space - offset should be 0.
type SignalProcessor interface {
	// Fed a byte, to return a bool when a marker has been identified.
	Process(dataItem byte) (endOfMarker bool, err error)
	// Initialises the processor to the given search space and length.
	// For search space size, you can use byte('z'-'a') + 1
	// note the +1 as z-a is 25!
	Initialise(markerLength uint64, searchSpaceSize byte, searchSpaceOffset byte) (err error)
	// Returns the state of the internal counter tracking how many bytes have
	// been processed.
	GetCharactersprocessed() uint64
}
