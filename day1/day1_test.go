package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSumConsequtive(t *testing.T) {
	var scenarios = []struct {
		numbers []int
		out     int
	}{
		{[]int{1, 2, 3, 4}, 0 },
		{[]int{1, 1, 2, 2}, 3},
		{[]int{1, 1, 1, 1}, 4},
		{[]int{9, 1, 2, 1, 2, 1, 2, 9}, 9},
	}
	for _, test := range scenarios {
		assert.Equal(t, test.out, sum_nth_consequtive(test.numbers, 1))
	}

}

func TestSumNthConsequtive(t *testing.T) {
	var scenarios = []struct {
		numbers []int
		out     int
	}{
		{[]int{1, 2, 1, 2}, 6 },
		{[]int{1, 2, 2, 1}, 0},
		{[]int{1, 2, 3, 1, 2, 3}, 12},
		{[]int{1, 2, 1, 3, 1, 4, 1, 5}, 4},
	}
	for _, test := range scenarios {
		assert.Equal(t, test.out, sum_nth_consequtive(test.numbers, len(test.numbers) / 2))
	}

}
