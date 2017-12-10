package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	text_scanner "text/scanner"
	"strings"
)

func TestParseRow(t *testing.T) {
	scenarios := []struct {
		row  string
		inst *Instruction
	}{
		{"b inc 5 if a > 1", &Instruction{
			register:  "b",
			value:     5,
			operation: INC_OP,
			condition: &BinOp{
				left:  &Expr{value: "a", exprType: VARIABLE},
				right: &Expr{value: "1", exprType: NUMBER},
				op:    GREATER_THAN,
			},
		}},
		{"a inc 1 if b < 5", &Instruction{
			register:  "a",
			value:     1,
			operation: INC_OP,
			condition: &BinOp{
				left:  &Expr{value: "b", exprType: VARIABLE},
				right: &Expr{value: "5", exprType: NUMBER},
				op:    LESS_THAN,
			},
		}},
		{"c dec -10 if a >= 1", &Instruction{
			register:  "c",
			value:     -10,
			operation: DEC_OP,
			condition: &BinOp{
				left:  &Expr{value: "a", exprType: VARIABLE},
				right: &Expr{value: "1", exprType: NUMBER},
				op:    GREATER_THAN_EQUAL,
			},
		}},
		{"c inc -20 if c == 10", &Instruction{
			register:  "c",
			value:     -20,
			operation: INC_OP,
			condition: &BinOp{
				left:  &Expr{value: "c", exprType: VARIABLE},
				right: &Expr{value: "10", exprType: NUMBER},
				op:    EQUAL,
			},
		}},
		{"c dec -20 if c <= 10", &Instruction{
			register:  "c",
			value:     -20,
			operation: DEC_OP,
			condition: &BinOp{
				left:  &Expr{value: "c", exprType: VARIABLE},
				right: &Expr{value: "10", exprType: NUMBER},
				op:    LESS_THAN_EQUAL,
			},
		}},
		{"c dec -20 if c != 10", &Instruction{
			register:  "c",
			value:     -20,
			operation: DEC_OP,
			condition: &BinOp{
				left:  &Expr{value: "c", exprType: VARIABLE},
				right: &Expr{value: "10", exprType: NUMBER},
				op:    NOT_EQUAL,
			},
		}},
	}

	for i, s := range scenarios {
		inst, err := ParseRow(s.row)
		assert.NoError(t, err)
		assert.Equal(t, s.inst, inst, "Row %d", i)
	}
}

func TestIsChar(t *testing.T) {
	scenarios := []struct {
		in  string
		out bool
	}{
		{"a", true},
		{"b", true},
		{"1", false},
		{"0", false},
		{"A", true},
		{"<", false},
		{"=", false},
	}

	for _, s := range scenarios {
		assert.Equal(t, s.out, IsChar(s.in), "In: %s", s.in)
	}
}

func TestParseChars(t *testing.T) {
	scenarios := []struct {
		in  string
		out string
	}{
		{"a", "a"},
		{"aaa", "aaa"},
		{"aaa1", "aaa"},
	}

	for _, s := range scenarios {
		var scanner text_scanner.Scanner
		scanner.Init(strings.NewReader(s.in))
		expr := ParseChars(&scanner)

		assert.Equal(t, s.out, expr)
	}
}

func TestParseBinaryOp(t *testing.T) {
	scenarios := []struct {
		in  string
		out *BinOp
	}{
		{"a>1", &BinOp{
			left:  &Expr{value: "a", exprType: VARIABLE},
			right: &Expr{value: "1", exprType: NUMBER},
			op:    GREATER_THAN,
		}},
		{"b<5", &BinOp{
			left:  &Expr{value: "b", exprType: VARIABLE},
			right: &Expr{value: "5", exprType: NUMBER},
			op:    LESS_THAN,
		}},
		{"a>=1", &BinOp{
			left:  &Expr{value: "a", exprType: VARIABLE},
			right: &Expr{value: "1", exprType: NUMBER},
			op:    GREATER_THAN_EQUAL,
		}},
		{"c==10", &BinOp{
			left:  &Expr{value: "c", exprType: VARIABLE},
			right: &Expr{value: "10", exprType: NUMBER},
			op:    EQUAL,
		}},
		{"c<=10", &BinOp{
			left:  &Expr{value: "c", exprType: VARIABLE},
			right: &Expr{value: "10", exprType: NUMBER},
			op:    LESS_THAN_EQUAL,
		}},
		{"c!=10", &BinOp{
			left:  &Expr{value: "c", exprType: VARIABLE},
			right: &Expr{value: "10", exprType: NUMBER},
			op:    NOT_EQUAL,
		}},
	}

	for i, s := range scenarios {
		binOp, _ := ParseBinaryOp(s.in)
		assert.Equal(t, s.out, binOp, "Row %d", i)
	}
}
