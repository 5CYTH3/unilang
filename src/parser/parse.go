package parser

import (
	"fmt"
	"os"
	"strings"

	t "scythe.com/uni/src/tokens"
	"scythe.com/uni/src/util"
)

func InfixToRPN(arr []t.Tokens) []t.Tokens {
	var stack util.Stack
	var stackArray []t.Tokens
	for i := 0; i < len(arr); i++ {
		if arr[i].GetOp() == t.OP_PUSH {
			j := i
			for ; j < len(arr) && arr[j].GetOp() == t.OP_PUSH; j++ {
				fmt.Println(arr[i])
				stackArray = append(stackArray, arr[j])
			}
			i = j - 1

		} else if arr[i].GetOp() == t.L_PAREN {
			stack.Push(arr[i])
		} else if arr[i].GetOp() == t.R_PAREN {
			for !stack.Empty() {
				lrp := stack.Top()
				if lrp.GetOp() == t.L_PAREN {
					break
				}
				stackArray = append(stackArray, lrp)
				stack.Pop()

			}
			stack.Pop()
		} else if arr[i].GetOp() == t.OP_DUMP {
			fmt.Println(arr[i])
			stackArray = append(stackArray, arr[i])
		} else {
			for !stack.Empty() {
				top := stack.Top()
				if top == t.L_Paren() || !(top.GetPriority() >= arr[i].GetPriority()) {
					fmt.Println(arr[i])
					break
				}
				stackArray = append(stackArray, top)
				stack.Pop()
			}
			stack.Push(arr[i])
		}
	}
	fmt.Println(stackArray)
	return stackArray
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
