package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	scenarios := []struct {
		in  []int
		out []int
	}{
		{[]int{0, 1, 2}, []int{0, 1, 2}},
		{[]int{0, 0, 1}, []int{0, 1}},
	}

	for _, s := range scenarios {
		unique := Uniq(s.in)
		assert.Len(t, unique, len(s.out))

		for _, u := range unique {
			assert.Contains(t, s.out, u)
		}
	}
}

func TestParseRow(t *testing.T) {
	scenarios := []struct {
		in      string
		house   int
		pipedTo []int
	}{
		{"0 <-> 82", 0, []int{82, 0}},
		{"1 <-> 1", 1, []int{1}},
		{"2 <-> 0, 3, 4", 2, []int{0, 2, 3, 4}},
		{"3 <-> 2, 4", 3, []int{3, 2, 4}},
		{"4 <-> 2, 3, 6", 4, []int{4, 2, 3, 6}},
		{"5 <-> 6", 5, []int{5, 6}},
		{"6 <-> 4, 5", 6, []int{6, 4, 5}},
	}

	for _, s := range scenarios {
		house, err := ParseRow(s.in)
		assert.NoError(t, err)
		assert.Equal(t, s.house, house.id)
		assert.Len(t, house.connectsTo, len(s.pipedTo))

		for _, p := range house.connectsTo {
			assert.Contains(t, s.pipedTo, p)
		}
	}
}

func TestFindConnectedGroup(t *testing.T) {
	housesById := map[int]*House{
		0: {id: 0, connectsTo: []int{0, 2}},
		1: {id: 1, connectsTo: []int{1}},
		2: {id: 2, connectsTo: []int{0, 2, 3, 4}},
		3: {id: 3, connectsTo: []int{2, 4, 3}},
		4: {id: 4, connectsTo: []int{4, 2, 3, 6}},
		5: {id: 5, connectsTo: []int{5, 6}},
		6: {id: 6, connectsTo: []int{6, 4, 5}},
	}
	//explored := make(map[int]*House)

	connectedGroup := FindConnectedGroupForHouseZero(housesById)
	assert.Len(t, connectedGroup, 6)
}
