package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPassphraseValid(t *testing.T) {
	scenarios := []struct {
		in  string
		out bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}

	for _, s := range scenarios {
		assert.Equal(t, s.out, IsPassphraseValid(s.in))
	}
}

func TestIsAnagram(t *testing.T) {
	scenarios := []struct {
		in  [2]string
		out bool
	}{
		{[2]string{"ecdab", "abcde"}, true},
		{[2]string{"abcde", "fghij"}, false},
		{[2]string{"oiii", "ioii"}, true},
	}

	for _, s := range scenarios {
		assert.Equal(t, s.out, IsAnagram(s.in[0], s.in[1]), "Expected IsAnagram(%s, %s) == %t", s.in[0], s.in[1], s.out)
	}
}

func TestIsPassphraseValidWithoutAnagrams(t *testing.T) {
	scenarios := []struct {
		in  string
		out bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
	}

	for _, s := range scenarios {
		assert.Equal(t, s.out, IsPassphraseValidWithoutAnagrams(s.in), "expected", s.in)
	}
}
