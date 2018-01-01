package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"container/list"
)

func TestTuringMachine_MoveRight_WithNewMachine(t *testing.T) {
	tm := NewTuringMachine()
	tm.MoveRight()

	assert.Equal(t, 2, tm.tape.Len())
	assert.Equal(t, 0, tm.cursor.Value)
	assert.Equal(t, 0, tm.cursor.Prev().Value)
}

func TestTuringMachine_MoveRight_WithExistingMachine(t *testing.T) {
	tape := list.New()
	tape.PushFront(0)
	tape.PushFront(1)
	cursor := tape.Front()

	tm := &TuringMachine{
		cursor: cursor,
		tape:   tape,
		state:  stateA,
	}

	// does not generate a new element
	tm.MoveRight()
	assert.Equal(t, 2, tm.tape.Len())
	assert.Equal(t, 0, tm.cursor.Value, "value does not match")

	// generated a new element
	tm.MoveRight()
	assert.Equal(t, 3, tm.tape.Len())
	assert.Equal(t, 0, tm.cursor.Value)
}

func TestTuringMachine_MoveLeft_WithNewMachine(t *testing.T) {
	tm := NewTuringMachine()
	tm.MoveLeft()

	assert.Equal(t, 2, tm.tape.Len())
	assert.Equal(t, 0, tm.cursor.Value)
	assert.Equal(t, 0, tm.cursor.Next().Value)
}

func TestTuringMachine_MoveLeft_WithExistingMachine(t *testing.T) {
	tape := list.New()
	tape.PushFront(0)
	tape.PushFront(1)
	// tape = [1, 0]
	cursor := tape.Back()

	tm := &TuringMachine{
		cursor: cursor,
		tape:   tape,
		state:  stateA,
	}

	// does not generate a new element
	tm.MoveLeft()
	assert.Equal(t, 2, tm.tape.Len())
	assert.Equal(t, 1, tm.cursor.Value, "value does not match")

	// generated a new element
	tm.MoveLeft()
	assert.Equal(t, 3, tm.tape.Len())
	assert.Equal(t, 0, tm.cursor.Value)
}

func TestTuringMachine_Movement(t *testing.T) {
	tm := NewTuringMachine()
	tm.MoveLeft()
	tm.MoveLeft()
	tm.MoveRight()
	tm.MoveRight() // where we originally started
	tm.MoveRight()
	assert.Equal(t, 4, tm.tape.Len())
}

func TestTuringMachine_Execute(t *testing.T) {
	scenarios := []struct {
		startState    int
		startValue    int
		expectedTape  []int
		expectedState int
	}{
		// State A
		{stateA, 0, []int{1, 0}, stateB},
		{stateA, 1, []int{0, 0}, stateC},
		// State B
		{stateB, 0, []int{0, 1}, stateA},
		{stateB, 1, []int{1, 0}, stateD},
		// State C
		{stateC, 0, []int{1, 0}, stateA},
		{stateC, 1, []int{0, 0}, stateE},
		// State D
		{stateD, 0, []int{1, 0}, stateA},
		{stateD, 1, []int{0, 0}, stateB},
		// State E
		{stateE, 0, []int{0, 1}, stateF},
		{stateE, 1, []int{0, 1}, stateC},
		// State F
		{stateF, 0, []int{1, 0}, stateD},
		{stateF, 1, []int{1, 0}, stateA},
	}

	for _, s := range scenarios {
		tm := NewTuringMachine()
		tm.state = s.startState
		tm.cursor.Value = s.startValue

		tm.Execute()

		assert.Equal(t, s.expectedTape, tm.Tape())
		assert.Equal(t, s.expectedState, tm.state)
	}
}
