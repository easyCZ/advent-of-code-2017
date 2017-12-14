package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"io"
	"bytes"
)

var (
	suffix = []int{17, 31, 73, 47, 23}
)

func Reverse(arr []int) []int {
	var reversed []int
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
}

func KnotHashCycle(ring []int, length, position, skip int) ([]int, int, int) {
	from := position % len(ring)
	to := (position + length) % len(ring)

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

	return ring, position + length + skip, skip + 1
}

func KnotHashRound(lengths []int) []int {
	position := 0
	skip := 0
	ring := make([]int, 256)
	for i := 0; i < 256; i++ {
		ring[i] = i
	}

	for _, length := range lengths {
		ring, position, skip = KnotHashCycle(ring, length, position, skip)
	}

	return ring
}

func XorSegment(segment []int) int {
	result := segment[0]
	for i := 1; i < len(segment); i++ {
		result = result ^ segment[i]
	}
	return result
}

func IntToHex(val, size int) string {
	hex := fmt.Sprintf("%x", val)
	if len(hex) < size {
		for i := 0; i < size-len(hex); i++ {
			hex = "0" + hex
		}
	}
	return hex
}

func KnotHash(lengths []int) string {
	withSuffix := append(lengths, suffix...)

	ring := make([]int, 256)
	for i := 0; i < 256; i++ {
		ring[i] = i
	}

	position := 0
	skip := 0
	for i := 0; i < 64; i++ {
		for _, length := range withSuffix {
			ring, position, skip = KnotHashCycle(ring, length, position, skip)
		}
	}

	var denseHash []int
	for i := 16; i <= len(ring); i += 16 {
		segment := ring[i-16: i]
		xor := XorSegment(segment)
		denseHash = append(denseHash, xor)
	}

	var buffer bytes.Buffer
	for _, d := range denseHash {
		buffer.WriteString(IntToHex(d, 2))
	}
	hex := buffer.String()
	return hex
}

func ReadLine(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	return scanner.Text()
}

func ParseInts(line string) ([]int, error) {
	tokens := strings.Split(line, ",")
	var parsed []int
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			return parsed, err
		}
		parsed = append(parsed, num)
	}
	return parsed, nil
}

func ParseAcii(line string) []int {
	var ascii []int
	for _, b := range []byte(line) {
		ascii = append(ascii, int(b))
	}
	return ascii
}

func main() {
	line := ReadLine(os.Stdin)
	inputAsInt, err := ParseInts(line)
	if err != nil {
		panic(err)
	}

	firstRound := KnotHashRound(inputAsInt)
	fmt.Println("First and second multiplied:", firstRound[0]*firstRound[1])

	ascii := ParseAcii(line)
	hash := KnotHash(ascii)

	fmt.Println("Hash:", hash)
}
