package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	b "scythe.com/uni/src/build"
	p "scythe.com/uni/src/parser"
	t "scythe.com/uni/src/tokens"
)

// Pop and return last element of a list
func pop(alist *[]int) int {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}

func interpreter(entry []t.Tokens) {
	arr := make([]int, 0)
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
		} else if i.GetOp() == t.OP_MUL {
			a := pop(&arr)
			b := pop(&arr)
			arr = append(arr, a*b)
		} else if i.GetOp() == t.OP_DIV {
			a := pop(&arr)
			b := pop(&arr)
			arr = append(arr, b/a)
		} else {
			fmt.Printf("Invalid operator")
		}
	}
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

func Usage() {
	fmt.Println(`
Usage: uni <command> [argument]

Commands:
	- run
	- build [file]`)
}

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		// Build command
		case "build":
			if len(os.Args) >= 3 {
				if strings.HasSuffix(os.Args[2], ".uf") || strings.HasSuffix(os.Args[2], ".uo") {
					b.GenerateAssembly(p.ParseFile(os.Args[2]))
				} else {
					fmt.Println("err: Please provide a valid file. (.uo, .uf)")
					os.Exit(1)
				}
				// Error
			} else {
				fmt.Println("err: Please provide a file for the parsing.")
				fmt.Println("-> Usage: uni build <file>")
			}
		// Interpret
		case "run":
			sim()
		default:
			fmt.Println("err: The command " + "\"" + os.Args[1] + "\"" + " is not valid.")
			os.Exit(1)
		}
	} else {
		Usage()
	}
}
