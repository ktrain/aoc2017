package main

import "fmt"
import "math"

const input = 265149

func main() {
	fmt.Println("Input:", input);
	// square width = the sqrt of the input, rounded up to the nearest odd number
	squareWidth := ((int(math.Ceil(math.Sqrt(input))) / 2) * 2) + 1
	maxNum := squareWidth * squareWidth

	lastEdgeNum := maxNum - (squareWidth / 2)
	distanceToMiddleOfEdge := 0

	if input >= lastEdgeNum {
		distanceToMiddleOfEdge = input - lastEdgeNum
	} else {
		distanceToMiddleOfEdge = (lastEdgeNum - input) % (squareWidth / 2)
	}

	distanceFromMiddleOfEdgeToCenter := squareWidth / 2
	manhattanDistance := distanceToMiddleOfEdge + distanceFromMiddleOfEdgeToCenter

	fmt.Println("Square width:", squareWidth)
	fmt.Println("Max num:", maxNum)
	fmt.Println("Distance to middle of edge:", distanceToMiddleOfEdge)
	fmt.Println("Manhattan distance", manhattanDistance)
}
