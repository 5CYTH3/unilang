package parser

import (
	"fmt"
	"os"
	"strings"

	t "scythe.com/uni/src/tokens"
)

// Pop and return last element of a list
func pop(alist *[]int) int {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}

func InfixToRPN(arr []t.Tokens) []t.Tokens {
	stack := make([]t.Tokens, 0)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i].GetOp() == t.OP_PUSH {

		} else if arr[i].GetOp() == t.OP_PLUS {
			
		} else if arr[i].GetOp() == t.OP_DUMP {

		} else if arr[i].GetOp() == t.OP_MIN {

		} else if arr[i].GetOp() == t.OP_MUL {

		} else if arr[i].GetOp() == t.OP_DIV {

		} else {
			fmt.Printf("Invalid operator")
		}
	}

	return arr
}

// Parse a file and return the an array of tokens from a splitted string
func ParseFile(file string) []t.Tokens {

	// File reading -> splitted file
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	t_file := string(f)
	t_file = strings.Replace(t_file, "\n", " ", -1)
	splitted := strings.Split(t_file, " ")

	// Array that'll hold all operators
	var arr []t.Tokens

	// Appending to "arr" all the tokens associated to words
	for _, i := range splitted {
		arr = append(arr, t.Tokenize(i))
	}
	return arr
}

// Parse a line (string) and return the Parse function with the line passed as parameter
func ParseLine(line string) []t.Tokens {
	trimmed := strings.Split(line, " ")
	var arr []t.Tokens

	for _, i := range trimmed {
		arr = append(arr, t.Tokenize(i))
	}
	return arr
}
