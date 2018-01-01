package main

import (
	"container/list"
	"fmt"
	"bytes"
)

const (
	stateA = iota
	stateB
	stateC
	stateD
	stateE
	stateF
)

type TuringMachine struct {
	state  int
	tape   *list.List
	cursor *list.Element
}

func (tm *TuringMachine) Tape() []int {
	var result []int
	for e := tm.tape.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(int))
	}
	return result
}

func (tm *TuringMachine) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("<%d> -> %d [", tm.state, tm.cursor.Value))
	for e := tm.tape.Front(); e != nil; e = e.Next() {
		buffer.WriteString(fmt.Sprintf("%d, ", e.Value))
	}
	buffer.WriteString("]")
	return buffer.String()
}

func (tm *TuringMachine) MoveRight() {
	next := tm.cursor.Next()
	if next == nil {
		tm.tape.PushBack(0)
	}

	tm.cursor = tm.cursor.Next()
}

func (tm *TuringMachine) MoveLeft() {
	prev := tm.cursor.Prev()
	if prev == nil {
		tm.tape.PushFront(0)
	}

	tm.cursor = tm.cursor.Prev()
}

func (tm *TuringMachine) Write(val int) {
	tm.cursor.Value = val
}

func (tm *TuringMachine) ActivateState(state int) {
	tm.state = state
}

func (tm *TuringMachine) Checksum() int {
	count := 0
	for elem := tm.tape.Front(); elem != nil; elem = elem.Next() {
		if elem.Value == 1 {
			count += 1
		}
	}
	return count
}

func (tm *TuringMachine) ExecuteN(numberOfInstructions int) {
	if numberOfInstructions < 0 {
		return
	}
	for i := 0; i < numberOfInstructions; i++ {
		tm.Execute()
	}
}

// Execute the immediate instruction
func (tm *TuringMachine) Execute() {
	switch tm.state {
	case stateA:
		tm.executeStateA()
	case stateB:
		tm.executeStateB()
	case stateC:
		tm.executeStateC()
	case stateD:
		tm.executeStateD()
	case stateE:
		tm.executeStateE()
	case stateF:
		tm.executeStateF()
	default:
		panic("Unexpected state encountered")
	}
}

func (tm *TuringMachine) executeStateA() {
	switch tm.cursor.Value {
	case 0:
		tm.Write(1)
		tm.MoveRight()
		tm.ActivateState(stateB)
	case 1:
		tm.Write(0)
		tm.MoveLeft()
		tm.ActivateState(stateC)
	default:
		panic("Unexpected value of cursor")
	}
}

func (tm *TuringMachine) executeStateB() {
	switch tm.cursor.Value {
	case 0:
		tm.Write(1)
		tm.MoveLeft()
		tm.ActivateState(stateA)
	case 1:
		tm.Write(1)
		tm.MoveRight()
		tm.ActivateState(stateD)
	default:
		panic("Unexpected value of cursor")
	}
}

func (tm *TuringMachine) executeStateC() {
	switch tm.cursor.Value {
	case 0:
		tm.Write(1)
		tm.MoveRight()
		tm.ActivateState(stateA)
	case 1:
		tm.Write(0)
		tm.MoveLeft()
		tm.ActivateState(stateE)
	default:
		panic("Unexpected value of cursor")
	}
}

func (tm *TuringMachine) executeStateD() {
	switch tm.cursor.Value {
	case 0:
		tm.Write(1)
		tm.MoveRight()
		tm.ActivateState(stateA)
	case 1:
		tm.Write(0)
		tm.MoveRight()
		tm.ActivateState(stateB)

	default:
		panic("Unexpected value of cursor")
	}
}

func (tm *TuringMachine) executeStateE() {
	switch tm.cursor.Value {
	case 0:
		tm.Write(1)
		tm.MoveLeft()
		tm.ActivateState(stateF)
	case 1:
		tm.Write(1)
		tm.MoveLeft()
		tm.ActivateState(stateC)
	default:
		panic("Unexpected value of cursor")
	}
}

func (tm *TuringMachine) executeStateF() {
	switch tm.cursor.Value {
	case 0:
		tm.Write(1)
		tm.MoveRight()
		tm.ActivateState(stateD)
	case 1:
		tm.Write(1)
		tm.MoveRight()
		tm.ActivateState(stateA)

	default:
		panic("Unexpected value of cursor")
	}
}

func NewTuringMachine() *TuringMachine {
	tape := list.New()
	tape.PushFront(0)

	return &TuringMachine{
		state:  stateA,
		cursor: tape.Front(),
		tape:   tape,
	}
}

func main() {
	machine := NewTuringMachine()
	machine.ExecuteN(12919244)
	checksum := machine.Checksum()

	fmt.Println("Checksum:", checksum)
}
