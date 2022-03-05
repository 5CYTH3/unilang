package tokens

type Operator int64
type Token struct {
	op    Operator
	value int
}

func (t *Token) GetOp() Operator {
	return t.op
}

func (t *Token) GetValue() int {
	return t.value
}

const (
	OP_PLUS Operator = iota
	OP_MIN
	OP_PUSH
	OP_DUMP
)

func Plus() Token {
	return Token{OP_PLUS, 0}
}

func Min() Token {
	return Token{OP_MIN, 0}
}

func Push(value int) Token {
	return Token{OP_PUSH, value}
}

func Dump() Token {
	return Token{OP_DUMP, 0}
}
