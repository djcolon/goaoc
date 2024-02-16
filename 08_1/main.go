package main

import (
	"fmt"
	"log"
	"os"
)

// Mark all trees from the west.
func markTreesFromWest(matrix [][]uint8) {
	// Iterate over the matrix rows starting from the top.
	for ri := 0; ri < len(matrix); ri++ {
		tallestTreeSoFar := -1
		// Then over the columns in each row, starting from the left.
		for ci := 0; ci < len(matrix[ri]); ci++ {
			// Get info about the tree.
			tree := matrix[ri][ci]
			treeHeight := GetTreeHeight(tree)
			// Mark it.
			matrix[ri][ci] = SetTreeVisibleFromWest(tree, treeHeight > tallestTreeSoFar)
			// Remember the tallest tree.
			if treeHeight > tallestTreeSoFar {
				tallestTreeSoFar = treeHeight
			}
		}
	}
}

// Mark all trees from the north.
func markTreesFromNorth(matrix [][]uint8) {
	// Matrix is square, so using matrix[0] is fine (we can assume the matrix
	// is never empty).
	// Iterate over matrix columns starting from the left.
	for ci := 0; ci < len(matrix[0]); ci++ {
		tallestTreeSoFar := -1
		// Then iterate over the column starting from the top.
		for ri := 0; ri < len(matrix); ri++ {
			// Get info about the tree.
			tree := matrix[ri][ci]
			treeHeight := GetTreeHeight(tree)
			// Mark it.
			matrix[ri][ci] = SetTreeVisibleFromNorth(tree, treeHeight > tallestTreeSoFar)
			// Remember the tallest tree.
			if treeHeight > tallestTreeSoFar {
				tallestTreeSoFar = treeHeight
			}
		}
	}
}

// Mark all trees from the west.
func markTreesFromEast(matrix [][]uint8) {
	// Iterate over the matrix rows starting from the top.
	for ri := 0; ri < len(matrix); ri++ {
		tallestTreeSoFar := -1
		// Then over the columns in each row, starting from the right.
		for ci := len(matrix[ri]) - 1; ci >= 0; ci-- {
			// Get info about the tree.
			tree := matrix[ri][ci]
			treeHeight := GetTreeHeight(tree)
			// Mark it.
			matrix[ri][ci] = SetTreeVisibleFromEast(tree, treeHeight > tallestTreeSoFar)
			// Remember the tallest tree.
			if treeHeight > tallestTreeSoFar {
				tallestTreeSoFar = treeHeight
			}
		}
	}
}

// Mark all trees from the south and count how many trees are visible.
func markTreesFromSouthAndCount(matrix [][]uint8) int {
	// Matrix is square, so using matrix[0] is fine (we can assume the matrix
	// is never empty).
	// Iterate over matrix columns starting from the left.
	visibleTrees := 0
	for ci := 0; ci < len(matrix[0]); ci++ {
		tallestTreeSoFar := -1
		// Then iterate over the column starting from the bottom.
		for ri := len(matrix) - 1; ri >= 0; ri-- {
			// Get info about the tree.
			tree := matrix[ri][ci]
			treeHeight := GetTreeHeight(tree)
			// Remember the tallest tree.
			if treeHeight > tallestTreeSoFar {
				tallestTreeSoFar = treeHeight
				// Visible form the south.
				visibleTrees++
			} else if getTreeVisible(tree) {
				// Is the tree visible at all?
				visibleTrees++
			}
		}
	}
	return visibleTrees
}

// Computes the number of visible trees from the map in the given file path.
func computeVisibleTrees(filePath string, matrixSize int) int {
	matrix := LoadMatrix(filePath, matrixSize)
	markTreesFromWest(matrix)
	markTreesFromNorth(matrix)
	markTreesFromEast(matrix)
	result := markTreesFromSouthAndCount(matrix)

	return result
}

// Parses cli args and kicks off the program.
func main() {
	log.Println("Advent of Code 2022 - 8.1")

	// Checks args.
	path := ""
	if len(os.Args) < 2 {
		path = "./input.txt"
	} else {
		path = os.Args[1]
	}

	// Pass it to our main function.
	result := computeVisibleTrees(path, 99)

	fmt.Printf("Result: %d", result)
}
