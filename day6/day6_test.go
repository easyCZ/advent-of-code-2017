package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRedistributeBanks(t *testing.T) {
	scenarios := []struct {
		in  []int
		out []int
	}{
		{[]int{0, 2, 7, 0}, []int{2, 4, 1, 2}},
		{[]int{2, 4, 1, 2}, []int{3, 1, 2, 3}},
		{[]int{3, 1, 2, 3}, []int{0, 2, 3, 4}},
		{[]int{0, 2, 3, 4}, []int{1, 3, 4, 1}},
		{[]int{1, 3, 4, 1}, []int{2, 4, 1, 2}},
	}

	for i, s := range scenarios {
		assert.Equal(t, s.out, RedistributeBanks(s.in), "Failed on row %d", i)
	}
}

func TestStepsToRepetition(t *testing.T) {
	banks := []int{0, 2, 7, 0}
	steps, loopSize := StepsToRepetition(banks)
	assert.Equal(t, 5, steps)
	assert.Equal(t, 4, loopSize)
}
