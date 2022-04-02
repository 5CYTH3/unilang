package tokens

import (
	"strconv"

	"scythe.com/uni/src/util"
)

type Token int64
type Tokens struct {
	op       Token
	value    int64
	priority int8
	isOp     bool
	rAssoc   bool
}

var tokenmap = map[string]Tokens{
	"*":   Mul(),
	"/":   Div(),
	"+":   Plus(),
	"-":   Min(),
	"(":   L_Paren(),
	")":   R_Paren(),
	"dmp": Dump(),
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
func (t *Tokens) GetValue() int64 {
	return t.value
}

func (t *Tokens) GetPriority() int8 {
	return t.priority
}

func Default() Tokens {
	return Tokens{OP_DEFAULT, 0, -1, false, false}
}

// Return a Token with plus operator and value code 0
func Plus() Tokens {
	return Tokens{OP_PLUS, 0, 2, true, false}
}

// Return a Token with minus operator and value code 0
func Min() Tokens {
	return Tokens{OP_MIN, 0, 2, true, false}
}

// Return a Token with multiplication operator and value code 0
func Mul() Tokens {
	return Tokens{OP_MUL, 0, 3, true, false}
}

// Return a Token with division operator and value code 0
func Div() Tokens {
	return Tokens{OP_DIV, 0, 3, true, false}
}

// Return a Token with push operator and value passed in parameter
func Push(value int64) Tokens {
	return Tokens{OP_PUSH, value, 0, false, false}
}

// Return a Token with DUMP operator and value code 0
func Dump() Tokens {
	return Tokens{OP_DUMP, 0, 8, false, false}
}

func L_Paren() Tokens {
	return Tokens{L_PAREN, 0, 0, false, false}
}

func R_Paren() Tokens {
	return Tokens{R_PAREN, 0, 0, false, false}
}

func InfixToPostfix(arr []Tokens) []Tokens {
	operatorStack := make([]Tokens, 0)
	postFixTerms := make([]Tokens, 0)

	for i := 0; i <= len(arr)-1; i++ {
		token := arr[i]
		switch token.op {
		case L_PAREN:
			operatorStack = append(operatorStack, token)
		case R_PAREN:
			for {
				a := util.Pop(&operatorStack)
				if a.op == L_PAREN {
					break
				}
				postFixTerms = append(postFixTerms, a)
			}
			util.Pop(&operatorStack)
		default:
			if token.isOp {
				for len(operatorStack) > 0 {
					top := operatorStack[len(operatorStack)-1]
					if !top.isOp || token.priority > top.priority || token.priority == top.priority && token.rAssoc {
						break
					}
					a := util.Pop(&operatorStack)
					postFixTerms = append(postFixTerms, a)
				}
				operatorStack = append(operatorStack, token)
			} else {
				postFixTerms = append(postFixTerms, token)
			}
		}
	}

	for len(operatorStack) > 0 {
		a := util.Pop(&operatorStack)
		postFixTerms = append(postFixTerms, a)
	}
	return postFixTerms
}

// Parse a string array and append for each chars an operator to an array. Then, return the array
func ParseTokenAsOperator(word string) Tokens {
	num, err := strconv.Atoi(word)
	if err != nil {
		return tokenmap[word]
	} else {
		return Push(int64(num))
	}
}
