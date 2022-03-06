package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	p "scythe.com/uni/parser"
	t "scythe.com/uni/tokens"
)

// Pop and return last element of a list
func pop(alist *[]int) int {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}
	
// Compiling task. Takes array of Token as entry and prints out the result.
func interpreter(entry []t.Tokens) {
	var arr []int
	for _, i := range entry {
		if i.GetOp() == t.OP_PUSH {
			arr = append(arr, i.GetValue())
		} else if i.GetOp() == t.OP_PLUS {
			a := pop(&arr)
			b := pop(&arr)
			arr = append(arr, a+b)
		} else if i.GetOp() == t.OP_DUMP {
			a := pop(&arr)
			fmt.Println(a)
		} else if i.GetOp() == t.OP_MIN {
			a := pop(&arr)
			b := pop(&arr)
			arr = append(arr, a-b)
		} else {
			fmt.Println("parsing error: Invalid syntax.")
			os.Exit(1)
		}
	}
}

func GenerateAssembly(entry []t.Tokens) {
	// var arr []int
	for _, i := range entry {
		if i.GetOp() == t.OP_PUSH {
			// Asm
		} else if i.GetOp() == t.OP_PLUS {
			// Asm
			fmt.Println()	
		} else if i.GetOp() == t.OP_DUMP {
			// Asm
		} else if i.GetOp() == t.OP_MIN {
			// Asm
		} else {
			fmt.Println("parsing error: Invalid syntax.")
			os.Exit(1)
		}
	}
}

// Interpret the file from ParseFile function
func compile(arg string) {
	interpreter(p.ParseFile(arg))
}

// Interpret the user input from ParseLine function
func sim() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("$uni-> ")
		reader.Scan()
		stack := p.ParseLine(reader.Text())
		interpreter(stack)
	}
}

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "compile":
			if len(os.Args) >= 3 {
				if strings.HasSuffix(os.Args[3], ".uf") || strings.HasSuffix(os.Args[3], ".uo") {
					compile(os.Args[3])
				} else {
					fmt.Println("err: Please provide a valid file. (.uo, .uf)")
				}
			} else {
				fmt.Println("err: Please provide a file for the parsing.")
				fmt.Println("-> Usage: uni compile <file>")
			}
		case "test":
			sim()
		}
	} else {
		fmt.Println(`
Usage: uni <command> [argument]

Commands:
	- test
	- compile [file]
	`)
	}
}
