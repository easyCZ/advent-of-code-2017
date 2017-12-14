package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"io"
)

func Reverse(arr []int) []int {
	var reversed []int
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
}

func HashDigest(ring []int, size, position, skip int) ([]int, int, int) {
	from := position % len(ring)
	to := (position + size) % len(ring)

	var toReverse []int
	if to < from {
		toReverse = append(toReverse, ring[from:]...)
		toReverse = append(toReverse, ring[:to]...)
	} else {
		toReverse = ring[from:to]
	}
	reversed := Reverse(toReverse)
	for i := 0; i < len(reversed); i++ {
		ring[(position+i)%len(ring)] = reversed[i%len(ring)]
	}

	return ring, position + size + skip, skip + 1
}

func Hash(input []int) []int {
	position := 0
	skip := 0
	ring := make([]int, 256)
	for i := 0; i < 256; i++ {
		ring[i] = i
	}

	for _, in := range input {
		ring, position, skip = HashDigest(ring, in, position, skip)
	}

	return ring
}

func Parse(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	return strings.Split(scanner.Text(), ",")
}

func ParseInts(in []string) ([]int, error) {
	var parsed []int
	for _, token := range in {
		num, err := strconv.Atoi(token)
		if err != nil {
			return parsed, err
		}
		parsed = append(parsed, num)
	}
	return parsed, nil
}

func main() {
	input := Parse(os.Stdin)
	inputAsInt, err := ParseInts(input)
	if err != nil {
		panic(err)
	}

	hashed := Hash(inputAsInt)
	fmt.Println("First and second multiplied:", hashed[0]*hashed[1])
}
