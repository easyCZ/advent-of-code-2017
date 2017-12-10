package main

import (
	"bufio"
	"io"
	"os"
)

type Operation = int

const (
	INC_OP = iota
	DEC_OP
)

type Register = string

type Condition = string

type Instruction struct {
	register  Register
	operation int
	condition Condition
}

func ParseInput(reader io.Reader) []Instruction {
	scanner := bufio.NewScanner(reader)
}

func main() {
	instructions := ParseInput(os.StdIn)

}
