package main

import (
	"fmt"
	"bytes"
)

func ToBinary(val int) string {
	return fmt.Sprintf("%b", val)
}

type Generator struct {
	divisor int
	factor int
	value  int
}

func (g *Generator) Yield() string {
	// keep generating until the value is divisible
	for val := g.Generate(); val % g.divisor != 0; val = g.Generate() {}
	return g.ValueAsBinary()
}

func (g *Generator) Generate() int {
	next := g.factor * g.value
	newValue := next % 2147483647
	g.value = newValue
	return newValue
}

func (g *Generator) ValueAsBinary() string {
	var buffer bytes.Buffer
	binary := ToBinary(g.value)
	for i := 0; i < 32-len(binary); i++ {
		buffer.WriteString("0")
	}
	buffer.WriteString(binary)

	return buffer.String()
}

type Judge struct{}

func (j *Judge) Observe(genA, genB *Generator, iterations int) int {
	count := 0
	for i := 0; i < iterations; i++ {
		valA := genA.Yield()
		valB := genB.Yield()

		if valA[16:] == valB[16:] {
			count += 1
		}
	}
	return count
}

func main() {
	generatorA := &Generator{
		value:  783,
		factor: 16807,
		divisor: 1,
	}

	generatorB := &Generator{
		value:  325,
		factor: 48271,
		divisor: 1,
	}

	judge := &Judge{}
	count := judge.Observe(generatorA, generatorB, 40000000)
	fmt.Println("count after 40m iterations of simple generators:", count) // 650

	complexGenA := &Generator{
		value:  783,
		factor: 16807,
		divisor: 4,
	}

	complexGenB := &Generator{
		value:  325,
		factor: 48271,
		divisor: 8,
	}

	complexCount := judge.Observe(complexGenA, complexGenB, 5000000)
	fmt.Println("count after 5m complex iterations:", complexCount)
}
