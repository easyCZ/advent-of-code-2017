package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"strconv"
	"fmt"
	text_scanner "text/scanner"
	"regexp"
	"errors"
	"math"
)

const (
	INC_OP = iota
	DEC_OP
)

const (
	EQUAL              = iota
	NOT_EQUAL
	LESS_THAN
	GREATER_THAN
	LESS_THAN_EQUAL
	GREATER_THAN_EQUAL
)

const (
	VARIABLE = iota
	NUMBER
)

type Register = string

type Expr struct {
	exprType int
	value    string
}

func (e *Expr) Evaluate(environment map[string]int) (int, error) {
	if e.exprType == NUMBER {
		return strconv.Atoi(e.value)
	}
	val, ok := environment[e.value]
	if !ok {
		return 0, errors.New(fmt.Sprintf("Var '%s' does not exist in environment", e.value))
	}
	return val, nil
}

func (e *Expr) AsNumber() (int, error) {
	return strconv.Atoi(e.value)
}

func (e *Expr) String() string {
	return fmt.Sprintf("%s", e.value)
}

type BinOp struct {
	left  *Expr
	right *Expr
	op    int
}

func (b *BinOp) String() string {
	var op string
	switch b.op {
	case EQUAL:
		op = "=="
	case NOT_EQUAL:
		op = "!="
	case LESS_THAN:
		op = "<"
	case LESS_THAN_EQUAL:
		op = "<="
	case GREATER_THAN:
		op = ">"
	case GREATER_THAN_EQUAL:
		op = ">="
	}
	return fmt.Sprintf("%s %s %s", b.left.String(), op, b.right.String())
}

func (b *BinOp) Evaluate(environment map[string]int) (bool, error) {
	left, err := b.left.Evaluate(environment)
	if err != nil {
		return false, err
	}
	right, err := b.right.Evaluate(environment)
	if err != nil {
		return false, err
	}

	switch b.op {
	case EQUAL:
		return left == right, nil
	case NOT_EQUAL:
		return left != right, nil
	case LESS_THAN:
		return left < right, nil
	case LESS_THAN_EQUAL:
		return left <= right, nil
	case GREATER_THAN:
		return left > right, nil
	case GREATER_THAN_EQUAL:
		return left >= right, nil
	default:
		return false, errors.New(fmt.Sprintf("Failed to match OP %s", b.op))
	}
}

type Instruction struct {
	register  Register
	operation int
	value     int
	condition *BinOp
}

func NewInstruction(register Register, op int, val int, cond *BinOp) *Instruction {
	return &Instruction{
		value:     val,
		operation: op,
		register:  register,
		condition: cond,
	}
}

func ParseDigits(scanner *text_scanner.Scanner) string {
	digits := ""

	keepParsing := true
	for keepParsing {
		next := scanner.Peek()
		if IsDigit(string(next)) {
			digits += string(scanner.Next())
		} else {
			keepParsing = false
		}
	}

	return digits
}

func ParseChars(scanner *text_scanner.Scanner) string {
	chars := ""
	keepParsing := true
	for keepParsing {
		next := scanner.Peek()
		if IsChar(string(next)) {
			chars += string(scanner.Next())
		} else {
			keepParsing = false
		}
	}

	return chars
}

func ParseOp(scanner *text_scanner.Scanner) int {
	current := scanner.Next()
	next := scanner.Peek()

	var op string
	if IsOp(string(next)) {
		scanner.Next()
		op = string(current) + string(next)
	} else {
		op = string(current)
	}

	switch op {
	case "==":
		return EQUAL
	case "!=":
		return NOT_EQUAL
	case "<":
		return LESS_THAN
	case "<=":
		return LESS_THAN_EQUAL
	case ">":
		return GREATER_THAN
	case ">=":
		return GREATER_THAN_EQUAL
	default:
		panic(fmt.Sprintf("Unknown OP encountered %s", op))
	}
}

func IsOp(current string) bool {
	return current == "<" || current == ">" || current == "!" || current == "="
}

func IsDigit(current string) bool {
	return regexp.MustCompile(`\d`).Find([]byte(current)) != nil || current == "-"
}

func IsChar(current string) bool {
	return regexp.MustCompile(`[a-zA-Z]`).Find([]byte(current)) != nil
}

//func ParseOp(scanner *text_scanner.Scanner) (int, error) {
//	//current := scanner.
//}

func ParseBinaryOp(raw string) (*BinOp, error) {
	binOp := &BinOp{}

	var scanner text_scanner.Scanner
	scanner.Init(strings.NewReader(raw))

	var left *Expr
	nextChar := string(scanner.Peek())
	if IsChar(nextChar) {
		left = &Expr{
			exprType: VARIABLE,
			value:    ParseChars(&scanner),
		}
	} else if IsDigit(nextChar) {
		left = &Expr{
			value:    ParseDigits(&scanner),
			exprType: NUMBER,
		}
	} else {
		panic(fmt.Sprintf("Unknown char encountered %s", string(nextChar)))
	}

	binOp.left = left

	nextChar = string(scanner.Peek())
	if IsOp(nextChar) {
		binOp.op = ParseOp(&scanner)
	} else {
		panic(fmt.Sprintf("Expected a boolean operation %s", string(nextChar)))
	}

	var right *Expr
	nextChar = string(scanner.Peek())
	if IsChar(nextChar) {
		right = &Expr{
			exprType: VARIABLE,
			value:    ParseChars(&scanner),
		}
	} else if IsDigit(nextChar) {
		right = &Expr{
			value:    ParseDigits(&scanner),
			exprType: NUMBER,
		}
	} else {
		panic(fmt.Sprintf("Unknown char encountered %s", string(nextChar)))
	}
	binOp.right = right

	return binOp, nil
}

func ParseRow(row string) (*Instruction, error) {
	fields := strings.Fields(row)

	op := fields[1]
	var operation int
	if op == "inc" {
		operation = INC_OP
	} else {
		operation = DEC_OP
	}

	val, err := strconv.Atoi(fields[2])
	if err != nil {
		return nil, err
	}

	cond := strings.Join(fields[4:], "")

	binOp, err := ParseBinaryOp(cond)
	if err != nil {
		return nil, err
	}
	return NewInstruction(fields[0], operation, val, binOp), nil
}

func ParseInput(reader io.Reader) ([]*Instruction, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var instructions []*Instruction
	for scanner.Scan() {
		row := scanner.Text()
		inst, err := ParseRow(row)
		if err != nil {
			return instructions, err
		}

		instructions = append(instructions, inst)
	}

	return instructions, nil
}

func Execute(instructions []*Instruction, env map[string]int) (map[string]int, int, error) {
	max := MaxEnvValue(env)
	for _, inst := range instructions {
		shouldExecute, err := inst.condition.Evaluate(env)
		if err != nil {
			return nil, max, err
		}

		if shouldExecute {

			_, ok := env[inst.register]
			if !ok {
				return nil, max, errors.New("Failed to access env for" + inst.register)
			}

			switch inst.operation {
			case DEC_OP:
				env[inst.register] -= inst.value
			case INC_OP:
				env[inst.register] += inst.value
			}

			currentMax := MaxEnvValue(env)
			if currentMax > max {
				max = currentMax
			}

		}
	}

	return env, max, nil
}

func BuildEnvironment(instructions []*Instruction) map[string]int {
	env := make(map[string]int)
	for _, inst := range instructions {
		env[inst.register] = 0
	}

	return env
}

func MaxEnvValue(env map[string]int) int {
	max := math.MinInt32
	for _, val := range env {
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	instructions, err := ParseInput(os.Stdin)
	if err != nil {
		panic(err)
	}
	environment := BuildEnvironment(instructions)

	env, max, err := Execute(instructions, environment)
	if err != nil {
		panic(err)
	}

	fmt.Println("Final max", MaxEnvValue(env))
	fmt.Println("Execution max", max)
}
