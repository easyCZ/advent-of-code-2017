package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindNextSquare(t *testing.T) {
	scenarios := []struct {
		addr int
		out  int
	}{
		{1, 1},
		{5, 3},
		{7, 3},
		{25, 5},
		{24, 5},
		{17, 5},
	}

	for _, scenario := range scenarios {
		assert.Equal(t, scenario.out, findNextSquare(scenario.addr))
	}
}

func TestDistanceToAccessPort(t *testing.T) {
	scenarios := []struct {
		address int
		out     int
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}

	for _, scenario := range scenarios {
		assert.Equal(t, scenario.out, ManhattanDistanceToAccessPort(scenario.address))
	}
}

func TestSumAdjacent(t *testing.T) {
	scenarios := []struct {
		in  [][]int
		out int
	}{
		{[][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
			8,
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 40,
		},
	}

	for _, sc := range scenarios {
		assert.Equal(t, sc.out, sumAdjacent(sc.in, 1, 1))
	}
}
