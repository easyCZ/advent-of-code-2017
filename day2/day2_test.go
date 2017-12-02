package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseRow(t *testing.T) {
	scenarios := []struct {
		row string
		out []int
	}{
		{"5 1 9 5", []int{5, 1, 9, 5}},
		{"7	5	3", []int{7, 5, 3}},
		{"2\t4\t6\t8", []int{2, 4, 6, 8}},
	}

	for _, scenario := range scenarios {
		assert.Equal(t, scenario.out, ParseRow(scenario.row))
	}
}

func TestRowDifference(t *testing.T) {
	scenarios := []struct {
		row []int
		out int
	}{
		{[]int{5, 1, 9, 5}, 8},
		{[]int{7, 5, 3}, 4},
		{[]int{2, 4, 6, 8}, 6},
	}

	for _, scenario := range scenarios {
		assert.Equal(t, scenario.out, RowDifference(scenario.row))
	}
}

func TestEvenlyDivisible(t *testing.T) {
	scenarios := []struct {
		row []int
		out int
	}{
		{[]int{5, 9, 2, 8}, 4},
		{[]int{9, 4, 7 ,3}, 3},
		{[]int{3, 8, 6, 5}, 2},
	}

	for _, scenario := range scenarios {
		assert.Equal(t, scenario.out, EvenlyDivisible(scenario.row))
	}
}
