package main

import (
	"math"
	"fmt"
)

// get the next full square, must be an odd int
// 27: sqrt(27) => 4.xyz
//	* The grid is square and must have an odd width since we start with a 1 in the middle
// 	* Look for the smallest square which contains the requested address
// 	* Gives us the width & height of the square
func findNextSquare(address int) int {
	//
	upperBound := math.Ceil(math.Sqrt(float64(address)))
	nextSquare := int(upperBound)

	if nextSquare%2 == 0 {
		return nextSquare + 1
	}
	return nextSquare
}

func getDirectPathsForSize(size int) []int {
	width := size
	height := size

	// offsets from the smallest bigger square to the center point of each side
	//	37  36  35  34  33  32  31
	//	38  17  16  15  14  13  30
	//	39  18   5   4   3  12  29
	//	40  19   6   1   2  11  28
	//	41  20   7   8   9  10  27
	//	42  21  22  23  24  25  26
	//	43  44  45  46  47  48  49
	//
	// for size of 5, left: 19, top: 16, right: 11, bot: 23
	return []int{
		width*height - width - (height-3)/2,                  // left
		width*height - width - (height - 1) - (width-3)/2,    // top
		width*height - 2*width - (height - 2) - (height-3)/2, // right
		width*height - (width-1)/2,                           // bottom
	}
}

func ManhattanDistanceToAccessPort(address int) int {

	squareSize := findNextSquare(address)
	directTargets := getDirectPathsForSize(squareSize)

	// find nearest to address
	nearest := math.MaxInt32
	for _, target := range directTargets {
		if math.Abs(float64(target-address)) < math.Abs(float64(nearest-address)) {
			nearest = target
		}
	}

	movesToNearest := int(math.Abs(float64(address - nearest)))
	directMoves := (squareSize - 1) / 2

	return directMoves + movesToNearest
}

func main() {
	fmt.Println(ManhattanDistanceToAccessPort(347991))
}
