package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sumTopElves(topElves []int) (sum int) {
	sum = 0
	for _, calories := range topElves {
		sum += calories
	}
	return sum
}

func updateRanking(topElves *[]int, elfTotal, topElvesCount int) {
	// We're done with the elf's inventory.
	// Compare the elf to the top elves and insert if appropriate.
	for i, calories := range *topElves {
		if elfTotal > calories {
			// Insert into topElves
			// Shift current rankings down, unless we're already at the last
			// index.
			if i < topElvesCount {
				// Starting at the last elf, shift all our elves down a rank - up to
				// our current el;f which will be replaced.
				for j := topElvesCount - 1; j > i; j-- {
					(*topElves)[j] = (*topElves)[j-1]
				}
			}
			// After bumping all elves down we can insert our new top elf.
			(*topElves)[i] = elfTotal
			return
		}
	}
}
func getElfWithMostCalories(inputPath string, topElvesCount int) (topElves []int, err error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	// Defer defers the execution of a funciton until the surrounding function
	// returns.
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Go over all the elves.
	// Sum the calories, when we end the elf, compare the total to the highest
	// encountered so far. if it is higher - store it.
	// We'll be left with the highest elf at the end.
	elfTotal := 0
	topElves = make([]int, topElvesCount, topElvesCount)

	text := ""
	for scanner.Scan() {
		text = scanner.Text()
		if text != "" {
			// We're still in an elf's inventory.
			// Add the snack's calories.
			entry, err := strconv.Atoi(text)
			if err != nil {
				return nil, err
			}
			elfTotal += entry
		} else {
			updateRanking(&topElves, elfTotal, topElvesCount)
			elfTotal = 0
		}
	}
	// Do a last update with our final elf (there may not be a blank line at the
	// end of the file)
	updateRanking(&topElves, elfTotal, topElvesCount)

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return topElves, nil
}

func main() {
	// Get the filename for the input.
	if len(os.Args) < 2 {
		log.Fatal("Please pass a file name to read as input.")
	}

	topElves, err := getElfWithMostCalories(os.Args[1], 3)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(topElves)
	fmt.Println("Total:", sumTopElves(topElves))
}
