package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestDance_Spin(t *testing.T) {
	dance := NewDance(5)

	dance.Spin(0)
	assert.Equal(t, []uint8{'a', 'b', 'c', 'd', 'e'}, dance.programs)

	dance.Spin(1)
	assert.Equal(t, []uint8{'e', 'a', 'b', 'c', 'd'}, dance.programs)

	dance.Spin(2)
	assert.Equal(t, []uint8{'c', 'd', 'e', 'a', 'b'}, dance.programs)

	dance.Spin(3)
	assert.Equal(t, []uint8{'e', 'a', 'b', 'c', 'd'}, dance.programs)

	dance.Spin(4)
	assert.Equal(t, []uint8{'a', 'b', 'c', 'd', 'e'}, dance.programs)

	dance.Spin(5)
	assert.Equal(t, []uint8{'a', 'b', 'c', 'd', 'e'}, dance.programs)

	dance.Spin(6)
	assert.Equal(t, []uint8{'e', 'a', 'b', 'c', 'd'}, dance.programs)
}

func TestDance_Exchange(t *testing.T) {
	dance := NewDance(5)

	dance.Exchange(0, 1)
	assert.Equal(t, []uint8{'b', 'a', 'c', 'd', 'e'}, dance.programs)

	dance.Exchange(0, 3)
	assert.Equal(t, []uint8{'d', 'a', 'c', 'b', 'e'}, dance.programs)

	dance.Exchange(4, 3)
	assert.Equal(t, []uint8{'d', 'a', 'c', 'e', 'b'}, dance.programs)
}

func TestDance_Partner(t *testing.T) {
	dance := NewDance(5)

	dance.Partner('a', 'b')
	assert.Equal(t, []uint8{'b', 'a', 'c', 'd', 'e'}, dance.programs)

	dance.Partner('e', 'a')
	assert.Equal(t, []uint8{'b', 'e', 'c', 'd', 'a'}, dance.programs)
}

func BenchmarkDance_Spin(b *testing.B) {
	file, err := os.Open("/Users/milan/workspace/advent-of-code-2017/day16/input.txt")
	if err != nil {
		b.Error(err)
	}
	moves, err := ParseMoves(file)
	if err != nil {
		b.Error(err)
	}
	dance := NewDance(16)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dance.PerformMoves(moves)
	}
}
