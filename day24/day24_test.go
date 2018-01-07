package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func GetParts() []*Part {
	return []*Part{
		{0,2},
		{2,2},
		{2,3},
		{3,4},
		{3,5},
		{0,1},
		{10,1},
		{9,10},
	}
}

func TestEnumerateBridges(t *testing.T) {
	bridges := EnumerateBridges(GetParts())

	assert.Len(t, bridges, 11)
	//0/1
	//0/1--10/1
	//0/1--10/1--9/10
	//0/2
	//0/2--2/3
	//0/2--2/3--3/4
	//0/2--2/3--3/5
	//0/2--2/2
	//0/2--2/2--2/3
	//0/2--2/2--2/3--3/4
	//0/2--2/2--2/3--3/5
}