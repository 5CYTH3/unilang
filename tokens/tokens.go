package tokens

type Operator int64
type Token struct {
	op    Operator
	value int
}

// Return the Operator of the "Token" struct
func (t *Token) GetOp() Operator {
	return t.op
}

// Return the int value of the "Token" struct
func (t *Token) GetValue() int {
	return t.value
}

// Enum of all operators
const (
	OP_PLUS Operator = iota
	OP_MIN
	OP_PUSH
	OP_DUMP
)

// Return a Token with plus operator and value code 0
func Plus() Token {
	return Token{OP_PLUS, 0}
}

// Return a Token with minus operator and value code 0
func Min() Token {
	return Token{OP_MIN, 0}
}

// Return a Token with push operator and value passed in parameter
func Push(value int) Token {
	return Token{OP_PUSH, value}
}

// Return a Token with DUMP operator and value code 0
func Dump() Token {
	return Token{OP_DUMP, 0}
}
