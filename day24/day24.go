package main

import (
	"io"
	"bufio"
	"strings"
	"errors"
	"strconv"
	"os"
	"fmt"
)

type Part struct {
	left  int
	right int
}

func (p *Part) Connects(part *Part) bool {
	return p.left == part.left || p.right == part.right
}

func FindEligibleParts(port int, parts []*Part) []*Part {
	var eligible []*Part
	for _, part := range parts {
		if part.left == port || part.right == port {
			eligible = append(eligible, part)
		}
	}

	return eligible
}

func FindBridgesForPort(port int, parts []*Part) ([][]*Part) {
	fmt.Println("Searching for", port, parts)
	eligible := FindEligibleParts(port, parts)

	fmt.Println("Eligible", eligible)

	var allBridges [][]*Part
	for _, part := range eligible {
		var remaining []*Part
		for _, p := range parts {
			if p != part {
				remaining = append(remaining, p)
			}
		}

		bridges := []*Part{part}
		var continuations [][]*Part
		if part.left == port {
			continuations = FindBridgesForPort(part.right, remaining)
		} else {
			continuations = FindBridgesForPort(part.left, remaining)
		}

		for _, c := range continuations {

		}

		for i := range bridges {
			bridges[i] = append([]*Part{part}, bridges[i]...)
		}
		fmt.Println("bridges", bridges)

		//for _, bridge := range bridges {
		//	enhanced := append([]*Part{part}, bridge...)
		//	allBridges = append(allBridges, enhanced)
		//	fmt.Println("all bridges", enhanced)
		//}

	}

	return allBridges
}

func EnumerateBridges(parts []*Part) [][]*Part {
	return FindBridgesForPort(0, parts)
}

func Parse(reader io.Reader) ([]*Part, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var parts []*Part
	for scanner.Scan() {
		row := scanner.Text()
		tokens := strings.Split(row, "/")
		if len(tokens) != 2 {
			return nil, errors.New("Malformed input " + row)
		}

		left, err := strconv.Atoi(tokens[0])
		if err != nil {
			return nil, err
		}
		right, err := strconv.Atoi(tokens[1])
		if err != nil {
			return nil, err
		}

		parts = append(parts, &Part{left, right})
	}

	return parts, nil
}

func main() {
	parts, err := Parse(os.Stdin)
	if err != nil {
		panic(err)
	}

	//var path []Part
	bridges := FindBridgesForPort(0, parts)
	fmt.Println("There are n bridges available", len(bridges))

}
