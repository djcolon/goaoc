package main

import "log"

// A badgeFinder is an elf that has been assigned to finding the badges in sets
// of groupSize bags. They'll take groups of bags from the bagGroups channel,
// find which item is the group's badge, and then write the score into the
// itemScores channel.
// TODO: figure out error handling in goroutines.
func badgeFinder(id int, bagGroups <-chan []string, groupSize int, itemScores chan<- int) {
	groupsProcessed := 0
	log.Printf("badgeFinder elf #%d, reporting for duty!", id)
	for bagGroup := range bagGroups {
		// Get the duplicate item.
		itemPresentInEachBag, _ := findItemPresentInEachBag(&bagGroup, groupSize)
		// get the score.
		itemPriority, _ := getItemPriority(itemPresentInEachBag)
		// Add it.
		itemScores <- itemPriority
		groupsProcessed++
	}
	log.Printf("badgeFinder elf #%d, signing off after processing %d group's bags!", id, groupsProcessed)
}

// scoreAdder is an elf with a calculator. As badgeFinders call out scores for
// their bagGroups, this elf adds them up. When all the badgeFinders are done
// they'll report back the total score.
func scoreAdder(itemScores <-chan int, totalScore chan<- int) {
	total := 0
	for score := range itemScores {
		total += score
	}
	totalScore <- total
}
