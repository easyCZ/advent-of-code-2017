package main

import (
	"strconv"
	"strings"
	"errors"
	"os"
	"io"
	"bufio"
	"bytes"
	"fmt"
)

const (
	spin     = iota
	exchange
	partner
)

func parsePartner(value string) (string, string, error) {
	tokens := strings.Split(value, "/")
	if len(tokens) != 2 {
		return "", "", errors.New("Unexpected number of partner attributes " + value)
	}
	return tokens[0], tokens[1], nil
}

func parseExchange(value string) (int, int, error) {
	tokens := strings.Split(value, "/")
	if len(tokens) != 2 {
		return 0, 0, errors.New("Unexpected number of exchange attributes " + value)
	}

	left, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, 0, err
	}

	right, err := strconv.Atoi(tokens[1])
	if err != nil {
		return 0, 0, err
	}

	return left, right, nil
}

type Dance struct {
	programs []string
}

func (d *Dance) PerformMoves(moves []Move) error {
	for _, move := range moves {
		err := d.PerformMove(move)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dance) PerformMove(move Move) error {
	switch move.move {
	case spin:
		val, err := strconv.Atoi(move.value)
		if err != nil {
			return err
		}
		d.Spin(val)
	case exchange:
		left, right, err := parseExchange(move.value)
		if err != nil {
			return err
		}
		d.Exchange(left, right)
	case partner:
		left, right, err := parsePartner(move.value)
		if err != nil {
			return err
		}
		d.Partner(left, right)
	}

	return nil
}

func (d *Dance) Spin(size int) {
	size = size % len(d.programs)
	split := len(d.programs) - size
	head := d.programs[: split]
	tail := d.programs[split:]

	d.programs = append(tail, head...)
}

func (d *Dance) Exchange(a, b int) {
	a = a % len(d.programs)
	b = b % len(d.programs)
	temp := d.programs[a]
	d.programs[a] = d.programs[b]
	d.programs[b] = temp
}

func (d *Dance) Partner(a string, b string) {
	var left int
	var right int
	for i, p := range d.programs {
		if p == a {
			left = i
		}
		if p == b {
			right = i
		}
	}
	d.Exchange(left, right)
}

func (d *Dance) Order() string {
	var buffer bytes.Buffer
	for _, program := range d.programs {
		buffer.WriteString(program)
	}
	return buffer.String()
}

type Move struct {
	move  int
	value string
}

func ParseMoves(reader io.Reader) []Move {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var buffer bytes.Buffer
	for scanner.Scan() {
		text := scanner.Text()
		buffer.WriteString(text)
	}

	tokens := strings.Split(buffer.String(), ",")
	var moves []Move
	for _, token := range tokens {
		var move Move
		move.value = token[1:]
		if token[0] == 's' {
			move.move = spin
		} else if token[0] == 'x' {
			move.move = exchange
		} else if token[0] == 'p' {
			move.move = partner
		} else {
			panic("unexpected move: " + token)
		}

		moves = append(moves, move)

	}
	return moves
}

func main() {
	moves := ParseMoves(os.Stdin)
	fmt.Println(moves)
	dance := &Dance{
		programs: []string{
			"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
		},
	}

	dance.PerformMoves(moves)
	fmt.Println("Order after moves: ", dance.Order())

}
