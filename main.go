package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	t "scythe.com/uni/tokens"
)

func pop(alist *[]int) int {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}

// TEMPORARY
/*
func parse(file string) Operator {
	var stack []interface{}
	trimmed := strings.Split(file, " ")

	for _, i := range trimmed {
		switch i {
		case "+":
			plus()
		case "-":
			min()
		case "":
		}
	}
}
*/

// Compiling task. Takes array of Token as entry and prints out the result.
func compile(entry []t.Token) {
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

func parseFile(file string) {
	/* f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	t_file := string(f)
	trimmed := strings.Split(t_file, " ")
	fmt.Println(trimmed)*/

	program := []t.Token{
		t.Push(3),
		t.Push(18),
		t.Plus(),
		t.Dump(),
	}

	compile(program)

}

func parseLine(line string) []t.Token {
	var stack []t.Token
	trimmed := strings.Split(line, " ")
	for _, i := range trimmed {
		switch i {
		case "+":
			fmt.Println("Plus")
			stack = append(stack, t.Plus())
		case "-":
			fmt.Println("Min")
			stack = append(stack, t.Min())
		case "dmp":
			fmt.Println("Dumped")
			stack = append(stack, t.Dump())
		default:
			fmt.Println("Num called")
			if num, err := strconv.Atoi(i); err == nil {
				fmt.Println("Num check passed")
				fmt.Println(num)
				stack = append(stack, t.Push(num))
			}
		}
	}

	return stack

}

func interpret() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("$uni-> ")
		reader.Scan()
		stack := parseLine(reader.Text())
		compile(stack)
	}
}

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "compile":
			if len(os.Args) >= 3 {
				parseFile(os.Args[2])
			} else {
				fmt.Println("err: Please provide a file for the parsing.")
				fmt.Println("-> Usage: uni compile <file>")
			}
		case "test":
			interpret()
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
