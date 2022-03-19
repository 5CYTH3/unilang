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

func CleanString(str string) []string {
	str = strings.Replace(str, "\n", " ", -1)
	strArr := strings.Split(str, " ")
	return strArr
}

func LexString(input []string) []t.Tokens {
	arr := make([]t.Tokens, 0)
	for _, i := range input {
		arr = append(arr, t.Tokenize(i))
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
		str := CleanString(t_file)
		arr = LexString(str)
	} else {
		panic(err)
	}
	return arr
}

// Parse a line (string) and return the Parse function with the line passed as parameter
func LexLine(line string) []t.Tokens {
	trimmed := strings.Split(line, " ")
	var arr []t.Tokens

	for _, i := range trimmed {
		arr = append(arr, t.Tokenize(i))
	}
	return arr
}
