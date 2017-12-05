package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestExecuteStrangeInstruction(t *testing.T) {
	scenarios := []struct {
		cursorIn  int
		jumpsIn   []int
		cursorOut int
		jumpsOut  []int
	}{
		{
			0, []int{0, 3, 0, 1, -3},
			0, []int{1, 3, 0, 1, -3},
		},
		{
			0, []int{1, 3, 0, 1, -3},
			1, []int{2, 3, 0, 1, -3},
		},
		{
			1, []int{2, 3, 0, 1, -3},
			4, []int{2, 4, 0, 1, -3},
		},
		{
			4, []int{2, 4, 0, 1, -3},
			1, []int{2, 4, 0, 1, -2},
		},
		{
			1, []int{2, 4, 0, 1, -2},
			5, []int{2, 5, 0, 1, -2},
		},
	}

	for i, s := range scenarios {
		nextCursor, nextJumps := ExecuteStrangeInstruction(s.cursorIn, s.jumpsIn)
		assert.Equal(t, s.cursorOut, nextCursor, "Cursor must match %d", i)
		assert.Equal(t, s.jumpsOut, nextJumps, "Cursor must match %d", i)
	}
}

func TestEscapeStrangeMaze(t *testing.T) {
	jumps := []int{0, 3, 0, 1, -3}
	steps := EscapeStrangeMaze(jumps)

	assert.Equal(t, 5, steps)
}

func TestExecuteStrangerInstruction(t *testing.T) {
	scenarios := []struct {
		cursorIn  int
		jumpsIn   []int
		cursorOut int
		jumpsOut  []int
	}{
		{
			0, []int{0, 3, 0, 1, -3},
			0, []int{1, 3, 0, 1, -3},
		}, {
			0, []int{1, 3, 0, 1, -3},
			1, []int{2, 3, 0, 1, -3},
		}, {
			1, []int{2, 3, 0, 1, -3},
			4, []int{2, 2, 0, 1, -3},
		}, {
			4, []int{2, 2, 0, 1, -3},
			1, []int{2, 2, 0, 1, -2},
		}, {
			1, []int{2, 2, 0, 1, -2},
			3, []int{2, 3, 0, 1, -2},
		}, {
			3, []int{2, 3, 0, 1, -2},
			4, []int{2, 3, 0, 2, -2},
		}, {
			4, []int{2, 3, 0, 2, -2},
			2, []int{2, 3, 0, 2, -1},
		}, {
			2, []int{2, 3, 0, 2, -1},
			2, []int{2, 3, 1, 2, -1},
		}, {
			2, []int{2, 3, 1, 2, -1},
			3, []int{2, 3, 2, 2, -1},
		}, {
			3, []int{2, 3, 2, 2, -1},
			5, []int{2, 3, 2, 3, -1},
		},
	}

	for i, s := range scenarios {
		nextCursor, nextJumps := ExecuteStrangerInstruction(s.cursorIn, s.jumpsIn)
		assert.Equal(t, s.cursorOut, nextCursor, "Cursor must match %d", i)
		assert.Equal(t, s.jumpsOut, nextJumps, "Cursor must match %d", i)
	}
}

func TestEscapeStrangerMaze(t *testing.T) {
	jumps := []int{0, 3, 0, 1, -3}
	steps := EscapeStrangerMaze(jumps)

	assert.Equal(t, 10, steps)
}
