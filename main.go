package main

import (
	"bufio"
	"fmt"
	"os"

	p "scythe.com/uni/parser"
	t "scythe.com/uni/tokens"
)

func pop(alist *[]int) int {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}

// Compiling task. Takes array of Token as entry and prints out the result.
func interpreter(entry []t.Token) {
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

func compile(arg string) {
	interpreter(p.ParseFile(arg))
}

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
				compile(os.Args[2])
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
