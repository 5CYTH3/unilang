package tokens

import (
	"strconv"

	c "github.com/fatih/color"
)

type Token int64
type Tokens struct {
	op    Token
	value int
}

// Return the Operator of the "Token" struct
func (t *Tokens) GetOp() Token {
	return t.op
}

// Return the int value of the "Token" struct
func (t *Tokens) GetValue() int {
	return t.value
}

// Enum of all operators
const (
	OP_PLUS Token = iota
	OP_MIN
	OP_PUSH
	OP_DUMP
)

// Return a Token with plus operator and value code 0
func Plus() Tokens {
	return Tokens{OP_PLUS, 0}
}

// Return a Token with minus operator and value code 0
func Min() Tokens {
	return Tokens{OP_MIN, 0}
}

// Return a Token with push operator and value passed in parameter
func Push(value int) Tokens {
	return Tokens{OP_PUSH, value}
}

// Return a Token with DUMP operator and value code 0
func Dump() Tokens {
	return Tokens{OP_DUMP, 0}
}

func Tokenize(data []string) []Tokens {
	var stack []Tokens
	for _, i := range data {
		if i == "+" {
			c.Red("Plus")
			stack = append(stack, Plus())
		} else if i == "-" {
			c.Red("Min")
			stack = append(stack, Min())
		} else if i == "dmp" {
			c.Red("Dumped")
			stack = append(stack, Dump())
		} else if num, err := strconv.Atoi(i); err == nil {
			c.Red("Num check passed")
			cyan := c.New(c.FgCyan).Add(c.Underline)
			cyan.Println(num)
			stack = append(stack, Push(num))
		}
	}
	return stack
}
