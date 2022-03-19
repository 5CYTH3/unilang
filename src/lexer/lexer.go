package parser

import (
	"os"
	"strings"

	t "scythe.com/uni/src/tokens"
)

/*
func InfixToRPN(arr []t.Tokens) []t.Tokens {
	var stack util.Stack
	var stackArray []t.Tokens
	for i := 0; i < len(arr); i++ {
		if arr[i].GetOp() == t.OP_PUSH {
			j := i
			for ; j < len(arr) && arr[j].GetOp() == t.OP_PUSH; j++ {
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
			stackArray = append(stackArray, arr[i])
		} else {
			for !stack.Empty() {
				top := stack.Top()
				if top == t.L_Paren() || !(top.GetPriority() >= arr[i].GetPriority()) {
					break
				}
				stackArray = append(stackArray, top)
				stack.Pop()
			}
			stack.Push(arr[i])
		}
	}
	return stackArray
}
*/

func CleanString(str string) []string {
	str = strings.Replace(str, "\n", " ", -1)
	strArr := strings.Split(str, " ")
	return strArr
}

func LexString(input []string) []t.Tokens {
	arr := make([]t.Tokens, 0)
	for _, i := range input {
		arr = append(arr, t.ParseTokenAsOperator(i))
	}
	return arr
}

// Parse a file and return the an array of tokens from a splitted string
func LexFile(file string) []t.Tokens {
	var arr []t.Tokens

	// File reading -> splitted file
	f, err := os.ReadFile(file)
	if err == nil {
		t_file := string(f)
		arr = LexString(CleanString(t_file))
	} else {
		panic(err)
	}
	return arr
}
