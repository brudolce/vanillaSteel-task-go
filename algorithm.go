package main

import (
	"math"
)

func minNumberOfSquares(n int) int {
	squares := 0
	spaceLeft := n
	var nextSquare int

	for spaceLeft > 0 {
		if spaceLeft < 4 {
			squares += spaceLeft
			break
		}
		nextSquare = int(math.Floor((math.Sqrt(float64(spaceLeft)))))
		spaceLeft -= int(math.Pow(float64(nextSquare), 2))
		squares += 1
	}
	return squares
}
