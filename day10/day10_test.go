package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	scenario := []struct {
		arr         []int
		currPos     int
		skipSize    int
		outArr      []int
		currPosOut  int
		skipSizeOut int
	}{
		{
			currPos:     0,
			arr:         []int{0, 1, 2, 3, 4},
			outArr:      []int{2, 1, 0, 3, 4},
			currPosOut:  3,
			skipSize:    0,
			skipSizeOut: 1,
		},
	}

	for _, s := range scenario {
		arr, posOut, skipSize := KnotHashCycle(s.arr, 3, s.currPos, s.skipSize)
		assert.Equal(t, s.outArr, arr)
		assert.Equal(t, s.currPosOut, posOut)
		assert.Equal(t, s.skipSizeOut, skipSize)
	}
}

func TestParseAcii(t *testing.T) {
	scenarios := []struct {
		in  string
		out []int
	}{
		{"1,2,3", []int{49, 44, 50, 44, 51}},
	}

	for _, s := range scenarios {
		assert.Equal(t, s.out, ParseAcii(s.in))
	}
}

func TestKnotHash(t *testing.T) {
	scenarios := []struct {
		input string
		hash  string
	}{
		{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
		{"1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}

	for _, s := range scenarios {
		parsed := ParseAcii(s.input)
		hash := KnotHash(parsed)

		assert.Len(t, hash, 32)
		assert.Equal(t, s.hash, hash)
	}
}

func TestXorSegment(t *testing.T) {

	assert.Equal(t, 64, XorSegment([]int{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22}))

}
