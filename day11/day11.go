package main

import (
	"io"
	"bufio"
	"strings"
	"os"
	"fmt"
	"math"
)

var (
	directionalMoves = map[string]*HexPoint{
		"n":  NewHexPoint(0, 1, -1),
		"s":  NewHexPoint(0, -1, 1),
		"nw": NewHexPoint(-1, 1, 0),
		"se": NewHexPoint(1, -1, 0),
		"sw": NewHexPoint(-1, 0, 1),
	}
)

// Treat a hex as a 3-dimensional point
// Slice a 3-dimensional space along x + y + z = 0
// to obtain a generalized representation of a hexpoint
// in which traditional cartesian operations work as usual
// In each point, x + y + z = 0 must hold
type HexPoint struct {
	x int
	y int
	z int
}

func (p *HexPoint) Add(toAdd *HexPoint) *HexPoint {
	return NewHexPoint(p.x+toAdd.x, p.y+toAdd.y, p.z+toAdd.z)
}

func (p *HexPoint) DistanceFromOrigin() float64 {
	origin := NewOrigin()
	return (math.Abs(float64(origin.x-p.x)) + math.Abs(float64(origin.y-p.y)) + math.Abs(float64(origin.z-p.z))) / 2.0
}

func NewHexPoint(x, y, z int) *HexPoint {
	if x+y+z != 0 {
		panic("HexPoint must satisfy x + y + z = 0")
	}
	return &HexPoint{x, y, z}
}

func NewOrigin() *HexPoint {
	return NewHexPoint(0, 0, 0)
}

func Move(step string, p *HexPoint) *HexPoint {
	move, ok := directionalMoves[step]
	if !ok {
		panic(fmt.Sprintf("Failed to find %sin the delta moves", step))
	}
	return p.Add(move)
}

func Parse(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	return strings.Split(scanner.Text(), ",")
}

func main() {
	path := Parse(os.Stdin)
	point := NewOrigin()

	for _, step := range path {
		point = Move(step, point)
	}

	fmt.Println("Went to ", point)
	fmt.Println("Distance to origin:", point.DistanceFromOrigin())
}
