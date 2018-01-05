package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestToBinary(t *testing.T) {
	scenarios := []struct {
		in  int
		out string
	}{
		{0, "0"},
		{1, "1"},
		{2, "10"},
		{3, "11"},
		{4, "100"},

		{1092455, "100001010101101100111"},
		{430625591, "11001101010101101001100110111"},
	}

	for _, s := range scenarios {
		assert.Equal(t, s.out, ToBinary(s.in))
	}
}

func TestGenerator_Generate(t *testing.T) {
	scenarios := []struct {
		in  *Generator
		out []string
	}{
		{&Generator{
			factor:  16807,
			value:   65,
			divisor: 1,
		}, []string{
			"00000000000100001010101101100111",
			"01000110011001001111011100111001",
			"00001110101000101110001101001010",
			"01100111111110000001011011000111",
			"01010000100111111001100000100100",
		}},
		{&Generator{
			factor:  48271,
			value:   8921,
			divisor: 1,
		}, []string{
			"00011001101010101101001100110111",
			"01001001100010001000010110001000",
			"01010101010100101110001101001010",
			"00001000001101111100110000000111",
			"00010001000000000010100000000100",
		}},
	}

	for _, s := range scenarios {
		var generated []string
		for i := 0; i < 5; i++ {
			generated = append(generated, s.in.Yield())
		}
		assert.Equal(t, s.out, generated)
	}
}

func TestJudge_Observe(t *testing.T) {
	judge := &Judge{}
	count := judge.Observe(
		&Generator{factor: 16807, value: 65, divisor: 1,},
		&Generator{factor: 48271, value: 8921, divisor: 1,},
		5)

	assert.Equal(t, 1, count)
}
