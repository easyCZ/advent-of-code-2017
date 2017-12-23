package main

import (
	"strings"
	"errors"
	"fmt"
	"strconv"
	"io"
	"bufio"
	"os"
)

const upDirection = 1
const downDirection = -1

type Layer struct {
	depth     int
	size      int
	position  int
	direction int // +1 for downwards movement, -1 for upwards
}

func (l *Layer) IsAtTheTop() bool {
	return l.position == 0
}

func (l *Layer) Move() {
	if l.size == 0 {
		return
	}

	next := l.position + l.direction
	if next == 0 {
		l.direction = upDirection
	} else if next == l.size-1 {
		l.direction = downDirection
	}

	l.position = next
}

func NewLayer(depth, size int) *Layer {
	return &Layer{position: 0, size: size, depth: depth, direction: upDirection}
}

type Firewall struct {
	layers []*Layer
	time   int
}

func (f *Firewall) Move() {
	for _, layer := range f.layers {
		if layer != nil {
			layer.Move()
		}
	}
	f.time += 1
}

func TraverseFirewall(firewall *Firewall, delay int) (int) {
	if delay < 0 {
		delay = 0
	}
	for i := 0; i < delay; i++ {
		firewall.Move()
	}

	severity := 0
	for _, layer := range firewall.layers {

		if layer != nil {
			if layer.IsAtTheTop() {
				fmt.Println("conflict", layer.position)
				severity += layer.depth * layer.size
			}
		}

		firewall.Move()
	}

	return severity
}

func TraverseFirewallWithoutGettingCought(firewall *Firewall) int {

	for delay := 0; ; delay++ {
		severity := TraverseFirewall(firewall, delay)
		//fmt.Println("Severity", delay, severity, )
		if delay > 15 {
			return delay
		}
		if severity == 0 {
			return delay
		}
	}
}

func ParseRow(row string) (*Layer, error) {
	tokens := strings.Split(row, ":")
	if len(tokens) != 2 {
		return nil, errors.New(fmt.Sprintf("Not enough tokens when parsing '%s'", row))
	}
	depth, err := strconv.Atoi(tokens[0])
	if err != nil {
		return nil, err
	}

	size, err := strconv.Atoi(strings.TrimSpace(tokens[1]))
	if err != nil {
		return nil, err
	}

	return NewLayer(depth, size), nil
}

func Parse(reader io.Reader) (*Firewall, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var parsedLayers []*Layer
	for scanner.Scan() {
		row := scanner.Text()
		layer, err := ParseRow(row)
		if err != nil {
			return nil, err
		}

		parsedLayers = append(parsedLayers, layer)
	}

	// fill in empty slots to have nice indexes
	layerCount := -1
	for _, layer := range parsedLayers {
		if layer.depth > layerCount {
			layerCount = layer.depth
		}
	}

	layers := make([]*Layer, layerCount+1)
	for _, layer := range parsedLayers {
		layers[layer.depth] = layer
	}

	return &Firewall{layers: layers, time: 0}, nil
}

func main() {
	firewall, err := Parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	//severity := TraverseFirewall(firewall, 0)
	//fmt.Println("Severity of immediate pass:", severity)

	delay := TraverseFirewallWithoutGettingCought(firewall)
	fmt.Println("Must delay to pass without detection [ps]:", delay)
}
