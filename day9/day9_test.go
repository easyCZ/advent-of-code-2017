package main

import (
	"testing"
	"strings"
	"github.com/stretchr/testify/assert"
	"bufio"
)

func TestConsumeGarbage(t *testing.T) {
	scenario := []struct {
		in  string
		out string
	}{
		{"<>", ""},
		{"<random characters>", ""},
		{"<<<<>", ""},
		{"<{!>}>", ""},
		{"<!!>", ""},
		{"<!!!>>", ""},
		{"<!!!>>aaa", "aaa"},
	}

	for _, s := range scenario {
		reader := bufio.NewReader(strings.NewReader(s.in))

		ConsumeGarbage(reader)

		var remaining []byte
		for b, err := reader.ReadByte(); err == nil; b, err = reader.ReadByte() {
			remaining = append(remaining, b)
		}

		assert.Equal(t, s.out, string(remaining))
	}
}

//func TestParseGroup(t *testing.T) {
//	scenarios := []struct{
//		in string
//		out int
//	}{
//		{"{}", 1},
//		{"{{{}}}", 6},
//		{"{{},{}}", 5},
//		{"{{{},{},{{}}}}", 16},
//		{"{<a>,<a>,<a>,<a>}", 1},
//		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
//		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
//		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
//	}
//
//	for _, s := range scenarios {
//		score := DropObsolete(bufio.NewReader(strings.NewReader(s.in)), 0)
//		assert.Equal(t, s.out, score)
//	}
//
//}

func TestDropObsolete(t *testing.T) {
	scenarios := []struct{
		in string
		out string
	}{
		{"{}", "{}"},
		{"{{{}}}", "{{{}}}"},
		{"{{},{}}", "{{}{}}"},
		{"{{{},{},{{}}}}", "{{{}{}{{}}}}"},
		{"{<a>,<a>,<a>,<a>}", "{}"},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", "{{}{}{}{}}"},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", "{{}{}{}{}}"},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", "{{}}"},
	}

	for _, s := range scenarios {
		bytes := DropObsolete(bufio.NewReader(strings.NewReader(s.in)))
		assert.Equal(t, s.out, string(bytes))
	}
}

func TestScoreGroups(t *testing.T) {
	scenarios := []struct {
		in string
		out int
	} {
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}

	for i, s := range scenarios {
		bytes := DropObsolete(bufio.NewReader(strings.NewReader(s.in)))
		score := ScoreGroups(bytes)
		assert.Equal(t, s.out, score, "Row", i)
	}
}
