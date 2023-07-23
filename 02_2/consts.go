package main

// Enum representing different moves
const (
	rock = iota
	paper
	scissors
)

// Enum representing desired outcomes.
const (
	loss = iota
	draw
	win
)

// Map to look up opponent moves.
var moveLookup map[byte]int = map[byte]int{
	'A': rock,
	'B': paper,
	'C': scissors,
}

// Map to look up desired outcomes.
var outcomeLookup map[byte]int = map[byte]int{
	'X': loss,
	'Y': draw,
	'Z': win,
}

// Map to look up outcome score.
// Could be an array, but this is nice and readble
var outcomeScoreLookup map[int]int = map[int]int{
	loss: 0,
	draw: 3,
	win:  6,
}

// Map to look up player move score.
// Could be an array, but this is nice and readble
var playerMoveScoreLookup map[int]int = map[int]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

// Truth table to determine player move from opponent move (A, B, C) and desired
// outcome (X, Y, Z) (Lose, Draw, Win).
//       R  P  S
// Lose  S  R  P
// Draw  R  P  S
// Win   P  S  R
var moveTable [][]int = [][]int{
	{scissors, rock, paper},
	{rock, paper, scissors},
	{paper, scissors, rock},
}
