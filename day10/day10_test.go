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
		arr, posOut, skipSize := HashDigest(s.arr, 3, s.currPos, s.skipSize)
		assert.Equal(t, s.outArr, arr)
		assert.Equal(t, s.currPosOut, posOut)
		assert.Equal(t, s.skipSizeOut, skipSize)
	}
}
