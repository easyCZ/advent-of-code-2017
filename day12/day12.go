package main

import (
	"strings"
	"strconv"
	"errors"
	"io"
	"bufio"
	"os"
	"fmt"
)

type House struct {
	id         int
	connectsTo []int
}

func Uniq(in []int) []int {
	uniq := make(map[int]bool)
	for _, i := range in {
		uniq[i] = true
	}
	var acc []int
	for key, val := range uniq {
		if val {
			acc = append(acc, key)
		}
	}
	return acc
}

func Parse(reader io.Reader) (map[int]*House, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	houses := make(map[int]*House)
	for scanner.Scan() {
		row := scanner.Text()
		house, err := ParseRow(row)
		if err != nil {
			return nil, err
		}
		houses[house.id] = house
	}

	return houses, nil
}

func ParseRow(row string) (*House, error) {
	// 2 <-> 0, 3, 4
	tokens := strings.SplitN(row, " <-> ", 2)
	if len(tokens) != 2 {
		return nil, errors.New("could not split")
	}
	house, err := strconv.Atoi(tokens[0])
	if err != nil {
		return nil, err
	}

	var pipes []int
	pipeTokens := strings.Split(tokens[1], ",")
	for _, p := range pipeTokens {
		num, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return nil, err
		}
		pipes = append(pipes, num)
	}

	pipes = append(pipes, house)

	return &House{connectsTo: Uniq(pipes), id: house}, nil
}

func FindConnectedGroup(id int, housesById map[int]*House, explored map[int]bool) []*House {

	start := housesById[id]
	connectedTo := start.connectsTo
	explored[id] = true

	connected := make([]*House, 0)
	connected = append(connected, start)

	for _, houseId := range connectedTo {
		_, hasBeenExplored := explored[houseId]
		explored[houseId] = true

		if !hasBeenExplored && houseId != id {
			neighbours := FindConnectedGroup(houseId, housesById, explored)
			connected = append(connected, neighbours...)
		}

	}

	return connected
}

func FindConnectedGroupForHouseZero(houses map[int]*House) []*House {
	explored := make(map[int]bool)
	return FindConnectedGroup(0, houses, explored)
}

func main() {
	houses, err := Parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	connectedGroup := FindConnectedGroupForHouseZero(houses)
	fmt.Println("Connected group of House 0:", len(connectedGroup))
}
