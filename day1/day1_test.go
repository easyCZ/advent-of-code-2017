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
		assert.Equal(t, test.out, sum_consequtive(test.numbers))
	}

}
