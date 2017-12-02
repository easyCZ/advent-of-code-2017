package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func sum_nth_consequtive(numbers []int, n int) int {
	total := 0

	for index := 0; index < len(numbers); index++ {
		current := numbers[index]
		next := numbers[(index + n) % len(numbers)]

		if current == next {
			total += current
		}
	}

	return total
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	numbers := make([]int, 0)
	for _, char := range line {
		if char == '\n' {
			continue
		}
		num, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}

	fmt.Println(sum_nth_consequtive(numbers, len(numbers) / 2))
}

