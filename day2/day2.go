package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"math"
)

func ParseRow(row string) []int {
	tokens := strings.Fields(row)
	numbers := make([]int, 0)

	for _, token := range tokens {
		number, err := strconv.Atoi(token)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func RowDifference(numbers []int) int {
	if len(numbers) == 0 {
		panic("Expected at least one number in a row")
	}
	min := math.MaxInt32
	max := math.MinInt32

	for _, number := range numbers {
		if number <= min {
			min = number
		}
		if number >= max {
			max = number
		}
	}

	return max - min
}

func EvenlyDivisible(numbers []int) int {
	if len(numbers) <= 1 {
		panic("Expected at least two numbers in a row")
	}

	for i, numerator := range numbers {
		for j, denominator := range numbers {
			if i == j {
				continue
			}

			remainder := math.Remainder(float64(numerator), float64(denominator))
			if remainder == 0 {
				return numerator / denominator
			}
		}
	}

	panic("Failed to find two numbers which are divisible")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	checksum := 0

	for scanner.Scan() {
		row := scanner.Text()
		parsedRow := ParseRow(row)
		rowDiff := EvenlyDivisible(parsedRow)

		checksum += rowDiff
	}
	fmt.Println(checksum)
}
