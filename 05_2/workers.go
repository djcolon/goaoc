package main

// Worker that instantiates stacks from the definitionLinesChannel, then
// processes moves from the moveChannel, and finally returns the result into the
// topCratesChannel.
func stacksWorker(definitionLinesChannel <-chan string, moveChannel <-chan move, topCratesChannel chan<- string) (err error) {
	// Build up definition strings untilt he channel closes, then use them to
	// instantiate our stacks.
	definitionLines := make([]string, 0, 10)
	for definitionLine := range definitionLinesChannel {
		definitionLines = append(definitionLines, definitionLine)
	}
	// The channel closed, build our stacks.
	resultStacks, err := makeStacks(definitionLines)
	if err != nil {
		return err
	}
	// Then start processing moves.
	for move := range moveChannel {
		resultStacks.moveCrates(move)
	}
	// Once the moveChannel is closed and empty we are done.
	// Return the result
	topCratesChannel <- resultStacks.getTopCrates()
	close(topCratesChannel)
	return nil
}

// Worker that turns move strings into moves for rpocessing by the stacksWorker.
func movesWorker(moveStringChannel <-chan string, moveChannel chan<- move) (err error) {
	for moveString := range moveStringChannel {
		// Parse mvoes as long as they come.
		move, err := makeMove(moveString)
		if err != nil {
			return err
		}
		moveChannel <- move
	}
	// Close the output channel when they are no longer coming.
	close(moveChannel)
	return nil
}
