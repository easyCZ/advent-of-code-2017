package main

import (
	"strings"
	"errors"
	"fmt"
	"strconv"
	"io"
	"bufio"
	"os"
	"bytes"
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

func (l *Layer) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i := 0; i < l.size; i++ {
		if i == l.position {
			buffer.WriteString("[S]")
		} else {
			buffer.WriteString("[ ]")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

func (l *Layer) Clone() *Layer {
	return &Layer{
		position:  l.position,
		size:      l.size,
		depth:     l.depth,
		direction: l.direction,
	}
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

func (f *Firewall) String() string {
	var buffer bytes.Buffer
	for _, l := range f.layers {
		if l == nil {
			buffer.WriteString("[]")
		} else {
			buffer.WriteString(l.String())
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}

func (f *Firewall) Clone() *Firewall {
	layers := make([]*Layer, 0)
	for _, layer := range f.layers {
		if layer == nil {
			layers = append(layers, nil)
		} else {
			cloned := layer.Clone()
			layers = append(layers, cloned)
		}
	}
	return &Firewall{
		layers: layers,
		time:   f.time,
	}
}

func TraverseFirewall(firewall *Firewall, delay int) (int) {
	//fmt.Println("initial", firewall.String())
	for i := 0; i < delay; i++ {
		firewall.Move()
	}
	//fmt.Println("after delay", delay)
	//fmt.Println(firewall.String())
	severity := 0
	for _, layer := range firewall.layers {

		if layer != nil {
			if layer.IsAtTheTop() {
				//fmt.Println("conflict", layer.String())
				severity += layer.depth * layer.size
			}
		}

		firewall.Move()
	}

	return severity
}

func ParseFromFile(filepath string) (*Firewall, error) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	return Parse(file)
}

func TraverseFirewallWithoutGettingCought(firewall *Firewall) int {
	cloned := firewall.Clone()
	found := false
	delay := 0
	for !found {
		in := cloned.Clone()
		fmt.Println("Delay", delay)
		fmt.Println(in)
		severity := TraverseFirewall(in, delay)
		if severity == 0 {
			found = true
			return delay
		}
		delay = delay + 1
	}
	return delay
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
