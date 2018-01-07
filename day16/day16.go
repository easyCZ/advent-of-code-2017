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

type Move struct {
	moveType int
	spin     Spin
	exchange Exchange
	partner  Partner
}

type Spin struct {
	size int
}

type Exchange struct {
	left  int
	right int
}

type Partner struct {
	left  uint8
	right uint8
}

func NewDance(size int) *Dance {
	var programs []uint8
	for i := 0; i < size; i++ {
		val := 'a' + uint8(i)
		programs = append(programs, val)
	}
	return &Dance{programs}
}

type Dance struct {
	programs []uint8
}

func (d *Dance) PerformMovesRepeatedly(count int, moves []*Move) error {
	cache := make(map[string]int)
	var first int
	var second int
	for i := 1; i <= count; i++ {
		//if i%10000 == 0 {
		//	fmt.Println("Iteration: ", i)
		//}
		err := d.PerformMoves(moves)
		if err != nil {
			return err
		}
		order := d.Order()
		fmt.Println(i, order)
		index, ok := cache[order]
		if ok {
			first = index
			second = i
			//fmt.Println("first:", first, "second", second)
			//fmt.Println(order, cache)
			break
		}
		cache[order] = i
	}

	if first == 0 && second == 0 {
		return nil
	}

	loopSize := second - first + 1
	fmt.Println("count - first", count-second)
	remaining := (count - first) % loopSize
	fmt.Println("Remaining", remaining)
	for i := 0; i < remaining; i++ {
		fmt.Println("Remaining iteration: ", i)
		err := d.PerformMoves(moves)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dance) PerformMoves(moves []*Move) error {
	for _, move := range moves {
		err := d.PerformMove(move)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dance) PerformMove(move *Move) error {
	switch move.moveType {
	case spin:
		d.Spin(move.spin.size)
	case partner:
		d.Partner(move.partner.left, move.partner.right)
	case exchange:
		d.Exchange(move.exchange.left, move.exchange.right)
	default:
		fmt.Println(fmt.Sprintf("%v", move))
		panic("Unexpected type of move")
	}
	return nil
}

func (d *Dance) Spin(size int) {
	size = size % len(d.programs)
	split := len(d.programs) - size

	var programs []uint8
	programs = append(programs, d.programs[split:]...)
	programs = append(programs, d.programs[:split]...)
	d.programs = programs
}

func (d *Dance) Exchange(a, b int) {
	a = a % len(d.programs)
	b = b % len(d.programs)

	tempA := d.programs[a]

	d.programs[a] = d.programs[b]
	d.programs[b] = tempA
}

func (d *Dance) Partner(a uint8, b uint8) {
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
		buffer.WriteString(string(program))
	}
	return buffer.String()
}

func ParseMoves(reader io.Reader) ([]*Move, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var buffer bytes.Buffer
	for scanner.Scan() {
		text := scanner.Text()
		buffer.WriteString(text)
	}

	tokens := strings.Split(buffer.String(), ",")
	var moves []*Move
	for _, token := range tokens {
		var move *Move
		switch token[0] {
		case 's':
			val, err := strconv.Atoi(token[1:])
			if err != nil {
				return nil, err
			}
			move = &Move{
				moveType: spin,
				spin:     Spin{size: val},
			}
		case 'x':
			left, right, err := parseExchange(token[1:])
			if err != nil {
				return nil, err
			}
			move = &Move{
				moveType: exchange,
				exchange: Exchange{right: right, left: left},
			}
		case 'p':
			left, right, err := parsePartner(token[1:])
			if err != nil {
				return nil, err
			}
			move = &Move{
				moveType: partner,
				partner:  Partner{right: right, left: left},
			}
		default:
			panic("unexpected move: " + token)
		}
		moves = append(moves, move)

	}
	return moves, nil
}

func parsePartner(value string) (uint8, uint8, error) {
	tokens := strings.Split(value, "/")
	if len(tokens) != 2 {
		return 0, 0, errors.New("Unexpected number of partner attributes " + value)
	}
	return tokens[0][0], tokens[1][0], nil
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

func main() {
	moves, err := ParseMoves(os.Stdin)
	if err != nil {
		panic(err)
	}
	dance := NewDance(16)

	dance.PerformMoves(moves)
	fmt.Println("Order after moves: ", dance.Order())

	longDance := NewDance(16)
	//longDance.PerformMovesRepeatedly(60, moves)
	//fmt.Println("after 60", longDance.Order())
	//longDance.PerformMovesRepeatedly(60, moves)
	//fmt.Println("after 120", longDance.Order())

	longDance.PerformMovesRepeatedly(1000000000, moves)
	fmt.Println("Order after 1billion dances: ", longDance.Order())

}
