package main

import (
	"io"
	"bufio"
	"strings"
	"strconv"
	"os"
	"fmt"
)

func ParseInput(reader io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	row := strings.Fields(scanner.Text())

	var parsed []int
	for _, p := range row {
		num, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, num)
	}
	return parsed, nil
}

func BankToDistribute(banks []int) int {
	index := 0
	max := banks[0]

	for i, b := range banks {
		if b > max {
			max = b
			index = i
		}
	}
	return index
}

func RedistributeBanks(banks []int) []int {
	toDistribute := BankToDistribute(banks)
	newBanks := make([]int, len(banks))
	copy(newBanks[:], banks[:])

	itemsToDistribute := newBanks[toDistribute]
	newBanks[toDistribute] = 0
	for i := 0; i < itemsToDistribute; i++ {
		newBanks[(toDistribute+i+1)%len(newBanks)] += 1
	}

	return newBanks
}

// get steps to repetition as the number of steps and the loop size
func StepsToRepetition(banks []int) (int, int) {
	seen := make(map[string]int)
	var i int
	stop := false
	for i = 0; !stop; {
		newBanks := RedistributeBanks(banks)
		i += 1

		key := fmt.Sprint(newBanks)
		iteration, ok := seen[key]
		if ok {
			return i, i - iteration
		} else {
			seen[key] = i
		}
		banks = newBanks

	}
	return i, i
}

func main() {
	banks, err := ParseInput(os.Stdin)
	if err != nil {
		panic(err)
	}

	steps, loopSize := StepsToRepetition(banks)
	fmt.Println("Takes steps:", steps)
	fmt.Println("Loop size:", loopSize)
}
