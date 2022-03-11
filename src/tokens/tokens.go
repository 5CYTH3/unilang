package tokens

import (
	"strconv"
)

type Token int64
type Tokens struct {
	op    Token
	value int
}

// Enum of all operators
const (
	OP_PLUS Token = iota // addition
	OP_MIN               // minus
	OP_MUL               // multiplication
	OP_DIV               // division
	OP_PUSH
	OP_DUMP
	OP_DEFAULT
	L_PAREN
	R_PAREN
)

// Return the Operator of the "Token" struct
func (t *Tokens) GetOp() Token {
	return t.op
}

// Return the int value of the "Token" struct
func (t *Tokens) GetValue() int {
	return t.value
}

// Return a Token with plus operator and value code 0
func Plus() Tokens {
	return Tokens{OP_PLUS, 1}
}

func Default() Tokens {
	return Tokens{OP_DEFAULT, 0}
}

// Return a Token with minus operator and value code 0
func Min() Tokens {
	return Tokens{OP_MIN, 1}
}

// Return a Token with multiplication operator and value code 0
func Mul() Tokens {
	return Tokens{OP_MUL, 1}
}

// Return a Token with division operator and value code 0
func Div() Tokens {
	return Tokens{OP_DIV, 1}
}

// Return a Token with push operator and value passed in parameter
func Push(value int) Tokens {
	return Tokens{OP_PUSH, value}
}

// Return a Token with DUMP operator and value code 0
func Dump() Tokens {
	return Tokens{OP_DUMP, 1}
}

func L_Paren() Tokens {
	return Tokens{L_PAREN, 2}
}

func R_Paren() Tokens {
	return Tokens{R_PAREN, 2}
}

// Parse a string array and append for each chars an operator to an array. Then, return the array
func Tokenize(word string) Tokens {
	if word == "+" {
		return Plus()
	} else if word == "-" {
		return Min()
	} else if word == "*" {
		return Mul()
	} else if word == "/" {
		return Div()
	} else if word == "dmp" {
		return Dump()
	} else if word == "(" {
		return L_Paren()
	} else if word == ")" {
		return R_Paren()
	} else {
		num, _ := strconv.Atoi(word)
		return Push(num)
	}
}
