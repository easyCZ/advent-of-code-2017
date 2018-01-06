package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDance_Spin(t *testing.T) {
	dance := &Dance{
		programs: []string{"a", "b", "c", "d", "e"},
	}

	dance.Spin(0)
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, dance.programs)

	dance.Spin(1)
	assert.Equal(t, []string{"e", "a", "b", "c", "d"}, dance.programs)

	dance.Spin(2)
	assert.Equal(t, []string{"c", "d", "e", "a", "b"}, dance.programs)

	dance.Spin(3)
	assert.Equal(t, []string{"e", "a", "b", "c", "d"}, dance.programs)

	dance.Spin(4)
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, dance.programs)

	dance.Spin(5)
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, dance.programs)

	dance.Spin(6)
	assert.Equal(t, []string{"e", "a", "b", "c", "d"}, dance.programs)
}

func TestDance_Exchange(t *testing.T) {
	dance := &Dance{
		programs: []string{"a", "b", "c", "d", "e"},
	}

	dance.Exchange(0, 1)
	assert.Equal(t, []string{"b", "a", "c", "d", "e"}, dance.programs)

	dance.Exchange(0, 3)
	assert.Equal(t, []string{"d", "a", "c", "b", "e"}, dance.programs)

	dance.Exchange(4, 3)
	assert.Equal(t, []string{"d", "a", "c", "e", "b"}, dance.programs)
}

func TestDance_Partner(t *testing.T) {
	dance := &Dance{
		programs: []string{"a", "b", "c", "d", "e"},
	}

	dance.Partner("a", "b")
	assert.Equal(t, []string{"b", "a", "c", "d", "e"}, dance.programs)

	dance.Partner("e", "a")
	assert.Equal(t, []string{"b", "e", "c", "d", "a"}, dance.programs)
}
