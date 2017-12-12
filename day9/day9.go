package main

import (
	"os"
	"bufio"
	"fmt"
)

func ConsumeIgnored(reader *bufio.Reader) {
	reader.Discard(1)
}

func ConsumeGarbage(reader *bufio.Reader) int {
	size := 0
	for b, err := reader.ReadByte(); err == nil; b, err = reader.ReadByte() {
		current := string(b)

		switch current {
		case "!":
			ConsumeIgnored(reader)
		case ">":
			return size
		default:
			size += 1
		}
	}

	return size
}

func DropObsolete(reader *bufio.Reader) ([]byte, int) {
	var bytes []byte
	junkSize := 0

	for b, err := reader.ReadByte(); err == nil; b, err = reader.ReadByte() {
		current := string(b)

		if current == "!" {
			ConsumeIgnored(reader)
		}

		if current == "{" {
			x, junk := DropObsolete(reader)
			junkSize += junk
			bytes = append(bytes, []byte("{")...)
			bytes = append(bytes, x...)
		}

		if current == "}" {
			bytes = append(bytes, []byte("}")...)
		}

		if current == "<" {
			junk := ConsumeGarbage(reader)
			junkSize += junk

		}
	}

	return bytes, junkSize
}

func ScoreGroups(in []byte) (int) {

	depth := 0
	score := 0

	for i := 0; i < len(in); i++ {
		current := string(in[i])

		switch current {
		case "{":
			depth += 1

		case "}":
			score += depth
			depth -= 1
		}
	}
	return score
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	out, junk := DropObsolete(reader)
	score := ScoreGroups(out)
	fmt.Println("Score:", score)
	fmt.Println("Junk:", junk)
}
