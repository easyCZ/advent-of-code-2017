package main

import (
	"bufio"
	"io"
	"strconv"
	"os"
	"fmt"
)

func Parse(file io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var jumps []int

	for scanner.Scan() {
		row := scanner.Text()
		num, err := strconv.Atoi(row)
		if err != nil {
			return jumps, err
		}
		jumps = append(jumps, num)
	}

	return jumps, nil
}

func ExecuteStrangeInstruction(cursor int, jumps []int) (int, []int) {
	jumpValue := jumps[cursor]
	jumps[cursor] += 1
	return cursor + jumpValue, jumps
}

func EscapeStrangeMaze(jumps []int) int {
	steps := 0

	for cursor := 0; cursor >= 0 && cursor < len(jumps); steps++ {
		cursor, jumps = ExecuteStrangeInstruction(cursor, jumps)
	}

	return steps
}

func ExecuteStrangerInstruction(cursor int, jumps []int) (int, []int) {
	jumpValue := jumps[cursor]

	if jumpValue >= 3{
		jumps[cursor] -= 1
	} else {
		jumps[cursor] += 1
	}

	return cursor + jumpValue, jumps
}

func EscapeStrangerMaze(jumps []int) int {
	steps := 0

	for cursor := 0; cursor >= 0 && cursor < len(jumps); steps++ {
		cursor, jumps = ExecuteStrangerInstruction(cursor, jumps)
	}

	return steps
}

func main() {
	jumps, err := Parse(os.Stdin)
	if err != nil {
		panic(err)
	}

	// since we're modifying the jumps in place
	jumpsCopy := make([]int, len(jumps))
	copy(jumpsCopy[:], jumps[:])

	fmt.Println("Steps to escape strange maze:", EscapeStrangeMaze(jumps))
	fmt.Println("Steps to escape stranger maze:", EscapeStrangerMaze(jumpsCopy))
}
