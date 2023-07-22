package main

// Consts representing opponent moves - rock, paper or scissors as A, B, C.
// Followed by player moves - rock, paper or scissors as X, Y, Z.
const (
	A = iota
	B
	C
	X
	Y
	Z
)

// Const representing the scores for each type of player move.
const (
	XScore = 1
	YScore = 2
	ZScore = 3
)

// Const representing the scores for each outcome.
const (
	LossScore = 0
	DrawScore = 3
	WinScore  = 6
)

// Truth table to calculate score from the opponent (A, B, C) and player
// (X, Y, Z) moves.
//    A  B  C
// X  3  6  0
// Y  0  3  6
// Z  6  0  3
var winTable [][]int = [][]int{
	{DrawScore, WinScore, LossScore},
	{LossScore, DrawScore, WinScore},
	{WinScore, LossScore, DrawScore},
}
