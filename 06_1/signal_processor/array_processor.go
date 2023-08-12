package signalprocessor

import (
	"errors"
	"fmt"
	"math"
)

// Array based implementaiton of a Signalprocessor.
type ArrayProcessor struct {
	// An array to keep track of the bytes currently in out window.
	windowArray []byte
	// An array to keep track of our search space. Give that the maximum size
	searchSpaceArray []uint8
	// Counter to keep track of how far we are in out processing.
	charactersProcessed uint64
	// Lenght of the marker.
	markerLength uint64
	// Size of the search space.
	searchSpaceSize byte
	// Offset.
	searchSpaceOffset byte
	// The number of duplicates currently in the window.
	duplicatesInWindow int
}

// Updates the window array with the new item, and retruns the oldest item that
// is now shifting out of the window.
func (processor *ArrayProcessor) updateWindowArray(newItem byte) (oldItem byte) {
	processor.charactersProcessed++
	// charactersProcessed mod markerLength will allow us to loop over the items
	// in the window in circular fashion.
	index := processor.charactersProcessed % processor.markerLength
	oldItem = processor.windowArray[index]
	processor.windowArray[index] = newItem
	return oldItem
}

// Fed a byte, to return a bool when a marker has been identified.
// Will throw an error if there are more than 255 of the same character within
// the window.
func (processor *ArrayProcessor) Process(dataItem byte) (endOfBuffer bool, err error) {
	// Make sure we don't overflow.
	if processor.charactersProcessed == math.MaxUint64 {
		return false, errors.New("charactersProcessed reached overflow limit")
	}
	// And that our byte isn't out of bounds.
	if dataItem < processor.searchSpaceOffset {
		return false, fmt.Errorf("received dataItem %c smaller than searchSpaceOffset %c in ArrayProcessor", dataItem, processor.searchSpaceOffset)
	}
	if dataItem-processor.searchSpaceOffset > byte(processor.searchSpaceSize) {
		return false, fmt.Errorf("received dataItem %c larger than searchSpaceSize %d after shifting by searchSpaceOffset %c in ArrayProcessor", dataItem, processor.searchSpaceSize, processor.searchSpaceOffset)
	}
	// Then do the actual prcoessing.
	return processor.processUnsafe(dataItem)
}

// Process functions, but without up-front safety checks.
func (processor *ArrayProcessor) processUnsafe(dataItem byte) (endOfBuffer bool, err error) {
	// Shift the dataItem into the search space.
	shiftedDataItem := dataItem - processor.searchSpaceOffset

	// Update the search space array for our new item.
	dataItemArrayEntry := &processor.searchSpaceArray[shiftedDataItem]
	if *dataItemArrayEntry == 1 {
		// The entry is now a duplicate until it shift below 1
		processor.duplicatesInWindow++
	}
	// Well, mostly unsafe...
	if *dataItemArrayEntry == math.MaxUint8 {
		return false, fmt.Errorf("detected overflow in searchSpace. More than MaxUint8 characters of %c in searchSpaceArray", dataItem)
	}
	(*dataItemArrayEntry)++

	// Update our window array.
	oldItem := processor.updateWindowArray(shiftedDataItem)
	// And remove the oldest character as long as we're past the initial length
	// of our marker.
	if processor.charactersProcessed > processor.markerLength {
		oldItemArrayEntry := &processor.searchSpaceArray[oldItem]
		(*oldItemArrayEntry)--
		if *oldItemArrayEntry == 1 {
			// The entry is now a duplicate until it shift below 1
			processor.duplicatesInWindow--
		}

		// Finally test for duplicates.
		return processor.duplicatesInWindow == 0, nil
	}
	// Always false until we've filled our window.
	return false, nil
}

// Initialises the processor.
func (processor *ArrayProcessor) Initialise(markerLength uint64, searchSpaceSize byte, searchSpaceOffset byte) (err error) {
	processor.windowArray = make([]byte, markerLength)
	processor.searchSpaceArray = make([]uint8, searchSpaceSize)
	processor.charactersProcessed = 0
	processor.markerLength = markerLength
	processor.searchSpaceSize = searchSpaceSize
	processor.searchSpaceOffset = searchSpaceOffset
	processor.duplicatesInWindow = 0
	return nil
}

// Returns private charactersProcessed.
func (processor *ArrayProcessor) GetCharactersprocessed() (charactersProcessed uint64) {
	return charactersProcessed
}
